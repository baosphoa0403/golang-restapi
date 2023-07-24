package services

import (
	"example/todogolang/model"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetItems(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		type DataPaging struct {
			Page  int   `json:"page" form:"page"`
			Limit int   `json:"limit" form:"limit"`
			Total int64 `json:"total" form:"-"`
		}

		var paging DataPaging

		if err := ctx.ShouldBind(&paging); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if paging.Page <= 0 {
			paging.Page = 1
		}

		if paging.Limit <= 0 {
			paging.Limit = 10
		}

		offset := (paging.Page - 1) * paging.Limit

		var result []model.ToDoItem

		db.Table(model.ToDoItem{}.TableName()).Limit(paging.Limit).
			Count(&paging.Total).
			Limit(paging.Limit).
			Offset(offset).
			Order("id desc").
			Find(&result)

		ctx.JSON(http.StatusOK, gin.H{"data": result})
	}
}

func CreateItem(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var dataItem model.ToDoItem

		if err := ctx.ShouldBind(&dataItem); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// preprocess title - trim all spaces
		dataItem.Title = strings.TrimSpace(dataItem.Title)

		if dataItem.Title == "" {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "title cannot be blank"})
			return
		}

		// do not allow "finished" status when creating a new task
		dataItem.Status = "Doing" // set to default

		if err := db.Create(&dataItem).Error; err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"data": dataItem.Id})
	}
}

func GetItemById(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var dataItem model.ToDoItem

		id, err := strconv.Atoi(ctx.Param("id"))

		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := db.First(&dataItem, id).Error; err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"data": dataItem})
	}
}

func UpdateItemToDo(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))

		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		var dataItem model.ToDoItem

		if err := ctx.ShouldBind(&dataItem); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		fmt.Println(dataItem.Title)

		if err := db.Where("id = ?", id).Updates(&dataItem).Error; err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"data": true})
	}
}

func DeleteItemById(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := db.Table(model.ToDoItem{}.TableName()).
			Where("id = ?", id).
			Delete(nil).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"data": true})
	}
}
