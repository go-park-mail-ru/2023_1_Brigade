package repository

import (
	"context"
	"github.com/labstack/echo/v4"
	"github.com/minio/minio-go/v7"
	log "github.com/sirupsen/logrus"
	"mime/multipart"
	"net/url"
	"project/internal/images"
	myErrors "project/internal/pkg/errors"
	"time"
)

func NewImagesMemoryRepository(minioClient *minio.Client) images.Repostiory {
	bucketname := "avatars"
	err := minioClient.MakeBucket(context.Background(), bucketname, minio.MakeBucketOptions{})
	if err != nil {
		log.Error(err)
	}

	return &repository{minio: minioClient, bucketname: bucketname}
}

type repository struct {
	minio      *minio.Client
	bucketname string
}

func (r *repository) GetImage(ctx echo.Context, filename string) (*url.URL, error) {
	expires := time.Second * 60
	presignedURL, err := r.minio.PresignedGetObject(context.Background(), r.bucketname, filename, expires, nil)
	if minio.ToErrorResponse(err).Code == "NoSuchKey" {
		return nil, myErrors.ErrAvatarNotFound
	}

	return presignedURL, err
}

func (r *repository) LoadImage(ctx echo.Context, file multipart.File, filename string) error {
	_, err := r.minio.PutObject(context.Background(), r.bucketname, filename, file, -1, minio.PutObjectOptions{})
	return err
}
