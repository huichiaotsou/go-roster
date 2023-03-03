package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/huichiaotsou/go-roster/pkg/handler"
	"github.com/huichiaotsou/go-roster/pkg/repository"
	"github.com/rs/cors"
)

func main() {
	// initialize router
	r := mux.NewRouter()

	// initialize repository
	repo := repository.NewRepository()

	// initialize handler
	h := handler.NewHandler(repo)

	// set up routes
	r.HandleFunc("/api/items", h.GetItems).Methods("GET")
	r.HandleFunc("/api/items/{id}", h.GetItem).Methods("GET")
	r.HandleFunc("/api/items", h.CreateItem).Methods("POST")
	r.HandleFunc("/api/items/{id}", h.UpdateItem).Methods("PUT")
	r.HandleFunc("/api/items/{id}", h.DeleteItem).Methods("DELETE")

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
