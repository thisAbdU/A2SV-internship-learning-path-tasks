package usecase

import (
	"context"
	"errors"

	"task-management-api/domain/entities"
	"task-management-api/domain/model"
	"task-management-api/mongo"
	"task-management-api/utils"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type authUseCase struct {
	userRepository entities.UserRepository
	utils          utils.Utils
    context        context.Context
}

func NewAuthUseCase(userRepo entities.UserRepository, utils utils.Utils) entities.AuthUseCase {
	return &authUseCase{
		userRepository: userRepo,
		utils: utils,
        context:       context.TODO(),
	}
}

func (uc *authUseCase) Login(userLogin *model.UserLogin) (string, error) {
	user, err := uc.userRepository.GetUserByUsername(uc.context, userLogin.Username)

	if err != nil {
		return "", errors.New("user Not Found")
	}
	if user.Password != userLogin.Password {
		return "", errors.New("invalid Password")
	}

	token, err := uc.utils.GenerateToken(user.ID.Hex())
	if err != nil {
		return "", errors.New("token Generation Failed")
	}

	return token, nil
}

func (uc *authUseCase) Register(userCreate *model.UserCreate) (*model.UserInfo, error) {
	if userCreate == nil || userCreate.Username == "" || userCreate.Password == "" {
		return nil, errors.New("invalid user data")
	}

	existingUser, err := uc.userRepository.GetUserByUsername(uc.context, userCreate.Username)
	if err != nil {
		if !errors.Is(err, mongo.ErrNoDocuments) {
			return nil, err
		}
	} else {
		if existingUser != nil && existingUser.UserName != "" {
			return nil, errors.New("username already exists")
		}
	}

	newUser := &model.UserCreate{
		ID:       primitive.NewObjectID(),
		Username: userCreate.Username,
		Password: userCreate.Password,
		Email:    userCreate.Email,
		Bio:      userCreate.Bio,
	}

	userInfo, err := uc.userRepository.CreateUser(uc.context, *newUser)
	if err != nil {
		return nil, errors.New("user Creation Unseccssfull")
	}

	return userInfo, nil
}

func (uc *authUseCase) AdminRegister(currUser entities.AuthenticatedUser, userCreate *model.UserCreate, param any) (*model.UserInfo, error) {
	if currUser.Role != "ADMIN" {
		return nil, errors.New("Unauthorized")
	}

	return uc.Register(userCreate)
}
