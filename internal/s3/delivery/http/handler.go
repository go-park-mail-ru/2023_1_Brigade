package http

import (
	"net/http"
	"project/internal/s3"

	"github.com/labstack/echo/v4"
)

type s3Handler struct {
	usecase s3.Usecase
}

func (sh *s3Handler) GetImageHandler(ctx echo.Context) error {
	image, err := sh.usecase.GetImageById(ctx, ctx.Param("imageID"))
	if err != nil {
		return err
	}
	info, err := image.Stat()
	if err != nil {
		return err
	}
	buf := make([]byte, 0, info.Size)
	image.Read(buf)
	return ctx.Blob(http.StatusOK, info.ContentType, buf)
}

func (sh *s3Handler) UploadImageHandler(ctx echo.Context) error {
	image, err := ctx.FormFile("image")
	if err != nil {
		return err
	}
	
}

func NewS3Handler(e echo.Echo, s3Usecase s3.Usecase) s3Handler {
	handler := s3Handler{}

	uploadImageUrl := "/image/"
	getImageUrl := "/image/:imageID"

	api := e.Group("api/v1")

	uploadImage := api.Group(uploadImageUrl)
	getImage := api.Group(getImageUrl)

	uploadImage.POST("", handler.UploadImageHandler)
	getImage.GET("", handler.GetImageHandler)

	return handler
}
