package main

import (
	"log"
	"task-management-api/config"
	"task-management-api/router"
	"time"

	"github.com/gin-gonic/gin"
)

func main()  {
	
	env, db, err := config.Initialize()
	if err !=  nil{
		log.Fatal(err)
	}

	route := gin.Default()
	
	router.NewRouter(env, time.Second * 5, db, route)

	log.Println("")
	log.Println("-----------------------------------")
	log.Println("Server is running on port: " + env.GetPort())
	log.Println("Press CTRL + C to stop the server")
	log.Println("-----------------------------------")
	log.Println("")

	route.Run(":" + env.GetPort())

}