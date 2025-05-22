package main

import (
	"context"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

const (
	MinioEndpoint  = "localhost:9000"
	MinioAccessKey = "minioadmin"
	MinioSecretKey = "minioadmin"
	MinioBucket    = "raw_data"
)

func main() {

}

func initMinIO() *minio.Client {

	client, err := minio.New(MinioEndpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(MinioAccessKey, MinioSecretKey, ""),
		Secure: false,
	})
	if err != nil {
		log.Fatalf("Ошибка создания MinIO клиента: %v", err)
	}

	ctx := context.Background()
	exist, err := client.BucketExists(ctx, MinioBucket)
	if err != nil {
		log.Fatalf("Ошибка проверки бакета: %v", err)
	}

	if !exist {
		err := client.MakeBucket(ctx, MinioBucket, minio.MakeBucketOptions{})
		if err != nil {
			log.Fatalf("Ошибка создания бакета: %v", err)
		}
		log.Printf("Бакет %s успешно создан", MinioBucket)
	}

	return client
}

func setupRouter(minioClient *minio.Client) *gin.Engine {
	router := gin.Default()

	router.POST("/upload", func(c *gin.Context) {
		file, header, err := c.Request.FormFile("file")
		if err != nil {
			sendError(c, http.StatusBadRequest, "file is not found in order")
		}
		defer file.Close()
	})

	return router
}

func sendError(c *gin.Context, code int, message string) {
	c.AbortWithStatusJSON(code, gin.H{
		"error":   true,
		"message": message,
	})
}
