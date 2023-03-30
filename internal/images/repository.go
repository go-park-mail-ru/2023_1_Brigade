package images

import (
	"github.com/labstack/echo/v4"
	"mime/multipart"
	"net/url"
)

type Repostiory interface {
	GetImage(ctx echo.Context, filename string) (*url.URL, error)
	LoadImage(ctx echo.Context, file multipart.File, filename string) error
}
