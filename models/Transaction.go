package models

type Transaction struct {
	ID              uint   `json:"id"`
	UserId          uint   `json:"user_id"`
	Total           uint   `json:"total"`
	TransactionCode string `json:"transaction_code"`
}
