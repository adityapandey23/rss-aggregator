package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/adityapandey23/rss-aggregator/internal/database"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type apiConfig struct {
	DB *database.Queries
}

func main() {
	godotenv.Load(".env")

	// For PORT
	portString := os.Getenv("PORT")
	if portString == "" {
		log.Fatal("Port is not found in the environment") // Exits the program
	}

	// For Database connection
	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		log.Fatal("DB_URL is not found in the environment") // Exits the program
	}

	conn, conn_err := sql.Open("postgres", dbURL)
	if conn_err != nil {
		log.Fatal("Can't connect to database")
	}

	apiCfg := apiConfig{
		DB: database.New(conn),
	}

	router := chi.NewRouter() // Creating a router

	router.Use(cors.Handler(cors.Options{ // Technically wouldn't be so permissive
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	// New routers
	v1Router := chi.NewRouter()
	v1Router.Get("/healthz", handlerReadiness) // Used the Get function to scope the request to only get requests
	v1Router.Get("/err", handlerErr)
	v1Router.Post("/users", apiCfg.handlerCreateUser)

	// Mounting the routes
	router.Mount("/v1", v1Router)

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
