package auth

import (
	"crypto/sha256"
	"encoding/hex"
	"net/http"
	"os"
	"time"

	"github.com/Paschalolo/reddit-recipie-aggregator/internal/repository"
	"github.com/Paschalolo/reddit-recipie-aggregator/pkg"
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

func (h *AuthHandler) SignUpHandler(c *gin.Context) {
	if c.GetHeader("Authorization") != "" {
		c.JSON(http.StatusOK, gin.H{
			"message": "please log out",
		})
		return
	}

	var user pkg.AuthUser
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"error": err.Error(),
		})
		return
	}
	hash := sha256.New()
	hashedSting := hash.Sum([]byte(user.Password))
	user.Password = hex.EncodeToString(hashedSting)

	if err := h.db.AddUser(c.Request.Context(), &user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "cannot sign up user exists ",
		})
		return
	}

	// jwt response
	expirationTime := time.Now().Add(24 * time.Hour)
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

	c.JSON(http.StatusOK, gin.H{
		"message": "Added you as user ",
		"jwtInfo": jwtOutput,
	})

}

// @Summary Sign in a user
// @Description Authenticates a user and returns a JWT token.
// @Tags auth
// @Accept json
// @Produce json
// @Param   body     body    user     true  "User credentials"
// @Success 200 {object} jWTOutput "Successfully signed in"
// @Failure 400 {object} map[string]string "Invalid request payload"
// @Failure 401 {object} map[string]string "Invalid username or password"
// @Failure 500 {object} map[string]string "Failed to generate token"
// @Router /signin [post]
func (h *AuthHandler) SignInHandler(c *gin.Context) {
	if c.GetHeader("Authorization") != "" {
		c.JSON(http.StatusOK, gin.H{
			"message": "already signed in ",
		})
		return
	}
	hash := sha256.New()
	var user user
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"error": err.Error(),
		})
		return
	}
	hashedSting := hash.Sum([]byte(user.Password))
	user.Password = hex.EncodeToString(hashedSting)
	err := h.db.FindUser(c.Request.Context(), user.Username, user.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid username or password ",
			// "error": err.Error(),
		})
		return
	}

	expirationTime := time.Now().Add(24 * time.Hour)
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

// @Summary Refresh JWT token
// @Description Refreshes an expired JWT token.
// @Tags auth
// @Produce json
// @Param   Authorization   header   string    true  "Bearer token"
// @Success 200 {object} jWTOutput "Token refreshed successfully"
// @Failure 401 {object} map[string]string "Invalid token"
// @Failure 400 {object} map[string]string "Token is not expired yet"
// @Failure 500 {object} map[string]string "Failed to generate token"
// @Router /refresh [post]
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
	expirationTime := time.Now().Add(time.Hour * 24)
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

func (h *AuthHandler) SignOutHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Remove token in Authorization Header ",
	})
}
