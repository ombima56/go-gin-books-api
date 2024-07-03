package main

import "go-gin-books-api/routes"

func main() {
	r := routes.SetupRoutes()
	r.Run("localhost:8080")
}
