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

func getMongoClient(uri string, db string) (*mongo.Database, error) {
    once.Do(func() {
        clientOpitons := options.Client().ApplyURI(uri)
        c, cancel := context.WithTimeout(context.Background(), 10*time.Second)
        defer cancel()

        client, err := mongo.Connect(c, clientOpitons)
        if err != nil {
            log.Fatal(err)
        }

        err = client.Ping(c, nil)
        if err != nil {
            log.Println("Failed to connect to MongoDB")
            log.Fatal(err)
        }

        log.Println("Connected to MongoDB")

    })

    return mongoClient.Database(db), nil
}

// func ConnectToDatabase() *mongo.Database {
//     uri := "mongodb+srv://taskmanager:task123@cluster0.jkxryyl.mongodb.net/"

//     clientOptions := options.Client().ApplyURI(uri)
//     client, err := mongo.Connect(context.TODO(), clientOptions)
//     if err != nil {
//         log.Fatal(err)
//     }

//     err = client.Ping(context.TODO(), nil)
//     if err != nil {
//         log.Fatal(err)
//     }

//     log.Println("Connected to MongoDB")

//     db = client.Database("taskManagementDatabase")

//     return db
// }