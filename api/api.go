package api

import "github.com/gorilla/mux"

type API struct {
	Router *mux.Router
}

func NewApi(r *mux.Router) *API {
	return &API{
		Router: r,
	}
}

func (a *API) RegisterAllRoutes() {
	a.SetUserRoutes(a.Router)
}
