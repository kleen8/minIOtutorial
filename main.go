package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"strings"

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

    // initializing a new minio client
    minioClient, err := minio.New(endpoint, &minio.Options{
        Creds: credentials.NewStaticV4(accessKey, secretKey, ""),
        Secure: false,
    })
    if err != nil {
        log.Fatalln(err)
    }
    
    // Checking if the fucket exists
    ctx := context.Background();
    exists, err := minioClient.BucketExists(ctx, bucket)
    if err != nil {
        log.Fatalf("error: %s\n", err.Error())
    }
    if !exists {
        fmt.Println("Bucket does not exist, creating...")
        err = minioClient.MakeBucket(ctx, bucket, minio.MakeBucketOptions{})
        if err != nil {
            log.Fatalf("error: %s\n", err.Error())
        }
        fmt.Println("Bucket created succesfully")
    }

    // Upload a text file to this bucket
    objectName := "example.txt"
    textContext := "Hello, the file should be updated"
    err = uploadText(minioClient, objectName, textContext)
    if err != nil {
        log.Fatalf("error: %s\n", err.Error())
    }

    // Retrieve the text file
    retrievedText, err := getText(minioClient, objectName)
    if err != nil {
        log.Fatalf("error: %s\n", err.Error())
    }

    fmt.Println("Retrieved text: ", retrievedText)

    objectName2 := "testTekst.txt"
    // Retrieve the text file
    retrievedText1, err := getText(minioClient, objectName2)
    if err != nil {
        log.Fatalf("error: %s\n", err.Error())
    }
    fmt.Println("Retrieved text: ", retrievedText1)


}

func uploadText(client *minio.Client, objectName, content string) error {
    ctx := context.Background()
    reader := strings.NewReader(content)
    _, err := client.PutObject(ctx, bucket, objectName, reader, int64(reader.Len()), minio.PutObjectOptions{
        ContentType: "text/plain",
    })
    return err
}

func getText(client *minio.Client, objectName string) (string, error) {
    ctx := context.Background()
    object, err := client.GetObject(ctx, bucket, objectName, minio.GetObjectOptions{})
    if err != nil {
        return "", err
    }
    defer object.Close()
    content, err := io.ReadAll(object)
    if err != nil {
        log.Fatalf("error: %s\n", err.Error())
    }
    return string(content), nil
}
