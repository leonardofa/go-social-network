package auth

import (
	"api/src/config"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
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

// ValidateToken validates the JWT token in the request headers.
func ValidateToken(r *http.Request) error {
	tokenString, err := extractTokenFromRequest(r)
	if err != nil {
		return err
	}

	token, err := jwt.Parse(tokenString, getSecretKey)
	if err != nil {
		return err
	}

	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return nil
	}

	return errors.New("invalid token")
}

// extractTokenFromRequest extracts the JWT token from the request headers.
func extractTokenFromRequest(r *http.Request) (string, error) {
	tokenString := r.Header.Get("Authorization")
	if tokenString == "" {
		return "", errors.New("missing authorization header")
	}

	return strings.TrimPrefix(tokenString, "Bearer "), nil
}

// getSecretKey returns the secret key used to sign the JWT token.
func getSecretKey(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, errors.New(fmt.Sprintf("unexpected signing method %v", token.Header["alg"]))
	}

	return config.SecretKey, nil
}
