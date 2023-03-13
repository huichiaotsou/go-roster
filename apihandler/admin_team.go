package apihandler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/huichiaotsou/go-roster/types"
)

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

func (a *APIHandler) handleAssignUserTeams(w http.ResponseWriter, r *http.Request) {
	// Parse request body to Team slice
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
		err = fmt.Errorf("error while writing response in handleCreateTeams: %s", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
