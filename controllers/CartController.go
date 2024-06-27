package controllers

import (
	"go-marketplace/config"
	"go-marketplace/models"

	"github.com/gin-gonic/gin"
)

type results struct {
	ID           uint   `json:"id"`
	ProductId    uint   `json:"product_id"`
	UserId       uint   `json:"user_id"`
	ProductName  string `json:"product_name"`
	CartQuantity uint   `json:"cart_quantity"`
	ProductPrice uint   `json:"product_price"`
}

func IndexCart(c *gin.Context) {
	var cart []models.Cart
	var results []results

	currentUser, _ := c.Get("currentUser")
	user, _ := currentUser.(models.User)

	query := config.DB.Model(&cart).Select("carts.id as id, carts.quantity as cart_quantity, carts.user_id, products.id as product_id, products.name as product_name, products.price as product_price").Where("carts.user_id = ?", user.ID).Joins("inner join products ON products.id = carts.product_id")

	if c.Query("search") != "" {
		query.Where("products.name LIKE ?", "%"+c.Query("search")+"%")
	}

	query.Scan(&results)

	response := gin.H{
		"status":  200,
		"message": "Data cart anda, " + user.Name,
		"data":    results,
	}

	c.JSON(200, response)
}

func ShowCart(c *gin.Context) {
	var cart models.Cart
	var result results

	currentUser, _ := c.Get("currentUser")
	user, _ := currentUser.(models.User)

	if err := config.DB.Model(&cart).Select("carts.id as id, carts.quantity as cart_quantity, carts.user_id, products.id as product_id, products.name as product_name, products.price as product_price").Where("carts.user_id = ?", user.ID).Where("carts.id = ?", c.Param("id")).Joins("inner join products ON products.id = carts.product_id").First(&result).Error; err != nil {
		response := gin.H{
			"status":  404,
			"message": "Data tidak ditemukan",
		}

		c.JSON(404, response)
		return
	}

	response := gin.H{
		"status":  200,
		"message": "Data cart anda, " + user.Name,
		"data":    result,
	}

	c.JSON(200, response)
}

func CreateCart(c *gin.Context) {
	var input struct {
		ProductId uint `json:"product_id" binding:"required"`
		Quantity  uint `json:"quantity" binding:"required"`
	}

	currentUser, _ := c.Get("currentUser")
	user, _ := currentUser.(models.User)

	if err := c.ShouldBindJSON(&input); err != nil {
		response := gin.H{
			"status":  500,
			"message": "Error, " + err.Error(),
			"error":   err.Error(),
		}
		c.JSON(500, response)
		return
	}

	Cart := models.Cart{
		ProductId: input.ProductId,
		UserId:    user.ID,
		Quantity:  input.Quantity,
	}

	if err := config.DB.Model(&Cart).Create(&Cart).Error; err != nil {
		response := gin.H{
			"status":  500,
			"message": "Error, " + err.Error(),
			"error":   err.Error(),
		}
		c.JSON(500, response)
		return
	}

	response := gin.H{
		"status":  201,
		"message": "Data berhasil dibuat",
		"data":    Cart,
	}
	c.JSON(201, response)
}

func DeleteCart(c *gin.Context) {
	var cart models.Cart
	currentUser, _ := c.Get("currentUser")
	user, _ := currentUser.(models.User)

	if err := config.DB.Model(&cart).Where("carts.user_id = ?", user.ID).Where("carts.id = ?", c.Param("id")).First(&cart).Error; err != nil {
		response := gin.H{
			"status":  404,
			"message": "Data tidak ditemukan",
		}

		c.JSON(404, response)
		return
	}

	if err := config.DB.Model(&cart).Where("carts.user_id = ?", user.ID).Where("carts.id = ?", c.Param("id")).Delete(&cart).Error; err != nil {
		response := gin.H{
			"status":  500,
			"message": "Error, " + err.Error(),
			"error":   err.Error(),
		}
		c.JSON(500, response)
		return
	}

	response := gin.H{
		"status":  200,
		"message": "Cart anda berhasil dihapus",
		"data":    cart,
	}

	c.JSON(200, response)
}
