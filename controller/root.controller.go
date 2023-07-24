package controller

import (
	"example/todogolang/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RootRouter(db *gorm.DB) {
	router := gin.Default()

	v1 := router.Group("/v1")

	{
		v1.GET("/items", services.GetItems(db))
		v1.POST("/items", services.CreateItem(db))
		v1.GET("/items/:id", services.GetItemById(db)) // get an item by ID
		v1.PUT("/items/:id", services.UpdateItemToDo(db))
		v1.DELETE("/items/:id", services.DeleteItemById(db))
	}

	router.Run()
}
