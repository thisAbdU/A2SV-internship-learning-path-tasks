package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Environment interface {
	GetJwtKey() string
	GetDbURL() string
	GetDbName() string
	GetPort() string
}

type environment struct {
	jwtKey string
	dbURL  string
	dbName string
	port   string
}

func (e *environment) GetJwtKey() string {
	return e.jwtKey
}

func (e *environment) GetDbURL() string {
	return e.dbURL
}

func (e *environment) GetDbName() string {
	return e.dbName
}

func (e *environment) GetPort() string {
	return e.port
}

func NewEnvironment() (Environment, error) {
		log.Println("Loading .env file")
		err := godotenv.Load()
		if err != nil {
			log.Println(err)
			return nil, fmt.Errorf("error loading .env file: %w", err)
		}

	return &environment{
		dbURL:  os.Getenv("DbURL"),
		dbName: os.Getenv("DbName"),
		port:   os.Getenv("Port"),
		jwtKey: os.Getenv("jwtKey"),
	}, nil
}

func GetJwtKey() []byte {
	env, err := NewEnvironment()
	if err != nil {
		log.Println(err)
		log.Fatal(err)
	}

	return []byte(env.GetJwtKey())
}
