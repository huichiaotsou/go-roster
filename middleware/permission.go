package middleware

import (
	"net/http"

	"github.com/huichiaotsou/go-roster/types"
	"github.com/huichiaotsou/go-roster/utils"
)

// User Permission
func (m *Middleware) UserPerm(next http.Handler) http.Handler {
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
func (m *Middleware) SuperPerm(next http.Handler) http.Handler {
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
