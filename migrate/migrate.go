package main

import (
	"github.com/shihabahmed/go-learn/initializers"
	"github.com/shihabahmed/go-learn/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.ToDo{})
}