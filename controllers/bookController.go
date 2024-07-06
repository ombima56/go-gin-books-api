package controllers

import (
	"net/http"
	"strconv"

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

func GetBookByID(c *gin.Context) {
	id := c.Param("id")

	bookId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid book ID"})
		return
	}

	for _, book := range books {
		if book.ID == uint(bookId) {
			c.IndentedJSON(http.StatusOK, book)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book not found"})
}

func CreateBook(c *gin.Context) {
	var newBook models.Book

	if err := c.BindJSON(&newBook); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid request body"})
		return
	}

	newBook.ID = uint(len(books) + 1)
	books = append(books, newBook)
	c.IndentedJSON(http.StatusCreated, newBook)
}

func UpdateBook(c *gin.Context) {
	id := c.Param("id")

	bookID, err := strconv.Atoi(id)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid book ID"})
		return
	}

	var updateBook models.Book
	if err := c.BindJSON(&updateBook); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid request body"})
		return
	}

	for i, book := range books {
		if book.ID == uint(bookID) {
			books[i] = updateBook
			c.IndentedJSON(http.StatusOK, updateBook)
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book not found"})
}

func DeleteBook(c *gin.Context) {
	id := c.Param("id")
	bookID, err := strconv.Atoi(id)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid Book ID"})
		return
	}

	for i, book := range books {
		if book.ID == uint(bookID) {
			books = append(books[:i], books[i+1:]...)
			c.IndentedJSON(http.StatusOK, gin.H{"messsage": "Book deleted"})
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book not found"})
}
