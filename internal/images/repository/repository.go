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

func NewImagesMemoryRepository(s3_user_avatars *minio.Client, s3_chat_avatars *minio.Client, s3_chat_images *minio.Client) images.Repository {
	return &repository{s3_user_avatars: s3_user_avatars, s3_chat_avatars: s3_chat_avatars, s3_chat_images: s3_chat_images}
}

type repository struct {
	s3_user_avatars *minio.Client
	s3_chat_avatars *minio.Client
	s3_chat_images  *minio.Client
}

func (r repository) UploadImage(ctx context.Context, file io.Reader, bucketName string, filename string) error {
	err := errors.New("")
	switch bucketName {
	case configs.User_avatars_bucket:
		_, err = r.s3_user_avatars.PutObject(context.Background(), bucketName, filename, file, -1, minio.PutObjectOptions{})
	case configs.Chat_avatars_bucket:
		_, err = r.s3_chat_avatars.PutObject(context.Background(), bucketName, filename, file, -1, minio.PutObjectOptions{})
	case configs.Chat_images_bucket:
		_, err = r.s3_chat_images.PutObject(context.Background(), bucketName, filename, file, -1, minio.PutObjectOptions{})
	}

	return err
}

func (r repository) GetImage(ctx context.Context, bucketName string, filename string) (string, error) {
	expires := 24 * time.Hour
	url := &url.URL{}
	err := errors.New("")

	switch bucketName {
	case configs.User_avatars_bucket:
		url, err = r.s3_user_avatars.PresignedGetObject(ctx, bucketName, filename, expires, nil)
	case configs.Chat_avatars_bucket:
		url, err = r.s3_chat_avatars.PresignedGetObject(ctx, bucketName, filename, expires, nil)
	case configs.Chat_images_bucket:
		url, err = r.s3_chat_images.PresignedGetObject(ctx, bucketName, filename, expires, nil)
	}
	if err != nil {
		return "", err
	}

	return url.String(), nil
}
