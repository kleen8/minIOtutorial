package main

import (
	"context"
	"fmt"
	"log"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/minio/minio-go/v7/pkg/lifecycle"
)

const (
	endpoint  = "localhost:9000"
	accessKey = "admin"
	secretKey = "password"
	bucket    = "test-bucket-jelle"
)

func main() {
	fmt.Println("Updating lifecycle of test bucket")

	// Initialize the MinIO client
	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKey, secretKey, ""),
		Secure: false,
	})
	if err != nil {
		log.Fatalf("error: %s\n", err.Error())
	}

	// Create lifecycle configuration
	config := lifecycle.NewConfiguration()
	config.Rules = []lifecycle.Rule{
		{
			ID:     "ExpireArchive",
			Status: "Enabled",
			RuleFilter: lifecycle.Filter{Prefix: "archive/"},
			Expiration: lifecycle.Expiration{
				Days: 30,
			},
		},
	}

	// Apply lifecycle configuration
	err = minioClient.SetBucketLifecycle(context.Background(), bucket, config)
	if err != nil {
		log.Fatalf("error: %s\n", err.Error())
	}

	fmt.Println("Lifecycle rule applied successfully!")
}

