package config_test

import (
	"task-management-api/config"
	"task-management-api/mongo"
	"task-management-api/mongo/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func MockNewClient(connection string) (Client, error) {
    mockClient := &mongoClient{}
    return mockClient, nil
}

func TestGetMongoClient(t *testing.T) {
    mockClient := new(mocks.Client)
    mockDatabase := new(mocks.Database)

    env := config.Environment{
        DbURL:  "mongodb://localhost:27017",
        DbName: "testdb",
    }

    mockClient.On("Connect", mock.Anything).Return(nil)
    mockClient.On("Ping", mock.Anything).Return(nil)
    mockClient.On("Database", env.DbName).Return(mockDatabase)

   // Mock the NewClient function
   originalNewClient := NewClient
   NewClient = MockNewClient
   defer func() { NewClient = originalNewClient }()

    db, err := config.GetMongoClient(&env)
    assert.NoError(t, err, "Expected no error while getting MongoDB client")
    assert.NotNil(t, db, "Expected non-nil MongoDB database instance")

    mockClient.AssertExpectations(t)

    db2, err := config.GetMongoClient(&env)
    assert.NoError(t, err, "Expected no error while getting MongoDB client")
    assert.Equal(t, db, db2, "Expected the same MongoDB client instance")
}