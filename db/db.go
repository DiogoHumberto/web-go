package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func Connect() *sql.DB {
	// Connect to your postgres DB.
	conn := "host=localhost port=5432 user=docker password=docker dbname=api-produto sslmode=disable"

	db, err := sql.Open("postgres", conn)

	if err != nil {
		panic(err.Error())
	}

	return db

}
