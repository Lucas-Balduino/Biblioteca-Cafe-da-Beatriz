package database

import (
	"database/sql"
)

func ConnectToDatabase() *sql.DB {
	connection := "user=root dbname=databasezuda password=root host=localhost sslmode=disable"
	db, err := sql.Open("databasezuda", connection)
	if err != nil {
		panic(err.Error())
	}
	return db
}
