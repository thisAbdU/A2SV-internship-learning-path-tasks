package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Environment struct {
	TokenSecret string
	DbURL string
	DbName string
	JwtSecret   string
	JwtExpiration int
	Port string
}

func NewEnvironment()(*Environment, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	} 

	return &Environment{
		DbURL: os.Getenv("DbURL"),
		DbName: os.Getenv("DbName"),
		Port: os.Getenv("Port"),
	}, err
}

