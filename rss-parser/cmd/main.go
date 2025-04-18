package main

import (
	"github.com/Paschalolo/reddit-recipie-aggregator/rss-parser/internal"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.POST("/parse", internal.ParseHandler)
	router.Run(":8081")
}
