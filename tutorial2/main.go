package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

const (
    endpoint = "localhost:9000"
    accessKey = "admin"
    secretKey = "password"
    bucket = "test-bucket-jelle"
)

func main() {
    
    minioClient, err := minio.New(endpoint, &minio.Options{
        Creds: credentials.NewStaticV4(accessKey, secretKey, ""),
        Secure: false,
        })
    if err != nil {
        log.Fatalf("error: %s\n", err.Error())
    }

    objectName := "example.txt"
    expiry := time.Minute * 10

    presignedURL, err := minioClient.PresignedGetObject(context.Background(),
        bucket, objectName, expiry, nil)
    if err != nil {
        log.Fatalf("error: %s\n", err.Error())
    }
    fmt.Println("Presigned GET URL:", presignedURL)
    fmt.Println()

    putObjectName := "tutorial2TestFile.txt"

    presignedPutUrl, err := minioClient.PresignedPutObject(context.Background(), bucket, putObjectName, expiry)
    if err != nil {
        log.Fatalf("error: %s\n", err.Error())
    }
    fmt.Println("Presigned PUT URL:", presignedPutUrl)
}

