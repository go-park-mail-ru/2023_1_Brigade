package usecase

import (
	"io"
	"project/internal/s3"

	"github.com/labstack/echo/v4"
	"github.com/minio/minio-go/v7"
)

type usecase struct {
	repo s3.Repository
}

func NewS3Usecase(s3Repo s3.Repository) s3.Usecase {
	return &usecase{repo: s3Repo}
}

func (u *usecase) GetImageById(ctx echo.Context, imageID string) (*minio.Object, error) {
	return u.repo.GetImageById(ctx, imageID)
}

func (u *usecase) UploadImage(ctx echo.Context, image io.Reader, imageSize int64) error {
	return u.repo.UploadImage(ctx, image, imageSize)
}
