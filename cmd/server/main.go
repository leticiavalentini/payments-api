package main

import (
	//"log"

	//"payments-api/internal/db"
	//"payments-api/internal/routes"

	"log"
	"os"
	"payments-api/internal/db"
	"payments-api/internal/routes"

	"github.com/gin-gonic/gin"
)

func main() {

	log.SetOutput(os.Stdout)
	log.Println("startup log test")

	database := db.NewPostgres()
	defer database.Close()

	gin.SetMode(gin.DebugMode)
	router := gin.Default()

	routes.Register(router, database)

	err := router.Run(":8080")

	if err != nil {
		log.Fatal("And issue occurred on router startup")
	}

	log.Println("server running on :8080")

}
