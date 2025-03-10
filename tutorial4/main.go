package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

const (
	endpoint  = "localhost:9000"
	accessKey = "admin"
	secretKey = "password"
	bucket    = "test-bucket-jelle"
)

func main() {
    fmt.Println("Trying to get files and metadata")

    minioClient, err := minio.New(endpoint, &minio.Options{
        Creds: credentials.NewStaticV4(accessKey, secretKey, ""),
        Secure: false,
    })
    if err != nil {
        log.Fatalf("error: %s\n", err.Error())
    }

    r := gin.Default()

    r.GET("/files", func(c *gin.Context) {
        files, err := listFiles(minioClient, bucket)
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
        }
        c.JSON(http.StatusOK, files)
    })

    r.Run(":8080")

}

// listFiles retrieves all objects in the bucket
func listFiles(client *minio.Client, bucket string) ([]map[string]interface{}, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	objects := client.ListObjects(ctx, bucket, minio.ListObjectsOptions{Recursive: true})

	var files []map[string]interface{}
	for obj := range objects {
		if obj.Err != nil {
			return nil, obj.Err
		}

		files = append(files, map[string]interface{}{
			"Name":         obj.Key,
			"Size":         obj.Size,
			"LastModified": obj.LastModified,
			"ContentType":  getObjectContentType(client, bucket, obj.Key),
		})
	}
	return files, nil
}

// getObjectContentType fetches content type from the object metadata
func getObjectContentType(client *minio.Client, bucket, objectName string) string {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	objInfo, err := client.StatObject(ctx, bucket, objectName, minio.StatObjectOptions{})
	if err != nil {
		return "unknown"
	}

	return objInfo.ContentType
}
