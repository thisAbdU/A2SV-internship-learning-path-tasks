package controller

import (
	"context"
	"log"
	"task-management-api/config"
	"task-management-api/domain/entities"
	"task-management-api/domain/model"
	"task-management-api/usecase"

	"net/http"

	"github.com/gin-gonic/gin"
)

type usercontroller struct {
	UserUsecase usecase.UserUsecase
	newEnvironment config.Environment
}

func NewUserController(newEnvironment config.Environment, userUsecase usecase.UserUsecase) *usercontroller {
	return &usercontroller{
		UserUsecase:   userUsecase,
		newEnvironment: newEnvironment,
	}
}

func (uc *usercontroller) GetUsers(c *gin.Context) {

	users, err := uc.UserUsecase.GetUsers(context.Background(), c.Query("param"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "error retrieving users"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"users": users})
}

func (uc *usercontroller) GetUserByID(c *gin.Context) {
	
	id := c.Param("id")

	user, err := uc.UserUsecase.GetUserByID(context.TODO(), id)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "error retrieving user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": user})
}

func (uc *usercontroller) UpdateUser(c *gin.Context) {
	
	id := c.Param("id")
	var updatedUser entities.User
	if err := c.BindJSON(&updatedUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Bad Request"})
		return
	}

	err := uc.UserUsecase.UpdateUser(id, updatedUser)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Not Found"})
		return
	}
}

func (uc *usercontroller) DeleteUser(c *gin.Context) {
	
	id := c.Param("id")
	err := uc.UserUsecase.DeleteUser(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error"})
		return
	}
}

func (uc *usercontroller) CreateUser(c *gin.Context) {
	
	var newUser model.UserCreate
	if err := c.BindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Bad Request"})
		return
	}

	err := uc.UserUsecase.CreateUser(newUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User created successfully", "user": newUser})
}