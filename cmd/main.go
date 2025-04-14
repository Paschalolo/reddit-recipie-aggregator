package main

import (
	ginHandler "github.com/Paschalolo/reddit-recipie-aggregator/internal/handler/http"
	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()
	router.GET("/", ginHandler.HomeHandler)
	router.Run(":8081")
}
