package controllers

import (
	"go-marketplace/config"
	"go-marketplace/models"

	"github.com/gin-gonic/gin"
)

func IndexTransaction(c *gin.Context) {
	currentUser, _ := c.Get("currentUser")
	user, _ := currentUser.(models.User)

	var transaction []models.Transaction
	var results []struct {
		ID              uint   `json:"id"`
		UserId          uint   `json:"user_id"`
		UserName        string `json:"user_name"`
		TransactionCode string `json:"transaction_code"`
		Total           uint   `json:"total"`
	}

	query := config.DB.Model(&transaction).Select("transactions.*, users.id as user_id, users.name as user_name").Joins("inner join users on users.id = transactions.user_id").Where("transactions.user_id = ?", user.ID)

	if c.Query("search") != "" {
		query.Where("transactions.transaction_code LIKE ?", "%"+c.Query("search")+"%")
	}

	query.Scan(&results)

	response := gin.H{
		"status":  200,
		"message": "Data transaksi anda, " + user.Name,
		"data":    results,
	}

	c.JSON(200, response)
}

func ShowTransaction(c *gin.Context) {
	currentUser, _ := c.Get("currentUser")
	user, _ := currentUser.(models.User)

	var transaction models.Transaction
	var transactionDetail []models.TransactionDetail

	if err := config.DB.Model(&transaction).Where("transaction_code = ?", c.Param("transactionCode")).Where("user_id = ?", user.ID).First(&transaction).Error; err != nil {
		response := gin.H{
			"status":  404,
			"message": "Data tidak ditemukan",
		}

		c.JSON(404, response)
		return
	}

	if err := config.DB.Model(&transactionDetail).Where("transaction_id = ?", transaction.ID).Find(&transactionDetail).Error; err != nil {
		response := gin.H{
			"status":  404,
			"message": "Data tidak ditemukan",
		}

		c.JSON(404, response)
		return
	}

	response := gin.H{
		"status":  200,
		"message": "Data detail transaksi anda, " + user.Name,
		"data":    transactionDetail,
	}

	c.JSON(200, response)

}
