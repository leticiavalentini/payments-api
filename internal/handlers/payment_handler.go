package handlers

import (
	"log"
	"net/http"
	"payments-api/internal/models"
	"payments-api/internal/repository"

	"github.com/gin-gonic/gin"
)

type PaymentHandler struct {
	repo *repository.PaymentRepository
}

func NewPaymentHandler(repo *repository.PaymentRepository) *PaymentHandler {
	return &PaymentHandler{repo: repo}
}

func (h *PaymentHandler) ListPayments(c *gin.Context) {
	ctx := c.Request.Context()

	payments, err := h.repo.List(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "database error"})
		return
	}

	c.JSON(http.StatusOK, payments)
}

func (h *PaymentHandler) CreatePayment(c *gin.Context) {
	log.Println("on create payment")
	key := c.GetHeader("Idempotency-key")
	if key == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "missing idempotency key",
		})
		return
	}

	var req models.CreatePaymentRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request body",
		})
		return
	}

	models.HashRequest(req)
	//id := uuid.New().String()

	ctx := c.Request.Context()

	payment, err := h.repo.CreatePayment(ctx, req, key)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, payment)
}

func (h *PaymentHandler) GetPayment(c *gin.Context) {
	log.Println("on get payment")
	payment, err := h.repo.GetPayment(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, payment)

}
