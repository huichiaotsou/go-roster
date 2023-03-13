package apihandler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/huichiaotsou/go-roster/config"
	"github.com/huichiaotsou/go-roster/types"
)

var (
	ADMIN_API = "/admin"
	TEAM_API  = "/team"
)

// /api/v1/admin/team
func (a *APIHandler) SetAdminTeamRoutes() {
	apiVersion := fmt.Sprintf("/api/%s", config.GetApiVersion())
	adminTeamApi := apiVersion + ADMIN_API + TEAM_API

	adminPermRouter := a.router.PathPrefix(adminTeamApi).Subrouter()
	adminPermRouter.Use(a.mw.CheckSuperPerm)

	// Handle create teams
	adminPermRouter.HandleFunc("/create", a.handleCreateTeams).Methods(http.MethodPost)

	// // Handle assign user teams
	adminPermRouter.HandleFunc("/assign_user", a.handleAssignTeams).Methods(http.MethodDelete)
}

func (a *APIHandler) handleCreateTeams(w http.ResponseWriter, r *http.Request) {
	// Parse request body to Team slice
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
