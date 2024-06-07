package database

import (
	"database/sql"
)

var DB *sql.DB // vai ser usado por todos, sera aberto uma vez apenas

func ConnectToDatabase() error {
	connection := "user=root dbname=databasezuda password=root host=localhost sslmode=disable"

	DB, err := sql.Open("databasezuda", connection)
	if err != nil {
		panic(err.Error())
	}

	err = DB.Ping()
	if err != nil {
		panic(err.Error())
	}

	return nil
}
