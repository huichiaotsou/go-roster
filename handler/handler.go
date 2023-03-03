package handler

import (
	"github.com/gorilla/mux"
	"github.com/huichiaotsou/go-roster/model"
)

type Handler struct {
	Model  *model.Model
	Router *mux.Router
}

func NewHandler(router *mux.Router) *Handler {
	return &Handler{
		Model:  nil, // TO-DO
		Router: router,
	}
}
