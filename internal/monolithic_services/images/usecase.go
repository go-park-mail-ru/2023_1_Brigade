package images

import (
	"context"
	"io"
)

type Usecase interface {
	UploadGeneratedImage(ctx context.Context, bucketName string, filename string, firstCharacterName string) error
	UploadImage(ctx context.Context, file io.Reader, bucketName string, filename string) error
	GetImage(ctx context.Context, bucketName string, filename string) (string, error)
}
