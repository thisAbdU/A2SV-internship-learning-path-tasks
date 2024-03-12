package main

import(
	"errors"
	"net/http"
	"github.com/gin-gonic/gin"
)

type book struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Author   string `json:"author"`
	Quantity int    `json:"quantity"`
}

var books = []book{
	{ID: "1", Title: "In Search of Lost Time", Author: "Marcel Proust", Quantity: 2},
	{ID: "2", Title: "The Great Gatsby", Author: "F. Scott Fitzgerald", Quantity: 5},
	{ID: "3", Title: "War and Peace", Author: "Leo Tolstoy", Quantity: 6},
}

func getBook(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, books)
}

func addBook(c *gin.Context)  {
	var newBook book
	
	if err := c.BindJSON(&newBook); err != nil{
		return 
	}

	books = append(books, newBook)
	c.IndentedJSON(http.StatusCreated, newBook)
}

func bookById(c *gin.Context){
    id := c.Param("id")
    book, err := getBookById(id)

    if err != nil {
        c.IndentedJSON(http.StatusNotFound, gin.H{"message": "book not found"})
        return
    }
    
    c.IndentedJSON(http.StatusOK, book)
}

func getBookById(id string) (*book, error) {
    for _, b := range books {
        if b.ID == id {
            return &b, nil
        }
    }
    return nil, errors.New("book not found")
}

func checkoutBook(c *gin.Context){
	id, ok := c.GetQuery("id")

	if !ok{
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "missing id parameter"})
		return 
	}

	book, err := getBookById(id)

	if err != nil{
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "book not found"})
		return
	}

	if book.Quantity <= 0{
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "book not available"})
		return
	}

	book.Quantity -= 1
	c.IndentedJSON(http.StatusOK, book)
}

