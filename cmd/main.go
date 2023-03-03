package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/huichiaotsou/go-roster/api"
	"github.com/huichiaotsou/go-roster/config"
	"github.com/rs/cors"
)

func main() {
	// initialize router
	r := mux.NewRouter()

	// register all routes
	api.NewApi(r).RegisterAllRoutes()

	// set up CORS
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"},
	})

	// wrap router with CORS middleware
	handler := c.Handler(r)

	// start server
	addr := fmt.Sprintf(":%s", config.GetServerPort())
	log.Fatal(http.ListenAndServe(addr, handler))
}
