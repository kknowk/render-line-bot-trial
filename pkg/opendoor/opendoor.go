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

	// toUserID, err := useaws.FindUserIDByMacAddress(status.KeyMacID) // MACアドレスに対応するユーザーIDを取得
	// if err != nil {
	// 	fmt.Println("Failed to find user ID by MAC address:", err)
	// 	return
	// }

	// c.JSON(http.StatusOK, gin.H{"message": toUserID + " " + status.KeyStatus + " " + status.CurrentTime + " " + status.KeyMacID})

	// key_statusがOPENの場合にLINE BOTにメッセージを送信
	if status.KeyStatus == "OPEN" {
		sendmessage.SendMessageToLineBot("OPEN!\n" + status.CurrentTime, status.KeyMacID) // LINE BOTにメッセージを送信する関数（実装は別途必要）
		c.JSON(http.StatusOK, gin.H{"message": "Door opened and message sent to LINE BOT"})
	} else if status.KeyStatus == "CLOSE" {
		sendmessage.SendMessageToLineBot("CLOSE!\n" + status.CurrentTime, status.KeyMacID) // LINE BOTにメッセージを送信する関数（実装は別途必要）
		c.JSON(http.StatusOK, gin.H{"message": "Door opened and message sent to LINE BOT"})
	} else if status.KeyStatus == "Warning_Open" {
		sendmessage.SendMessageToLineBot("鍵が開けっぱなしですよ！", status.KeyMacID) // LINE BOTにメッセージを送信する関数（実装は別途必要）
		c.JSON(http.StatusOK, gin.H{"message": "Door opened and message sent to LINE BOT"})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid key_status"})
	}
}
