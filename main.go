package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/shihabahmed/go-learn/controllers"
)

func main()  {
	godotenv.Load()
	server := gin.Default()
	server.GET("/", controllers.GetTodos)
	server.POST("/", controllers.CreateTodos)
	server.Run()
}