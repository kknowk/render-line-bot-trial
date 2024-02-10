package useaws

import (
	"context"
	"fmt"
	"io"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"os"
)

// AWS S3クライアントの初期化
func initS3Client() *s3.Client {

	accessKeyID := os.Getenv("AWS_ACCESS_KEY_ID")
	secretAccessKey := os.Getenv("AWS_SECRET_ACCESS_KEY")

	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion("ap-northeast-1"),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(accessKeyID, secretAccessKey, "")),
	)
	if err != nil {
		fmt.Println("Error loading AWS configuration:", err)
	} else {
		fmt.Println("Successfully loaded AWS configuration")
	}
	return s3.NewFromConfig(cfg)
}

// ユーザーIDとMACアドレスの紐付けをS3に保存
func AssociateUserWithMacAddress(userID, macAddress string) {
	client := initS3Client()

	bucket := "magickeybucket"                       // 適切なバケット名に設定
	key := fmt.Sprintf("UserDevices/%s.txt", "testUserID") // ファイル名（ここではuserIDを使用）
	body := fmt.Sprintf("UserID: %s, MacAddress: %s", userID, macAddress)

	_, err := client.PutObject(context.TODO(), &s3.PutObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
		Body:   strings.NewReader(body),
	})
	if err != nil {
		fmt.Println("Failed to upload to S3:", err)
	} else {
		fmt.Println("Successfully uploaded to S3")
	}
}

// 指定されたMACアドレスに対応するユーザーIDをS3バケットから取得
func FindUserIDByMacAddress(macAddress string) (string, error) {
	client := initS3Client()
	if client == nil {
		return "", fmt.Errorf("Failed to initialize S3 client")
	}

	bucket := "magickeybucket"                             // バケット名を設定
	key := fmt.Sprintf("UserDevices/%s.txt", "testUserID") // MACアドレスに基づくファイルパス

	// S3からファイルを取得
	output, err := client.GetObject(context.TODO(), &s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	})
	if err != nil {
		return "", fmt.Errorf("Failed to get object from S3: %v", err)
	}
	defer output.Body.Close()

	// ファイルの内容を読み込む
	body, err := io.ReadAll(output.Body)
	if err != nil {
		return "", fmt.Errorf("Failed to read object body: %v", err)
	}

	// ファイルの内容を文字列に変換
	content := string(body)

	// 文字列を解析してUserIDを取得
	var userID string
	parts := strings.Split(content, ",")
	for _, part := range parts {
		// "UserID: "で始まる部分を探す
		if strings.HasPrefix(part, "UserID: ") {
			userID = strings.TrimSpace(strings.TrimPrefix(part, "UserID: "))
			break
		}
	}

	if userID == "" {
		return "", fmt.Errorf("UserID not found in object body")
	}

	return userID, nil
}
