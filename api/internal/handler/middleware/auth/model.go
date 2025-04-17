package auth

import (
	"time"

	"github.com/golang-jwt/jwt"
)

type claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

type jWTOutput struct {
	Token   string    `json:"token"`
	Expires time.Time `json:"expires"`
}

// User represents user credentials for authentication.
type user struct {
	Password string `json:"password"` // User's password.
	Username string `json:"username"` // User's username.
}
