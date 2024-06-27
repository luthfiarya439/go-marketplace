package models

type TransactionDetail struct {
	ID              uint   `json:"id"`
	TransactionId   uint   `json:"transaction_id"`
	ProductId       uint   `json:"product_id"`
	ProductName     string `json:"product_name"`
	ProductPrice    uint   `json:"product_price"`
	ProductQuantity uint   `json:"product_quantity"`
}
