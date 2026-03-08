package main

import (
	"database/sql"
	"log"
	"os"
	"payments-api/internal/db"
	"payments-api/internal/routes"

	"github.com/gin-gonic/gin"
)

var paymentQueue = make(chan string, 100)

func main() {

	log.SetOutput(os.Stdout)
	log.Println("startup log test")

	database := db.NewPostgres()
	defer database.Close()

	gin.SetMode(gin.DebugMode)
	router := gin.Default()

	routes.Register(router, database, paymentQueue)
	go paymentWorker(database, paymentQueue)

	err := router.Run(":8080")

	if err != nil {
		log.Fatal("And issue occurred on router startup")
	}

	log.Println("server running on :8080")

}

func paymentWorker(db *sql.DB, paymentQueue <-chan string) {
	for paymentID := range paymentQueue {
		log.Println("processing payment", paymentID)
	}
}
