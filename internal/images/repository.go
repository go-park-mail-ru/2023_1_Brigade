package images

import (
	"context"
	"io"
)

type Repostiory interface {
	UploadImage(ctx context.Context, file io.Reader, bucketName string, filename string) error
	GetImage(ctx context.Context, bucketName string, filename string) (string, error)
}
