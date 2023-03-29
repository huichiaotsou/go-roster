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
	FUNCTION_ROUTE = "/function"
)

func (a *APIHandler) SetFuncRoutes() {
	apiVersion := fmt.Sprintf("/api/%s", config.GetApiVersion())
	functionAPI := apiVersion + FUNCTION_ROUTE // /api/v1/service_type

	// Requires superuser permission
	superPermRouter := a.router.PathPrefix(functionAPI).Subrouter()
	superPermRouter.Use(a.mw.SuperPerm)
	superPermRouter.HandleFunc("", a.handleCreateFunctions).Methods(http.MethodPost)
	superPermRouter.HandleFunc("/user/{user_id}", a.handleSetUserFuncs).Methods(http.MethodPost)
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

func (a *APIHandler) handleSetUserFuncs(w http.ResponseWriter, r *http.Request) {
	userID, _ := strconv.ParseInt(mux.Vars(r)["user_id"], 10, 64)

	if err := a.db.ClearUserFuncs(userID); err != nil {
		handleError(w, err, "error while clearing user_funcs in handleSetUserFuncs", http.StatusInternalServerError)
		return
	}

	var f types.FuncIDs
	if err := json.NewDecoder(r.Body).Decode(&f); err != nil {
		handleError(w, err, "error while decoding funcs", http.StatusBadRequest)
		return
	}

	if err := a.db.InsertUserFuncs(userID, f.FuncIDs); err != nil {
		handleError(w, err, "error while inserting user_funcs in handleSetUserFuncs", http.StatusInternalServerError)
		return
	}

	respondWithJSON(w, http.StatusOK, map[string]string{
		"message": "User_funcs set",
	})
}
