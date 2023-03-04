package api

import (
	"fmt"

	"github.com/huichiaotsou/go-roster/config"
	"github.com/huichiaotsou/go-roster/handler"
)

func SetUserRoutes(h *handler.Handler) {
	// example: /api/v1/user
	api := fmt.Sprintf("/api/%s", config.GetApiVersion())
	userApi := api + "/user"

	// create user
	h.Router.HandleFunc(userApi, h.CreateUser).Methods("POST")

	// apply checkUserPermission middleware to user permission sub router
	apiWithID := fmt.Sprintf(userApi + "/{id}")
	userRouter := h.Router.PathPrefix(apiWithID).Subrouter()
	userRouter.Use(h.CheckUserPermission)

	// modify & delete user with user ID
	userRouter.HandleFunc(apiWithID, h.UpdateUser).Methods("PUT")
	userRouter.HandleFunc(apiWithID, h.DeleteUser).Methods("DELETE")
}
