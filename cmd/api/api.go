package api

import (
	"fmt"
	"log"

	"github.com/gorilla/mux"
	"github.com/huichiaotsou/go-roster/config"
	"github.com/huichiaotsou/go-roster/handler"
	"github.com/huichiaotsou/go-roster/model"
	"github.com/huichiaotsou/go-roster/utils"
)

var (
	API_VERSION = fmt.Sprintf("/api/%s", config.GetApiVersion())
)

func RegisterAllRoutes(router *mux.Router) {
	db, err := utils.GetDb(config.GetDBConfig())
	if err != nil {
		log.Fatalf("error while getting database: %s", err)
	}

	handler := handler.NewHandler(router, model.NewModel(db))

	SetUserRoutes(handler)
	// Set___Routers(router)
	// Set___Routers(router)
	// Set___Routers(router)
}
