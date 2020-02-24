package db_connect

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
)

var Db *sqlx.DB

const (
	host     = "postgres_db"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "postgres"
)

func InitDB(database string) *sqlx.DB {
	var err error
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	log.Printf("Database connect info: %s", psqlInfo)
	Db, err = sqlx.Open("postgres", psqlInfo)

	if err != nil {
		log.Fatalf("Error - connecting to database: %s", err.Error())
		return nil
	}

	log.Printf("Successfully connected to %v, database %v", host, database)
	return Db
}
