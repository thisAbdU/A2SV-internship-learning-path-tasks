package router

import (
	"task-management-api/config"
	"task-management-api/controller"
	"task-management-api/repository"
	"task-management-api/usecase"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewAuthRouter(environment *config.Environment, timeout time.Duration, db *mongo.Database, r *gin.RouterGroup) {
	userRepository := repository.NewUserRepository(db, "user")
	// userUseCase := usecase.NewUserUsecase(environment, userRepository)
	// // userController := controller.NewUserController(*environment, *userUseCase)

	authUsecase :=  usecase.NewAuthorizationUsecase(environment, &userRepository)
	authController := controller.NewAuthController(*environment, *authUsecase)

	r.POST("/register", authController.Register)
	r.POST("/login", authController.Login)
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


func userRouter(environment *config.Environment, timeout time.Duration, db *mongo.Database, r *gin.RouterGroup) {

	userRepository := repository.NewUserRepository(db, "user")
	userUseCase := usecase.NewUserUsecase(environment, userRepository)
	userController := controller.NewUserController(*environment, *userUseCase)

	r.GET("/", userController.GetUsers)
	r.GET("/:id", userController.GetUserByID)
	r.PUT("/:id", userController.UpdateUser)
	r.DELETE("/:id", userController.DeleteUser)
}

func NewRouter(environment *config.Environment, timeout time.Duration, db *mongo.Database, r *gin.Engine) {
	authRouter := r.Group("/auth")
	NewAuthRouter(environment, timeout, db, authRouter)

	taskGroup := r.Group("/task")
	taskRouter(environment, timeout, db, taskGroup)
}

