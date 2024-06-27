package models

type Cart struct {
	ID        uint `json:"id"`
	ProductId uint `json:"product_id"`
	UserId    uint `json:"user_id"`
	Quantity  uint `json:"quantity"`
}
