package apihandler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/huichiaotsou/go-roster/config"
	"github.com/huichiaotsou/go-roster/types"
)

var (
	FUNCTION_ROUTE = "/function"
)

func (a *APIHandler) SetFuncRoutes() {
	apiVersion := fmt.Sprintf("/api/%s", config.GetApiVersion())
	functionAPI := apiVersion + FUNCTION_ROUTE // /api/v1/service_type

	// Requires team admin or superuser permission
	functionRouter := a.router.PathPrefix(functionAPI).Subrouter()
	functionRouter.Use(a.mw.CheckSuperPerm)

	functionRouter.HandleFunc("", a.handleCreateFunctions).Methods(http.MethodPost)
}

func (a *APIHandler) handleCreateFunctions(w http.ResponseWriter, r *http.Request) {
	var funcs types.Functions
	if err := json.NewDecoder(r.Body).Decode(&funcs); err != nil {
		handleError(w, err, "error while decoding funcs", http.StatusBadRequest)
		return
	}

	if err := a.db.InsertFunctions(funcs); err != nil {
		handleError(w, err, "error while inserting functions in handleCreateFunctions", http.StatusInternalServerError)
		return
	}

	respondWithJSON(w, http.StatusOK, map[string]string{
		"message": "Functions created",
	})
}
