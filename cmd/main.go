package main

import (
	recipeRouter "github.com/Paschalolo/reddit-recipie-aggregator/internal"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	recipeRouter.Module(router)
	router.Run(":8081")
}
