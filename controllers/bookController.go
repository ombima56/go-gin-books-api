package controllers

import (
	"net/http"

	"go-gin-books-api/models"

	"github.com/gin-gonic/gin"
)

var books = []models.Book{
	{ID: 1, Title: "Pride and Prejudice", Author: "Jane Austine"},
	{ID: 2, Title: "The Ransom Trilogy", Author: "C.S.Lewis"},
	{ID: 3, Title: "To Kill a Mockingbird", Author: "Harper Lee"},
	{ID: 4, Title: "1984", Author: "George Orwell"},
	{ID: 5, Title: "The Catcher in the Rye", Author: "J.D. Salinger"},
	{ID: 6, Title: "Harvard Address", Author: "Alexander Solzhenitsyn"},
}

func GetBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, books)
}
