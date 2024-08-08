package repository_test

import (
	"context"
	"log"
	"task-management-api/domain/entities"
	"task-management-api/mongo/mocks"
	"task-management-api/repository"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
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

func TestGetTaskByID(t *testing.T) {
    mockCollection := new(mocks.Collection)
    mockDatabase := new(mocks.Database)
    
    tr := repository.NewTaskRepository(mockDatabase, "tasks")
    
    ctx := context.TODO()
    taskID := "60c72b2f9b1d4c3d88b8e5e6"
    userID := "12345"
    
    objectID, _ := primitive.ObjectIDFromHex(taskID)
    
    task := entities.Task{Title: "Task 1", Description: "Description 1"}
    
    expectedFilter := bson.M{"$and": []bson.M{{"_id": objectID}, {"userid": userID}}}
    
    mockDatabase.On("Collection", "tasks").Return(mockCollection)
    
    mockSingleResult := new(mocks.SingleResult)
    mockCollection.On("FindOne", ctx, expectedFilter, mock.Anything).Return(mockSingleResult)
    
    mockSingleResult.On("Decode", mock.AnythingOfType("*entities.Task")).Run(func(args mock.Arguments) {
        arg := args.Get(0).(*entities.Task)
        *arg = task
    }).Return(nil)
    
    result, err := tr.GetTaskByID(ctx, taskID, userID)
    
    assert.NoError(t, err)
    assert.Equal(t, "Task 1", result.Title)
    assert.Equal(t, "Description 1", result.Description)
    
    mockCollection.AssertExpectations(t)
    mockSingleResult.AssertExpectations(t)
}

func TestCreateTask(t *testing.T) {
    mockCollection := new(mocks.Collection)
    mockDatabase := new(mocks.Database)

    tr := repository.NewTaskRepository(mockDatabase, "tasks")

    ctx := context.TODO()
    userID := "12345"

    task := entities.Task{UserID: userID, Title: "Task 1", Description: "Description 1"}

    mockDatabase.On("Collection", "tasks").Return(mockCollection)

    mockCollection.On("InsertOne", ctx, mock.Anything).Return(nil, nil)
    

    err := tr.CreateTask(ctx, task)

    assert.NoError(t, err)

    mockCollection.AssertExpectations(t)
}

func TestUpdateTask(t *testing.T) {
    mockCollection := new(mocks.Collection)
    mockDatabase := new(mocks.Database)

    tr := repository.NewTaskRepository(mockDatabase, "tasks")

    ctx := context.TODO()
    taskID := "60c72b2f9b1d4c3d88b8e5e6"
    userID := "12345"

    task := entities.Task{Title: "Task 1", Status:"Done", Description: "Description 1"}

    objectID, _ := primitive.ObjectIDFromHex(taskID)

    expectedFilter := bson.M{"_id": objectID, "userid": userID}

    expectedUpdate := bson.M{"$set": bson.M{"title": task.Title, "status": task.Status, "description": task.Description}}

    mockDatabase.On("Collection", "tasks").Return(mockCollection)
    log.Println("Mock collection", mockCollection)

    mockUpdateResult := &mongo.UpdateResult{
        MatchedCount: 1,
        ModifiedCount: 1,
    }

    mockCollection.On("UpdateOne", ctx, expectedFilter, expectedUpdate).Return(mockUpdateResult, nil).Once()

    err := tr.UpdateTask(ctx, taskID, task, userID)

    assert.NoError(t, err)

    mockCollection.AssertExpectations(t)
}

func TestDeleteTask(t *testing.T){
    mockCollection := new(mocks.Collection)
    mockDatabase := new(mocks.Database)

    tr := repository.NewTaskRepository(mockDatabase, "tasks")

    ctx := context.TODO()
    taskID := "60c72b2f9b1d4c3d88b8e5e6"
    userID := "12345"

    objectID, _ := primitive.ObjectIDFromHex(taskID)

    expectedFilter := primitive.M{"$and": []primitive.M{
        {"_id": objectID},
        {"userid": userID},
    }}

    mockDatabase.On("Collection", "tasks").Return(mockCollection)

    mockDeleteCount := int64(1)
    mockCollection.On("DeleteMany", ctx, expectedFilter).Return(mockDeleteCount, nil)

    err := tr.DeleteTask(ctx, taskID, userID)

    assert.NoError(t, err)

    mockCollection.AssertExpectations(t)
}