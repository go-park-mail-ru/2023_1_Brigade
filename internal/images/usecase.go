package images

import (
	"github.com/labstack/echo/v4"
	"mime/multipart"
)

type Usecase interface {
	LoadImage(ctx echo.Context, file multipart.File, filename string, userID uint64) (string, error)
}
