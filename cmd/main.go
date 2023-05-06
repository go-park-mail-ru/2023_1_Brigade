package main

import (
	"context"
	"fmt"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"log"
	"time"
)

func main() {
	accessKey := "5C8YjViNM475zK7rafg8ut"
	secKey := "i1Prj7cjWGdDTQrEpbhX37wfcQRtAzAcvqsbtpRD6VG9"
	endpoint := "hb.bizmrg.com"
	ssl := true

	// Подключиться к VK Cloud S3.
	client, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKey, secKey, ""),
		Secure: ssl,
	})
	if err != nil {
		log.Fatal(err)
	}

	// Показать список всех бакетов.
	buckets, err := client.ListBuckets(context.TODO())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("List of all buckets for this access key:")
	for _, bucket := range buckets {
		fmt.Println(bucket.Name)
	}

	// Создаем ссылку на объект
	expires := 24 * time.Hour // ссылка будет действительна 24 часа
	url, err := client.PresignedGetObject(context.TODO(), "brigade_images", "Снимок экрана от 2023-05-05 14-59-26.png", expires, nil)
	if err != nil {
		log.Fatal(err)
	}

	// Выводим ссылку на экран
	fmt.Println(url)
}
