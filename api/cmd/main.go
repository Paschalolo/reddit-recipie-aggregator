package main

import (
	"log"

	recipeRouter "github.com/Paschalolo/reddit-recipie-aggregator/internal"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	apiV1 := router.Group("/v1")
	recipeRouter.Run(apiV1)

	if err := router.Run(); err != nil {
		log.Fatalln("error in tls server")
	}
	// if err := router.RunTLS(":8080", "certs/localhost.crt", "certs/localhost.key"); err != nil {
	// 	log.Fatalln("error in tls server")
	// }

}
