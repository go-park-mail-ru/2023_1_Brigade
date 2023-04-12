package repository

import (
	"context"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/minio/minio-go/v7"
	log "github.com/sirupsen/logrus"
	"io"
	"mime/multipart"
	"os"
	"project/internal/images"
)

func NewImagesMemoryRepository(db *sqlx.DB, minioClient *minio.Client) images.Repostiory {
	bucketname := "avatars"
	err := minioClient.MakeBucket(context.Background(), bucketname, minio.MakeBucketOptions{})
	if err != nil {
		log.Error(err)
	}

	return &repository{db: db, minio: minioClient, bucketname: bucketname}
}

type repository struct {
	db         *sqlx.DB
	bucketname string
	minio      *minio.Client
}

func (r repository) LoadImage(ctx context.Context, file multipart.File, filename string, userID uint64) (string, error) {
	hash := uuid.New().String()
	fileOnDisk, err := os.Create("/home/ubuntu/avatars/" + hash)
	if err != nil {
		return "", err
	}
	defer file.Close()

	_, err = io.Copy(fileOnDisk, file)
	if err != nil {
		return "", err
	}

	url := "https://technogramm.ru/avatars/" + hash
	_, err = r.db.Query("UPDATE profile SET avatar=$1 WHERE id=$2", url, userID)

	return url, nil

	//_, err := r.minio.PutObject(context.Background(), r.bucketname, filename, file, -1, minio.PutObjectOptions{})
	//if err != nil {
	//	return nil, err
	//}
	//
	//expires := time.Hour * 24 * 7 // 7 days
	//presignedURL, err := r.minio.PresignedGetObject(context.Background(), r.bucketname, filename, expires, nil)
	//if minio.ToErrorResponse(err).Code == "NoSuchKey" {
	//	return nil, myErrors.ErrAvatarNotFound
	//}
	//
	//_, err = r.db.Query("UPDATE profile SET avatar=$1 WHERE id=$2", presignedURL.String(), userID)
	//
	//return presignedURL, err
}
