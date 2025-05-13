package main

import (
	"context"
	"log"

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

	}

	return client
}
