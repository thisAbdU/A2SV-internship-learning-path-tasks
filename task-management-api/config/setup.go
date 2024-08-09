package config

import (
	"log"
	"task-management-api/mongo"
)

func Initialize() (*Environment, mongo.Database, error){
	env, err := NewEnvironment()
	if err != nil {
		log.Println(err)
		log.Fatal(err)
	}

	db, err := GetMongoClient(env)
	if err != nil {
		log.Fatal(err)
	}

    return env, db ,nil
}
