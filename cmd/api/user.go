package api

import (
	"fmt"

	"github.com/huichiaotsou/go-roster/config"
	"github.com/huichiaotsou/go-roster/handler"
)

func SetUserRoutes(h *handler.Handler) {
	apiVersion := fmt.Sprintf("/api/%s", config.GetApiVersion())
	userApi := apiVersion + "/user"

	// create user
	h.Router.HandleFunc(userApi, h.CreateUser).Methods("POST")

	// apply CheckUserPerm middleware to the sub router userPermRouter
	apiWithID := fmt.Sprintf(userApi + "/{id}")
	userPermRouter := h.Router.PathPrefix(apiWithID).Subrouter()
	userPermRouter.Use(h.CheckUserPerm)

	// modify & delete user with user ID
	userPermRouter.HandleFunc(apiWithID, h.UpdateUser).Methods("PUT")
	userPermRouter.HandleFunc(apiWithID, h.DeleteUser).Methods("DELETE")
}
