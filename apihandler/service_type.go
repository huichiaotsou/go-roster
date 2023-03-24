package apihandler

import (
	"encoding/json"
	"fmt"
	"net/http"

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
	apiWithID := serviceTypeAPI + "/team/{team_id}"

	// Requires team admin permission
	serviceTypeRouter := a.router.PathPrefix(apiWithID).Subrouter()
	serviceTypeRouter.Use(a.mw.TeamAdminOrSuperuserPerm)
	serviceTypeRouter.HandleFunc("", a.handleCreateServiceType).Methods(http.MethodPost)

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

	if err := a.db.InsertServiceType(newServiceType); err != nil {
		handleError(w, err, "error while inserting new service type in handleCreateServiceType", http.StatusInternalServerError)
		return
	}

	respondWithJSON(w, http.StatusOK, map[string]string{
		"message": "Servie type created",
	})
}
