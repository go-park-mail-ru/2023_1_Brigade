package main

import (
	"context"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	log "github.com/sirupsen/logrus"
	repositoryImages "project/internal/images/repository"
	usecaseImages "project/internal/images/usecase"
)

// import (
//
//	"context"
//	"fmt"
//	"github.com/minio/minio-go/v7"
//	"github.com/minio/minio-go/v7/pkg/credentials"
//	"log"
//	"os"
//	"project/internal/pkg/image_generation"
//	"time"
//
// )
func main() {
	accessKey := "hPGMCe6ZttM8VBVs7sXkFi"
	secKey := "9knxejdQVDA3J8YGchKjh2XvMzyupvakHJqG6kBwe15R"
	endpoint := "hb.bizmrg.com"
	ssl := true

	client, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKey, secKey, ""),
		Secure: ssl,
	})
	if err != nil {
		log.Fatal(err)
	}

	imagesRepostiory := repositoryImages.NewImagesMemoryRepository(client)
	imagesUsecase := usecaseImages.NewImagesUsecase(imagesRepostiory)

	err = imagesUsecase.UploadGeneratedImage(context.TODO(), "brigade_user_avatars", "test", "S")
	if err != nil {
		log.Fatal(err)
	}
	//
	//file, err := os.Open("../../avatars/background.png")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//err = imagesUsecase.UploadImage(context.TODO(), file, "brigade_user_avatars", "test.png")
	//log.Fatal(err)
}

//	accessKey := "hPGMCe6ZttM8VBVs7sXkFi"
//	secKey := "9knxejdQVDA3J8YGchKjh2XvMzyupvakHJqG6kBwe15R"
//	endpoint := "hb.bizmrg.com"
//	ssl := true
//
//	// Подключиться к VK Cloud S3.
//	client, err := minio.New(endpoint, &minio.Options{
//		Creds:  credentials.NewStaticV4(accessKey, secKey, ""),
//		Secure: ssl,
//	})
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	// Показать список всех бакетов.
//	buckets, err := client.ListBuckets(context.TODO())
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	fmt.Println("List of all buckets for this access key:")
//	for _, bucket := range buckets {
//		fmt.Println(bucket.Name)
//	}
//
//	_, err = image_generation.GenerateAvatar("a")
//	if err != nil {
//		log.Fatal(err)
//	}
//	//
//	//file, err := os.Open("../../avatars/background.png")
//	//if err != nil {
//	//	log.Fatal(err)
//	//}
//	//defer file.Close()
//	//file =
//	//
//	_, err = client.PutObject(context.Background(), "brigade_user_avatars", "a.png", file, -1, minio.PutObjectOptions{})
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	// Создаем ссылку на объект
//	expires := 24 * time.Hour // ссылка будет действительна 24 часа
//	url, err := client.PresignedGetObject(context.TODO(), "brigade_user_avatars", "a.png", expires, nil)
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	// Выводим ссылку на экран
//	fmt.Println(url)
//}
