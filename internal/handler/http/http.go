package http

import (
	"net/http"

	"github.com/Paschalolo/reddit-recipie-aggregator/internal/application"
	"github.com/Paschalolo/reddit-recipie-aggregator/pkg"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	App application.Application
}

func NewHandler(App application.Application) *Handler {
	return &Handler{App: App}
}
func (h *Handler) NewRecipeHandler(c *gin.Context) {
	var Recipe pkg.Recipe
	if err := c.ShouldBindJSON(&Recipe); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	rec, err := h.App.AddRecipe(c.Request.Context(), &Recipe)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, rec)

}

func (h *Handler) ListRecipeHandler(c *gin.Context) {
	list, err := h.App.ListRecipe(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, list)
}

func (h *Handler) UpdateRecipeHandler(c *gin.Context) {
	id := c.Param("id")
	var Recipe pkg.Recipe
	if err := c.ShouldBindJSON(&Recipe); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	rec, err := h.App.UpdateRecipe(c.Request.Context(), id, &Recipe)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, rec)

}
