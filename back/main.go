package main

import (
	"biblioteca/database"
	"biblioteca/routes"
	"log"
	"net/http"
)

func main() {
	routes.SetupRoutes()
	err := database.ConnectToDatabase()
	if err != nil {
		log.Fatalln("Erro ao conectar com a database: ", err)
	}

	http.ListenAndServe(":8000", nil)
}
