package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type task struct {
	Id          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status"`
	DueDate     time.Time `json:"dueDate"`
}

var tasks = []task{
	{Id:"1", Title:"Studying", Description:"Study", Status:"In progress", DueDate:time.Date(2002, 12, 21, 0, 0, 0, 0, time.UTC)},
	{Id: "2", Title: "Task 2", Description: "Second task", DueDate: time.Now().AddDate(0, 0, 1), Status: "In Progress"},
    {Id: "3", Title: "Task 3", Description: "Third task", DueDate: time.Now().AddDate(0, 0, 2), Status: "Completed"},
}

func getTasks(c * gin.Context){
	c.IndentedJSON(http.StatusOK, tasks)
}

func getTasksByID(c *gin.Context)  {
	id := c.Param("id")
	for _, task := range tasks{
		if task.Id == id{
			c.IndentedJSON(http.StatusOK, task)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "couldn't find id"})
}

func updateTask(c *gin.Context)  {
	id := c.Param("id")

	var udpatedTask task

	if err := c.BindJSON(&udpatedTask); err != nil{
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	for _, task := range tasks{
		if task.Id == id{
			if task.Description != ""{
				task.Description = udpatedTask.Description
			}
			if task.Title != ""{
				task.Title = udpatedTask.Title
			}
			if task.Status != ""{
				task.Status = udpatedTask.Status
			}

			c.IndentedJSON(http.StatusOK, gin.H{"message": "Task updated!"})
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Task Not found"})

}

func deleteTask(c *gin.Context)  {
	id := c.Param("id")

	for i, task := range tasks{
		if task.Id == id{
			tasks = append(tasks[:i], tasks[i+1:]...)
		}
	}
}

func addTask(c *gin.Context) {

	var newTask task

	if err := c.BindJSON(&newTask); err != nil{
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	tasks = append(tasks, newTask)
	c.IndentedJSON(http.StatusCreated, tasks)

}
func main()  {
	router := gin.Default()
	router.GET("/tasks", getTasks)
	router.GET("/tasks/:id", getTasksByID)
	router.PUT("/tasks/:id", updateTask)
	router.DELETE("/tasks/:id", deleteTask)
	router.POST("/tasks/", addTask)
	
	router.Run("localhost:8080")
}