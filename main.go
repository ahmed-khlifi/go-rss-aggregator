package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("Hello, World!")

	// Get the value of the PORT environment variable
	godotenv.Load()
	portString := os.Getenv("PORT")
	if portString == "" {
		// log.Fatal will print the error and then call os.Exit(1)
		log.Fatal("PORT environment variable not set")
	}

	// `router := chi.NewRouter()` is creating a new instance of a router from the `chi` package. This
	// router will be used to define the routes and handle the incoming HTTP requests.
	router := chi.NewRouter()

	// The code `srv := &http.Server{ Handler: router, Addr: ":" + portString, }` is creating a new
	// instance of the `http.Server` struct.
	srv := &http.Server{
		Handler: router,
		Addr:   ":" + portString,
	}

	// The code `err := srv.ListenAndServe()` is starting the HTTP server and listening for incoming
	// requests. If there is an error starting the server, it will be assigned to the `err` variable.
	fmt.Printf("Server is listening on port %s", portString)
	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
  
}