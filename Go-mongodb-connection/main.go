package main

import (
	"context"
	"fmt"
	"log"

	// "go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Trainer struct {
    Name string
    Age  int
    Place string
}

func main() {
    clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

    client, err := mongo.Connect(context.TODO(), clientOptions)

    if err != nil{
        log.Fatal(err)
    }

    err = client.Ping(context.TODO(), nil)

    if err != nil{
        log.Fatal(err)
    }

    fmt.Println("connected to mongo")

    collection := client.Database("test").Collection("trainers")

    // emre := Trainer{"Emre", 35, "Abrehot"}
    // semir := Trainer{"Semir", 23, "AAiT"}
    // aman := Trainer{"Aman", 22, "ASTU"}
    // simon := Trainer{"Simon", 23, "Abrehot"}

    // insertResult, err := collection.InsertOne(context.TODO(), emre)

    // if err != nil{
    //     log.Fatal(err)
    // }

    // trainers := []interface{}{semir, aman, simon}


    // insertManyResult, err := collection.InsertMany(context.TODO(), trainers)

    // if err != nil{
    //     log.Fatal(err)
    // }

    filter := bson.D{{Key: "name", Value: "Emre"}}

    update := bson.D{
        {Key: "$inc", Value: bson.D{
         {Key: "age", Value: 1},
        }},
     }
    updateResult, err := collection.UpdateOne(context.TODO(), filter, update)

    if err != nil{
        log.Fatal(err)
    }

    fmt.Println("updated documents", updateResult.UpsertedID)
}
