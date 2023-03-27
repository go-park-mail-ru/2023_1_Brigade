package s3

import (
	"io"

	"github.com/labstack/echo/v4"
	"github.com/minio/minio-go/v7"
)

type Repository interface {
	GetImageById(ctx echo.Context, imageID string) (image *minio.Object, err error)
	UploadImage(ctx echo.Context, image io.Reader, imageSize int64) error
}
