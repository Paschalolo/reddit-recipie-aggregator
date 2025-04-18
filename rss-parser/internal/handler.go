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
	if err := h.App.PushToQueue(&request); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Error while publish to Rabbit mq ",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Success ",
	})
}

func (h *Handler) GetRecipesHandler(c *gin.Context) {
	recipes, err := h.App.FindRecipes(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "could not get recipes",
		})
		return
	}
	c.JSON(http.StatusOK, recipes)
}
