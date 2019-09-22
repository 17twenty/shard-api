package main

type apiSimpleResponse struct {
	Error      string `json:"error,omitempty"`
	Message    string `json:"message,omitempty"`
	statusCode int
}
