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
	SUPERUSER_ROUTE = "/superuser"
)

// /api/v1/super
func (a *APIHandler) SetSuperUserRoutes() {
	apiVersion := fmt.Sprintf("/api/%s", config.GetApiVersion())
	superuserAPI := apiVersion + SUPERUSER_ROUTE // /api/v1/superuser

	// Requires super user permission
	superPermRouter := a.router.PathPrefix(superuserAPI).Subrouter()
	superPermRouter.Use(a.mw.CheckSuperPerm)

	superPermRouter.HandleFunc("/{user_id}", a.handleEnableSuperuser).Methods(http.MethodPut)
	superPermRouter.HandleFunc("/{user_id}", a.handleDisableSuperuser).Methods(http.MethodDelete)

	superPermRouter.HandleFunc("/create_teams", a.handleCreateTeams).Methods(http.MethodPost)
	superPermRouter.HandleFunc("/create_campus", a.handleCreateCampus).Methods(http.MethodPost)

	superPermRouter.HandleFunc("/define_permissions", a.handleDefinePerms).Methods(http.MethodPost)
	superPermRouter.HandleFunc("/assign_team_permissions", a.handleAssignTeamPerms).Methods(http.MethodPost)
	// superPermRouter.HandleFunc("/assign_user_permissions", a.handleAssignUserPerms).Methods(http.MethodPost)
}

func (a *APIHandler) handleEnableSuperuser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID := vars["user_id"]

	if err := a.db.UpdateIsSuperuser(userID, true); err != nil {
		handleError(w, err, "error while enabling super user in handleEnableSuperuser", http.StatusInternalServerError)
		return
	}

	respondWithJSON(w, http.StatusOK, map[string]string{
		"message": "Enable superuser with success",
	})
}

func (a *APIHandler) handleDisableSuperuser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID := vars["user_id"]

	if err := a.db.UpdateIsSuperuser(userID, false); err != nil {
		handleError(w, err, "error while disabling super user in handleDisableSuperuser", http.StatusInternalServerError)
		return
	}

	respondWithJSON(w, http.StatusOK, map[string]string{
		"message": "Disable superuser with success",
	})
}

func (a *APIHandler) handleCreateTeams(w http.ResponseWriter, r *http.Request) {
	var teams types.Teams
	if err := json.NewDecoder(r.Body).Decode(&teams); err != nil {
		handleError(w, err, "error while decoding teams in handleCreateTeams", http.StatusBadRequest)
		return
	}

	if err := a.db.InsertTeams(teams); err != nil {
		handleError(w, err, "error while creating teams in handleCreateTeams", http.StatusInternalServerError)
		return
	}

	respondWithJSON(w, http.StatusOK, map[string]string{
		"message": "Teams created successfully",
	})
}

func (a *APIHandler) handleCreateCampus(w http.ResponseWriter, r *http.Request) {
	var campuses types.Campuses
	if err := json.NewDecoder(r.Body).Decode(&campuses); err != nil {
		handleError(w, err, "error while decoding campuses in handleCreateCampus", http.StatusBadRequest)
		return
	}

	if err := a.db.InsertCampus(campuses); err != nil {
		handleError(w, err, "error while creating campuses in handleCreateCampus", http.StatusInternalServerError)
		return
	}

	respondWithJSON(w, http.StatusOK, map[string]string{
		"message": "Campus created successfully",
	})
}

func (a *APIHandler) handleDefinePerms(w http.ResponseWriter, r *http.Request) {
	// Parse request body to perms struct
	var perms types.Permissions
	if err := json.NewDecoder(r.Body).Decode(&perms); err != nil {
		handleError(w, err, "error while decoding perms in handleDefinePerms", http.StatusBadRequest)
		return
	}

	if err := a.db.InsertPerms(perms); err != nil {
		handleError(w, err, "error while insert perms in handleDefinePerms", http.StatusInternalServerError)
		return
	}

	respondWithJSON(w, http.StatusOK, map[string]string{
		"message": "Permissions created successfully",
	})

}

func (a *APIHandler) handleAssignTeamPerms(w http.ResponseWriter, r *http.Request) {
	var userTeamPerm []types.UserTeamPerm
	if err := json.NewDecoder(r.Body).Decode(&userTeamPerm); err != nil {
		handleError(w, err, "error while decoding user teams in handleAssignUserPerms", http.StatusBadRequest)
		return
	}

	if err := a.db.InsertUserTeamPerms(userTeamPerm); err != nil {
		handleError(w, err, "error while assigning user teams in handleAssignUserPerms", http.StatusInternalServerError)
		return
	}

	respondWithJSON(w, http.StatusOK, map[string]string{
		"message": "Users permissions assigned successfully",
	})
}
