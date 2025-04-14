package main

import (
	recipeRouter "github.com/Paschalolo/reddit-recipie-aggregator/internal"
	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()
	apiV1 := router.Group("/api/v1")
	recipeRouter.Module(apiV1)
	router.Run(":8081")
}
