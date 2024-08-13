package controller_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"task-management-api/controller"
	"task-management-api/domain/entities"
	"task-management-api/domain/mocks"
	"task-management-api/domain/model"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestCreateTask(t *testing.T) {

	t.Run("success", func(t *testing.T) {
		mockUsecase := new(mocks.TaskUsecase)
		mockEnvironment := new(mocks.Environment)
	
		router := gin.Default()
	
		tc := controller.NewTaskController(mockEnvironment, mockUsecase)
	
		router.POST("/tasks", tc.CreateTask)

		mockUsecase.On("CreateTask", mock.AnythingOfType("entities.Task")).Return(nil).Once()

		newTask := entities.Task{
			Title:       "Test Task",
			Description: "This is a test task",
		}

		taskJSON, _ := json.Marshal(newTask)

		req, _ := http.NewRequest(http.MethodPost, "/tasks", bytes.NewBuffer(taskJSON))
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Request = req
		c.Set("user_id", "test_user_id")

		tc.CreateTask(c)

		assert.Equal(t, http.StatusCreated, w.Code)
		assert.JSONEq(t, `{"message": "Task created successfully"}`, w.Body.String())

		mockUsecase.AssertCalled(t, "CreateTask", mock.MatchedBy(func(task entities.Task) bool {
			return task.Title == newTask.Title && task.Description == newTask.Description && task.UserID == "test_user_id"
		}))
	})
	t.Run("error", func(t *testing.T) {
		t.Run("unauthorized", func(t *testing.T) {
			mockUsecase := new(mocks.TaskUsecase)
			mockEnvironment := new(mocks.Environment)
		
			router := gin.Default()
		
			tc := controller.NewTaskController(mockEnvironment, mockUsecase)
		
			router.POST("/tasks", tc.CreateTask)

			newTask := entities.Task{
				Title:       "Test Task",
				Description: "This is a test task",
			}
			
			mockUsecase.On("CreateTask", mock.AnythingOfType("entities.Task")).Return(errors.New("Please sign up to create a task")).Once()

			taskJSON, _ := json.Marshal(newTask)

			req, _ := http.NewRequest(http.MethodPost, "/tasks", bytes.NewBuffer(taskJSON))
			req.Header.Set("Content-Type", "application/json")
			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)
			c.Request = req

			tc.CreateTask(c)

			assert.Equal(t, http.StatusUnauthorized, w.Code)
			assert.JSONEq(t, `{"message": "Please sign up to create a task"}`, w.Body.String())
		})

		t.Run("bad_request", func(t *testing.T) {
			mockUsecase := new(mocks.TaskUsecase)
			mockEnvironment := new(mocks.Environment)
		
			router := gin.Default()
		
			tc := controller.NewTaskController(mockEnvironment, mockUsecase)
		
			router.POST("/tasks", tc.CreateTask)

			mockUsecase.On("CreateTask", mock.AnythingOfType("entities.Task")).Return(errors.New("Bad Request")).Once()

			invalidTaskJSON := `{"ttle": ""}`

			req, _ := http.NewRequest(http.MethodPost, "/tasks", strings.NewReader(invalidTaskJSON))
			req.Header.Set("Content-Type", "application/json")

			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)
			c.Request = req
			c.Set("user_id", "test_user_id")

			tc.CreateTask(c)

			assert.Equal(t, http.StatusBadRequest, w.Code)
			assert.JSONEq(t, `{"message": "Bad Request"}`, w.Body.String())
		})

		t.Run("internal_server_error", func(t *testing.T) {
			mockUsecase := new(mocks.TaskUsecase)
			mockEnvironment := new(mocks.Environment)
		
			router := gin.Default()
		
			tc := controller.NewTaskController(mockEnvironment, mockUsecase)
		
			router.POST("/tasks", tc.CreateTask)

			newTask := entities.Task{
				Title:       "Test Task",
				Description: "This is a test task",
			}

			mockUsecase.On("CreateTask", mock.AnythingOfType("entities.Task")).Return(errors.New("Internal Server Error")).Once()

			taskJSON, _ := json.Marshal(newTask)

			req, _ := http.NewRequest(http.MethodPost, "/tasks", bytes.NewBuffer(taskJSON))
			req.Header.Set("Content-Type", "application/json")

			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)
			c.Request = req
			c.Set("user_id", "test_user_id")

			tc.CreateTask(c)
			
			assert.Equal(t, http.StatusInternalServerError, w.Code)
			assert.JSONEq(t, `{"message": "Internal Server Error"}`, w.Body.String())
		})
	})
}

func TestGetTasks(t *testing.T) {
	t.Run("unauthorized", func(t *testing.T) {
		mockUsecase := new(mocks.TaskUsecase)
		mockEnvironment := new(mocks.Environment)
		
		gin.SetMode(gin.TestMode)
		router := gin.Default()
	
		tc := controller.NewTaskController(mockEnvironment, mockUsecase)
	
		router.GET("/tasks", tc.GetTasks)

		req, _ := http.NewRequest(http.MethodGet, "/tasks", nil)
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)
		assert.Equal(t, http.StatusUnauthorized, w.Code)
		assert.JSONEq(t, `{"error": "Please log in to acess tasks"}`, w.Body.String())
	})

	t.Run("internal server error - GetTasks failure", func(t *testing.T) {
		mockUsecase := new(mocks.TaskUsecase)
		mockEnvironment := new(mocks.Environment)
	
		gin.SetMode(gin.TestMode)
		router := gin.Default()

		testUserID := "123"
	
		tc := controller.NewTaskController(mockEnvironment, mockUsecase)

		mockUsecase.On("GetTasks", testUserID).Return(nil, errors.New("error retrieving tasks"))

		req, _ := http.NewRequest(http.MethodGet, "/tasks", nil)
		w := httptest.NewRecorder()

		c, _ := gin.CreateTestContext(w)
		c.Request = req

		router.Use(func(c *gin.Context) {
			c.Set("user_id", testUserID)
			c.Next()
		})

		router.GET("/tasks", tc.GetTasks)
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
		assert.JSONEq(t, `{"message": "error retrieving tasks"}`, w.Body.String())
	})

	t.Run("success", func(t *testing.T) {
		mockUsecase := new(mocks.TaskUsecase)
		mockEnvironment := new(mocks.Environment)

		gin.SetMode(gin.TestMode)
		router := gin.Default()
		
		router.Use(func(c *gin.Context) {
			c.Set("user_id", "test_user_id")
			c.Next()
		})

		tc := controller.NewTaskController(mockEnvironment, mockUsecase)
	
		router.GET("/tasks", tc.GetTasks)

		mockTasks := []*model.TaskInfo{
			{Title: "Task 1", Description: "Description 1", DueDate: "Today"},
			{Title: "Task 2", Description: "Description 2", DueDate: "Tomorrow"},
		}
		
		mockUsecase.On("GetTasks", "test_user_id").Return(mockTasks, nil)

		req, _ := http.NewRequest(http.MethodGet, "/tasks", nil)
		w := httptest.NewRecorder()

		log.Println("Request: ", req.Body)
		log.Println("response: ", w.Code)
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
		expectedResponse, _ := json.Marshal(gin.H{"tasks": mockTasks})
		log.Println("Response: ", w.Body.String())
		log.Println("Expected Response: ", string(expectedResponse))
		assert.JSONEq(t, string(expectedResponse), w.Body.String())
	})
}

func TestGetTaskByID(t *testing.T) {
    testUserID := "test_user_id"

    t.Run("unauthorized", func(t *testing.T) {
        mockUsecase := new(mocks.TaskUsecase)
        mockEnvironment := new(mocks.Environment)
        tc := controller.NewTaskController(mockEnvironment, mockUsecase)

        router := gin.Default() 
        router.GET("/tasks/:id", tc.GetTaskByID)

        req, _ := http.NewRequest(http.MethodGet, "/tasks/1", nil)
        w := httptest.NewRecorder()
        router.ServeHTTP(w, req)
        assert.Equal(t, http.StatusUnauthorized, w.Code)
        assert.JSONEq(t, `{"error": "Please log in to acess tasks"}`, w.Body.String())
    })

    t.Run("task not found", func(t *testing.T) {
        mockUsecase := new(mocks.TaskUsecase)
        mockEnvironment := new(mocks.Environment)
        tc := controller.NewTaskController(mockEnvironment, mockUsecase)

        router := gin.Default() // Create a new router for each sub-test
        router.Use(func(c *gin.Context) {
            c.Set("user_id", testUserID)
            c.Next()
        })
        router.GET("/tasks/:id", tc.GetTaskByID)

        mockUsecase.On("GetTaskByID", "1", "test_user_id").Return(nil, errors.New("mongo: no documents in result"))

        req, _ := http.NewRequest(http.MethodGet, "/tasks/1", nil)
        w := httptest.NewRecorder()
        router.ServeHTTP(w, req)
        assert.Equal(t, http.StatusNotFound, w.Code)
        assert.JSONEq(t, `{"message": "Task not found"}`, w.Body.String())
    })

    t.Run("internal server error", func(t *testing.T) {
        mockUsecase := new(mocks.TaskUsecase)
        mockEnvironment := new(mocks.Environment)
        tc := controller.NewTaskController(mockEnvironment, mockUsecase)

        router := gin.Default() // Create a new router for each sub-test
        router.Use(func(c *gin.Context) {
            c.Set("user_id", testUserID)
            c.Next()
        })
        router.GET("/tasks/:id", tc.GetTaskByID)

        mockUsecase.On("GetTaskByID", "1", "test_user_id").Return(nil, errors.New("some error"))
        req, _ := http.NewRequest(http.MethodGet, "/tasks/1", nil)
        w := httptest.NewRecorder()
        router.ServeHTTP(w, req)
        assert.Equal(t, http.StatusInternalServerError, w.Code)
        assert.JSONEq(t, `{"message": "error retrieving task"}`, w.Body.String())
    })

    t.Run("success", func(t *testing.T) {
        mockUsecase := new(mocks.TaskUsecase)
        mockEnvironment := new(mocks.Environment)
        tc := controller.NewTaskController(mockEnvironment, mockUsecase)

        router := gin.Default() // Create a new router for each sub-test
        router.Use(func(c *gin.Context) {
            c.Set("user_id", testUserID)
            c.Next()
        })

        router.GET("/tasks/:id", tc.GetTaskByID)

        taskID := primitive.NewObjectID()
        mockTask := &model.TaskInfo{Title: "Task 1", Description: "Description 1"}
        mockUsecase.On("GetTaskByID", taskID.Hex(), "test_user_id").Return(mockTask, nil)

        req, _ := http.NewRequest(http.MethodGet, "/tasks/"+taskID.Hex(), nil)
        w := httptest.NewRecorder()

        router.ServeHTTP(w, req)
        assert.Equal(t, http.StatusOK, w.Code)
        expectedResponse, _ := json.Marshal(gin.H{"task": mockTask})
        assert.JSONEq(t, string(expectedResponse), w.Body.String())
    })
}
func TestUpdateTask(t *testing.T) {

    gin.SetMode(gin.TestMode)

    t.Run("Unauthorized", func(t *testing.T) {
        mockUsecase := new(mocks.TaskUsecase)
        mockEnvironment := new(mocks.Environment)

        router := gin.Default()
        tc := controller.NewTaskController(mockEnvironment, mockUsecase)
        router.PUT("/tasks/:id", tc.UpdateTask)

        req, _ := http.NewRequest(http.MethodPut, "/tasks/1", nil)
        w := httptest.NewRecorder()
        router.ServeHTTP(w, req)

        assert.Equal(t, http.StatusUnauthorized, w.Code)
        assert.JSONEq(t, `{"error": "Please log in to acess tasks"}`, w.Body.String())
    })

    t.Run("Bad Request", func(t *testing.T) {
        mockUsecase := new(mocks.TaskUsecase)
        mockEnvironment := new(mocks.Environment)

        router := gin.Default()
        router.Use(func(c *gin.Context) {
            c.Set("user_id", "test_user_id")
            c.Next()
        })

        tc := controller.NewTaskController(mockEnvironment, mockUsecase)
        router.PUT("/tasks/:id", tc.UpdateTask)

        req, _ := http.NewRequest(http.MethodPut, "/tasks/1", bytes.NewBuffer([]byte("invalid json")))
        req.Header.Set("Content-Type", "application/json")
        w := httptest.NewRecorder()
        router.ServeHTTP(w, req)

        assert.Equal(t, http.StatusBadRequest, w.Code)
        assert.JSONEq(t, `{"message": "Bad Request"}`, w.Body.String())
    })

    t.Run("Task Not Found", func(t *testing.T) {
        mockUsecase := new(mocks.TaskUsecase)
        mockEnvironment := new(mocks.Environment)

        router := gin.Default()
        router.Use(func(c *gin.Context) {
            c.Set("user_id", "test_user_id")
            c.Next()
        })

        tc := controller.NewTaskController(mockEnvironment, mockUsecase)
        router.PUT("/tasks/:id", tc.UpdateTask)

        mockTaskID := primitive.NewObjectID()
        mockTask := entities.Task{
            ID:          mockTaskID,
            UserID:      "test_user_id",
            UserName:    "test_user",
            Password:    "test_password",
            Title:       "Test Task",
            Description: "This is a test task",
            Status:      "Pending",
        }

        taskJSON, _ := json.Marshal(mockTask)
        mockUsecase.On("UpdateTask", "1", mock.AnythingOfType("entities.Task"), "test_user_id").Return(errors.New("no documents updated"))

        req, _ := http.NewRequest(http.MethodPut, "/tasks/1", bytes.NewBuffer(taskJSON))
        req.Header.Set("Content-Type", "application/json")
        w := httptest.NewRecorder()
        router.ServeHTTP(w, req)

        assert.Equal(t, http.StatusNotFound, w.Code)
        assert.JSONEq(t, `{"message": "Task not found"}`, w.Body.String())
    })

    t.Run("Internal Server Error", func(t *testing.T) {
        mockUsecase := new(mocks.TaskUsecase)
        mockEnvironment := new(mocks.Environment)

        router := gin.Default()
        router.Use(func(c *gin.Context) {
            c.Set("user_id", "test_user_id")
            c.Next()
        })

        tc := controller.NewTaskController(mockEnvironment, mockUsecase)
        router.PUT("/tasks/:id", tc.UpdateTask)

        mockTaskID := primitive.NewObjectID()
        mockTask := entities.Task{
            ID:          mockTaskID,
            UserID:      "test_user_id",
            UserName:    "test_user",
            Password:    "test_password",
            Title:       "Test Task",
            Description: "This is a test task",
            Status:      "Pending",
        }

        taskJSON, _ := json.Marshal(mockTask)
        mockUsecase.On("UpdateTask", "1", mock.AnythingOfType("entities.Task"), "test_user_id").Return(errors.New("some internal error"))

        req, _ := http.NewRequest(http.MethodPut, "/tasks/1", bytes.NewBuffer(taskJSON))
        req.Header.Set("Content-Type", "application/json")
        w := httptest.NewRecorder()
        router.ServeHTTP(w, req)

        assert.Equal(t, http.StatusInternalServerError, w.Code)
        assert.JSONEq(t, `{"message": "Internal Server Error"}`, w.Body.String())
    })

    t.Run("Task Updated Successfully", func(t *testing.T) {
        mockUsecase := new(mocks.TaskUsecase)
        mockEnvironment := new(mocks.Environment)

        router := gin.Default()
        router.Use(func(c *gin.Context) {
            c.Set("user_id", "test_user_id")
            c.Next()
        })

        tc := controller.NewTaskController(mockEnvironment, mockUsecase)
        router.PUT("/tasks/:id", tc.UpdateTask)

        mockTaskID := primitive.NewObjectID()
        mockTask := entities.Task{
            ID:          mockTaskID,
            UserID:      "test_user_id",
            UserName:    "test_user",
            Password:    "test_password",
            Title:       "Test Task",
            Description: "This is a test task",
            Status:      "Pending",
        }

        taskJSON, _ := json.Marshal(mockTask)
        mockUsecase.On("UpdateTask", "1", mock.AnythingOfType("entities.Task"), "test_user_id").Return(nil)

        req, _ := http.NewRequest(http.MethodPut, "/tasks/1", bytes.NewBuffer(taskJSON))
        req.Header.Set("Content-Type", "application/json")
        w := httptest.NewRecorder()
        router.ServeHTTP(w, req)

        assert.Equal(t, http.StatusOK, w.Code)
        assert.JSONEq(t, `{"message": "Task updated successfully"}`, w.Body.String())
    })
}

func TestDeleteTask(t *testing.T) {
    gin.SetMode(gin.TestMode)

    t.Run("Unauthorized", func(t *testing.T) {
        mockUsecase := new(mocks.TaskUsecase)
        mockEnvironment := new(mocks.Environment)

        router := gin.Default()
        tc := controller.NewTaskController(mockEnvironment, mockUsecase)
        router.DELETE("/tasks/:id", tc.DeleteTask)

        req, _ := http.NewRequest(http.MethodDelete, "/tasks/1", nil)
        w := httptest.NewRecorder()
        router.ServeHTTP(w, req)

        assert.Equal(t, http.StatusNotFound, w.Code)
        assert.JSONEq(t, `{"error": "Please log in to acess tasks"}`, w.Body.String())
    })

    t.Run("Task Not Found", func(t *testing.T) {
        mockUsecase := new(mocks.TaskUsecase)
        mockEnvironment := new(mocks.Environment)

        router := gin.Default()
        router.Use(func(c *gin.Context) {
            c.Set("user_id", "test_user_id")
            c.Next()
        })

        tc := controller.NewTaskController(mockEnvironment, mockUsecase)
        router.DELETE("/tasks/:id", tc.DeleteTask)

        mockUsecase.On("DeleteTask", "1", "test_user_id").Return(errors.New("no documents deleted"))

        req, _ := http.NewRequest(http.MethodDelete, "/tasks/1", nil)
        w := httptest.NewRecorder()
        router.ServeHTTP(w, req)

        assert.Equal(t, http.StatusNotFound, w.Code)
        assert.JSONEq(t, `{"message": "Task not found"}`, w.Body.String())
    })

    t.Run("Internal Server Error", func(t *testing.T) {
        mockUsecase := new(mocks.TaskUsecase)
        mockEnvironment := new(mocks.Environment)

        router := gin.Default()
        router.Use(func(c *gin.Context) {
            c.Set("user_id", "test_user_id")
            c.Next()
        })

        tc := controller.NewTaskController(mockEnvironment, mockUsecase)
        router.DELETE("/tasks/:id", tc.DeleteTask)

        mockUsecase.On("DeleteTask", "1", "test_user_id").Return(errors.New("some internal error"))

        req, _ := http.NewRequest(http.MethodDelete, "/tasks/1", nil)
        w := httptest.NewRecorder()
        router.ServeHTTP(w, req)

        assert.Equal(t, http.StatusInternalServerError, w.Code)
        assert.JSONEq(t, `{"message": "Internal Server Error"}`, w.Body.String())
    })

    t.Run("Task Deleted Successfully", func(t *testing.T) {
        mockUsecase := new(mocks.TaskUsecase)
        mockEnvironment := new(mocks.Environment)

        router := gin.Default()
        router.Use(func(c *gin.Context) {
            c.Set("user_id", "test_user_id")
            c.Next()
        })

        tc := controller.NewTaskController(mockEnvironment, mockUsecase)
        router.DELETE("/tasks/:id", tc.DeleteTask)

        mockUsecase.On("DeleteTask", "1", "test_user_id").Return(nil)

        req, _ := http.NewRequest(http.MethodDelete, "/tasks/1", nil)
        w := httptest.NewRecorder()
        router.ServeHTTP(w, req)

        assert.Equal(t, http.StatusOK, w.Code)
        assert.JSONEq(t, `{"message": "Task deleted successfully"}`, w.Body.String())
    })
}