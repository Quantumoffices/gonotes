package main

import (
	"context"
	"log"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

func main() {
	endpoint := "116.62.245.150:9000"
	accessKey := "minioadmin"
	secretAccessKey := "minioadmin"
	useSSL := false

	// Initialize minio client object.
	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKey, secretAccessKey, ""),
		Secure: useSSL,
	})
	if err != nil {
		log.Fatalln(err)
	}

	// Make a new bucket called mymusic.
	bucketName := "image"
	location := "us-east-1"
	exists, err := minioClient.BucketExists(context.Background(), bucketName)
	if err != nil {
		panic(err)
	}
	if !exists {
		err = minioClient.MakeBucket(context.Background(), bucketName, minio.MakeBucketOptions{Region: location})
		if err != nil {
			panic(err)
		}
	}
	// Upload the zip file
	objectName := "11.png"
	filePath := "./minio/image/basketball.png"
	contentType := "application/octet-stream"

	// Upload the zip file with FPutObject
	uploadInfo, err := minioClient.FPutObject(context.Background(), bucketName, objectName, filePath, minio.PutObjectOptions{ContentType: contentType})
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("Successfully uploaded %s of size %+v\n", objectName, uploadInfo)

	//minioClient.FGetObject(context.Background(), bucketName, objectName, filePath, minio.PutObjectOptions{ContentType: contentType})
}
