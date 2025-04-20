package internal

import (
	"github.com/Paschalolo/reddit-recipie-aggregator/rss-parser/internal/repository/mongo"
	"github.com/Paschalolo/reddit-recipie-aggregator/rss-parser/internal/repository/queue"
	"github.com/gin-gonic/gin"
)

func Producer(router *gin.Engine) {
	Repo := mongo.NewMongo()
	Queue := queue.NewQueue()
	App := NewApp(Repo, Queue)
	handler := NewHandler(App)
	router.POST("/parse", handler.ParseHandler)
	router.GET("/recipes", handler.GetRecipesHandler)
}
