package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"line/pkg/opendoor"
	"line/pkg/postcallback"
)

func main() {
	engine := gin.Default()
	engine.GET("/", getTop)
	engine.POST("/callback", postcallback.PostCallback)
	engine.POST("/open_door", opendoor.OpenDoor)
	engine.Run(":" + "8080")
}

func getTop(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "hello world",
	})
}
