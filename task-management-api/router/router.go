package router

import (
	"task-management-api/config"
	"task-management-api/controller"
	"task-management-api/middleware"
	"task-management-api/repository"
	"task-management-api/usecase"
	"task-management-api/utils"
	"time"

	"task-management-api/mongo"

	"github.com/gin-gonic/gin"
)

func NewAuthRouter(environment *config.Environment, timeout time.Duration, db mongo.Database, r *gin.RouterGroup) {
	userRepository := repository.NewUserRepository(db, "user")
	tokenUtil := utils.NewTokenUtil(environment)
	authUsecase :=  usecase.NewAuthUseCase(userRepository, tokenUtil)
	authController := controller.NewAuthController(authUsecase)

	r.POST("/register", authController.Register)
	r.POST("/login", authController.Login)
}

func taskRouter(environment *config.Environment, timeout time.Duration, db mongo.Database, r *gin.RouterGroup) {
	taskRepository := repository.NewTaskRepository(db, "task")
	taskUseCase := usecase.NewTaskUsecase(taskRepository)
	taskController := controller.NewTaskController(*environment, taskUseCase)

	r.GET("/", taskController.GetTasks)
	r.POST("/", taskController.CreateTask)
	r.GET("/:id", taskController.GetTaskByID)
	r.PATCH("/:id", taskController.UpdateTask)
	r.DELETE("/:id", taskController.DeleteTask)
}


func userRouter(environment *config.Environment, timeout time.Duration, db mongo.Database, r *gin.RouterGroup) {

	userRepository := repository.NewUserRepository(db, "user")
	userUseCase := usecase.NewUserUsecase(userRepository)
	userController := controller.NewUserController(*environment, userUseCase)

	r.GET("/", userController.GetUsers)
	r.GET("/:id", userController.GetUserByID)
	r.PATCH("/:id", userController.UpdateUser).Use(middleware.AuthMiddleware())
	r.DELETE("/:id", userController.DeleteUser).Use(middleware.AuthMiddleware())
}

func NewRouter(environment config.Environment, timeout time.Duration, db mongo.Database, r *gin.Engine) {
	authRouter := r.Group("/auth")
	NewAuthRouter(&environment, timeout, db, authRouter)

	taskGroup := r.Group("/task")
	taskGroup.Use(middleware.AuthMiddleware())
	taskRouter(&environment, timeout, db, taskGroup)

	userGroup := r.Group("/")
	userRouter(&environment, timeout, db, userGroup)
}
