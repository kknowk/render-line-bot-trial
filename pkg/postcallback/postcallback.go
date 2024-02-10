package postcallback

import (
	"fmt"
	"regexp"
	// "net/http"

	"github.com/gin-gonic/gin"
	"github.com/line/line-bot-sdk-go/v7/linebot"

	"line/pkg/useaws"
)

func PostCallback(c *gin.Context) {
	// bot作成
	bot, err := linebot.New(
		"f7f28f6ac6442036faebd8c24419b3c3",
		"LYKvf90XMx1tZlGa47tYUhR/DcwfxilyHo72WdmhoRrhuZKYfJk8BI6gk72x3gqClGpsKwnF77JuggRQq4U+jH3N1cX0JpXClSG2m0vbAeMCH3tsN4z+teiLBilOI2XAi6pOlUNkqIkoO7JCI2mrmAdB04t89/1O/w1cDnyilFU=",
	)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// リクエスト処理
	events, berr := bot.ParseRequest(c.Request)
	if berr != nil {
		fmt.Println(berr.Error())
		return
	}

	for _, event := range events {
		if event.Type == linebot.EventTypeMessage {
			switch message := event.Message.(type) {
			case *linebot.TextMessage:

				// メッセージからMACアドレスを抽出
				macAddress := message.Text
				if !IsOption(macAddress) {
					// スルー
					continue
				}
				// ユーザーIDとMACアドレスを紐付け
				if !IsMacAddress(macAddress) {
					// エラーメッセージをユーザーに送信
					_, err := bot.ReplyMessage(
						event.ReplyToken,
						linebot.NewTextMessage("MACアドレス "+macAddress+" は不正です。"),
					).Do()
					if err != nil {
						fmt.Print(err)
					}
					return
				}
				userID := event.Source.UserID
				useaws.AssociateUserWithMacAddress(userID, macAddress)
				// 確認メッセージをユーザーに送信
				_, err := bot.ReplyMessage(
					event.ReplyToken,
					linebot.NewTextMessage("MACアドレス "+macAddress+" を登録しました。"),
				).Do()
				if err != nil {
					fmt.Print(err)
				}
			}
		}
	}

	// // リクエストボディからデータを抽出するための構造体を定義
	// var requestData struct {
	// 	UserID     string `json:"user_id"`
	// 	MacAddress string `json:"mac_address"`
	// }

	// // リクエストボディを解析
	// if err := c.BindJSON(&requestData); err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
	// 	return
	// }

	// // 抽出したデータをS3に保存
	// useaws.AssociateUserWithMacAddress(requestData.UserID, requestData.MacAddress)

	// // 応答を送信
	// c.JSON(http.StatusOK, gin.H{"message": "Data successfully saved to S3"})

}

func IsMacAddress(input string) bool {
	// MACアドレスの正規表現パターン
	macAddressPattern := `^([0-9A-Fa-f]{2}[:-]){5}([0-9A-Fa-f]{2})$`
	re := regexp.MustCompile(macAddressPattern)

	return re.MatchString(input)
}

func IsOption(input string) bool {

	// 一致を確認したい文字列のリスト
	options := []string{
		"今日の鍵占い",
		"初期設定・使い方",
		"お問い合わせ",
		"鍵を紛失した",
	}

	// 入力された文字列がリストのいずれかと一致するか確認
	for _, option := range options {
		if input == option {
			return true
		}
	}

	// 一致する文字列がない場合はfalseを返す
	return false
}
