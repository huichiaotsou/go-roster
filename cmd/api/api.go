package api

import (
	"github.com/gorilla/mux"
	"github.com/huichiaotsou/go-roster/handler"
)

func RegisterAllRoutes(router *mux.Router) {
	handler := handler.NewHandler(router)

	SetUserRoutes(handler)
	// Set___Routers(router)
	// Set___Routers(router)
	// Set___Routers(router)
}
