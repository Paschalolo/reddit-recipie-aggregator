package internal

import (
	"os"

	_ "github.com/Paschalolo/reddit-recipie-aggregator/docs"
	"github.com/Paschalolo/reddit-recipie-aggregator/internal/application"
	"github.com/Paschalolo/reddit-recipie-aggregator/internal/handler/http"
	"github.com/Paschalolo/reddit-recipie-aggregator/internal/handler/middleware/auth"
	mongoRepo "github.com/Paschalolo/reddit-recipie-aggregator/internal/repository/mongo"
	"github.com/Paschalolo/reddit-recipie-aggregator/internal/repository/redis"
	"github.com/gin-contrib/cors"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func Run(router *gin.RouterGroup) {
	router.Use(cors.Default())
	client, _ := mongo.Connect(options.Client().ApplyURI(os.Getenv("MONGO_URI")))
	repo := mongoRepo.NewMongoDB(client)
	AuthRepo := mongoRepo.NewAuthMongoDB(client)
	// utils.AddAuthUser(AuthRepo)
	cache := redis.NewRedis(repo)
	App := application.New(repo, cache)
	Handler := http.NewHandler(*App)
	authorised := router.Group("/")
	authorised.Use(auth.AuthMiddleware())
	authHandler := auth.NewAuthHandler(AuthRepo)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.GET("/recipes", Handler.ListRecipeHandler)
	router.POST("/signup", authHandler.SignUpHandler)
	router.POST("/signin", authHandler.SignInHandler)
	router.POST("/signout", authHandler.SignOutHandler)
	router.POST("/refresh", authHandler.RefreshHandler)
	// protected routes
	authorised.POST("/recipes", Handler.NewRecipeHandler)
	authorised.GET("/recipes/search", Handler.SearchRecipeHandler)
	authorised.GET("/recipes/:id", Handler.ListOneRecipeHandler)
	authorised.PUT("/recipes/:id", Handler.UpdateRecipeHandler)
	authorised.DELETE("/recipes/:id", Handler.DeleteRecipeHandler)
}
