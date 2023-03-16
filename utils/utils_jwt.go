package utils

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/huichiaotsou/go-roster/config"
	"github.com/huichiaotsou/go-roster/types"
)

func GenerateJWTToken(userID int64, teamPerms []types.TeamPermission, email string) (string, error) {
	// Set token claims
	claims := jwt.MapClaims{}
	claims[types.UserIDclaim] = userID
	claims[types.TeamPermsClaim] = teamPerms
	claims[types.Emailclaim] = email
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix() // Token expires after 24 hours

	// Generate token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(config.GetJwtKey())) // Replace with your own secret key
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func VerifyJWTToken(r *http.Request) (claims jwt.MapClaims, verified bool) {
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
