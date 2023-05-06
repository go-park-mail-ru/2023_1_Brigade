package repository

import (
	"context"
	"github.com/minio/minio-go/v7"
	"io"
	"project/internal/images"
	"time"
)

func NewImagesMemoryRepository(s3 *minio.Client) images.Repostiory {
	return &repository{s3: s3}
}

type repository struct {
	s3 *minio.Client
}

func (r repository) UploadImage(ctx context.Context, file io.Reader, bucketName string, filename string) error {
	_, err := r.s3.PutObject(context.Background(), bucketName, filename, file, -1, minio.PutObjectOptions{})
	return err
}

func (r repository) GetImage(ctx context.Context, bucketName string, filename string) (string, error) {
	expires := 24 * time.Hour // ссылка будет действительна 24 часа
	url, err := r.s3.PresignedGetObject(ctx, bucketName, filename, expires, nil)
	if err != nil {
		return "", err
	}

	return url.String(), nil
}

//func (r repository) UploadImage(ctx context.Context, file multipart.File, filename string, userID uint64) (string, error) {
//	hash := uuid.New().String()
//	fileOnDisk, err := os.Create("../../avatars/" + hash + ".jpg")
//	if err != nil {
//		return "", err
//	}
//	defer file.Close()
//
//	_, err = io.Copy(fileOnDisk, file)
//	if err != nil {
//		return "", err
//	}
//
//	url := "https://technogramm.ru/avatars/" + hash + ".jpg"
//	rows, err := r.db.Query("UPDATE profile SET avatar=$1 WHERE id=$2", url, userID)
//	defer rows.Close()
//
//	return url, nil
//}
