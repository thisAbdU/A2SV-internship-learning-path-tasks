package controller

import (
	"net/http"
	"task-management-api/config"
	"task-management-api/domain/entities"
	"task-management-api/usecase"

	"github.com/gin-gonic/gin"
)

type taskcontroller struct {
	TaskUsecase usecase.TaskUsecase
	newEnvironment config.Environment
}

func NewTaskController(newEnvironment config.Environment,taskUsecase usecase.TaskUsecase) *taskcontroller {
	return &taskcontroller{
		TaskUsecase: taskUsecase,
		newEnvironment: newEnvironment,
	}
}

func (tc *taskcontroller) GetTasks(c *gin.Context) {
   
    param := c.Query("param")
    if param == "" {
        param = c.PostForm("param")
    }

    tasks, err := tc.TaskUsecase.GetTasks(param)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"message": "error retrieving tasks"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"tasks": tasks})
}


func (tc *taskcontroller) GetTaskByID(c *gin.Context){

	id := c.Param("id")

	task, err := tc.TaskUsecase.GetTaskByID(id)	
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "error retrieving task"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"task": task})
}

func (tc *taskcontroller) UpdateTask(c *gin.Context){
	
	id := c.Param("id")
	var updatedTask entities.Task
	if err := c.BindJSON(&updatedTask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Bad Request"})
		return
	}

	err := tc.TaskUsecase.UpdateTask(id, updatedTask)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Not Found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Task updated successfully"})
}

func (tc *taskcontroller) DeleteTask(c *gin.Context){
	
	id := c.Param("id")
	err := tc.TaskUsecase.DeleteTask(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Task deleted successfully"})
}

func (tc *taskcontroller) CreateTask(c *gin.Context){
	
	var newTask entities.Task
	if err := c.BindJSON(&newTask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Bad Request"})
		return
	}

	err := tc.TaskUsecase.CreateTask(newTask)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Task created successfully"})
}

func (tc *taskcontroller) GetEnvironment(c *gin.Context){
	c.JSON(http.StatusOK, gin.H{"environment": tc.newEnvironment})
}

func (tc *taskcontroller) GetHealth(c *gin.Context){
	c.JSON(http.StatusOK, gin.H{"status": "healthy"})
}

