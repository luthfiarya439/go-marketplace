package controllers

import "github.com/gin-gonic/gin"

func Authenticate(c *gin.Context) {
	var input struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		response := gin.H{
			"status":  500,
			"message": "Error, " + err.Error(),
		}
		c.JSON(500, response)
		return
	}

}
