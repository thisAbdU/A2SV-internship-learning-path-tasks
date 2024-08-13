package usecase_test

import (
	"errors"
	"task-management-api/domain/entities"
	"task-management-api/domain/mocks"
	"task-management-api/domain/model"
	"task-management-api/mongo"
	"task-management-api/usecase"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestRegister(t *testing.T) {

    t.Run("successful registration", func(t *testing.T) {
		mockUserRepository := mocks.NewUserRepository(t)
		mockUtils := mocks.NewUtils(t)
		uc := usecase.NewAuthUseCase(mockUserRepository, mockUtils)

        userCreate := &model.UserCreate{
			Username: "testuser",
            Password: "password",
            Email:    "test@example.com",
            Bio:      "A test user",
        }

        mockUserRepository.On("GetUserByUsername", mock.Anything, userCreate.Username).Return(nil, mongo.ErrNoDocuments)
        mockUserRepository.On("CreateUser", mock.Anything, mock.Anything).Return(&model.UserInfo{ID: primitive.NewObjectID().Hex(), Username: userCreate.Username}, nil)

        userInfo, err := uc.Register(userCreate)

        assert.NoError(t, err)
        assert.NotNil(t, userInfo)
        assert.Equal(t, userCreate.Username, userInfo.Username)
        mockUserRepository.AssertExpectations(t)
    })

    t.Run("user already exists", func(t *testing.T) {
		mockUserRepository := mocks.NewUserRepository(t)
		mockUtils := mocks.NewUtils(t)
		uc := usecase.NewAuthUseCase(mockUserRepository, mockUtils)

        userCreate := &model.UserCreate{
            Username: "testuser",
            Password: "password",
            Email:    "test@example.com",
            Bio:      "A test user",
        }

        existingUser := &entities.User{ID: primitive.NewObjectID(), UserName: userCreate.Username}
        mockUserRepository.On("GetUserByUsername", mock.Anything, userCreate.Username).Return(existingUser, nil)

        userInfo, err := uc.Register(userCreate)

        assert.Error(t, err)
        assert.Nil(t, userInfo)
        assert.Equal(t, "username already exists", err.Error())
        mockUserRepository.AssertExpectations(t)
    })

    t.Run("invalid user data", func(t *testing.T) {
		mockUserRepository := mocks.NewUserRepository(t)
		mockUtils := mocks.NewUtils(t)
		uc := usecase.NewAuthUseCase(mockUserRepository, mockUtils)

        userCreate := &model.UserCreate{
            Username: "",
            Password: "",
        }

        userInfo, err := uc.Register(userCreate)

        assert.Error(t, err)
        assert.Nil(t, userInfo)
        assert.Equal(t, "invalid user data", err.Error())
    })

    t.Run("repository error", func(t *testing.T) {
		mockUserRepository := mocks.NewUserRepository(t)
		mockUtils := mocks.NewUtils(t)
		uc := usecase.NewAuthUseCase(mockUserRepository, mockUtils)

        userCreate := &model.UserCreate{
            Username: "testuser",
            Password: "password",
            Email:    "test@example.com",
            Bio:      "A test user",
        }

        mockUserRepository.On("GetUserByUsername", mock.Anything, userCreate.Username).Return(nil, errors.New("repository error"))

        userInfo, err := uc.Register(userCreate)

        assert.Error(t, err)
        assert.Nil(t, userInfo)
        assert.Equal(t, "repository error", err.Error())
        mockUserRepository.AssertExpectations(t)
    })
}

func TestLogin(t *testing.T) {
    t.Run("successful login", func(t *testing.T) {
		mockUserRepository := mocks.NewUserRepository(t)
		mockUtils := mocks.NewUtils(t)
		uc := usecase.NewAuthUseCase(mockUserRepository, mockUtils)

		
        userLogin := &model.UserLogin{
            Username: "testuser",
            Password: "password",
        }

        user := &entities.User{
            ID:       primitive.NewObjectID(),
            UserName: userLogin.Username,
			Password: userLogin.Password,
        }

        mockUserRepository.On("GetUserByUsername", mock.Anything, userLogin.Username).Return(user, nil)
        mockUtils.On("GenerateToken", mock.AnythingOfType("string")).Return("mockToken", nil)
        token, err := uc.Login(userLogin)

        assert.NoError(t, err)
        assert.Equal(t, "mockToken", token)
        mockUserRepository.AssertExpectations(t)
    })

    t.Run("user not found", func(t *testing.T) {
		mockUserRepository := mocks.NewUserRepository(t)
		mockUtils := mocks.NewUtils(t)
		uc := usecase.NewAuthUseCase(mockUserRepository, mockUtils)

        userLogin := &model.UserLogin{
            Username: "nonexistentuser",
            Password: "password",
        }

        mockUserRepository.On("GetUserByUsername", mock.Anything, userLogin.Username).Return(nil, errors.New("user Not Found"))

        token, err := uc.Login(userLogin)

        assert.Error(t, err)
        assert.Equal(t, "", token)
        assert.Equal(t, "user Not Found", err.Error())
        mockUserRepository.AssertExpectations(t)
    })

    t.Run("invalid password", func(t *testing.T) {
		mockUserRepository := mocks.NewUserRepository(t)
		mockUtils := mocks.NewUtils(t)
		uc := usecase.NewAuthUseCase(mockUserRepository, mockUtils)

        userLogin := &model.UserLogin{
            Username: "testuser",
            Password: "wrongpassword",
        }

        user := &entities.User{
            ID:       primitive.NewObjectID(),
            UserName: userLogin.Username,
        }

        mockUserRepository.On("GetUserByUsername", mock.Anything, userLogin.Username).Return(user, nil)

        token, err := uc.Login(userLogin)

        assert.Error(t, err)
        assert.Equal(t, "", token)
        assert.Equal(t, "invalid Password", err.Error())
        mockUserRepository.AssertExpectations(t)
    })

	t.Run("token generation failed", func(t *testing.T) {
		mockUserRepository := mocks.NewUserRepository(t)
		mockUtils := mocks.NewUtils(t)
		uc := usecase.NewAuthUseCase(mockUserRepository, mockUtils)

		userLogin := &model.UserLogin{
			Username: "testuser",
			Password: "password",
		}
	
		user := &entities.User{
			ID:       primitive.NewObjectID(),
			UserName: userLogin.Username,
			Password: userLogin.Password,
		}
	
		mockUserRepository.On("GetUserByUsername", mock.Anything, userLogin.Username).Return(user, nil)
		mockUtils.On("GenerateToken", mock.AnythingOfType("string")).Return("", errors.New("token Generation Failed"))
	
		token, err := uc.Login(userLogin)
	
		assert.Error(t, err)
		assert.Equal(t, "", token)
		assert.Equal(t, "token Generation Failed", err.Error())
		mockUserRepository.AssertExpectations(t)
		mockUtils.AssertExpectations(t)
	})
}