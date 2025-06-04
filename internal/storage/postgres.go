package storage

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func NewPostgresDB() *sql.DB {
	dsn := os.Getenv("POSTGRES_DSN")
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal("DB error:", err)
	}
	return db
}
