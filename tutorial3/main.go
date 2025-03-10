// In this Go script we move a file from the src bucket to an archive bucket, doing something that is named "soft delete"

package main

import (
	"context"
	"fmt"
    "log"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

const (
    endpoint = "localhost:9000"
    accesKey = "admin"
    secretKey = "password"
    bucket = "test-bucket-jelle"
)

func main() {

    minioClient, err := minio.New(endpoint, &minio.Options{
        Creds: credentials.NewStaticV4(accesKey, secretKey, ""),
        Secure: false,
    })
    if err != nil {
        log.Fatalf("error: %s\n", err.Error())
    }

    srcBucket := "test-bucket-jelle"
    srcObject := "example.txt"
    dstBucket := srcBucket
    dstObject := "archive/" +  srcObject

    src := minio.CopySrcOptions{Bucket: srcBucket, Object: srcObject}
    dst := minio.CopyDestOptions{Bucket: dstBucket, Object: dstObject}

    _, err = minioClient.CopyObject(context.Background(), dst, src)
    if err != nil {
        log.Fatalf("error: %s\n", err.Error())
    }
    
    err = minioClient.RemoveObject(context.Background(), srcBucket, srcObject, minio.RemoveObjectOptions{})
    if err != nil {
        log.Fatalf("error: %s\n", err.Error())
    }
    fmt.Println("file moved to archive succesfully")
}
