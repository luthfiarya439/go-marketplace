package main

import (
	"go-marketplace/config"
	"go-marketplace/controllers"
	"go-marketplace/middleware"
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

	adminRoute := router.Group("/api/admin")
	adminRoute.Use(middleware.AdminMiddleware)
	{
		adminRoute.GET("ping", controllers.GetProfile)

		// categories route
		adminRoute.GET("categories", controllers.IndexCategory)
		adminRoute.GET("categories/:id", controllers.ShowCategory)
		adminRoute.POST("categories", controllers.CreateCategory)
		adminRoute.PUT("categories/:id", controllers.UpdateCategory)
		adminRoute.DELETE("categories/:id", controllers.DeleteCategory)

	}

	port := os.Getenv("PORT")
	router.Run(":" + port)
}
