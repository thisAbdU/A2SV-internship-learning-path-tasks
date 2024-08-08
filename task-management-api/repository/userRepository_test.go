package repository_test

import (
	"context"
	"task-management-api/domain/entities"
	"task-management-api/domain/model"
	"task-management-api/mongo/mocks"
	"task-management-api/repository"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func TestGetUser(t *testing.T) {
    mockCursor := new(mocks.Cursor)
    mockCollection := new(mocks.Collection)
    mockDatabase := new(mocks.Database)

    ur := repository.NewUserRepository(mockDatabase, "user")

    ctx := context.TODO()

	usersArray := []entities.User{
		{ID: primitive.NewObjectID(), UserName: "test_1", Password: "1234"},
		{ID: primitive.NewObjectID(), UserName: "test_2", Password: "1234"},
	}

	decodedUsers := make([]*entities.User, 0)
    param := "test"

    expectedFilter := bson.M{
        "$or": []bson.M{
            {"username": primitive.Regex{Pattern: param, Options: "i"}},
            {"email": primitive.Regex{Pattern: param, Options: "i"}},
        },
    }

    mockDatabase.On("Collection", "user").Return(mockCollection)
    mockCollection.On("Find", ctx, expectedFilter).Return(mockCursor, nil).Once()

    mockCursor.On("Next", ctx).Return(true).Times(len(usersArray))

	for i := 0; i < 2; i++ {
		mockCursor.On("Decode", mock.Anything).Run(func(args mock.Arguments) {
			user := args.Get(0).(*entities.User)
			*user = usersArray[i]  
		}).Return(nil).Once()

		decodedUsers = append(decodedUsers, &usersArray[i])
	}
	
    mockCursor.On("Next", ctx).Return(false).Once() 

	mockCursor.On("Close", ctx).Return(nil).Once()

    users, err := ur.GetUser(ctx, param)
    if err != nil {
        t.Fatalf("unexpected error: %v", err)
    }

    if len(decodedUsers) != len(usersArray) {
        t.Fatalf("expected %d users, got %d", len(decodedUsers), len(usersArray))
    }

    if decodedUsers[0].UserName != "test_1" && decodedUsers[1].UserName != "test_2" {
        t.Fatalf("unexpected users: %+v", users)
    }

    mockCursor.AssertExpectations(t)
    mockCollection.AssertExpectations(t)
    mockDatabase.AssertExpectations(t)
}

func TestGetUserByID(t *testing.T)  {
    mockCollection := new(mocks.Collection)
    mockDatabase := new(mocks.Database)

    ur := repository.NewUserRepository(mockDatabase, "user")

    ctx := context.TODO()

	userID := primitive.NewObjectID()

	user := entities.User{ID: userID, UserName: "test_1", Password: "1234"}

	filter := bson.M{
		"_id": userID,
	}

	mockDatabase.On("Collection", "user").Return(mockCollection)
	mockSingleResult := new(mocks.SingleResult)
    mockCollection.On("FindOne", ctx, filter).Return(mockSingleResult, nil).Once()

	mockSingleResult.On("Decode", mock.AnythingOfType("*entities.User")).Run(func(args mock.Arguments) {
        arg := args.Get(0).(*entities.User)
        *arg = user
    }).Return(nil)

	result, err := ur.GetUserByID(ctx, userID.Hex())

	assert.NoError(t, err)
    assert.Equal(t, "test_1", result.UserName)
    assert.Equal(t, "1234", result.Password)
    
    mockCollection.AssertExpectations(t)
    mockSingleResult.AssertExpectations(t)
}

func TestUpdateUser(t *testing.T)  {
	mockCollection := new(mocks.Collection)
    mockDatabase := new(mocks.Database)

	ur := repository.NewUserRepository(mockDatabase, "user")

    ctx := context.TODO()

	userID := "60c72b2f9b1e8f3b5b7c8f8d"

	objectID, _ := primitive.ObjectIDFromHex(userID)
	updatedUser := entities.User{
		UserName: "newUsername",
		Password: "1234",
	}

	filter := bson.M{
		"_id": objectID,
	}

	update := bson.M{
		"$set": updatedUser,
	}

	mockDatabase.On("Collection", "user").Return(mockCollection)

	mockUpdateResult := &mongo.UpdateResult{
        MatchedCount: 1,
        ModifiedCount: 1,
    }

    mockCollection.On("UpdateOne", ctx, filter, update).Return(mockUpdateResult, nil).Once()

    err := ur.UpdateUser(ctx, userID, updatedUser)

    assert.NoError(t, err)

    mockCollection.AssertExpectations(t)
}

func TestDeleteUser(t *testing.T){
	mockCollection := new(mocks.Collection)
    mockDatabase := new(mocks.Database)

	ur := repository.NewUserRepository(mockDatabase, "user")

    ctx := context.TODO()

	userID := "60c72b2f9b1e8f3b5b7c8f8d"

	objectID, _ := primitive.ObjectIDFromHex(userID)

	filter := bson.M{
		"_id": objectID,
	}

	mockDatabase.On("Collection", "user").Return(mockCollection)
	
	mockDeleteCount := int64(1)

    mockCollection.On("DeleteOne", ctx, filter).Return(mockDeleteCount, nil).Once()

    err := ur.DeleteUser(ctx, userID)

    assert.NoError(t, err)

    mockCollection.AssertExpectations(t)
}

func TestCreateUser(t *testing.T){
	mockCollection := new(mocks.Collection)
    mockDatabase := new(mocks.Database)

    ur := repository.NewUserRepository(mockDatabase, "user")

    ctx := context.TODO()
    userID := primitive.NewObjectID()

    dummyUser := model.UserCreate{
		Username: "john_doe",
		Password: "securePassword123",
		Email:    "john.doe@example.com",
		Name:     "John Doe",
		Bio:      "A passionate software developer.",
		ID:       userID,
	}

    mockDatabase.On("Collection", "user").Return(mockCollection)

    mockCollection.On("InsertOne", ctx, mock.Anything).Return(nil, nil)
    

    _, err := ur.CreateUser(ctx, dummyUser)

    assert.NoError(t, err)

    mockCollection.AssertExpectations(t)
}