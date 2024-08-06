package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Environment struct {
	JwtKey string
	DbURL string
	DbName string
	Port string
}

func NewEnvironment()(*Environment, error) {
	err := godotenv.Load()
	if err != nil {
		log.Println(err)
		log.Fatal("Error loading .env file")
	}
	
	return &Environment{
		DbURL: os.Getenv("DbURL"),
		DbName: os.Getenv("DbName"),
		Port: os.Getenv("Port"),
		JwtKey: os.Getenv("jwtKey"),
	}, err
}

