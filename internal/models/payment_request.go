package models

type CreatePaymentRequest struct {
	Amount     int64  `json:"amount" binding:"required"`
	Currency   string `json:"currency" binding:"required"`
	MerchantID string `json:"merchant_id" binding:"required"`
}
