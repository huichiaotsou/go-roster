package api

import (
	"fmt"

	"github.com/huichiaotsou/go-roster/config"
	"github.com/huichiaotsou/go-roster/handler"
)

func SetUserRoutes(h *handler.Handler) {
	api := fmt.Sprintf("/api/%s/users", config.GetApiVersion())
	h.Router.HandleFunc(api, h.CreateUser).Methods("POST")
	h.Router.HandleFunc(api, h.ListUsers).Methods("GET")

	apiWithID := fmt.Sprintf(api + "/{id}")
	h.Router.HandleFunc(apiWithID, h.GetUser).Methods("GET")
	h.Router.HandleFunc(apiWithID, h.UpdateUser).Methods("PUT")
	h.Router.HandleFunc(apiWithID, h.DeleteUser).Methods("DELETE")
}
