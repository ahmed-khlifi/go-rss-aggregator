package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/ahmed-khlifi/go-rss-aggregator/internal/database"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"

	_ "github.com/lib/pq"
)

type apiConfig struct {
	DB *database.Queries
}

func main() {
	fmt.Println("Hello, World!")

	// Get the value of the PORT environment variable
	godotenv.Load()
	portString := os.Getenv("PORT")
	if portString == "" {
		// log.Fatal will print the error and then call os.Exit(1)
		log.Fatal("PORT environment variable not set")
	}

	db := os.Getenv("DB_URL")
	if db == "" {
		// log.Fatal will print the error and then call os.Exit(1)
		log.Fatal("Database conection URL is not set")
	}


	conn, err := sql.Open("postgres", db)
	if err != nil {
		log.Fatal("Can't connect to the database:",err)
	}

	apiCfg := apiConfig{
		DB:  database.New(conn),
	}
	// `router := chi.NewRouter()` is creating a new instance of a router from the `chi` package. This
	// router will be used to define the routes and handle the incoming HTTP requests.
	router := chi.NewRouter()

	// The code `router.Use(cors.Handler(cors.Options{ AllowedOrigins: []string{"*"}, })` is adding a
	// middleware to the router that will handle CORS (Cross-Origin Resource Sharing) requests. This
	// middleware will allow requests from any origin.
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"*"}, // Use this to allow specific origin sites. the * will allow all sites
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}, // Use this to allow all methods
		AllowedHeaders: []string{"*"},
		ExposedHeaders: []string{"Link"},
		AllowCredentials : false, 
		// `MaxAge: 300` is setting the maximum age (in seconds) for the CORS preflight response.
		MaxAge: 300,	
	}))


	v1Router := chi.NewRouter()
	v1Router.Get("/status", handlerReadiness)
	v1Router.Get("/error", handleError)
	v1Router.Post("/users", apiCfg.handleCreatUser)

	router.Mount("/v1", v1Router)
	// GET : /v1/status  => Should return 200 with empty JSON response

	// The code `srv := &http.Server{ Handler: router, Addr: ":" + portString, }` is creating a new
	// instance of the `http.Server` struct.
	srv := &http.Server{
		Handler: router,
		Addr:   ":" + portString,
	}

	// The code `err := srv.ListenAndServe()` is starting the HTTP server and listening for incoming
	// requests. If there is an error starting the server, it will be assigned to the `err` variable.
	fmt.Printf("Server is listening on port %s", portString)
	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
  
}