package utils

import (
	"database/sql"
	"log"

	_ "github.com/denisenkom/go-mssqldb"
)

var DB *sql.DB

// InitializeDB initializes the database connection
func InitializeDB() {
	var err error

	connectionString := "server=db-test.c32wo8s6qtn2.us-east-1.rds.amazonaws.com,1433;user id=admin;password=admin123;database=payment_platform"
	DB, err = sql.Open("sqlserver", connectionString)
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}
}

// CloseDB closes the database connection
func CloseDB() {
	if err := DB.Close(); err != nil {
		log.Fatalf("Error closing the database connection: %v", err)
	}
}
