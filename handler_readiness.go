package main

import "net/http"

func handlerReadiness(w http.ResponseWriter, r *http.Request) { // Make sure to follow this function definition
	respondWithJSON(w, 200, struct{}{})
}
