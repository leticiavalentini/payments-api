package db

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func NewPostgres() *sql.DB {
	connStr := "postgres://postgres:admin@localhost:5432/payments?sslmode=disable"

	db, err := sql.Open("postgres", connStr)

	if err != nil {
		log.Fatal(err)
	}

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	return db
}
