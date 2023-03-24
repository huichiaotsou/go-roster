package apihandler

import (
	"github.com/huichiaotsou/go-roster/middleware"
	"github.com/huichiaotsou/go-roster/model"

	"github.com/gorilla/mux"
)

type APIHandler struct {
	db     *model.Database
	mw     *middleware.Middleware
	router *mux.Router
}

func New(
	router *mux.Router,
	middleware *middleware.Middleware,
	db *model.Database,
) *APIHandler {
	return &APIHandler{
		db:     db,
		mw:     middleware,
		router: router,
	}
}

func (a *APIHandler) RegisterAllRoutes() {
	a.SetUserRoutes()
	a.SetSuperUserRoutes()
	a.SetServiceTypeRoutes()

	// Set___Routers(router)
	// Set___Routers(router)
	// Set___Routers(router)
}
