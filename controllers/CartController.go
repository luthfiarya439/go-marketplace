package controllers

import (
	"go-marketplace/config"
	"go-marketplace/helpers"
	"go-marketplace/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
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

func CheckoutCart(c *gin.Context) {
	var carts []models.Cart
	// var products []models.Product
	currentUser, _ := c.Get("currentUser")
	user, _ := currentUser.(models.User)

	if err := config.DB.Model(&carts).Where("user_id = ?", user.ID).Find(&carts).Error; err != nil {
		response := gin.H{
			"status":  500,
			"message": "Error: " + err.Error(),
		}

		c.JSON(500, response)
		return
	}

	if len(carts) == 0 {
		response := gin.H{
			"status":  404,
			"message": "Data cart tidak ada",
		}

		c.JSON(404, response)
		return
	}

	err := MakeTransaction(config.DB, carts, user)

	if err != nil {
		c.JSON(500, err.Error())
		return
	}

	response := gin.H{
		"status":  201,
		"message": "Berhasil melakukan checkout cart",
	}

	c.JSON(201, response)
}

func MakeTransaction(db *gorm.DB, carts []models.Cart, user models.User) error {
	return db.Transaction(func(tx *gorm.DB) error {
		productIds := helpers.GetProductId(carts)
		checkoutProducts := helpers.GetCheckoutProduct(carts)
		var txProduct []models.Product
		var txCart models.Cart
		var txTransactionDetails models.TransactionDetail
		if err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).Where("id IN ?", productIds).Find(&txProduct).Error; err != nil {
			return err
		}

		if err := helpers.CheckProductQuantity(txProduct, checkoutProducts); err != nil {
			return err
		}
		updateRawQuery := helpers.MakeUpdateStatement(checkoutProducts)
		if err := tx.Exec(updateRawQuery).Error; err != nil {
			return err
		}

		totalTransaction := helpers.CalculateTotalTransaction(txProduct, checkoutProducts)

		createTransaction := models.Transaction{
			Total:           uint(totalTransaction),
			TransactionCode: helpers.MakeTransactionCode(15),
			UserId:          user.ID,
		}

		if err := tx.Model(&createTransaction).Create(&createTransaction).Error; err != nil {
			return err
		}

		transactionDetails := helpers.MakeTransactionData(txProduct, createTransaction, checkoutProducts)

		if err := tx.Model(&txTransactionDetails).Create(&transactionDetails).Error; err != nil {
			return err
		}

		if err := tx.Model(&txCart).Where("user_id = ?", user.ID).Delete(&txCart).Error; err != nil {
			return err
		}

		return nil
	})
}
