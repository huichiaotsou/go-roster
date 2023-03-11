package apihandler

import (
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/huichiaotsou/go-roster/config"
)

func generateJWTToken(userID int64, teamIDs []int64, email string) (string, error) {
	// Set token claims
	claims := jwt.MapClaims{}
	claims["userId"] = userID
	claims["teamID"] = teamIDs
	claims["email"] = email
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix() // Token expires after 24 hours

	// Generate token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(config.GetJwtKey())) // Replace with your own secret key
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
