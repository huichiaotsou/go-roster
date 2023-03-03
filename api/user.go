package api

import (
	"fmt"

	"github.com/gorilla/mux"
	"github.com/huichiaotsou/go-roster/config"
	"github.com/huichiaotsou/go-roster/handler"
)

func (a *API) SetUserRoutes(router *mux.Router) {
	api := fmt.Sprintf("/api/%s/users", config.GetApiVersion())
	a.Router.HandleFunc(api, handler.CreateUser).Methods("POST")
	a.Router.HandleFunc(api, handler.ListUsers).Methods("GET")

	apiWithID := fmt.Sprintf(api + "/{id}")
	a.Router.HandleFunc(apiWithID, handler.GetUser).Methods("GET")
	a.Router.HandleFunc(apiWithID, handler.UpdateUser).Methods("PUT")
	a.Router.HandleFunc(apiWithID, handler.DeleteUser).Methods("DELETE")
}
