package auth

import (
	"api/src/config"
	"time"

	jwt "github.com/golang-jwt/jwt/v5"
)

// CreateToken creates a JWT token for the given user ID.
func CreateToken(userID uint64) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["exp"] = time.Now().Add(time.Hour * 6).Unix()
	claims["userID"] = userID
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(config.SecretKey)
}
