package internal

import (
	"encoding/json"
	"log"
	"os"

	_ "github.com/Paschalolo/reddit-recipie-aggregator/docs"
	"github.com/Paschalolo/reddit-recipie-aggregator/internal/application"
	"github.com/Paschalolo/reddit-recipie-aggregator/internal/handler/http"
	"github.com/Paschalolo/reddit-recipie-aggregator/internal/repository"
	mongoRepo "github.com/Paschalolo/reddit-recipie-aggregator/internal/repository/mongo"
	"github.com/Paschalolo/reddit-recipie-aggregator/internal/repository/redis"
	"github.com/Paschalolo/reddit-recipie-aggregator/pkg"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func AddBulkRecipe(repo repository.Repository) {
	file, err := os.ReadFile("recipes.json")
	if err != nil {
		log.Fatal("could not read file", err.Error())
	}
	var recipe []pkg.Recipe
	if err := json.Unmarshal(file, &recipe); err != nil {
		log.Fatalln("could not read file ", err.Error())
	}
	if err := repo.BulkAddRecipe(&recipe); err != nil {
		log.Fatalln("could not save file ", err.Error())
	}
}
func Run(router *gin.RouterGroup) {
	repo := mongoRepo.NewMongoDB()
	// AddBulkRecipe(repo)
	cache := redis.NewRedis(repo)
	App := application.New(repo, cache)
	Handler := http.NewHandler(*App)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.POST("/recipes", Handler.NewRecipeHandler)
	router.GET("/recipes", Handler.ListRecipeHandler)
	router.GET("/recipes/search", Handler.SearchRecipeHandler)
	router.GET("/recipes/:id", Handler.ListOneRecipeHandler)
	router.PUT("/recipes/:id", Handler.UpdateRecipeHandler)
	router.DELETE("/recipes/:id", Handler.DeleteRecipeHandler)
}
