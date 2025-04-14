// Recipes API
// @version 1.0.0
// @title Recipes API
// @description This is the API for managing recipes.
// @termsOfService http://swagger.io/terms/

// @contact.name Paschal Ahanmisi
// @contact.email pastechnology1@gmail.com

// @host localhost:8081
// @BasePath /api/v1

// @schemes http
// @produce application/json
// @consumes application/json

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

// @Summary Create a new recipe
// @Description Create a new recipe.
// @Accept json
// @Produce json
// @Param recipe body pkg.Recipe true "Recipe object to be created"
// @Success 200 {object} pkg.Recipe
// @Failure 400 {object} map[string]string "Invalid request payload"
// @Router /recipes [post]
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

// @Summary Get all recipes
// @Description Get a list of all available recipes.
// @Produce json
// @Success 200 {array} pkg.Recipe
// @Router /recipes [get]
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

// @Summary Return recipes by id
// @Description Return a recipe that contains a specific id.
// @Produce json
// @Param id path string true "ID of the recipe to list"
// @Accept json
// @Param recipe body pkg.Recipe true "Recipe object to be updated"
// @Success 200 {object} pkg.Recipe
// @Failure 400 {object} map[string]string "Invalid request payload"
// @Failure 404 {object} map[string]string "Recipe not found"
// @Router /recipes/{id} [get]
func (h *Handler) ListOneRecipeHandler(c *gin.Context) {
	id := c.Param("id")
	list, err := h.App.ListOneRecipe(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, list)
}

// @Summary Update recipes by id
// @Description Update a recipe that contains a specific id.
// @Produce json
// @Param id path string true "ID of the recipe to update"
// @Accept json
// @Param recipe body pkg.Recipe true "Recipe object to be updated"
// @Success 200 {object} pkg.Recipe
// @Failure 400 {object} map[string]string "Invalid request payload"
// @Failure 404 {object} map[string]string "Recipe not found"
// @Router /recipes/{id} [put]
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

// @Summary Delete recipes by id
// @Description Delete a recipe that contains a specific id.
// @Produce json
// @Param id path string true "ID of the recipe to delete"
// @Success 200 {object} map[string]string "message"
// @Failure 404 {object} map[string]string "Recipe not found"
// @Router /recipes/{id} [delete]
func (h *Handler) DeleteRecipeHandler(c *gin.Context) {
	id := c.Param("id")
	ok := h.App.DeleteRecipe(c.Request.Context(), id)
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "no recipe in the repository",
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "message has been deleted ",
	})
}

// @Summary Search recipes by tag
// @Description Search for recipes that contain a specific tag.
// @Produce json
// @Param tag query string true "Tag to search for"
// @Success 200 {array} pkg.Recipe
// @Failure 404 {object} map[string]string "No recipes found with that tag"
// @Router /recipes/search [get]
func (h *Handler) SearchRecipeHandler(c *gin.Context) {
	tag := c.Query("tag")
	recipe, err := h.App.SearchRecipe(c.Request.Context(), tag)
	if err != nil {

		c.JSON(http.StatusNotFound, gin.H{

			"message": err.Error(),
		})
	}
	c.JSON(http.StatusOK, recipe)

}
