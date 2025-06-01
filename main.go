package main

import (
	"context"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

const (
	Endpoint        = "localhost:9000"
	AccessKeyId     = "minioadmin"
	SecretAccessKey = "minioadmin"
	UseSSL          = false
	MinioBucketName = "raw-data"
)

func main() {

	r := gin.Default()

	r.POST("/upload", func(c *gin.Context) {
		file, err := c.FormFile("file")
		if err != nil {
			c.AbortWithStatusJSON(400, gin.H{
				"error": "Файл не найден",
			})
		}

		src, err := file.Open()
		if err != nil {
			c.AbortWithStatusJSON(500, gin.H{
				"error": "Не удалось открыть файл",
			})
			return
		}
		defer src.Close()
	})

	client, err := minio.New(Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(AccessKeyId, SecretAccessKey, ""),
		Secure: UseSSL,
	})
	if err != nil {
		fmt.Println("Error:", err)
	}

	exist, err := client.BucketExists(context.Background(), MinioBucketName)
	if err != nil {
		fmt.Println("Ошибка проверки бакета", err)
	}

	if !exist {
		err = client.MakeBucket(context.Background(), MinioBucketName, minio.MakeBucketOptions{})
		if err != nil {
			fmt.Println("Ошибка создания бакета:", err)
		}
	}

	r.Run(":8080")
}
