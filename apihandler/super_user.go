package apihandler

import (
	"fmt"
	"net/http"

	"github.com/huichiaotsou/go-roster/config"
)

var (
	ADMIN_API = "/super"
)

// /api/v1/super
func (a *APIHandler) SetSuperUserRoutes() {
	apiVersion := fmt.Sprintf("/api/%s", config.GetApiVersion())
	adminApi := apiVersion + ADMIN_API

	// Requires super user permission to creat teams and assign user teams
	superPermRouter := a.router.PathPrefix(adminApi).Subrouter()
	superPermRouter.Use(a.mw.CheckSuperPerm)
	superPermRouter.HandleFunc("/create_teams", a.handleCreateTeams).Methods(http.MethodPost)
	superPermRouter.HandleFunc("/assign_user_teams", a.handleAssignUserTeams).Methods(http.MethodPost)
}
