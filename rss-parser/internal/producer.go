package internal

import (
	"github.com/Paschalolo/reddit-recipie-aggregator/rss-parser/internal/repository/queue"
	"github.com/gin-gonic/gin"
)

func Producer(router *gin.Engine) {
	Queue := queue.NewQueue()
	App := NewApp(nil, Queue)
	handler := NewHandler(App)
	router.POST("/parse", handler.ParseHandler)
}
