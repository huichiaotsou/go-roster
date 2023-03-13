package apihandler

import (
	"fmt"
	"net/http"

	"github.com/huichiaotsou/go-roster/config"
)

var (
	ADMIN_API           = "/admin"
	WORSHIP_TEAM_API    = "/worship"
	SOUND_TEAM_API      = "/sound"
	PRODUCTION_TEAM_API = "/production"
)

// /api/v1/admin, /api/v1/worship, /api/v1/sound...
func (a *APIHandler) SetAdminRoutes() {
	apiVersion := fmt.Sprintf("/api/%s", config.GetApiVersion())
	adminApi := apiVersion + ADMIN_API

	// Requires super user permission to creat teams and assign user teams
	superPermRouter := a.router.PathPrefix(adminApi).Subrouter()
	superPermRouter.Use(a.mw.CheckSuperPerm)
	superPermRouter.HandleFunc("/create_teams", a.handleCreateTeams).Methods(http.MethodPost)
	superPermRouter.HandleFunc("/assign_user_teams", a.handleAssignUserTeams).Methods(http.MethodPost)
}
