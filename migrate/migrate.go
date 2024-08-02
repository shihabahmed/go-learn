package main

import (
	"github.com/shihabahmed/go-learn/internal/initializers"
	"github.com/shihabahmed/go-learn/internal/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.ToDo{})
}