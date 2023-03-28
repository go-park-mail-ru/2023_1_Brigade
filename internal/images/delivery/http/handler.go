package http

import (
	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
	"net/http"
	"project/internal/images"
	"project/internal/pkg/security"
)

type imagesHandler struct {
	imagesUsecase images.Usecase
}

func (h *imagesHandler) LoadCurrentUserAvatarHandler(ctx echo.Context) error {
	//session := ctx.Get("session").(model.Session)
	//userID := session.UserId

	userID := uint64(1) // заглушка

	maxSize := int64(64 << 20)
	err := ctx.Request().ParseMultipartForm(maxSize)
	if err != nil {
		return err
	}

	file, header, err := ctx.Request().FormFile("image")
	if err != nil {
		return err
	}
	filename := security.GenerateFilename(userID, header.Filename)
	defer func() {
		err := file.Close()
		if err != nil {
			log.Error(err)
		}
	}()

	_, err = h.imagesUsecase.LoadImage(ctx, file, filename)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusCreated, filename)
}

func NewImagesHandler(e *echo.Echo, imagesUsecase images.Usecase) imagesHandler {
	handler := imagesHandler{imagesUsecase: imagesUsecase}
	loadImagesUrl := "/images/"

	api := e.Group("api/v1")

	loadImages := api.Group(loadImagesUrl)

	loadImages.OPTIONS("", handler.LoadCurrentUserAvatarHandler)
	loadImages.POST("", handler.LoadCurrentUserAvatarHandler)

	return handler
}
