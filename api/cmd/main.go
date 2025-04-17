package main

import (
	"log"

	recipeRouter "github.com/Paschalolo/reddit-recipie-aggregator/internal"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	apiV1 := router.Group("/api/v1")
	recipeRouter.Run(apiV1)
	if err := router.RunTLS(":443", "certs/localhost.crt", "certs/localhost.key"); err != nil {
		log.Fatalln("error in tls server")
	}

}
