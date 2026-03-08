package routes

import (
	"database/sql"
	"payments-api/internal/handlers"
	"payments-api/internal/repository"

	"github.com/gin-gonic/gin"
)

func Register(router *gin.Engine, db *sql.DB, queue chan string) {
	repo := repository.NewPaymentRepository(db)
	handler := handlers.NewPaymentHandler(repo, queue)

	router.GET("/payments", handler.ListPayments)
	router.POST("/payments", handler.CreatePayment)
	router.GET("/payments/:id", handler.GetPayment)
}
