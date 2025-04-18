package main

import (
	RSS_Parser "github.com/Paschalolo/reddit-recipie-aggregator/rss-parser/internal"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	RSS_Parser.Run(router)
	router.Run(":8051")
}
