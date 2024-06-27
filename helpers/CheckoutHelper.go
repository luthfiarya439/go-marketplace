package helpers

import (
	"errors"
	"fmt"
	"go-marketplace/models"
	"math/rand"
)

func GetProductId(carts []models.Cart) []uint {
	productIds := make([]uint, 0, len(carts))

	checkProductId := func(slice []uint, value uint) bool {
		for _, item := range slice {
			if item == value {
				return true
			}
		}
		return false
	}

	for _, cart := range carts {
		if !checkProductId(productIds, cart.ProductId) {
			productIds = append(productIds, cart.ProductId)
		}
	}
	return productIds
}

func GetCheckoutProduct(carts []models.Cart) map[uint]uint {
	checkoutProducts := map[uint]uint{}

	for _, cart := range carts {
		var _, isExists = checkoutProducts[cart.ProductId]
		if isExists {
			checkoutProducts[cart.ProductId] = checkoutProducts[cart.ProductId] + cart.Quantity
		} else {
			checkoutProducts[cart.ProductId] = cart.Quantity
		}
	}

	return checkoutProducts
}

func CheckProductQuantity(products []models.Product, checkoutProducts map[uint]uint) error {

	for _, product := range products {
		if product.Quantity < checkoutProducts[product.ID] {
			return errors.New("stock ada yang habis")
		}
	}
	return nil
}

func MakeUpdateStatement(checkoutProducts map[uint]uint) string {
	caseStatements := ""
	for id, quantity := range checkoutProducts {
		caseStatements += fmt.Sprintf("WHEN %d THEN quantity - %d ", id, quantity)
	}
	rawQuery := fmt.Sprintf(`UPDATE products SET quantity = CASE id %s ELSE quantity END`, caseStatements)
	return rawQuery
}

func MakeTransactionCode(n int) string {
	const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func CalculateTotalTransaction(products []models.Product, checkoutProducts map[uint]uint) int {
	total := 0
	for _, product := range products {
		total += int(product.Price) * int(checkoutProducts[product.ID])
	}

	return total
}

func MakeTransactionData(products []models.Product, transaction models.Transaction, checkoutProducts map[uint]uint) []models.TransactionDetail {
	var transactionDetails []models.TransactionDetail

	for _, product := range products {
		transactionDetailData := models.TransactionDetail{
			ProductId:       product.ID,
			ProductName:     product.Name,
			ProductPrice:    product.Price,
			ProductQuantity: checkoutProducts[product.ID],
			TransactionId:   transaction.ID,
		}
		transactionDetails = append(transactionDetails, transactionDetailData)
	}

	return transactionDetails
}
