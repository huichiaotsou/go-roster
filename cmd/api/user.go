package api

import (
	"fmt"

	"github.com/huichiaotsou/go-roster/handler"
)

var (
	USER_API = API_VERSION + "/user"
)

func SetUserRoutes(h *handler.Handler) {
	// create user
	h.Router.HandleFunc(USER_API, h.CreateUser).Methods("POST")

	// apply CheckUserPerm middleware to the sub router userPermRouter
	apiWithID := fmt.Sprintf(USER_API + "/{id}")
	userPermRouter := h.Router.PathPrefix(apiWithID).Subrouter()
	userPermRouter.Use(h.CheckUserPerm)

	// modify & delete user with user ID
	userPermRouter.HandleFunc(apiWithID, h.UpdateUser).Methods("PUT")
	userPermRouter.HandleFunc(apiWithID, h.DeleteUser).Methods("DELETE")
}
