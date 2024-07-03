package routes

import (
	"go-gin-books-api/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes() *gin.Engine {
	r := gin.Default()

	r.GET("/books", controllers.GetBooks)

	return r
}
