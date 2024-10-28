package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	// Marshal = data structure -> JSON
	// Unmarshal = JSON -> data structure
	dat, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Failed to marshal JSON response: %v", err)
		w.WriteHeader(500) // Internal server error
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(dat)
}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	if code > 499 { // Means this is server side error (if in range of 400 means that it's client side error)
		log.Printf("Responding with %v error: %v", code, msg)
	}
	type errResponse struct {
		Error string `json:"error"` // Struct tags
	}
	respondWithJSON(w, code, errResponse{
		Error: msg,
	})
}
