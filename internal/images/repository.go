package images

import (
	"github.com/labstack/echo/v4"
	"mime/multipart"
	"net/url"
)

type Repostiory interface {
	LoadImage(ctx echo.Context, file multipart.File, filename string, userID uint64) (*url.URL, error)
}
