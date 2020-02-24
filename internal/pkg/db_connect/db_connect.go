package db_connect

import (
	"database/sql"
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
)

var Db *sqlx.DB

const (
	host = "postgres_db"
	port = 5432
	user = "postgres"
	password = "postgres"
	dbname = "postgres"
)

func InitDB(address, database string) *sqlx.DB{
	var err error
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	fmt.Println(psqlInfo)
	Db, err = sqlx.Open("postgres", psqlInfo)

	if err != nil {
		fmt.Println("Panic, while loading db", err)
		log.Fatal(err)
	} else {
		fmt.Printf("Successfully connected to %v, database %v", host, database)
	}
	return Db
}

func InitDBSQL(address, database string) *sql.DB {
	var err error

	//dsn := "localhost:5432/some-postgres?"
	//Db, err = sqlx.Open("postgres",
	//	"postgres://"+address+"/"+database+"?sslmode=disable")

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	Database, err := sql.Open("postgres", psqlInfo)

	if err != nil {
		fmt.Println("Panic, while loading db")
	}

	if err := Db.Ping(); err != nil {
		fmt.Println("Panic, while pinging db")
	}

	//Db, err := sqlx.Connect("postgres", "user=foo dbname=some-bar sslmode=disable")
	//if err != nil {
	//	log.Fatalln(err)
	//}

	fmt.Println("Successfully connected to %v, database %v", address, database)


	return Database
}