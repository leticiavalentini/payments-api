package models

type Payment struct {
	ID             string `json:"id"`
	MerchantID     string `json:"merchant_id"`
	IdempotencyKey string `json:"idempotency_key"`
	RequestHash    string `json:"-"`
	Amount         int    `json:"amount"`
	Currency       string `json:"currency"`
	Status         string `json:"status"`
}
