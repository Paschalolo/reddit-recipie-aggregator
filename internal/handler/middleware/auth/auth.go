package auth

import (
	"crypto/sha256"
	"encoding/hex"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/Paschalolo/reddit-recipie-aggregator/internal/repository"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

type AuthHandler struct {
	db repository.AuthRepo
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenValue := c.GetHeader("Authorization")
		claims := &claims{}
		tkn, err := jwt.ParseWithClaims(tokenValue, claims, func(t *jwt.Token) (any, error) {
			return []byte(os.Getenv("JWT_SECRET")), nil
		})
		if err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		if tkn == nil || !tkn.Valid {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		c.Next()
	}
}

func NewAuthHandler(db repository.AuthRepo) *AuthHandler {
	return &AuthHandler{db: db}
}

func (h *AuthHandler) SignInHandler(c *gin.Context) {
	hash := sha256.New()
	var user User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"error": err.Error(),
		})
		return
	}
	hashedSting := hash.Sum([]byte(user.Password))
	user.Password = hex.EncodeToString(hashedSting)
	log.Println(user.Password)
	err := h.db.FindUser(c.Request.Context(), user.Username, user.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			// "error": "Invalid username or password ",
			"error": err.Error(),
		})
		return
	}

	expirationTime := time.Now().Add(10 * time.Minute)
	claims := &claims{
		Username: user.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	jwtOutput := &jWTOutput{
		Token:   tokenString,
		Expires: expirationTime,
	}
	c.JSON(http.StatusOK, jwtOutput)
}

func (h *AuthHandler) RefreshHandler(c *gin.Context) {
	tokenValue := c.GetHeader("Authorization")
	claims := &claims{}
	tkn, err := jwt.ParseWithClaims(tokenValue, claims, func(t *jwt.Token) (any, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": err.Error(),
		})
		return
	}

	if tkn == nil || !tkn.Valid {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "invalid token",
		})
		return
	}

	if time.Until(time.Unix(claims.ExpiresAt, 0)) > 30*time.Second {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Token is not expired yet",
		})
		return
	}
	expirationTime := time.Now().Add(time.Minute * 10)
	claims.ExpiresAt = expirationTime.Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	jwtOutput := &jWTOutput{
		Token:   tokenString,
		Expires: expirationTime,
	}
	c.JSON(http.StatusOK, jwtOutput)
}
