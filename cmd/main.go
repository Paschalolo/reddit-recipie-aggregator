package main

import (
	"context"
	"os"

	recipeRouter "github.com/Paschalolo/reddit-recipie-aggregator/internal"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func main() {

	router := gin.Default()
	apiV1 := router.Group("/api/v1")
	client, _ := mongo.Connect(options.Client().ApplyURI(os.Getenv("MONGO_URI")))
	defer client.Disconnect(context.Background())
	recipeRouter.Module(apiV1, client)
	router.Run(":8081")
}
