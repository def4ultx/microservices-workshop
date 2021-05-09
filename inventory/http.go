package main

import (
	"encoding/json"
	"net/http"
)

func writeSuccessResponse(w http.ResponseWriter, resp interface{}) {
	w.Header().Add("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&resp)
}

func writeErrorResponse(w http.ResponseWriter, code int) {

	type Response struct {
		Message string
	}

	resp := Response{
		Message: http.StatusText(code),
	}

	w.Header().Add("Content-type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(&resp)
}
