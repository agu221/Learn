package storage

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var DB *sql.DB

func getConnectionString() (string, string, string, string, string) {
	err := godotenv.Load()
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	if err != nil {
		log.Printf("Warning: .env file note loaded: %v", err)
	}

	return host, port, user, password, dbname
}
func Connect() {
	host, port, user, password, database_name := getConnectionString()

	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, database_name)

	var err error
	DB, err = sql.Open("postgres", connStr)

	if err != nil {
		log.Fatal("Database Connectiong Failed: ", err)
	}
	err = DB.Ping()
	if err != nil {
		log.Fatal("Cannot reach database: ", err)
	}

	fmt.Println("Connested to PostgreSQL")
}
