package main

import (
	"go-marketplace/config"
	"go-marketplace/controllers"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func init() {
	config.LoadEnv()
	config.ConnectDatabase()
}

func main() {
	router := gin.Default()

	router.GET("ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})

	authRoute := router.Group("/auth")
	authRoute.POST("login", controllers.Authenticate)
	authRoute.POST("register", controllers.Register)

	port := os.Getenv("PORT")
	router.Run(":" + port)
}
