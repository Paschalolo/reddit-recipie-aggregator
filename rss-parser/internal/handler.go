package internal

import (
	"net/http"

	"github.com/Paschalolo/reddit-recipie-aggregator/rss-parser/pkg"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	App App
}

func NewHandler(app App) *Handler {
	return &Handler{App: app}
}

func (h *Handler) ParseHandler(c *gin.Context) {
	var request pkg.Request
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	if err := h.App.InsertOne(c.Request.Context(), request.URL); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Error while parsing rss feed ",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Succesfully added to database ",
	})
}
