package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/shihabahmed/go-learn/internal/constants"
	"github.com/shihabahmed/go-learn/internal/controllers"
)

func Init() *gin.Engine {
	godotenv.Load()
	server := gin.Default()
	initRoutes(server)
	return server
}

func initRoutes(server *gin.Engine) {
	server.GET("/", func (ctx *gin.Context) {
		ctx.String(http.StatusOK, "Service is alive...")
	})

	server.POST(constants.ToDoRoute, controllers.CreateToDo)
	server.PUT(constants.ToDoRouteWithParam, controllers.UpdateToDo)
	server.GET(constants.ToDoRoute, controllers.GetToDos)
	server.GET(constants.ToDoRouteWithParam, controllers.GetToDoByID)
	server.PATCH(constants.ToDoRouteWithParam, controllers.UpdateToDoStatus)
	server.DELETE(constants.ToDoRouteWithParam, controllers.DeleteToDo)
}