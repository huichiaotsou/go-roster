package api

import (
	"fmt"

	"github.com/gorilla/mux"
	"github.com/huichiaotsou/go-roster/config"
	"github.com/huichiaotsou/go-roster/handler"
)

func SetUserRoutes(router *mux.Router) {
	api := fmt.Sprintf("/api/%s/users", config.GetApiVersion())
	router.HandleFunc(api, handler.CreateUser).Methods("POST")
	router.HandleFunc(api, handler.ListUsers).Methods("GET")

	apiWithID := fmt.Sprintf(api + "/{id}")
	router.HandleFunc(apiWithID, handler.GetUser).Methods("GET")
	router.HandleFunc(apiWithID, handler.UpdateUser).Methods("PUT")
	router.HandleFunc(apiWithID, handler.DeleteUser).Methods("DELETE")
}
