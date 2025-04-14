package http

import (
	"net/http"

	"github.com/Paschalolo/reddit-recipie-aggregator/pkg"
	"github.com/gin-gonic/gin"
)

// Home function Handler
func HomeHandler(c *gin.Context) {
	s := &pkg.JsonFormat{Name: "Paschal", Age: 89}
	c.JSON(http.StatusOK, s)
}
