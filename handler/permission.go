package handler

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt"
	"github.com/huichiaotsou/go-roster/config"
	"github.com/huichiaotsou/go-roster/types"
)

// User Permission
func (h *Handler) CheckUserPermission(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Check user's permission based on the request context
		_, verified := verifyToken(r)
		if !verified {
			http.Error(w, "Forbidden", http.StatusForbidden)
			return

		}
		// If user has permission, call the next handler
		next.ServeHTTP(w, r)
	})
}

// Admin Permission
func (h *Handler) CheckAdminPermission(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Check user's permission based on the request context
		if !h.adminHasPermission(r) {
			http.Error(w, "Forbidden", http.StatusForbidden)
			return
		}
		// If user has admin permission, call the next handler
		next.ServeHTTP(w, r)
	})
}

func (h *Handler) adminHasPermission(r *http.Request) bool {
	claims, verified := verifyToken(r)
	if !verified {
		return false
	}

	// TO-DO: use userID to check admin role
	userID := claims["user_id"].(string)
	permissions, err := h.Model.GetPermissionsByUserID(userID)
	if err != nil {
		return false
	}

	// verify team and role
	teamID := claims["team_id"].(string)
	for _, p := range permissions {
		if p.TeamID == teamID && p.PermissionName == types.USER_ROLE_ADMIN {
			return true
		}
	}

	return false
}

func verifyToken(r *http.Request) (claims jwt.MapClaims, verified bool) {
	// Retrieve the JWT token from the request headers
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		return claims, false
	}
	tokenString := strings.Replace(authHeader, "Bearer ", "", 1)

	// Parse and validate the JWT token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Verify that the signing method is HMAC and use the secret key to validate the token
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(config.GetJwtKey()), nil
	})
	if err != nil || !token.Valid {
		return claims, false
	}

	// Retrieve the user ID from the JWT token
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return claims, false
	}

	return claims, true
}
