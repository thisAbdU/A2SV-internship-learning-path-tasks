package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

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

	db := connectToDatabase()
	defer db.Client().Disconnect(context.Background())

	var tasks []Task

	collection := db.Client().Database("taskManagementDatabase").Collection("task")

	filter := bson.M{}

	cursor, err := collection.Find(context.Background(), filter)
	if err != nil{
		log.Fatal(err)
	}
	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()){
		var task Task
		if err := cursor.Decode(&task); err != nil{
			log.Fatal(err)
		}
		tasks = append(tasks, task)
	}

	if err := cursor.Err(); err != nil{
		log.Fatal(err)
	}
	c.IndentedJSON(http.StatusOK, tasks)
}

func getTasksByID(c *gin.Context) {
	db := connectToDatabase()
	defer db.Client().Disconnect(context.Background())

	collection := db.Client().Database("taskManagementDatabase").Collection("task")

	Id := c.Param("id")

	filter := bson.M{"id": Id}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Error finding task: %v", err)})
		return
	}
	defer cursor.Close(ctx)

	var task Task
	if cursor.Next(ctx) {
		err := cursor.Decode(&task)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Error decoding task: %v", err)})
			return
		}
	} else {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}

	c.IndentedJSON(http.StatusOK, task.Description)
}

func updateTask(c *gin.Context) {
    // Connect to the MongoDB database
    db := connectToDatabase()
    defer db.Client().Disconnect(context.Background())

    // Get the task collection
    collection := db.Client().Database("taskManagementDatabase").Collection("task")

    // Extract task ID from request parameters
    id := c.Param("id")

    // Define filter to find the task by its ID
    filter := bson.M{"id": id}

    // Create context with a timeout
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    // Find the task by ID
    var task Task
    if err := collection.FindOne(ctx, filter).Decode(&task); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Error finding task: %v", err)})
        return
    }

    // Parse the incoming JSON data to get updated task details
    var updatedTask Task
    if err := c.BindJSON(&updatedTask); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Error parsing request body: %v", err)})
        return
    }

    // Update the task with new details
    update := bson.M{"$set": updatedTask}
    if _, err := collection.UpdateOne(ctx, filter, update); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Error updating task: %v", err)})
        return
    }

    // Respond with the updated task
    c.JSON(http.StatusOK, updatedTask)
}

func deleteTask(c *gin.Context) {
    // Connect to the MongoDB database
    db := connectToDatabase()
    defer db.Client().Disconnect(context.Background()) // Disconnect from the database when function exits

    // Get the task collection
    collection := db.Client().Database("taskManagementDatabase").Collection("task")
    
    // Extract task ID from request parameters
    id := c.Param("id")

    // Define filter to find the task by its ID
    filter := bson.M{"id": id}
    
    // Create context with a timeout
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    // Find the task by ID
    var task Task
    if err := collection.FindOne(ctx, filter).Decode(&task); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Error finding task: %v", err)})
        return
    }

    // Delete the task
    if _, err := collection.DeleteOne(ctx, filter); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Error deleting task: %v", err)})
        return
    }

    // Respond with success message
    c.JSON(http.StatusOK, gin.H{"message": "Task deleted successfully"})
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
