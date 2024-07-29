package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shihabahmed/go-learn/initializers"
	"github.com/shihabahmed/go-learn/models"
	"github.com/shihabahmed/go-learn/types"
)

func GetTodos(ctx *gin.Context) {
	var todos []models.ToDo
	result := initializers.DB.Find(&todos)

	if result.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to retrieve todos",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"todos": todos,
	})
}

func CreateTodos(ctx *gin.Context) {
	var data types.ToDo
	if err := ctx.ShouldBind(&data); err == nil {
		log.Printf("Error: %v", err)
	}
	item := models.ToDo{
		Title: data.Title,
		Note:  data.Note,
	}
	result := initializers.DB.Create(&item)
	if result.Error != nil {
		log.Printf("Error: %v", result.Error)
		ctx.Status(http.StatusBadRequest)
		return
	}
	ctx.Status(http.StatusCreated)
}
