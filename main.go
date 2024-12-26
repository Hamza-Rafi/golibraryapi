package main

import (
	"net/http"
	"strconv"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type book struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

var books = []book{
	{ID: "0", Title: "Maths book", Author: "Sabeeh"},
	{ID: "1", Title: "English book", Author: "Hassan"},
	{ID: "2", Title: "Science book", Author: "Tuan"},
}

func main() {
	router := gin.Default()

	router.Use(cors.Default())
	gin.SetMode(gin.ReleaseMode)

	router.GET("/books", getBooks)
	router.POST("/books", postBooks)
	router.PUT("/books", updateBook)

	router.GET("/books/:id", getBookByID)
	router.DELETE("/books/:id", deleteBookByID)

	router.Run("0.0.0.0:8080")
}

func getBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, books)
}

func postBooks(c *gin.Context) {
	var newBook book

	if err := c.BindJSON(&newBook); err != nil {
		return
	}

	if newBook.ID != "" {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "cannot assign book id"})
		return
	}

	newId := len(books)
	newBook.ID = strconv.Itoa(newId)

	books = append(books, newBook)
	c.IndentedJSON(http.StatusCreated, newBook)
}

func getBookByID(c *gin.Context) {
	id := c.Param("id")

	for _, book := range books {
		if book.ID == id {
			c.IndentedJSON(http.StatusOK, book)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"error": "book not found"})
}

func deleteBookByID(c *gin.Context) {
	id := c.Param("id")

	for i, book := range books {
		if book.ID != id {
			continue
		}

		books = append(books[:i], books[i+1:]...)
		c.IndentedJSON(http.StatusOK, gin.H{"message": "book deleted successfully"})
		return
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"error": "book not found"})
}

func updateBook(c *gin.Context) {
	var newBookData book

	if err := c.BindJSON(&newBookData); err != nil {
		return
	}

	for i, book := range books {
		if book.ID != newBookData.ID {
			continue
		}

		books[i] = newBookData
		c.IndentedJSON(http.StatusOK, gin.H{"message": "book updated"})
		return
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"error": "book not found"})
}
