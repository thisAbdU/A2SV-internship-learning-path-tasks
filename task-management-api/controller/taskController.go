package controller

import (
	"net/http"
	"task-management-api/config"
	"task-management-api/domain/entities"

	"github.com/gin-gonic/gin"
)

type taskcontroller struct {
	TaskUsecase entities.TaskUsecase
	newEnvironment config.Environment
}

func NewTaskController(newEnvironment config.Environment,taskUsecase entities.TaskUsecase) *taskcontroller {
	return &taskcontroller{
		TaskUsecase: taskUsecase,
		newEnvironment: newEnvironment,
	}
}

func (tc *taskcontroller) GetTasks(c *gin.Context) {
    userID , exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Please log in to acess tasks"})
		return
	}

	userIDStr, ok := userID.(string)
    if !ok {
        c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal server error"})
        return
    }

    tasks, err := tc.TaskUsecase.GetTasks(userIDStr)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"message": "error retrieving tasks"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"tasks": tasks})
}


func (tc *taskcontroller) GetTaskByID(c *gin.Context){
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Please log in to acess tasks"})
		return
	}

	id := c.Param("id")

	task, err := tc.TaskUsecase.GetTaskByID(id, userID.(string))

	if err != nil {

		if err.Error() == "mongo: no documents in result" {
			c.JSON(http.StatusNotFound, gin.H{"message": "Task not found"})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"message": "error retrieving task"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"task": task})
}

func (tc *taskcontroller) UpdateTask(c *gin.Context){
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Please log in to acess tasks"})
		return
	}

	id := c.Param("id")

	var updatedTask entities.Task

	if err := c.BindJSON(&updatedTask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Bad Request"})
		return
	}

    err := tc.TaskUsecase.UpdateTask(id, updatedTask, userID.(string))
    if err != nil {
        if err.Error() == "no documents updated" {
            c.JSON(http.StatusNotFound, gin.H{"message": "Task not found"})
            return
        }
        c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error"})
        return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Task updated successfully"})
}

func (tc *taskcontroller) DeleteTask(c *gin.Context){
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "Please log in to acess tasks"})
		return
	}

	id := c.Param("id")

	err := tc.TaskUsecase.DeleteTask(id, userID.(string))
	if err != nil {
		if err.Error() == "no documents deleted"{
			c.JSON(http.StatusNotFound, gin.H{"message": "Task not found"})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Task deleted successfully"})
}
func (tc *taskcontroller) CreateTask(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists || userID == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Please sign up to create a task"})
		return
	}

	var newTask entities.Task

	if err := c.BindJSON(&newTask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Bad Request"})
		return
	}

	newTask.UserID = userID.(string)

	err := tc.TaskUsecase.CreateTask(newTask)
	if err != nil {
		switch err.Error() {
		case "Please sign up to create a task":
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Please sign up to create a task"})
			return
		case "Bad Request":
			c.JSON(http.StatusBadRequest, gin.H{"message": "Bad Request"})
			return
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error"})
		}
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Task created successfully"})
}


func (tc *taskcontroller) GetEnvironment(c *gin.Context){
	c.JSON(http.StatusOK, gin.H{"environment": tc.newEnvironment})
}
