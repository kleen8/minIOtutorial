package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

const (
    endpoint = "localhost:9000"
    accessKey = "admin"
    secretKey = "password"
    bucket = "test-bucket-jelle"
)

var minioClient *minio.Client

func initMinio() {
    var err error
    minioClient, err = minio.New(endpoint, &minio.Options{
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
}

func uploadFile(c *gin.Context) {
    file, err := c.FormFile("file")
    if err != nil {
        log.Fatalf("error: %s\n", err.Error())
        c.JSON(http.StatusNotFound, gin.H{"error": "File not found"})
    }

    src, err := file.Open()
    if err != nil {
        log.Fatalf("error: %s\n", err.Error())
        c.JSON(http.StatusNotFound, gin.H{"error": "File not found"})
    }

    defer src.Close()
    objectName := file.Filename
    ctx := context.Background()
    _, err = minioClient.PutObject(ctx, bucket, objectName, src, file.Size, minio.PutObjectOptions{
            ContentType: "application/octet-stream"})
    if err != nil {
        log.Fatalf("error: %s\n", err.Error())
        c.JSON(http.StatusNotFound, gin.H{"error": "File not found"})
    }

    c.JSON(http.StatusOK, gin.H{"message" : "File uploaded succesfully"})
}

func getFile(c *gin.Context) {
    objectName := c.Param("filename")
    ctx := context.Background()
    object, err := minioClient.GetObject(ctx, bucket, objectName, minio.GetObjectOptions{})
    if err != nil {
        log.Fatalf("error: %s\n", err.Error())
        c.JSON(http.StatusNotFound, gin.H{"error": "File not found"})
    }
    defer object.Close()

    content, err := io.ReadAll(object)
    if err != nil {
        log.Fatalf("error: %s\n", err.Error())
        c.JSON(http.StatusNotFound, gin.H{"error": "File not found"})
    }

    c.String(http.StatusOK, string(content))

}


func main() {
    initMinio()

    r := gin.Default()

    r.POST("/upload", uploadFile)
    r.GET("/files/:filename", getFile)

    if err := r.Run(":8080"); err != nil {
        log.Fatalf("error: %s\n", err.Error())
    }
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
