package controller

import (
	"log"
	"net/http"
	"task-management-api/config"
	"task-management-api/domain/model"
	"task-management-api/usecase"

	"github.com/gin-gonic/gin"
)


type authcontroller struct {
	AuthorizationUsecase usecase.AuthorizationUsecase
	newEnvironment config.Environment
}

func NewAuthController(newEnvironment config.Environment, authorizationUsecase usecase.AuthorizationUsecase) *authcontroller {
	return &authcontroller{
		AuthorizationUsecase: authorizationUsecase,
		newEnvironment: newEnvironment,
	}
}

func (au *authcontroller) Register(c *gin.Context){

	var newUser *model.UserCreate
	if err := c.BindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Bad Request"})
		return
	}

	_ , err := au.AuthorizationUsecase.Register(newUser)	
	if err != nil {
		
		if err.Error() == "username already exists" {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User created successfully"})
}

func (uc *authcontroller) Login(c *gin.Context){
	var userLogin *model.UserLogin

	if err := c.BindJSON(&userLogin); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Bad Request"})
		return
	}

	token, err := uc.AuthorizationUsecase.Login(userLogin)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})

}