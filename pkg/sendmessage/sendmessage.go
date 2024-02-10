package sendmessage

import (
	"fmt"
	"github.com/line/line-bot-sdk-go/v7/linebot"

	"line/pkg/useaws"
)


func SendMessageToLineBot(message string, macAddress string) {
	// LINE BOTのチャネルシークレットとアクセストークン
	channelSecret := "f7f28f6ac6442036faebd8c24419b3c3"                                                                                                                                            // 例: "YOUR_CHANNEL_SECRET"
	channelToken := "LYKvf90XMx1tZlGa47tYUhR/DcwfxilyHo72WdmhoRrhuZKYfJk8BI6gk72x3gqClGpsKwnF77JuggRQq4U+jH3N1cX0JpXClSG2m0vbAeMCH3tsN4z+teiLBilOI2XAi6pOlUNkqIkoO7JCI2mrmAdB04t89/1O/w1cDnyilFU=" // 例: "YOUR_CHANNEL_ACCESS_TOKEN"

	// LINE BOTクライアントの初期化
	bot, err := linebot.New(channelSecret, channelToken)
	if err != nil {
		fmt.Println("LINE BOTの初期化に失敗:", err)
		return
	}

	toUserID, err := useaws.FindUserIDByMacAddress(macAddress) // MACアドレスに対応するユーザーIDを取得
	if err != nil {
		fmt.Println("Failed to find user ID by MAC address:", err)
		return
	}

	// ここにメッセージを送信するユーザーのIDを設定
	// toUserID := "U36b4b97e241ad297ccfb51fd50f3f14c" // 送信先のユーザーID

	// メッセージの送信
	if _, err := bot.PushMessage(toUserID, linebot.NewTextMessage(message)).Do(); err != nil {
		fmt.Println("メッセージの送信に失敗:", err)
	} else {
		fmt.Println("メッセージを送信しました")
	}

	var imageURL string = "https://i.pinimg.com/originals/aa/b1/ef/aab1efb29ccf7d6fae94997f33548f18.jpg"
	// 画像メッセージを送信
	if imageURL != "" && message == "鍵が開けっぱなしですよ！" {
		if _, err := bot.PushMessage(toUserID, linebot.NewImageMessage(imageURL, imageURL)).Do(); err != nil {
			fmt.Println("Error sending image message:", err)
		}
	}
}