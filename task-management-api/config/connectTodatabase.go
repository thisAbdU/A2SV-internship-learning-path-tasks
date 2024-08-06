package config

import (
	"context"
	"log"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)
var (
    mongoClient *mongo.Client
    once        sync.Once
)

func GetMongoClient(env *Environment) (*mongo.Database, error) {
    var err error

    once.Do(func() {
        clientOptions := options.Client().ApplyURI(env.DbURL)
        ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
        defer cancel()

        client, err := mongo.Connect(ctx, clientOptions)
        if err != nil {
            log.Fatal(err)
        }

        err = client.Ping(ctx, nil)
        if err != nil {
            log.Println("Failed to connect to MongoDB")
            log.Fatal(err)
        }

        log.Println("Connected to MongoDB")
        mongoClient = client
    })

    return mongoClient.Database(env.DbName), err
}
