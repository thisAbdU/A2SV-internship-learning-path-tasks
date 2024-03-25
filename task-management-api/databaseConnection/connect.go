package main

import (
    "context"
    "log"

    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

var db *mongo.Database

func ConnectToDatabase() *mongo.Database {
    uri := "mongodb+srv://taskmanager:task123@cluster0.jkxryyl.mongodb.net/"

    clientOptions := options.Client().ApplyURI(uri)
    client, err := mongo.Connect(context.TODO(), clientOptions)
    if err != nil {
        log.Fatal(err)
    }

    err = client.Ping(context.TODO(), nil)
    if err != nil {
        log.Fatal(err)
    }

    log.Println("Connected to MongoDB")

    db = client.Database("taskManagementDatabase")

    return db
}