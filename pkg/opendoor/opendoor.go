package opendoor

import (
	// "fmt"
	// "line/pkg/useaws"
	"net/http"
	"line/pkg/sendmessage"

	"github.com/gin-gonic/gin"
)


func OpenDoor(c *gin.Context) {

	type DoorStatus struct {
		KeyStatus   string `json:"key_status"`
		CurrentTime string `json:"time"`
		KeyMacID    string `json:"key_id"`
	}

	var status DoorStatus

	// リクエストボディからJSONをパース
	if err := c.BindJSON(&status); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// key_statusがOPENの場合にLINE BOTにメッセージを送信
	if status.KeyStatus == "OPEN" {
		// 鍵が開いた時のメッセージ
		sendmessage.SendMessageToLineBot("🔓 開いたよ！\n" + "時刻：" + status.CurrentTime, status.KeyMacID)
		c.JSON(http.StatusOK, gin.H{"message": "Door opened and message sent to LINE BOT"})
	} else if status.KeyStatus == "CLOSE" {
		// 鍵が閉まった時のメッセージ
		sendmessage.SendMessageToLineBot("🔒 閉まったよ！\n" + "時刻：" + status.CurrentTime, status.KeyMacID)
		c.JSON(http.StatusOK, gin.H{"message": "Door closed and message sent to LINE BOT"})
	} else if status.KeyStatus == "Warning_Open" {
		// 鍵が開けっぱなしの警告メッセージ
		sendmessage.SendMessageToLineBot("⚠️ 鍵が開けっぱなしですよ！気をつけて！\n" + "MACアドレス：" + status.KeyMacID, status.KeyMacID)
		c.JSON(http.StatusOK, gin.H{"message": "Warning: Door left open and message sent to LINE BOT"})
	} else {
		// 不正な鍵の状態が指定された場合
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid key_status"})
	}	
}
