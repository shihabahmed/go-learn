package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shihabahmed/go-learn/internal/initializers"
	"github.com/shihabahmed/go-learn/internal/models"
	"github.com/shihabahmed/go-learn/internal/types"
)

func GetToDos(ctx *gin.Context) {
	var toDos []models.ToDo
	log.Println(*initializers.DB)
	result := initializers.DB.Find(&toDos)

	if result.Error != nil {
		log.Println(result.Error)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to retrieve ToDos",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"toDos": toDos,
	})
}

func CreateToDo(ctx *gin.Context) {
	var data types.ToDo
	if err := ctx.ShouldBind(&data); err == nil {
		log.Printf("Error: %v", err)
	}
	item := models.ToDo{
		Title: data.Title,
		Note:  data.Note,
		Complete: false,
	}
	result := initializers.DB.Create(&item)
	if result.Error != nil {
		log.Printf("Error: %v", result.Error)
		ctx.Status(http.StatusBadRequest)
		return
	}
	ctx.Status(http.StatusCreated)
}

func UpdateToDo(ctx *gin.Context) {
	id := ctx.Param("id")
	var data types.ToDo
	if err := ctx.ShouldBind(&data); err == nil {
		log.Printf("Error: %v", err)
	}
	var item models.ToDo
	initializers.DB.First(&item, id)

	initializers.DB.Model(&item).Updates(models.ToDo{
		Title: data.Title,
		Note:  data.Note,
		Complete: data.Complete,
	})
	ctx.JSON(http.StatusOK, gin.H{
		"toDo": item,
	})
}

func DeleteToDo(ctx *gin.Context) {
	id := ctx.Param("id")

	initializers.DB.Delete(&models.ToDo{}, id)

	ctx.Status(http.StatusOK)
}

func UpdateToDoStatus(ctx *gin.Context) {
	id := ctx.Param("id")

	var data struct { Complete bool }
	if err := ctx.ShouldBind(&data); err == nil {
		log.Printf("Error: %v", err)
	}

	var item models.ToDo
	initializers.DB.First(&item, id)

	initializers.DB.Model(&item).UpdateColumn("complete", data.Complete)

	ctx.JSON(http.StatusOK, gin.H{
		"toDo": item,
	})
}