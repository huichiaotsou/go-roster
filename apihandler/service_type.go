package apihandler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/huichiaotsou/go-roster/config"
	"github.com/huichiaotsou/go-roster/types"
)

var (
	SERVICE_TYPE_ROUTE = "/service_type"
)

func (a *APIHandler) SetServiceTypeRoutes() {
	apiVersion := fmt.Sprintf("/api/%s", config.GetApiVersion())
	serviceTypeAPI := apiVersion + SERVICE_TYPE_ROUTE // /api/v1/service_type

	// Requires team admin permission
	serviceTypeRouter := a.router.PathPrefix(serviceTypeAPI).Subrouter()
	serviceTypeRouter.Use(a.mw.SuperPerm)

	serviceTypeRouter.HandleFunc("/team/{team_id}", a.handleCreateServiceType).Methods(http.MethodPost)
	serviceTypeRouter.HandleFunc("/{service_type_id}", a.handleDeleteServiceType).Methods(http.MethodDelete)
	serviceTypeRouter.HandleFunc("/{service_type_id}/funcs", a.handleSetServiceTypeFuncs).Methods(http.MethodPost)
}

func (a *APIHandler) handleCreateServiceType(w http.ResponseWriter, r *http.Request) {
	var newServiceType types.ServiceType
	if err := json.NewDecoder(r.Body).Decode(&newServiceType); err != nil {
		handleError(w, err, "error while decoding newServiceType", http.StatusBadRequest)
		return
	}

	if newServiceType.TeamID != mux.Vars(r)["team_id"] {
		handleError(w, nil, "team id not identical", http.StatusBadRequest)
		return
	}

	if err := a.db.UpsertServiceType(newServiceType); err != nil {
		handleError(w, err, "error while inserting new service type in handleCreateServiceType", http.StatusInternalServerError)
		return
	}

	respondWithJSON(w, http.StatusOK, map[string]string{
		"message": "Service type created/modified",
	})
}

func (a *APIHandler) handleDeleteServiceType(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.ParseInt(mux.Vars(r)["service_type_id"], 10, 64)

	if err := a.db.DeleteServiceType(id); err != nil {
		handleError(w, err, "error while inserting new service type in handleCreateServiceType", http.StatusInternalServerError)
		return
	}

	respondWithJSON(w, http.StatusOK, map[string]string{
		"message": "Servie type deleted",
	})
}

func (a *APIHandler) handleSetServiceTypeFuncs(w http.ResponseWriter, r *http.Request) {

}
