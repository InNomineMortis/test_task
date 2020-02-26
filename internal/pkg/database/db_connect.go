package database

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
	"os"
)

var DB *sqlx.DB

var (
	host     = "localhost"
	port     = "5432"
	user     = "postgres"
	password = "postgres"
	name     = "postgres"
)

func init() {
	dbHost := os.Getenv("DB_HOST")
	if dbHost != "" {
		host = dbHost
	}

	dbPort := os.Getenv("DB_PORT")
	if dbPort != "" {
		port = dbPort
	}

	dbUser := os.Getenv("DB_USER")
	if dbUser != "" {
		user = dbUser
	}

	dbPassword := os.Getenv("DB_PASSWORD")
	if dbPassword != "" {
		password = dbPassword
	}

	dbName := os.Getenv("DB_NAME")
	if dbName != "" {
		name = dbName
	}
}

func InitPostgres(database string) {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, name)
	log.Printf("Database connect info: %s", psqlInfo)
	var err error
	DB, err = sqlx.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatalf("Error - connecting to database: %s", err.Error())
	}

	log.Printf("Successfully connected to %v, database %v", host, database)
}
