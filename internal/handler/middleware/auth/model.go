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

type User struct {
	Password string `json:"password"`
	Username string `json:"username"`
}
