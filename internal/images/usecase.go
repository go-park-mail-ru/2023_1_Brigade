package images

import (
	"context"
	"mime/multipart"
)

type Usecase interface {
	LoadImage(ctx context.Context, file multipart.File, filename string, userID uint64) (string, error)
}
