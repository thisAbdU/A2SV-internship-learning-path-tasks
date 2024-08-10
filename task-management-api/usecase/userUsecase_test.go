package usecase_test

import (
	"context"
	"errors"
	"task-management-api/domain/entities"
	"task-management-api/domain/mocks"
	"task-management-api/domain/model"
	"task-management-api/usecase"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestGetUsers(t *testing.T) {
	mockUserRepository := new(mocks.UserRepository)
	ctx := context.TODO()
	param := "testParam"

	t.Run("success", func(t *testing.T) {
		mockUsers := []*entities.User{
			{UserName: "user1"},
			{UserName: "user2"},
		}

		mockUserRepository.On("GetUser", mock.Anything, param).Return(mockUsers, nil).Once()

		u := usecase.NewUserUsecase(mockUserRepository)

		users, err := u.GetUsers(ctx, param)

		assert.NoError(t, err)
		assert.Equal(t, mockUsers, users)

		mockUserRepository.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		expectedErr := errors.New("repository error")

		mockUserRepository.On("GetUser", mock.Anything, param).Return(nil, expectedErr).Once()

		u := usecase.NewUserUsecase(mockUserRepository)

		users, err := u.GetUsers(ctx, param)

		assert.Nil(t, users)
		assert.Error(t, err)
		assert.Equal(t, expectedErr, err)

		mockUserRepository.AssertExpectations(t)
	})
}

func TestGetUserByID(t *testing.T) {
	mockUserRepository := new(mocks.UserRepository)
	ctx := context.TODO()
	userID := primitive.NewObjectID()

	t.Run("success", func(t *testing.T) {
		mockUser := &entities.User{
			ID:       userID,
			UserName: "user1",
			Password: "123",
		}

		mockUserRepository.On("GetUserByID", mock.Anything, userID).Return(mockUser, nil).Once()

		u := usecase.NewUserUsecase(mockUserRepository)

		user, err := u.GetUserByID(ctx, userID.Hex())

		assert.NoError(t, err)
		assert.Equal(t, mockUser, user)

		mockUserRepository.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		expectedErr := errors.New("repository error")

		mockUserRepository.On("GetUserByID", mock.Anything, userID).Return(nil, expectedErr).Once()

		u := usecase.NewUserUsecase(mockUserRepository)

		user, err := u.GetUserByID(ctx, userID.Hex())

		assert.Nil(t, user)
		assert.Error(t, err)
		assert.Equal(t, expectedErr, err)

		mockUserRepository.AssertExpectations(t)
	})
}

func TestUpdateUser(t *testing.T) {
	mockUserRepository := new(mocks.UserRepository)
	ctx := context.TODO()
	userID := primitive.NewObjectID()

	updatedUser := entities.User{
		ID:       userID,
		UserName: "updatedUser",
		Password: "123",
	}

	t.Run("success", func(t *testing.T) {
		mockUserRepository.On("UpdateUser", ctx, userID, updatedUser).Return(nil).Once()

		u := usecase.NewUserUsecase(mockUserRepository)

		err := u.UpdateUser(context.TODO(), userID.Hex(), updatedUser)

		assert.NoError(t, err)

		mockUserRepository.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		expectedErr := errors.New("repository error")

		mockUserRepository.On("UpdateUser", ctx, userID, updatedUser).Return(expectedErr).Once()

		u := usecase.NewUserUsecase(mockUserRepository)

		err := u.UpdateUser(context.TODO(), userID.Hex(), updatedUser)

		assert.Error(t, err)
		assert.Equal(t, expectedErr, err)

		mockUserRepository.AssertExpectations(t)
	})
}

func TestDeleteUser(t *testing.T) {
	mockUserRepository := new(mocks.UserRepository)
	ctx := context.TODO()
	userID := "testUserID"

	t.Run("success", func(t *testing.T) {
		mockUserRepository.On("DeleteUser", ctx, userID).Return(nil).Once()

		u := usecase.NewUserUsecase(mockUserRepository)

		err := u.DeleteUser(context.TODO(), userID)

		assert.NoError(t, err)

		mockUserRepository.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		expectedErr := errors.New("repository error")

		mockUserRepository.On("DeleteUser", ctx, userID).Return(expectedErr).Once()

		u := usecase.NewUserUsecase(mockUserRepository)

		err := u.DeleteUser(context.TODO(), userID)

		assert.Error(t, err)
		assert.Equal(t, expectedErr, err)

		mockUserRepository.AssertExpectations(t)
	})
}

func TestCreateUser(t *testing.T) {
	mockUserRepository := new(mocks.UserRepository)
	ctx := context.TODO()
	newUser := model.UserCreate{
		Username: "john_doe",
		Password: "P@ssw0rd123",
		Email: "john.doe@example.com",
		Name: "John Doe",
		Bio: "Software engineer with a passion for open-source projects and tech innovations.",
	}

	t.Run("success", func(t *testing.T) {
		mockUserRepository.On("CreateUser", ctx, newUser).Return("newUserID", nil).Once()

		u := usecase.NewUserUsecase(mockUserRepository)

		err := u.CreateUser(context.TODO(), newUser)

		assert.NoError(t, err)

		mockUserRepository.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		expectedErr := errors.New("repository error")

		mockUserRepository.On("CreateUser", ctx, newUser).Return("", expectedErr).Once()

		u := usecase.NewUserUsecase(mockUserRepository)

		err := u.CreateUser(context.TODO(), newUser)

		assert.Error(t, err)
		assert.Equal(t, expectedErr, err)

		mockUserRepository.AssertExpectations(t)
	})
}
