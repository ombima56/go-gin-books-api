package main

import (
	"go-gin-books-api/models"
	"go-gin-books-api/routes"
)

func main() {
	models.ConnectDataBase()

	r := routes.SetupRoutes()
	r.Run("localhost:8080")
}
