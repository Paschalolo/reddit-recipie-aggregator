package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func IndexHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"msg": "hello world",
	})
}
func SetUpSever() *gin.Engine {
	router := gin.Default()
	router.GET("/", IndexHandler)
	return router
}
func main() {
	SetUpSever().Run()
}
