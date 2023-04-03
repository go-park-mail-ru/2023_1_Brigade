package images

import (
	"context"
	"mime/multipart"
	"net/url"
)

type Repostiory interface {
	LoadImage(ctx context.Context, file multipart.File, filename string, userID uint64) (*url.URL, error)
}
