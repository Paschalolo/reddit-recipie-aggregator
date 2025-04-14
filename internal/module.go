package internal

import (
	"github.com/Paschalolo/reddit-recipie-aggregator/internal/application"
	"github.com/Paschalolo/reddit-recipie-aggregator/internal/handler/http"
	"github.com/Paschalolo/reddit-recipie-aggregator/internal/repository/memory"
	"github.com/gin-gonic/gin"
)

func Module(router *gin.Engine) {
	repo := memory.NewRepository()
	App := application.New(repo)
	Handler := http.NewHandler(*App)
	router.POST("/recipes", Handler.NewRecipeHandler)
	router.GET("/recipes", Handler.ListRecipeHandler)
}
