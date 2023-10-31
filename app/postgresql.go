package app

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"time"
)

func NewPostgresqlDatabase(configuration Config) *sql.DB {

	dbConnection, err := sql.Open("postgres", configuration.Get("DATABASE_URL"))
	if err != nil {
		fmt.Print("\nCannot connect to database", configuration.Get("POSTGRES_DB"))
		panic(err)
	}

	fmt.Println("Connect to database", configuration.Get("POSTGRES_DB"))
	dbConnection.SetMaxIdleConns(5)
	dbConnection.SetMaxOpenConns(20)
	dbConnection.SetConnMaxLifetime(60 * time.Minute)
	dbConnection.SetConnMaxIdleTime(10 * time.Minute)

	return dbConnection
}
