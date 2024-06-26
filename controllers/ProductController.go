package controllers

import (
	"go-marketplace/config"
	"go-marketplace/models"

	"github.com/gin-gonic/gin"
)

func IndexProduct(c *gin.Context) {
	var products []models.Product
	var results []struct {
		ID           uint   `json:"id"`
		CategoryId   uint   `json:"category_id"`
		Name         string `json:"name"`
		Price        uint   `json:"price"`
		Quantity     uint   `json:"quantity"`
		CategoryName string `json:"category_name"`
	}

	query := config.DB.Model(&products).Select("products.*, categories.name as category_name").Joins("inner join categories on categories.id = products.category_id")

	if c.Query("search") != "" {
		query.Where("products.name LIKE ?", "%"+c.Query("search")+"%")
	}

	if c.Query("category") != "" {
		query.Where("categories.name = ?", c.Query("category"))
	}

	query.Scan(&results)

	response := gin.H{
		"status":  200,
		"message": "Data produk",
		"data":    results,
	}

	c.JSON(200, response)
}

func ShowProduct(c *gin.Context) {
	var Product models.Product

	if err := config.DB.Model(&Product).Where("id = ?", c.Param("id")).First(&Product).Error; err != nil {
		response := gin.H{
			"status":  404,
			"message": "Data tidak ditemukan",
		}

		c.JSON(404, response)
		return
	}

	response := gin.H{
		"status":  200,
		"message": "Data detail produk",
		"data":    Product,
	}

	c.JSON(200, response)
}

func CreateProduct(c *gin.Context) {
	var input struct {
		Name       string `json:"name" binding:"required"`
		CategoryId uint   `json:"category_id" binding:"required"`
		Price      uint   `json:"price" binding:"required"`
		Quantity   uint   `json:"quantity" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		response := gin.H{
			"status":  500,
			"message": "Error, " + err.Error(),
			"error":   err.Error(),
		}
		c.JSON(500, response)
		return
	}

	Product := models.Product{
		Name:       input.Name,
		CategoryId: input.CategoryId,
		Price:      input.Price,
		Quantity:   input.Quantity,
	}

	if err := config.DB.Model(&Product).Create(&Product).Error; err != nil {
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
		"data":    Product,
	}
	c.JSON(201, response)
}

func UpdateProduct(c *gin.Context) {
	var input struct {
		Name       string `json:"name" binding:"required"`
		CategoryId uint   `json:"category_id" binding:"required"`
		Price      uint   `json:"price" binding:"required"`
		Quantity   uint   `json:"quantity" binding:"required"`
	}

	var product models.Product

	if err := c.ShouldBindJSON(&input); err != nil {
		response := gin.H{
			"status":  500,
			"message": "Error, " + err.Error(),
			"error":   err.Error(),
		}
		c.JSON(500, response)
		return
	}

	if err := config.DB.Model(&product).Where("id = ?", c.Param("id")).First(&product).Error; err != nil {
		response := gin.H{
			"status":  404,
			"message": "Data tidak ditemukan",
		}

		c.JSON(404, response)
		return
	}

	if err := config.DB.Model(&product).Where("id = ?", c.Param("id")).Updates(&input).Error; err != nil {
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
		"message": "Data berhasil diperbaharui",
		"data":    product,
	}
	c.JSON(201, response)
}

func DeleteProduct(c *gin.Context) {
	var product models.Product

	if err := config.DB.Model(&product).Where("id = ?", c.Param("id")).First(&product).Error; err != nil {
		response := gin.H{
			"status":  404,
			"message": "Data tidak ditemukan",
		}

		c.JSON(404, response)
		return
	}

	if err := config.DB.Model(&product).Where("id = ?", c.Param("id")).Delete(&product).Error; err != nil {
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
		"message": "Data berhasil dihapus",
		"data":    product,
	}
	c.JSON(200, response)
}
