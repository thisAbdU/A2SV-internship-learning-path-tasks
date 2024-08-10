package config

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"

	"task-management-api/mongo"
)
var (
    mongoClient mongo.Client
    once        sync.Once
)

func GetMongoClient(env Environment) (mongo.Database, error) {
	once.Do(func() {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		clientOptions, err := mongo.NewClient(env.GetDbURL())
		err = clientOptions.Connect(ctx)
		if err != nil {
			log.Fatal(err)
		}

		err = clientOptions.Ping(ctx)
		if err != nil {
			log.Println("Failed to connect to MongoDB")
			log.Fatal(err)
		}

		log.Println("Connected to MongoDB")
		mongoClient = clientOptions
	})

	if mongoClient == nil {
		return nil, fmt.Errorf("failed to create MongoDB client")
	}

	return mongoClient.Database(env.GetDbName()), nil
}