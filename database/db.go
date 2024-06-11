package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

var DB *sql.DB // vai ser usado por todos, sera aberto uma vez apenas

func ConnectToDatabase() error {
	connection := "user=root password=root dbname=databasezuda host=localhost sslmode=disable"

	var err error
	DB, err = sql.Open("postgres", connection)
	if err != nil {
		panic(err.Error())
	}

	err = DB.Ping()
	if err != nil {
		panic(err.Error())
	}

	fmt.Println("Connectado!")

	return nil
}
