package usecase

import (
	"context"
	"errors"
	"task-management-api/config"
	"task-management-api/domain/entities"
	"task-management-api/domain/model"
	"task-management-api/middleware"
)

type AuthorizationUsecase struct {
	context context.Context
	userRepository entities.UserRepository
	Environment   *config.Environment
}

func NewAuthorizationUsecase(environment *config.Environment, userRepository *entities.UserRepository) *AuthorizationUsecase {
	return &AuthorizationUsecase{
		userRepository: *userRepository,
		Environment:    environment,
	}
}

func (uc *AuthorizationUsecase) Login(currUser *entities.AuthenticatedUser, userLogin *model.UserLogin, param any) (string, error) {
    user, err := uc.userRepository.GetUserByID(uc.context, userLogin.Id)
    if err != nil {
        return "", errors.New("user Not Found")
    }
    if user.Password != userLogin.Password {
        return "", errors.New("invalid Password")
    }
    token, err := middleware.GenerateToken(user.ID)
    if err != nil {
        return "", errors.New("token Generation Failed")
    }
    return token, nil
}


func (uc *AuthorizationUsecase) Register(currUser *entities.AuthenticatedUser, userCreate *model.UserCreate, param any) (*model.UserInfo, string, error) {
    if userCreate == nil || userCreate.Username == "" || userCreate.Password == "" {
        return nil, "Invalid user data", errors.New("invalid user data")
    }

    existingUser, err := uc.userRepository.GetUserByID(uc.context, userCreate.Id)
    if err != nil {
        return nil, "Failed to check existing user", err
    }
if existingUser != nil {
	return nil, "Username already exists", errors.New("username already exists")
}


newUser := &entities.User{
	ID:       userCreate.Id,
	UserName: userCreate.Username,
	Password: userCreate.Password,
}

userInfo, err := uc.userRepository.CreateUser(uc.context, *newUser)
    if err != nil {
        return nil, "Failed to create user", err
    }

    return userInfo, "", nil
}


func (uc *AuthorizationUsecase) AdminRegister(currUser *entities.AuthenticatedUser, userCreate *model.UserCreate, param any)(*model.UserInfo, string, error){
	if currUser.Role != "ADMIN" {
		return nil, "User Creation Unseccssfull", errors.New("Unauthorized")
	}
	return uc.Register(currUser, userCreate, param)
}