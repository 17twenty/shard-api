package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	port := flag.Int("port", 4000, "Port to listen on")
	flag.Parse()
	log.SetFlags(log.Lshortfile | log.Ldate | log.Lmicroseconds)

	r := mux.NewRouter()
	r.HandleFunc("/create", createRequest).Methods("POST", "OPTIONS")
	r.HandleFunc("/combine", combineRequest).Methods("POST", "OPTIONS")

	r.Use(jsonMiddleware, handlers.CORS(
		handlers.AllowedHeaders([]string{"content-type"}),
		handlers.AllowedOrigins([]string{"http://localhost:8001", "https://app.shardomatic.com"}),
		handlers.AllowCredentials()),
	)

	srv := &http.Server{
		Handler: r,
		Addr:    fmt.Sprintf("0.0.0.0:%d", *port),
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Println("Starting shard-api on port", *port, "...")
	log.Fatal(srv.ListenAndServe())
}
