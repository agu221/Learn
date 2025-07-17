package storage

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

var DB *sql.DB

func getConnectionString() (string, string, string) {
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: .env file note loaded: %v", err)
	}
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	database_name := os.Getenv("DB_USER")

	return user, password, database_name
}
func Connect() {
	user, password, database_name := getConnectionString()

	connStr := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disbale", user, password, database_name)

	DB, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Database Connectiong Failed: ", err)
	}

	err = DB.Ping()
	if err != nil {
		log.Fatal("Cannot reach database: ", err)
	}

	fmt.Println("Connested to PostgreSQL")
}
