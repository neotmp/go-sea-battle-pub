package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var DB *sql.DB

func Connect() {

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("unable to load .env file: %e", err)
	}

	conString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", os.Getenv("AWS_HOST"), os.Getenv("AWS_PORT"), os.Getenv("AWS_USER"), os.Getenv("AWS_PASSWORD"), os.Getenv("AWS_DBNAME"), os.Getenv("AWS_SSLMODE"))

	// Connect to database
	database, err := sql.Open("postgres", conString)
	if err != nil {
		log.Fatal(err)
	}
	//defer database.Close()

	if err := database.Ping(); err != nil {
		log.Fatal("Failed to ping db", err)
	}

	fmt.Println("Connected to DB")

	DB = database

}
