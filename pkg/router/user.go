package router

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/huichiaotsou/go-roster/config"
	"github.com/huichiaotsou/go-roster/pkg/handler"
)

type Router struct {
	router *mux.Router
}

func (r *Router) WithUserRouter(userHandler *handler.UserHandler) {
	addr := fmt.Sprintf("/api/%s/users", config.Cfg.APIVersion)
	r.router.HandleFunc(addr, userHandler.GetUsers).Methods(http.MethodGet)
	r.router.HandleFunc(addr, userHandler.CreateUser).Methods(http.MethodPost)
}
