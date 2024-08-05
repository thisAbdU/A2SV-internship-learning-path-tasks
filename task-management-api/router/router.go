package router

import (

	"example/GO-PRACTICE-EXERCISE/GO-API-exercise/config"
	"example/GO-PRACTICE-EXERCISE/GO-API-exercise/controller"
	"example/GO-PRACTICE-EXERCISE/GO-API-exercise/repository"
	"example/GO-PRACTICE-EXERCISE/GO-API-exercise/usecase"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)




func NewAuthRouter(environment *config.Environment, timeout time.Duration, db *mongo.Database, r *gin.RouterGroup) {
	userRepository := repository.NewUserRepository(db, "user")
	userUseCase := usecase.NewUserUsecase(environment, userRepository)
	userController := controller.NewUserController(*environment, *userUseCase)
	r.POST("/register", userController.Register)
	r.POST("/login", userController.Login)
}

func taskRouter(environment *config.Environment, timeout time.Duration, db *mongo.Database, r *gin.RouterGroup) {

	taskRepository := repository.NewTaskRepository(db, "task")
	taskUseCase := usecase.NewTaskUsecase(environment,&taskRepository)
	taskController := controller.NewTaskController(*environment, *taskUseCase)
	r.GET("/", taskController.GetTasks)
	r.GET("/:id", taskController.GetTaskByID)
	r.PUT("/:id", taskController.UpdateTask)
	r.DELETE("/:id", taskController.DeleteTask)
}

func NewRouter(environment *config.Environment, timeout time.Duration, db *mongo.Database, r *gin.Engine) {
	authRouter := r.Group("/auth")
	NewAuthRouter(environment, timeout, db, authRouter)

	taskGroup := r.Group("/task")
	taskRouter(environment, timeout, db, taskGroup)
}


func userRouter(environment *config.Environment, timeout time.Duration, db *mongo.Database, r *gin.RouterGroup) {

	userRepository := repository.NewUserRepository(db, "user")
	userUseCase := usecase.NewUserUsecase(environment, userRepository)
	userController := controller.NewUserController(*environment, *userUseCase)
	r.GET("/", userController.GetUsers)
	r.GET("/:id", userController.GetUserByID)
	r.PUT("/:id", userController.UpdateUser)
	r.DELETE("/:id", userController.DeleteUser)
}

