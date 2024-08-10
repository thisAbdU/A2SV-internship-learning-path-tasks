package controller_test

import (
	"bytes"
	"encoding/json"
	"errors"
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
	mockUsecase := new(mocks.UserUsecase)
	mockEnvironment := new(mocks.Environment)

	router := gin.Default()
	uc := controller.NewUserController(mockEnvironment, mockUsecase)
	router.GET("/users", uc.GetUsers)

	t.Run("internal server error", func(t *testing.T) {
		mockUsecase.On("GetUsers", mock.Anything, "param_value").Return(nil, errors.New("some error"))
		req, _ := http.NewRequest(http.MethodGet, "/users?param=param_value", nil)
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)
		assert.Equal(t, http.StatusInternalServerError, w.Code)
		assert.JSONEq(t, `{"message": "error retrieving users"}`, w.Body.String())
	})

	t.Run("success", func(t *testing.T) {
		mockUsers := []entities.User{
			{ID: primitive.NewObjectID(), UserName: "User 1"},
			{ID: primitive.NewObjectID(), UserName: "User 2"},
		}
		mockUsecase.On("GetUsers", mock.Anything, "param_value").Return(mockUsers, nil)
		req, _ := http.NewRequest(http.MethodGet, "/users?param=param_value", nil)
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)
		assert.Equal(t, http.StatusOK, w.Code)
		expectedResponse, _ := json.Marshal(gin.H{"users": mockUsers})
		assert.JSONEq(t, string(expectedResponse), w.Body.String())
	})
}

func TestGetUserByID(t *testing.T) {

	mockUsecase := new(mocks.UserUsecase)
	mockEnvironment := new(mocks.Environment)

	router := gin.Default()
	uc := controller.NewUserController(mockEnvironment, mockUsecase)
	router.GET("/users/:id", uc.GetUserByID)

	t.Run("internal server error", func(t *testing.T) {
		mockUsecase.On("GetUserByID", mock.Anything, "id_value").Return(nil, errors.New("some error"))
		req, _ := http.NewRequest(http.MethodGet, "/users/id_value", nil)
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)
		assert.Equal(t, http.StatusInternalServerError, w.Code)
		assert.JSONEq(t, `{"message": "error retrieving user"}`, w.Body.String())
	})

	t.Run("success", func(t *testing.T) {
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

	mockUsecase := new(mocks.UserUsecase)
	mockEnvironment := new(mocks.Environment)

	router := gin.Default()
	uc := controller.NewUserController(mockEnvironment, mockUsecase)
	router.PUT("/users/:id", uc.UpdateUser)

	t.Run("bad request", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodPut, "/users/id_value", nil)
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)
		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.JSONEq(t, `{"message": "Bad Request"}`, w.Body.String())
	})

	t.Run("not found", func(t *testing.T) {
		mockUser := entities.User{ID: primitive.NewObjectID(), UserName: "User 1"}
		mockUsecase.On("UpdateUser", mock.Anything, "id_value", mockUser).Return(errors.New("some error"))
		req, _ := http.NewRequest(http.MethodPut, "/users/id_value", nil)
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)
		assert.Equal(t, http.StatusNotFound, w.Code)
		assert.JSONEq(t, `{"message": "Not Found"}`, w.Body.String())
	})
}

func TestDeleteUser(t *testing.T) {

	mockUsecase := new(mocks.UserUsecase)
	mockEnvironment := new(mocks.Environment)

	router := gin.Default()
	uc := controller.NewUserController(mockEnvironment, mockUsecase)
	router.DELETE("/users/:id", uc.DeleteUser)

	t.Run("internal server error", func(t *testing.T) {
		mockUsecase.On("DeleteUser", mock.Anything, "id_value").Return(errors.New("some error"))
		req, _ := http.NewRequest(http.MethodDelete, "/users/id_value", nil)
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)
		assert.Equal(t, http.StatusInternalServerError, w.Code)
		assert.JSONEq(t, `{"message": "error deleting user"}`, w.Body.String())
	})

	t.Run("success", func(t *testing.T) {
		mockUsecase.On("DeleteUser", mock.Anything, "id_value").Return(nil)
		req, _ := http.NewRequest(http.MethodDelete, "/users/id_value", nil)
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)
		assert.Equal(t, http.StatusOK, w.Code)
	})
}

func TestCreateUser(t *testing.T) {
	mockUsecase := new(mocks.UserUsecase)
	mockEnvironemnt := new(mocks.Environment)
	router := gin.Default()
	uc := controller.NewUserController(mockEnvironemnt, mockUsecase)
	router.POST("/users", uc.CreateUser)

	t.Run("bad request", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodPost, "/users", bytes.NewBuffer([]byte(`invalid json`)))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)
		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.JSONEq(t, `{"message": "Bad Request"}`, w.Body.String())
	})

	t.Run("internal server error", func(t *testing.T) {
		newUser := model.UserCreate{Name: "Test User"}
		mockUsecase.On("CreateUser", mock.Anything, newUser).Return(errors.New("some error"))
		userJSON, _ := json.Marshal(newUser)
		req, _ := http.NewRequest(http.MethodPost, "/users", bytes.NewBuffer(userJSON))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)
		assert.Equal(t, http.StatusInternalServerError, w.Code)
		assert.JSONEq(t, `{"message": "Internal Server Error"}`, w.Body.String())
	})

	t.Run("success", func(t *testing.T) {
		newUser := model.UserCreate{Name: "Test User"}
		mockUsecase.On("CreateUser", mock.Anything, newUser).Return(nil)
		userJSON, _ := json.Marshal(newUser)
		req, _ := http.NewRequest(http.MethodPost, "/users", bytes.NewBuffer(userJSON))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)
		assert.Equal(t, http.StatusCreated, w.Code)
		expectedResponse, _ := json.Marshal(gin.H{"message": "User created successfully", "user": newUser})
		assert.JSONEq(t, string(expectedResponse), w.Body.String())
	})
}

func TestNewUserController(t *testing.T) {
	mockUsecase := new(mocks.UserUsecase)
	mockEnvironment := new(mocks.Environment)

	uc := controller.NewUserController(mockEnvironment, mockUsecase)

	assert.NotNil(t, uc)
}


