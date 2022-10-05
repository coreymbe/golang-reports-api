package main

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type AuthToken struct {
	TokenType string    `json:"token_type"`
	Token     string    `json:"access_token"`
	ExpiresAt time.Time `json:"expires_at"`
}

type AuthTokenClaim struct {
	*jwt.RegisteredClaims
}

type AuthError struct {
	Error string `json:"error"`
}

var JWTSecretKey = []byte(os.Getenv("JWTSecretKey"))

func authCheck(auth_token string) bool {
	token, err := jwt.Parse(auth_token, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWTSecretKey")), nil
	})

	if token.Valid {
		logAuth("Token is valid.")
		return true
	} else if errors.Is(err, jwt.ErrTokenMalformed) {
		logError(err)
		return false
	} else if errors.Is(err, jwt.ErrTokenExpired) {
		logError(err)
		return false
	} else {
		logError(err)
		return false
	}
}
