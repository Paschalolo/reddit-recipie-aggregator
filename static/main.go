package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

var recipes []Recipe

func init() {
	recipes = []Recipe{}
	f, err := os.ReadFile("recipes.json")
	if err != nil {
		log.Fatalln("couldnt open file ")
	}
	if err := json.Unmarshal(f, &recipes); err != nil {
		log.Fatalln("couldnt unmarsh json file ")
	}
}

func InderHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"recipes": recipes,
	})
}
func RecipeHandler(c *gin.Context) {
	for _, recipe := range recipes {
		if recipe.ID == c.Param("id") {
			c.HTML(http.StatusOK, "recipe.tmpl", gin.H{
				"recipe": recipe,
			})
			return
		}
	}
	c.File("404.html")
}

func main() {
	router := gin.Default()
	router.GET("/", InderHandler)
	router.GET("/recipes/:id", RecipeHandler)
	router.LoadHTMLGlob("templates/*")
	router.Static("/assets", "./assets")
	router.Run()
}
