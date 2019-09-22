package main

import (
	"encoding/json"
	"log"
	"net/http"

	sssa "github.com/SSSaaS/sssa-golang"
)

var (
	apiServerError       = apiSimpleResponse{Error: "Generic Server Error", statusCode: http.StatusInternalServerError}
	apiBadRequestError   = apiSimpleResponse{Error: "Bad Request", statusCode: http.StatusBadRequest}
	apiUnauthorizedError = apiSimpleResponse{Error: "Unauthorized", statusCode: http.StatusUnauthorized}
	apiNotFoundError     = apiSimpleResponse{Error: "Not Found", statusCode: http.StatusNotFound}
	apiOK                = apiSimpleResponse{Message: "success", statusCode: http.StatusOK}
)

// JSONRespondWith - handles JSON response with Status code
func JSONRespondWith(wr http.ResponseWriter, resp apiSimpleResponse) {
	wr.WriteHeader(resp.statusCode)
	j := json.NewEncoder(wr)
	if err := j.Encode(resp); err != nil {
		log.Printf("Error Encoding JSON: %s", err)
	}
}

func jsonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

func createRequest(wr http.ResponseWriter, req *http.Request) {
	v := struct {
		Minimum int    `json:"minimum,omitempty"`
		Shares  int    `json:"shares,omitempty"`
		Secret  string `json:"secret,omitempty"`
	}{}
	defer req.Body.Close()
	if err := json.NewDecoder(req.Body).Decode(&v); err != nil {
		log.Println("json.NewDecoder().Decode", err)
		JSONRespondWith(wr, apiBadRequestError)
		return
	}

	// JSON good - build shards
	shards, err := sssa.Create(v.Minimum, v.Shares, v.Secret)
	if err != nil {
		log.Println("sssa.Create", err)
		JSONRespondWith(wr, apiServerError)
	}

	json.NewEncoder(wr).Encode(struct {
		Shards []string `json:"shards,omitempty"`
	}{
		Shards: shards,
	})
}

func combineRequest(wr http.ResponseWriter, req *http.Request) {
	v := struct {
		Shards []string `json:"shards,omitempty"`
	}{}
	defer req.Body.Close()
	if err := json.NewDecoder(req.Body).Decode(&v); err != nil {
		log.Println("json.NewDecoder().Decode", err)
		JSONRespondWith(wr, apiBadRequestError)
		return
	}

	// The library sharts itself if the shard length != 88
	for i := range v.Shards {
		if len(v.Shards[i]) != 88 {
			JSONRespondWith(wr, apiBadRequestError)
			return
		}
	}

	// JSON good - combine shards
	secret, err := sssa.Combine(v.Shards)
	if err != nil {
		log.Println("sssa.Combine", err)
		JSONRespondWith(wr, apiServerError)
		return
	}

	json.NewEncoder(wr).Encode(struct {
		Secret string `json:"secret,omitempty"`
	}{
		Secret: secret,
	})
}
