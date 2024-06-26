package main

import (
	"go-marketplace/config"
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
	port := os.Getenv("PORT")
	router.Run(":" + port)
}
