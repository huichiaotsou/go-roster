package middleware

import (
	"net/http"

	"github.com/huichiaotsou/go-roster/types"
	"github.com/huichiaotsou/go-roster/utils"
)

// User Permission
func (m *Middleware) CheckUserPerm(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Check user's permission based on the request context
		_, verified := utils.VerifyJWTToken(r)
		if !verified {
			http.Error(w, "Forbidden", http.StatusForbidden)
			return

		}
		// If user has permission, call the next handler
		next.ServeHTTP(w, r)
	})
}

// Super User Permission
func (m *Middleware) CheckSuperPerm(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Check user's permission based on the request context
		if m.hasSuperUserPermission(r) {
			next.ServeHTTP(w, r)
			return
		}

		http.Error(w, "Forbidden", http.StatusForbidden)
	})
}

func (m *Middleware) hasSuperUserPermission(r *http.Request) bool {
	claims, verified := utils.VerifyJWTToken(r)
	if !verified {
		return false
	}

	userID := int64(claims[types.UserIDclaim].(float64))
	isSuperUser, err := m.Db.IsSuperUser(userID)
	if err != nil || !isSuperUser {
		return false
	}

	return true
}

// Admin Permission
func (m *Middleware) CheckAdminOrSuperPerm(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Check user's permission based on the request context
		if m.hasTeamAdminPermission(r) || m.hasSuperUserPermission(r) {
			next.ServeHTTP(w, r)
			return
		}

		http.Error(w, "Forbidden", http.StatusForbidden)
	})
}

func (m *Middleware) hasTeamAdminPermission(r *http.Request) bool {
	// claims, verified := utils.VerifyJWTToken(r)
	// if !verified {
	// 	return false
	// }
	// if claims[types.TeamPermsClaim] == nil {
	// 	return false
	// }

	// userID := claims[types.UserIDclaim].(float64)
	// teamPermsClaims := claims[types.TeamPermsClaim].([]types.TeamPermission)

	// // Parse request body to Teams slice
	// var teamPerms []types.TeamPermission
	// err := json.NewDecoder(r.Body).Decode(&teamPerms)
	// if err != nil {
	// 	err = fmt.Errorf("error while decoding teams in handleCreateTeams: %s", err)
	// 	return false
	// }

	return false
}
