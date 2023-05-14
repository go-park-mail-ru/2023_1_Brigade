package repository

import (
	"context"
	"errors"
	"github.com/minio/minio-go/v7"
	"io"
	"net/url"
	"project/internal/configs"
	"project/internal/images"
	"time"
)

func NewImagesMemoryRepository(s3UserAvatars *minio.Client, s3ChatAvatars *minio.Client, s3ChatImages *minio.Client) images.Repository {
	return &repository{s3UserAvatars: s3UserAvatars, s3ChatAvatars: s3ChatAvatars, s3ChatImages: s3ChatImages}
}

type repository struct {
	s3UserAvatars *minio.Client
	s3ChatAvatars *minio.Client
	s3ChatImages  *minio.Client
}

func (r repository) UploadImage(ctx context.Context, file io.Reader, bucketName string, filename string) error {
	err := errors.New("")
	switch bucketName {
	case configs.UserAvatarsBucket:
		_, err = r.s3UserAvatars.PutObject(ctx, bucketName, filename, file, -1, minio.PutObjectOptions{})
	case configs.ChatAvatarsBucket:
		_, err = r.s3ChatAvatars.PutObject(ctx, bucketName, filename, file, -1, minio.PutObjectOptions{})
	case configs.ChatImagesBucket:
		_, err = r.s3ChatImages.PutObject(ctx, bucketName, filename, file, -1, minio.PutObjectOptions{})
	}

	return err
}

func (r repository) GetImage(ctx context.Context, bucketName string, filename string) (string, error) {
	expires := time.Hour * 24 * 7
	url := &url.URL{}
	err := errors.New("")

	switch bucketName {
	case configs.UserAvatarsBucket:
		url, err = r.s3UserAvatars.PresignedGetObject(ctx, bucketName, filename, expires, nil)
	case configs.ChatAvatarsBucket:
		url, err = r.s3ChatAvatars.PresignedGetObject(ctx, bucketName, filename, expires, nil)
	case configs.ChatImagesBucket:
		url, err = r.s3ChatImages.PresignedGetObject(ctx, bucketName, filename, expires, nil)
	}
	if err != nil {
		return "", err
	}

	return url.String(), nil
}
