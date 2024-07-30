package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/shihabahmed/go-learn/controllers"
	"github.com/shihabahmed/go-learn/initializers"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func initRoutes(server *gin.Engine) {
	server.POST("/", controllers.CreateToDo)
	server.PUT("/:id", controllers.UpdateToDo)
	server.GET("/", controllers.GetToDos)
	server.PATCH("/:id", controllers.UpdateToDoStatus)
	server.DELETE("/:id", controllers.DeleteToDo)
}

func main()  {
	godotenv.Load()
	server := gin.Default()
	initRoutes(server)
	server.Run()
}