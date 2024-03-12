package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/books", getBook)
	router.POST("/books", addBook)
	router.GET("/books/:id", checkoutBook)
	router.GET("/books/:id", bookById)
	router.Run("localhost:8080")
}
