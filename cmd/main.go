package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	// initialize router
	r := mux.NewRouter()

	// set up CORS
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"},
	})

	// wrap router with CORS middleware
	handler := c.Handler(r)

	// start server
	log.Fatal(http.ListenAndServe(":8080", handler))
}
