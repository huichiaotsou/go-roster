package apihandler

import (
	"github.com/huichiaotsou/go-roster/middleware"
	"github.com/huichiaotsou/go-roster/model"
	log "github.com/sirupsen/logrus"

	"github.com/gorilla/mux"
)

type APIHandler struct {
	Logger     *log.Logger
	Router     *mux.Router
	Middleware *middleware.Middleware
	DB         *model.Database
}

func New(
	router *mux.Router,
	logger *log.Logger,
	middleware *middleware.Middleware,
	db *model.Database,
) *APIHandler {
	return &APIHandler{
		Logger:     logger,
		Router:     router,
		Middleware: middleware,
		DB:         db,
	}
}

func (a *APIHandler) RegisterAllRoutes() {

	a.SetUserRoutes()
	// Set___Routers(router)
	// Set___Routers(router)
	// Set___Routers(router)
}
