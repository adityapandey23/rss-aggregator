package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")
	portString := os.Getenv("PORT")
	if portString == "" {
		log.Fatal("Port is not found in the environment") // Exits the program
	}

	router := chi.NewRouter() // Creating a router

	srv := &http.Server{ // Making a server
		Handler: router,
		Addr:    ":" + portString,
	}

	log.Printf("Server starting on port:%v", portString)
	err := srv.ListenAndServe() // Making the server listen and serve to the incoming requests
	if err != nil {
		log.Fatal("Error", err)
	}
}
