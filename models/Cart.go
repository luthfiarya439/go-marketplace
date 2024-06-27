package models

type Cart struct {
	ProductId uint `json:"product_id"`
	UserId    uint `json:"user_id"`
	Quantity  uint `json:"quantity"`
}
