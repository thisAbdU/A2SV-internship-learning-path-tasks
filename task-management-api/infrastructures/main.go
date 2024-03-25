package databaseConnection

import (
	databaseConnection "example/GO-PRACTICE-EXERCISE/GO-API-exercise/databaseConnection"
	task "example/GO-PRACTICE-EXERCISE/GO-API-exercise/entities"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
    db := ConnectToDatabase()()

    taskRepository := task.Task.NewMongoTaskRepository(db)
    taskUsecase := &task.TaskUsecase{TaskRepository: taskRepository}

    router := gin.Default()
    router.GET("/tasks", func(c *gin.Context) {
        tasks, err := taskUsecase.GetTasks()
        if err != nil{
            c.JSON(http.StatusInternalServerError, gin.H{"message": "could not find task"})
            return
        }
        c.JSON(http.StatusOK, tasks)
    })
    router.GET("/tasks/:id", func(c *gin.Context) {
        id := c.Param("id")
        task, err := taskUsecase.GetTaskByID(id)
        // Handle errors and respond accordingly
    })
    router.PUT("/tasks/:id", func(c *gin.Context) {
        id := c.Param("id")
        var updatedTask task.Task
        if err := c.BindJSON(&updatedTask); err != nil {
            // Handle JSON parsing errors
        }
        err := taskUsecase.UpdateTask(id, updatedTask)
        // Handle errors and respond accordingly
    })
    router.DELETE("/tasks/:id", func(c *gin.Context) {
        id := c.Param("id")
        err := taskUsecase.DeleteTask(id)
        // Handle errors and respond accordingly
    })
    router.POST("/tasks", func(c *gin.Context) {
        var newTask task.Task
        if err := c.BindJSON(&newTask); err != nil {
            // Handle JSON parsing errors
        }
        err := taskUsecase.CreateTask(newTask)
        // Handle errors and respond accordingly
    })

    router.Run("localhost:8080")
}
