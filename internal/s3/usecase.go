package s3

import (
	"io"

	"github.com/labstack/echo/v4"
	"github.com/minio/minio-go/v7"
)

type Usecase interface {
	UploadImage(ctx echo.Context, image io.Reader, imageSize int64) error
	GetImageById(ctx echo.Context, imageID string) (*minio.Object, error)
}
