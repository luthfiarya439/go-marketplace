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

	// admin route
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

		// products route
		adminRoute.GET("products", controllers.IndexProduct)
		adminRoute.GET("products/:id", controllers.ShowProduct)
		adminRoute.POST("products", controllers.CreateProduct)
		adminRoute.PUT("products/:id", controllers.UpdateProduct)
		adminRoute.DELETE("products/:id", controllers.DeleteProduct)
	}

	// user route
	userRoute := router.Group("api/user")
	userRoute.Use(middleware.UserMiddleware)
	{
		userRoute.GET("ping", controllers.GetProfile)

		// categories route
		userRoute.GET("categories", controllers.IndexCategory)
		userRoute.GET("categories/:id", controllers.ShowCategory)

		// products route
		userRoute.GET("products", controllers.IndexProduct)
		userRoute.GET("products/:id", controllers.ShowProduct)

		// cart route
		userRoute.GET("carts", controllers.IndexCart)
		userRoute.GET("carts/:id", controllers.ShowCart)
		userRoute.POST("carts", controllers.CreateCart)
		userRoute.DELETE("carts/:id", controllers.DeleteCart)

		// transaction route
		userRoute.GET("transactions", controllers.IndexTransaction)
		userRoute.GET("transactions/:id", controllers.ShowTransaction)

		// checkout cart
		userRoute.POST("checkout", controllers.CheckoutCart)
	}

	port := os.Getenv("PORT")
	router.Run(":" + port)
}
