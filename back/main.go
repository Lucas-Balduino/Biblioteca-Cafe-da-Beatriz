package main

import (
	"biblioteca/database"
	"biblioteca/routes"
	"net/http"
)

func main() {
	routes.SetupRoutes()
	database.ConnectToDatabase()
	http.ListenAndServe(":8000", nil)
}
