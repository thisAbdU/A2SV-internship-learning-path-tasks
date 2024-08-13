package controller_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"task-management-api/controller"
	"task-management-api/domain/entities"
	"task-management-api/domain/mocks"
	"task-management-api/domain/model"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestGetUsers(t *testing.T) {
    gin.SetMode(gin.TestMode)

    t.Run("internal server error", func(t *testing.T) {
		mockUsecase := new(mocks.UserUsecase)
		mockEnvironment := new(mocks.Environment)
	
		router := gin.Default()
		uc := controller.NewUserController(mockEnvironment, mockUsecase)
		router.GET("/users", uc.GetUsers)

        mockUsecase.On("GetUsers", mock.Anything, "param_value").Return(nil, errors.New("some error"))
        req, _ := http.NewRequest(http.MethodGet, "/users?param=param_value", nil)
        w := httptest.NewRecorder()
        router.ServeHTTP(w, req)

        assert.Equal(t, http.StatusInternalServerError, w.Code)
        assert.JSONEq(t, `{"message": "error retrieving users"}`, w.Body.String())
    })

	t.Run("successful retrieval", func(t *testing.T) {
		mockUsecase := new(mocks.UserUsecase)
		mockEnvironment := new(mocks.Environment)
	
		router := gin.Default()
		uc := controller.NewUserController(mockEnvironment, mockUsecase)
		router.GET("/users", uc.GetUsers)

		useID1 := primitive.NewObjectID()
		useID2 := primitive.NewObjectID()
	
		users := []*entities.User{
			{ID: useID1, UserName: "John Doe"},
			{ID: useID2, UserName: "Jane Doe"},
		}
	
		mockUsecase.On("GetUsers", mock.Anything, "param_value").Return(users, nil)
		req, _ := http.NewRequest(http.MethodGet, "/users?param=param_value", nil)
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)
	
		assert.Equal(t, http.StatusOK, w.Code)
		assert.JSONEq(t, `{"users": [{"ID": "`+useID1.Hex()+`", "username": "John Doe", "password": ""}, {"ID": "`+useID2.Hex()+`", "username": "Jane Doe", "password": ""}]}`, w.Body.String())
	})
}

func TestGetUserByID(t *testing.T) {

    t.Run("internal server error", func(t *testing.T) {
        mockUsecase := new(mocks.UserUsecase)
        mockEnvironment := new(mocks.Environment)
    
        router := gin.Default()
        uc := controller.NewUserController(mockEnvironment, mockUsecase)
        router.GET("/users/:id", uc.GetUserByID)

        mockUsecase.On("GetUserByID", mock.Anything, "id_value").Return(nil, errors.New("some error"))
        req, _ := http.NewRequest(http.MethodGet, "/users/id_value", nil)
        w := httptest.NewRecorder()
        router.ServeHTTP(w, req)
        assert.Equal(t, http.StatusInternalServerError, w.Code)
        assert.JSONEq(t, `{"message": "error retrieving user"}`, w.Body.String())
    })

    t.Run("success", func(t *testing.T) {
        mockUsecase := new(mocks.UserUsecase)
        mockEnvironment := new(mocks.Environment)
    
        router := gin.Default()
        uc := controller.NewUserController(mockEnvironment, mockUsecase)
        router.GET("/users/:id", uc.GetUserByID)
        
        mockUser := entities.User{ID: primitive.NewObjectID(), UserName: "User 1"}
        mockUsecase.On("GetUserByID", mock.Anything, "id_value").Return(&mockUser, nil)
        req, _ := http.NewRequest(http.MethodGet, "/users/id_value", nil)
        w := httptest.NewRecorder()
        router.ServeHTTP(w, req)
        assert.Equal(t, http.StatusOK, w.Code)
        expectedResponse, _ := json.Marshal(gin.H{"user": mockUser})
        assert.JSONEq(t, string(expectedResponse), w.Body.String())
    })
}

func TestUpdateUser(t *testing.T) {

    t.Run("bad request", func(t *testing.T) {
        mockUsecase := new(mocks.UserUsecase)
        mockEnvironment := new(mocks.Environment)

        router := gin.Default()
        uc := controller.NewUserController(mockEnvironment, mockUsecase)
        router.PUT("/users/:id", uc.UpdateUser)

        req, _ := http.NewRequest(http.MethodPut, "/users/id_value", nil)
        w := httptest.NewRecorder()
        router.ServeHTTP(w, req)
        assert.Equal(t, http.StatusBadRequest, w.Code)
        assert.JSONEq(t, `{"message": "Bad Request"}`, w.Body.String())
    })

    t.Run("not found", func(t *testing.T) {
        mockUsecase := new(mocks.UserUsecase)
        mockEnvironment := new(mocks.Environment)

        router := gin.Default()
        uc := controller.NewUserController(mockEnvironment, mockUsecase)
        router.PUT("/users/:id", uc.UpdateUser)

        mockUser := entities.User{ID: primitive.NewObjectID(), UserName: "User 1"}
        mockUsecase.On("UpdateUser", mock.Anything, "id_value", mockUser).Return(errors.New("some error"))

        userJSON, _ := json.Marshal(mockUser)
        req, _ := http.NewRequest(http.MethodPut, "/users/id_value", bytes.NewBuffer(userJSON))
        w := httptest.NewRecorder()
        router.ServeHTTP(w, req)
        assert.Equal(t, http.StatusNotFound, w.Code)
        assert.JSONEq(t, `{"message": "Not Found"}`, w.Body.String())
    })

    t.Run("success", func(t *testing.T) {
        mockUsecase := new(mocks.UserUsecase)
        mockEnvironment := new(mocks.Environment)

        router := gin.Default()
        uc := controller.NewUserController(mockEnvironment, mockUsecase)
        router.PUT("/users/:id", uc.UpdateUser)

        mockUser := entities.User{UserName: "Updated User"}
        mockUsecase.On("UpdateUser", mock.Anything, "id_value", mockUser).Return(nil)

        userJSON, _ := json.Marshal(mockUser)
        req, _ := http.NewRequest(http.MethodPut, "/users/id_value", bytes.NewBuffer(userJSON))
        w := httptest.NewRecorder()
        router.ServeHTTP(w, req)
        assert.Equal(t, http.StatusOK, w.Code)
    })
}

func TestDeleteUser(t *testing.T) {
    gin.SetMode(gin.TestMode)

    t.Run("internal server error", func(t *testing.T) {
        mockUsecase := new(mocks.UserUsecase)
        mockEnvironment := new(mocks.Environment)

        router := gin.Default()
        uc := controller.NewUserController(mockEnvironment, mockUsecase)
        router.DELETE("/users/:id", uc.DeleteUser)

        mockUsecase.On("DeleteUser", mock.Anything, "id_value").Return(errors.New("some error"))

        req, _ := http.NewRequest(http.MethodDelete, "/users/id_value", nil)
        w := httptest.NewRecorder()
        router.ServeHTTP(w, req)
        assert.Equal(t, http.StatusInternalServerError, w.Code)
        assert.JSONEq(t, `{"message": "Internal Server Error"}`, w.Body.String())
    })

    t.Run("success", func(t *testing.T) {
        mockUsecase := new(mocks.UserUsecase)
        mockEnvironment := new(mocks.Environment)

        router := gin.Default()
        uc := controller.NewUserController(mockEnvironment, mockUsecase)
        router.DELETE("/users/:id", uc.DeleteUser)

        mockUsecase.On("DeleteUser", mock.Anything, "id_value").Return(nil)

        req, _ := http.NewRequest(http.MethodDelete, "/users/id_value", nil)
        w := httptest.NewRecorder()
        router.ServeHTTP(w, req)
        assert.Equal(t, http.StatusOK, w.Code)
    })
}

func TestCreateUser(t *testing.T) {
    gin.SetMode(gin.TestMode)

    t.Run("bad request", func(t *testing.T) {
        mockUsecase := new(mocks.UserUsecase)
        mockEnvironment := new(mocks.Environment)

        router := gin.Default()
        uc := controller.NewUserController(mockEnvironment, mockUsecase)

        router.POST("/users", uc.CreateUser)

        req, _ := http.NewRequest(http.MethodPost, "/users", bytes.NewReader([]byte(`{"invalid json"}`)))
        req.Header.Set("Content-Type", "application/json")

        w := httptest.NewRecorder()
        router.ServeHTTP(w, req)
        assert.Equal(t, http.StatusBadRequest, w.Code)
        assert.JSONEq(t, `{"message": "Bad Request"}`, w.Body.String())
    })

    t.Run("internal server error", func(t *testing.T) {
        mockUsecase := new(mocks.UserUsecase)
        mockEnvironment := new(mocks.Environment)

        router := gin.Default()
        uc := controller.NewUserController(mockEnvironment, mockUsecase)
        router.POST("/users", uc.CreateUser)
		
        newUser := model.UserCreate{
            Username: "testuser",
            Password: "password123",
            Email:    "testuser@example.com",
            Name:     "Test User",
            Bio:      "This is a test user",
        }

        reqBody, _ := json.Marshal(newUser)
        mockUsecase.On("CreateUser", mock.Anything, newUser).Return(errors.New("some error"))

        req, _ := http.NewRequest(http.MethodPost, "/users", bytes.NewBuffer(reqBody))
        w := httptest.NewRecorder()
        router.ServeHTTP(w, req)
        assert.Equal(t, http.StatusInternalServerError, w.Code)
        assert.JSONEq(t, `{"message": "Internal Server Error"}`, w.Body.String())
    })

    t.Run("success", func(t *testing.T) {
        mockUsecase := new(mocks.UserUsecase)
        mockEnvironment := new(mocks.Environment)

        router := gin.Default()
        uc := controller.NewUserController(mockEnvironment, mockUsecase)
        router.POST("/users", uc.CreateUser)

        newUser := model.UserCreate{
            Username: "testuser",
            Password: "password123",
            Email:    "testuser@example.com",
            Name:     "Test User",
            Bio:      "This is a test user",
        }

        reqBody, _ := json.Marshal(newUser)
        mockUsecase.On("CreateUser", mock.Anything, newUser).Return(nil)

        req, _ := http.NewRequest(http.MethodPost, "/users", bytes.NewBuffer(reqBody))
        w := httptest.NewRecorder()
        router.ServeHTTP(w, req)

        newUserJSON, _ := json.Marshal(newUser)

        expectedResponse := fmt.Sprintf(`{"message": "User created successfully", "user": %s}`, newUserJSON)

        assert.Equal(t, http.StatusCreated, w.Code)
        assert.JSONEq(t, expectedResponse, w.Body.String())
    })
}
func TestNewUserController(t *testing.T) {
	mockUsecase := new(mocks.UserUsecase)
	mockEnvironment := new(mocks.Environment)

	uc := controller.NewUserController(mockEnvironment, mockUsecase)

	assert.NotNil(t, uc)
}


