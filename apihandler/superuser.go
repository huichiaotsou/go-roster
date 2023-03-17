package apihandler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/huichiaotsou/go-roster/config"
	"github.com/huichiaotsou/go-roster/types"
)

var (
	SUPERUSER_API = "/superuser"
)

// /api/v1/super
func (a *APIHandler) SetSuperUserRoutes() {
	apiVersion := fmt.Sprintf("/api/%s", config.GetApiVersion())
	adminApi := apiVersion + SUPERUSER_API

	// Requires super user permission to creat teams and assign user teams
	superPermRouter := a.router.PathPrefix(adminApi).Subrouter()
	superPermRouter.Use(a.mw.CheckSuperPerm)
	superPermRouter.HandleFunc("/create_teams", a.handleCreateTeams).Methods(http.MethodPost)

	// Should merge the below 2 (assign_teams_permissions)
	// the table permission should contains only the permission ID
	superPermRouter.HandleFunc("/assign_user_teams", a.handleAssignUserTeams).Methods(http.MethodPost)
	superPermRouter.HandleFunc("/assign_user_permissions", a.handleAssignUserPerms).Methods(http.MethodPost)
}

func (a *APIHandler) handleCreateTeams(w http.ResponseWriter, r *http.Request) {
	// Parse request body to Teams slice
	var teams types.Teams
	err := json.NewDecoder(r.Body).Decode(&teams)
	if err != nil {
		err = fmt.Errorf("error while decoding teams in handleCreateTeams: %s", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Insert new user into database
	err = a.db.InsertTeams(teams)
	if err != nil {
		err = fmt.Errorf("error while creating teams in handleCreateTeams: %s", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Set token in Authorization header
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	// Return success response
	response := map[string]string{
		"message": "Teams created successfully",
	}
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		err = fmt.Errorf("error while writing response in handleCreateTeams: %s", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (a *APIHandler) handleAssignUserTeams(w http.ResponseWriter, r *http.Request) {
	// Parse request body to userTeams struct
	var userTeams types.UserTeams
	err := json.NewDecoder(r.Body).Decode(&userTeams)
	if err != nil {
		err = fmt.Errorf("error while decoding user teams in handleAssignTeams: %s", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = a.db.InsertUserTeams(userTeams)
	if err != nil {
		err = fmt.Errorf("error while assigning user teams in handleAssignTeams: %s", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Set token in Authorization header
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	// Return success response
	response := map[string]string{
		"message": "Teams assigned to users successfully",
	}
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		err = fmt.Errorf("error while writing response in handleAssignUserTeams: %s", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (a *APIHandler) handleAssignUserPerms(w http.ResponseWriter, r *http.Request) {
	// Parse request body to userPerms slice
	var userPerms []types.UserPerms
	err := json.NewDecoder(r.Body).Decode(&userPerms)
	if err != nil {
		err = fmt.Errorf("error while decoding user teams in handleAssignUserPerms: %s", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = a.db.InsertUserPerms(userPerms)
	if err != nil {
		err = fmt.Errorf("error while assigning user teams in handleAssignUserPerms: %s", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Set token in Authorization header
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	// Return success response
	response := map[string]string{
		"message": "Users permissions assigned successfully",
	}
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		err = fmt.Errorf("error while writing response in handleAssignUserPerms: %s", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
