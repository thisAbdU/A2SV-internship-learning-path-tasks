package repository_test

import (
	"context"
	"task-management-api/domain/entities"
	"task-management-api/mongo/mocks"
	"task-management-api/repository"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson"
)


func TestGetTasks(t *testing.T) {
    mockCursor := new(mocks.Cursor)
    mockCollection := new(mocks.Collection)
    mockDatabase := new(mocks.Database)

    tr := repository.NewTaskRepository(mockDatabase, "tasks")

    ctx := context.TODO()
    userID := "test-user-id"

    task1 := entities.Task{Title: "Task 1", Description: "Description 1"}
    task2 := entities.Task{Title: "Task 2", Description: "Description 2"}

    expectedFilter := bson.M{"userid": userID}

    mockDatabase.On("Collection", "tasks").Return(mockCollection)
    mockCollection.On("Find", ctx, expectedFilter).Return(mockCursor, nil)

    mockCursor.On("Next", ctx).Return(true).Once()
    mockCursor.On("Decode", mock.Anything).Run(func(args mock.Arguments) {
        arg := args.Get(0).(*entities.Task)
        *arg = task1 
    }).Return(nil).Once()

    mockCursor.On("Next", ctx).Return(true).Once() 
    mockCursor.On("Decode", mock.Anything).Run(func(args mock.Arguments) {
        arg := args.Get(0).(*entities.Task)
        *arg = task2 
    }).Return(nil).Once()

    mockCursor.On("Next", ctx).Return(false).Once() 
    mockCursor.On("Close", ctx).Return(nil) 

    result, err := tr.GetTasks(ctx, userID)

    assert.NoError(t, err)
    assert.Len(t, result, 2)
    assert.Equal(t, "Task 1", result[0].Title)
    assert.Equal(t, "Description 1", result[0].Description)
    assert.Equal(t, "Task 2", result[1].Title)
    assert.Equal(t, "Description 2", result[1].Description)

    mockCollection.AssertExpectations(t)
    mockCursor.AssertExpectations(t)
}
