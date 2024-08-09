package usecase_test

import (
	"errors"
	"task-management-api/domain/mocks"
	"task-management-api/domain/model"
	"task-management-api/usecase"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestRegister(t *testing.T) {
	mockUserRepository := mocks.NewUserRepository(t)
	
	au := usecase.NewAuthorizationUsecase(mockUserRepository)

	userCreate := &model.UserCreate{
		ID:       primitive.NewObjectID(),
		Username: "testuser",
		Password: "password",
		Email:    "test@example.com",
		Bio:      "A test user",
	}

	userInfo := &model.UserInfo{
		ID:       userCreate.ID.Hex(),
		Username: userCreate.Username,
		Email:    userCreate.Email,
	}

	t.Run("Successful Registration", func(t *testing.T) {
		mockUserRepository.On("GetUserByID", mock.Anything, userCreate.ID.Hex()).Return(nil, errors.New("mongo: no documents in result"))
		mockUserRepository.On("CreateUser", mock.Anything, *userCreate).Return(userInfo, nil)

		result, err := au.Register(userCreate)

		assert.NoError(t, err)
		assert.Equal(t, userInfo, result)
		mockUserRepository.AssertExpectations(t)
	})
}
