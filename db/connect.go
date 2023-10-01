package db

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

func ConnectToDB() *sql.DB {
	db, err := sql.Open("postgres",
		"user=postgres password=123 host=localhost dbname=bankdb sslmode=disable")
	if err != nil {
		log.Fatalf("Error: Unable to connect to database: %v", err)
	}
	defer db.Close()
	return db
}
