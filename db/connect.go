package db

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
	"os"
)

var DB *sql.DB

// Init sets up the global DB connection (call in main or init)
func Init() {
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		log.Fatal("DATABASE_URL not set")
	}
	var err error
	DB, err = sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal("Failed to open DB:", err)
	}
	if err = DB.Ping(); err != nil {
		log.Fatal("Failed to connect to DB:", err)
	}
	log.Println("Connected to PostgreSQL (db/connect.go)")
}
