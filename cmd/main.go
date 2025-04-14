package main

import (
	ginHandler "github.com/Paschalolo/reddit-recipie-aggregator/internal/handler/http"
	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()
	router.POST("/recipes", ginHandler.NewRecipeHandler)
	router.Run(":8081")
}
