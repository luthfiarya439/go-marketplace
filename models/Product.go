package models

type Product struct {
	ID         uint   `json:"id"`
	CategoryId uint   `json:"category_id"`
	Name       string `json:"name"`
	Price      uint   `json:"price"`
	Quantity   uint   `json:"quantity"`
}
