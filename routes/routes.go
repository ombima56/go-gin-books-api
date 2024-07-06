package routes

import (
	"go-gin-books-api/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes() *gin.Engine {
	r := gin.Default()

	r.GET("/books", controllers.GetBooks)
	r.GET("/books/:id", controllers.GetBookByID)
	r.POST("/books", controllers.CreateBook)
	r.PUT("/books/:id", controllers.UpdateBook)
	r.DELETE("/books/:id", controllers.DeleteBook)

	return r
}
