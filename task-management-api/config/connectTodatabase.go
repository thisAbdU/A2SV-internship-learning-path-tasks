package config

import (
	"context"
	"log"
	"sync"
	"time"

	"task-management-api/mongo"
)
var (
    mongoClient mongo.Client
    once        sync.Once
)

func GetMongoClient(env *Environment) (mongo.Database, error) {
    var err error

    once.Do(func() {
        ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
        defer cancel()

        client, err := mongo.NewClient(env.DbURL)
        if err != nil {
            log.Fatal(err)
        }

        err = client.Connect(ctx)
        if err != nil {
            log.Fatal(err)
        }

        err = client.Ping(ctx)
        if err != nil {
            log.Println("Failed to connect to MongoDB")
            log.Fatal(err)
        }

        log.Println("Connected to MongoDB")

        mongoClient = client
    })

    return mongoClient.Database(env.DbName), err
}
