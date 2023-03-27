package repository

import (
	"io"

	"project/internal/s3"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/minio/minio-go/v7"
)

type repository struct {
	client *minio.Client
}

func NewImageStorageRepository(mc *minio.Client) s3.Repository {
	return &repository{client: mc}
}

func (r *repository) UploadImage(ctx echo.Context, image io.Reader, imageSize int64) error {
	newId := uuid.New().String()
	_, err := r.client.PutObject(
		ctx.Request().Context(),
		"images",
		newId,
		image,
		imageSize,
		minio.PutObjectOptions{},
	)
	return err
}

func (r *repository) GetImageById(ctx echo.Context, imageID string) (*minio.Object, error) {
	return r.client.GetObject(
		ctx.Request().Context(),
		"images",
		imageID,
		minio.GetObjectOptions{},
	)
}
