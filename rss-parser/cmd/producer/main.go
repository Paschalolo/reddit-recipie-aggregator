package main

import (
	parser "github.com/Paschalolo/reddit-recipie-aggregator/rss-parser/internal"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	parser.Producer(router)
	router.Run(":8051")
}
