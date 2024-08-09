package usecase

import (
	"context"
	"errors"
	"log"

	// "log"
	"task-management-api/config"
	"task-management-api/domain/entities"
	"task-management-api/domain/model"
	"task-management-api/middleware"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AuthorizationUsecase struct {
	context context.Context
	userRepository entities.UserRepository
	Environment   *config.Environment
}

func NewAuthorizationUsecase(userRepository entities.UserRepository) *AuthorizationUsecase {
	return &AuthorizationUsecase{
		userRepository: userRepository,
	}
}

func (uc *AuthorizationUsecase) Login(userLogin *model.UserLogin) (string, error) {
    user, err := uc.userRepository.GetUserByUsername(uc.context, userLogin.Username)

    if err != nil {
        return "", errors.New("user Not Found")
    }
    if user.Password != userLogin.Password {
        return "", errors.New("invalid Password")
    }

    token, err := middleware.GenerateToken(user.ID.Hex())
    if err != nil {
        return "", errors.New("token Generation Failed")
    }
    
    return token, nil
}


func (uc *AuthorizationUsecase) Register(userCreate *model.UserCreate) (*model.UserInfo, error) {
    if userCreate == nil || userCreate.Username == "" || userCreate.Password == "" {
        return nil, errors.New("invalid user data")
    }

    existingUser, err := uc.userRepository.GetUserByUsername(uc.context, userCreate.Username)
    if err != nil{
        return nil, err
    }

    if existingUser.UserName != "" {
        return nil, errors.New("username already exists")
    }
    log.Println("runing")


    newUser := &model.UserCreate{
        ID:       primitive.NewObjectID(),
        Username: userCreate.Username,
        Password: userCreate.Password,
        Email:   userCreate.Email,
        Bio:    userCreate.Bio,
    }

    userInfo, err := uc.userRepository.CreateUser(uc.context, *newUser)
        if err != nil {
            return nil, errors.New("user Creation Unseccssfull")
        }

        return userInfo, nil
    }


    func (uc *AuthorizationUsecase) AdminRegister(currUser *entities.AuthenticatedUser, userCreate *model.UserCreate, param any)(*model.UserInfo, error){
        if currUser.Role != "ADMIN" {
            return nil, errors.New("Unauthorized")
        }
        
        return uc.Register(userCreate)
    }