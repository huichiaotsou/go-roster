package utils

import (
	"net/http"

	"github.com/huichiaotsou/go-roster/types"
)

func VerifyTeamAdminPermission(r *http.Request, teamID int64) bool {
	// Get team permission slice from user's jwt token
	claims, _ := VerifyJWTToken(r)
	tpc := claims[types.TeamPermsClaim].([]*types.TeamPermission)

	// Loop the slice and return true if the user is admin in that team
	for _, tp := range tpc {
		if tp.TeamID == teamID {
			return tp.PermissionID == types.Permission_Values["admin"]
		}
	}
	return false
}
