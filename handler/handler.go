package handler

import (
	"github.com/gorilla/mux"
	"github.com/huichiaotsou/go-roster/model"
)

type Handler struct {
	Model  *model.Model
	Router *mux.Router
}

func NewHandler(router *mux.Router, model *model.Model) *Handler {
	return &Handler{
		Model:  model,
		Router: router,
	}
}
