package api

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/huichiaotsou/go-roster/config"
)

func (a *API) SetUserRoutes(router *mux.Router) {
	api := fmt.Sprintf("/api/%s/users", config.GetApiVersion())
	a.Router.HandleFunc(api, createUser).Methods("POST")
	a.Router.HandleFunc(api, listUsers).Methods("GET")

	apiWithID := fmt.Sprintf(api + "/{id}")
	a.Router.HandleFunc(apiWithID, getUser).Methods("GET")
	a.Router.HandleFunc(apiWithID, updateUser).Methods("PUT")
	a.Router.HandleFunc(apiWithID, deleteUser).Methods("DELETE")
}

func createUser(w http.ResponseWriter, r *http.Request) {}
func listUsers(w http.ResponseWriter, r *http.Request)  {}
func getUser(w http.ResponseWriter, r *http.Request)    {}
func updateUser(w http.ResponseWriter, r *http.Request) {}
func deleteUser(w http.ResponseWriter, r *http.Request) {}
