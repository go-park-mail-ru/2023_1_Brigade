package http

import (
	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
	"net/http"
	"project/internal/images"
	"project/internal/model"
	"project/internal/user"
)

type imagesHandler struct {
	userUsecase   user.Usecase
	imagesUsecase images.Usecase
}

func (h imagesHandler) LoadCurrentUserAvatarHandler(ctx echo.Context) error {
	session := ctx.Get("session").(model.Session)
	userID := session.UserId

	maxSize := int64(64 << 20)
	err := ctx.Request().ParseMultipartForm(maxSize)
	if err != nil {
		return err
	}

	file, header, err := ctx.Request().FormFile("image")
	if err != nil {
		return err
	}

	defer func() {
		err := file.Close()
		if err != nil {
			log.Error(err)
		}
	}()

	url, err := h.imagesUsecase.LoadImage(ctx, file, header.Filename, userID)
	if err != nil {
		return err
	}

	user, err := h.userUsecase.GetUserById(ctx, session.UserId)
	if err != nil {
		return err
	}
	user.Avatar = url

	return ctx.JSON(http.StatusCreated, user)
}

func NewImagesHandler(e *echo.Echo, userUsecase user.Usecase, imagesUsecase images.Usecase) imagesHandler {
	handler := imagesHandler{userUsecase: userUsecase, imagesUsecase: imagesUsecase}
	loadImagesUrl := "/images/"

	api := e.Group("api/v1")

	loadImages := api.Group(loadImagesUrl)

	loadImages.POST("", handler.LoadCurrentUserAvatarHandler)

	return handler
}
