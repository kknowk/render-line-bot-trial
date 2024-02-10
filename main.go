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
	
	// 静的ファイルのルートを追加
	engine.StaticFile("/red", "./public/red.html") // publicディレクトリ内のred.htmlを提供
	engine.StaticFile("/blue", "./public/blue.html") // publicディレクトリ内のblue.htmlを提供
	engine.StaticFile("/yellow", "./public/yellow.html") // publicディレクトリ内のyellow.htmlを提供

	engine.Run(":" + "8080")
}

func getTop(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "hello world",
	})
}
