package db_connect

import (
	"fmt"
	"github.com/jackc/pgx"
	"github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
	"log"
	"test_task/internal/pkg/config"
	"time"
)

const (
	host     = "postgres"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "test"
)

func OpenSqlxViaPgxConnPool() *sqlx.DB {
	connConfig := pgx.ConnConfig{
		Host:     config.DBHostname,
		Database: config.Database,
		User:     config.User,
		Password: config.Password,
	}
	connPool, err := pgx.NewConnPool(pgx.ConnPoolConfig{
		ConnConfig:     connConfig,
		AfterConnect:   nil,
		MaxConnections: 20,
		AcquireTimeout: 30 * time.Second,
	})
	fmt.Printf("OpenSqlxViaPgxConnPool: Hostname %s\n", config.DBHostname)
	if err != nil {
		fmt.Printf("OpenSqlxViaPgxConnPool: Failed to create connections pool: error - %s\n", err.Error())
		//log.Fatalf("OpenSqlxViaPgxConnPool: Failed to create connections pool: error - %s\n", err.Error())
	}
	nativeDB := stdlib.OpenDBFromPool(connPool)
	err = nativeDB.Ping()
	if err != nil {
		fmt.Println("Errorrrr! Pinging", err)
	}
	log.Println("OpenSqlxViaPgxConnPool: the connection was created", nativeDB)
	return sqlx.NewDb(nativeDB, "pgx")
}