package controllers

import (
	"go-marketplace/config"
	"go-marketplace/models"

	"github.com/gin-gonic/gin"
)

func IndexCategory(c *gin.Context) {
	var categories []models.Category

	query := config.DB.Model(&categories)

	if c.Query("search") != "" {
		query.Where("name LIKE ?", "%"+c.Query("search")+"%")
	}

	query.Find(&categories)

	response := gin.H{
		"status":  200,
		"message": "Data kategori",
		"data":    categories,
	}

	c.JSON(200, response)
}

func ShowCategory(c *gin.Context) {
	var category models.Category

	if err := config.DB.Model(&category).Where("id = ?", c.Param("id")).First(&category).Error; err != nil {
		response := gin.H{
			"status":  404,
			"message": "Data tidak ditemukan",
		}

		c.JSON(404, response)
	}

	response := gin.H{
		"status":  200,
		"message": "Data detail kategori",
		"data":    category,
	}

	c.JSON(200, response)
}

func CreateCategory(c *gin.Context) {
	var input struct {
		Name string `json:"name" binding:"required"`
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

	category := models.Category{
		Name: input.Name,
	}

	if err := config.DB.Model(&category).Create(&category).Error; err != nil {
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
		"data":    category,
	}
	c.JSON(201, response)
}

func UpdateCategory(c *gin.Context) {
	var input struct {
		Name string `json:"name" binding:"required"`
	}

	var category models.Category

	if err := c.ShouldBindJSON(&input); err != nil {
		response := gin.H{
			"status":  500,
			"message": "Error, " + err.Error(),
			"error":   err.Error(),
		}
		c.JSON(500, response)
		return
	}

	if err := config.DB.Model(&category).Where("id = ?", c.Param("id")).First(&category).Error; err != nil {
		response := gin.H{
			"status":  404,
			"message": "Data tidak ditemukan",
		}

		c.JSON(404, response)
	}

	if err := config.DB.Model(&category).Where("id = ?", c.Param("id")).Updates(&input).Error; err != nil {
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
		"data":    category,
	}
	c.JSON(201, response)
}

func DeleteCategory(c *gin.Context) {
	var category models.Category

	if err := config.DB.Model(&category).Where("id = ?", c.Param("id")).First(&category).Error; err != nil {
		response := gin.H{
			"status":  404,
			"message": "Data tidak ditemukan",
		}

		c.JSON(404, response)
	}

	if err := config.DB.Model(&category).Where("id = ?", c.Param("id")).Delete(&category).Error; err != nil {
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
		"data":    category,
	}
	c.JSON(200, response)
}
