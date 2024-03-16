package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// album represents data about a record album.
type album struct {
    ID     string  `json:"id"`
    Title  string  `json:"title"`
    Artist string  `json:"artist"`
    Price  float64 `json:"price"`
}

// albums slice to seed record album data.
var albums = []album{
    {ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
    {ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
    {ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func getAlbums(c *gin.Context){
	c.IndentedJSON(http.StatusOK, albums)
}
func postAlbums(c *gin.Context){
	var newAlbum album
	if err := c.BindJSON(&newAlbum);err!=nil{
		return
	}

	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}
func getAlbumByID(c *gin.Context)  {
	id := c.Param("id")

	for _, a := range albums{
		if id == a.ID{
			c.IndentedJSON(http.StatusOK, a)
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "id not found!"})
}

func main() {
    router := gin.Default()
    router.GET("/albums", getAlbums)
	router.POST("/albums/", postAlbums)
	router.GET("/albums/:id", getAlbumByID)
    router.Run("localhost:8080")
}