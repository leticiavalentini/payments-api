package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"payments-api/internal/models"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type PaymentRepository struct {
	DB *sql.DB
}

func NewPaymentRepository(db *sql.DB) *PaymentRepository {
	return &PaymentRepository{DB: db}
}

func (r *PaymentRepository) List(ctx context.Context) ([]models.Payment, error) {
	rows, err := r.DB.QueryContext(ctx, "SELECT id, merchant_id, idempotency_key, request_hash, amount, currency, status FROM payments_v2")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var payments []models.Payment

	for rows.Next() {
		var p models.Payment
		err := rows.Scan(
			&p.ID,
			&p.MerchantID,
			&p.IdempotencyKey,
			&p.RequestHash,
			&p.Amount,
			&p.Currency,
			&p.Status,
		)
		if err != nil {
			return nil, err
		}

		payments = append(payments, p)
	}

	return payments, nil
}

func (r *PaymentRepository) CreatePayment(ctx context.Context, req models.CreatePaymentRequest, key string, queue chan string) (*models.Payment, error) {
	query := `
        INSERT INTO payments_v2
        (id, merchant_id, idempotency_key, request_hash, amount, currency, status)
        VALUES ($1,$2,$3,$4,$5,$6,'created')
        ON CONFLICT (merchant_id, idempotency_key)
        DO NOTHING
        RETURNING id, merchant_id, idempotency_key, request_hash, amount, currency, status
    `

	reqHash := models.HashRequest(req)
	id := uuid.New().String()
	var payment models.Payment

	err := r.DB.QueryRowContext(
		ctx,
		query,
		id,
		req.MerchantID,
		key,
		reqHash,
		req.Amount,
		req.Currency,
	).Scan(
		&payment.ID,
		&payment.MerchantID,
		&payment.IdempotencyKey,
		&payment.RequestHash,
		&payment.Amount,
		&payment.Currency,
		&payment.Status,
	)

	err = r.DB.QueryRowContext(ctx,
		`
            SELECT id, merchant_id, idempotency_key, request_hash,
                   amount, currency, status
            FROM payments_v2
            WHERE merchant_id=$1 AND idempotency_key=$2
        `,
		req.MerchantID,
		key,
	).Scan(
		&payment.ID,
		&payment.MerchantID,
		&payment.IdempotencyKey,
		&payment.RequestHash,
		&payment.Amount,
		&payment.Currency,
		&payment.Status,
	)

	if err != nil {
		return nil, err
	}

	if payment.RequestHash != reqHash {
		return nil, errors.New("idempotency key resused with different parameters")
	}

	queue <- id

	return &payment, err
}

func (r *PaymentRepository) GetPayment(ctx *gin.Context) (*models.Payment, error) {
	id := ctx.Param("id")
	var payment models.Payment

	err := r.DB.QueryRowContext(
		ctx,
		`SELECT id, merchant_id, idempotency_key, request_hash,
               amount, currency, status
        FROM payments
        WHERE id=$1`,
		id).Scan(
		&payment.ID,
		&payment.MerchantID,
		&payment.IdempotencyKey,
		&payment.RequestHash,
		&payment.Amount,
		&payment.Currency,
		&payment.Status,
	)

	fmt.Println("idempotency key", payment.IdempotencyKey)
	fmt.Println("merchant id", payment.MerchantID)

	if errors.Is(err, sql.ErrNoRows) {
		return nil, errors.New("payment not found")
	}

	if err != nil {
		return nil, errors.New("database error")
	}

	return &payment, nil
}
