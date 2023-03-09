package apihandler

import (
	"github.com/huichiaotsou/go-roster/middleware"
	"github.com/huichiaotsou/go-roster/model"
	log "github.com/sirupsen/logrus"

	"github.com/gorilla/mux"
)

type APIHandler struct {
	logger *log.Logger
	router *mux.Router
	mw     *middleware.Middleware
	db     *model.Database
}

func New(
	router *mux.Router,
	logger *log.Logger,
	middleware *middleware.Middleware,
	db *model.Database,
) *APIHandler {
	return &APIHandler{
		logger: logger,
		router: router,
		mw:     middleware,
		db:     db,
	}
}

func (a *APIHandler) RegisterAllRoutes() {

	a.SetUserRoutes()
	// Set___Routers(router)
	// Set___Routers(router)
	// Set___Routers(router)
}
