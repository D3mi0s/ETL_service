package main

import (
	"context"
	"fmt"

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
}
