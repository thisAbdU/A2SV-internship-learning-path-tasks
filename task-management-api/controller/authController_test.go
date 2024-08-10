package controller_test

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"task-management-api/controller"
	"task-management-api/domain/mocks"
	"task-management-api/domain/model"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestRegister(t *testing.T) {
	mockUsecase := new(mocks.AuthUseCase)
	router := gin.Default()

	ac := controller.NewAuthController(mockUsecase)

	router.POST("/register", ac.Register)

	t.Run("success", func(t *testing.T) {
		newUser := &model.UserCreate{
			Username: "newuser",
			Password: "password",
		}
		
		mockUsecase.On("Register", newUser).Return(nil).Once()

		body := `{"username":"newuser", "password":"password"}`
		req, _ := http.NewRequest("POST", "/register", strings.NewReader(body))
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusCreated, w.Code)
		assert.JSONEq(t, `{"message": "User created successfully"}`, w.Body.String())

		mockUsecase.AssertExpectations(t)
	})

	t.Run("bad request", func(t *testing.T) {
		body := `{"username": "newuser", "password": 1234}`
		req, _ := http.NewRequest("POST", "/register", strings.NewReader(body))
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.JSONEq(t, `{"message": "Bad Request"}`, w.Body.String())
	})

	t.Run("username already exists", func(t *testing.T) {
		newUser := &model.UserCreate{
			Username: "existinguser",
			Password: "password",
		}
		
		mockUsecase.On("Register", newUser).Return(errors.New("username already exists")).Once()

		body := `{"username":"existinguser", "password":"password"}`
		req, _ := http.NewRequest("POST", "/register", strings.NewReader(body))
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.JSONEq(t, `{"message": "username already exists"}`, w.Body.String())

		mockUsecase.AssertExpectations(t)
	})

	t.Run("internal server error", func(t *testing.T) {
		newUser := &model.UserCreate{
			Username: "erroruser",
			Password: "password",
		}
		
		mockUsecase.On("Register", newUser).Return(errors.New("unexpected error")).Once()

		body := `{"username":"erroruser", "password":"password"}`
		req, _ := http.NewRequest("POST", "/register", strings.NewReader(body))
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
		assert.JSONEq(t, `{"message": "Internal Server Error"}`, w.Body.String())

		mockUsecase.AssertExpectations(t)
	})
}

func TestLogin(t *testing.T) {
	mockUsecase := new(mocks.AuthUseCase)
	router := gin.Default()

	ac := controller.NewAuthController(mockUsecase)

	router.POST("/login", ac.Login)

	t.Run("success", func(t *testing.T) {
		userLogin := &model.UserLogin{
			Username: "newuser",
			Password: "password",
		}
		
		mockUsecase.On("Login", userLogin).Return("token", nil).Once()

		body := `{"username":"newuser", "password":"password"}`
		req, _ := http.NewRequest("POST", "/login", strings.NewReader(body))
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
		assert.JSONEq(t, `{"token": "token"}`, w.Body.String())

		mockUsecase.AssertExpectations(t)
	})

	t.Run("bad request", func(t *testing.T) {
		body := `{"username": "newuser", "password": 1234}`
		req, _ := http.NewRequest("POST", "/login", strings.NewReader(body))
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.JSONEq(t, `{"message": "Bad Request"}`, w.Body.String())
	})

	t.Run("unauthorized", func(t *testing.T) {
		userLogin := &model.UserLogin{
			Username: "existinguser",
			Password: "password",
		}
		
		mockUsecase.On("Login", userLogin).Return("", errors.New("unauthorized")).Once()

		body := `{"username":"existinguser", "password":"password"}`
		req, _ := http.NewRequest("POST", "/login", strings.NewReader(body))
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusUnauthorized, w.Code)
		assert.JSONEq(t, `{"message": "Unauthorized"}`, w.Body.String())

		mockUsecase.AssertExpectations(t)
	})

	t.Run("internal server error", func(t *testing.T) {
		userLogin := &model.UserLogin{
			Username: "erroruser",
			Password: "password",
		}
		
		mockUsecase.On("Login", userLogin).Return("", errors.New("unexpected error")).Once()
		body := `{"username":"erroruser", "password":"password"}`
		req, _ := http.NewRequest("POST", "/login", strings.NewReader(body))
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
		assert.JSONEq(t, `{"message": "Internal Server Error"}`, w.Body.String())

		mockUsecase.AssertExpectations(t)
	})
}

func TestNewAuthController(t *testing.T) {
	mockUsecase := new(mocks.AuthUseCase)
	ac := controller.NewAuthController(mockUsecase)

	assert.NotNil(t, ac)
}
