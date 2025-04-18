package internal

import (
	"github.com/Paschalolo/reddit-recipie-aggregator/rss-parser/internal/repository/mongo"
	"github.com/gin-gonic/gin"
)

func Run(router *gin.Engine) {
	Repo := mongo.NewMongo()
	App := NewApp(Repo)
	handler := NewHandler(App)
	router.POST("/parse", handler.ParseHandler)
}
