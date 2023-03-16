package middleware

import (
	"fmt"
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
		fmt.Println(err.Error())
		return false
	}

	return true
}

func (m *Middleware) hasTeamAdminPermission(r *http.Request) bool {
	claims, verified := utils.VerifyJWTToken(r)
	if !verified {
		return false
	}
	if claims[types.TeamPermsClaim] == nil {
		return false
	}

	userID := claims[types.UserIDclaim].(float64)
	teamPermsClaims := claims[types.TeamPermsClaim].([]interface{})
	var teamIDs = make([]int64, len(teamPermsClaims))
	for index, t := range teamPermsClaims {
		teamIDs[index] = int64(t.(float64))
	}

	r.Body.Read()

	fmt.Println("permissions: ", permissions)

	return false
}
