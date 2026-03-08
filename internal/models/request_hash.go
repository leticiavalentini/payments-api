package models

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

func HashRequest(req CreatePaymentRequest) string {
	data := fmt.Sprintf("%d:%s:%s",
		req.Amount,
		req.Currency,
		req.MerchantID,
	)

	h := sha256.Sum256([]byte(data))
	return hex.EncodeToString(h[:])
}
