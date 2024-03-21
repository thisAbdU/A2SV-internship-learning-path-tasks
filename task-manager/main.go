package main

import (
	"context"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var collection *mongo.Collection

type Task struct {
    ID          string `json:"id"`
    Title       string `json:"title"`
    Description string `json:"description"`
    Status      string `json:"status"`
}

func postTask(c *gin.Context) {

	db := connectToDatabase()

	var newTask Task

	if err := c.BindJSON(&newTask); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	db.Collection("task").InsertOne(context.TODO(), newTask)
	
}

func getTasks(c *gin.Context) {
	limit := c.Query("limit")
	sortBy := c.Query("sort_by")

	var err error
	limitInt := 10 // Default limit value
	if limit != "" {
		limitInt, err = strconv.Atoi(limit)
		if err != nil || limitInt < 1 {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Invalid limit value"})
			return
		}
	}

	sortByOption := 1
	if sortBy == "desc" {
		sortByOption = -1
	}

	var tasks []Task
	ct, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := collection.Find(ct, bson.M{}, options.Find().SetSort(bson.D{{Key: "created_at", Value: sortByOption}}).SetLimit(int64(limitInt)))
	if err != nil {
		log.Fatal(err)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve tasks"})
		return
	}
	defer cursor.Close(ct)

	for cursor.Next(ct) {
		var t Task
		if err := cursor.Decode(&t); err != nil {
			log.Fatal(err)
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Failed to get task"})
			return
		}

		tasks = append(tasks, t)
	}

	if err := cursor.Err(); err != nil {
		log.Fatal(err)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Error while retrieving tasks"})
		return
	}

	c.IndentedJSON(http.StatusOK, tasks)
}

func getTasksByID(c *gin.Context) {
    // Implement retrieval of a specific task by ID from MongoDB collection
}

func updateTask(c *gin.Context) {
    // Implement updating an existing task in MongoDB collection
}

func deleteTask(c *gin.Context) {
    // Implement deleting a task from MongoDB collection
}


func main() {
    connectToDatabase()
    router := gin.Default()

    router.GET("/tasks", getTasks)
    router.GET("/tasks/:id", getTasksByID)
    router.PUT("/tasks/:id", updateTask)
    router.DELETE("/tasks/:id", deleteTask)
    router.POST("/tasks", postTask)

    router.Run("localhost:8080")
}
