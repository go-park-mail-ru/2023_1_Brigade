package http

import (
	"context"
	"encoding/json"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
	"net/http"
	"project/internal/config"
	"project/internal/microservices/user"
	"project/internal/model"
	"project/internal/monolithic_services/images"
	myErrors "project/internal/pkg/errors"
)

type imagesHandler struct {
	userUsecase   user.Usecase
	imagesUsecase images.Usecase
}

func (h imagesHandler) UploadUserAvatarsHandler(ctx echo.Context) error {
	maxSize := int64(64 << 18)
	err := ctx.Request().ParseMultipartForm(maxSize)
	if err != nil {
		return myErrors.ErrBigFileSize
	}

	file, _, err := ctx.Request().FormFile("image")
	if err != nil {
		return err
	}

	defer func() {
		err := file.Close()
		if err != nil {
			log.Error(err)
		}
	}()

	filename := uuid.NewString()
	err = h.imagesUsecase.UploadImage(context.TODO(), file, config.UserAvatarsBucket, filename)
	if err != nil {
		return err
	}

	url, err := h.imagesUsecase.GetImage(context.TODO(), config.UserAvatarsBucket, filename)
	if err != nil {
		return err
	}

	data, err := json.Marshal(url)
	if err != nil {
		return err
	}

	return ctx.JSONBlob(http.StatusCreated, data)
}

func (h imagesHandler) UploadChatAvatarsHandler(ctx echo.Context) error {
	maxSize := int64(64 << 18)
	err := ctx.Request().ParseMultipartForm(maxSize)
	if err != nil {
		return myErrors.ErrBigFileSize
	}

	file, _, err := ctx.Request().FormFile("image")
	if err != nil {
		return err
	}

	defer func() {
		err := file.Close()
		if err != nil {
			log.Error(err)
		}
	}()

	filename := uuid.NewString()
	err = h.imagesUsecase.UploadImage(context.TODO(), file, config.ChatAvatarsBucket, filename)
	if err != nil {
		return err
	}

	url, err := h.imagesUsecase.GetImage(context.TODO(), config.ChatAvatarsBucket, filename)
	if err != nil {
		return err
	}

	data, err := json.Marshal(url)
	if err != nil {
		return err
	}

	return ctx.JSONBlob(http.StatusCreated, data)
}

func (h imagesHandler) UploadChatImagesHandler(ctx echo.Context) error {
	maxSize := int64(64 << 18)
	err := ctx.Request().ParseMultipartForm(maxSize)
	if err != nil {
		return myErrors.ErrBigFileSize
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

	err = h.imagesUsecase.UploadImage(context.TODO(), file, config.ChatImagesBucket, header.Filename)
	if err != nil {
		return err
	}

	url, err := h.imagesUsecase.GetImage(context.TODO(), config.ChatImagesBucket, header.Filename)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusCreated, model.File{
		Url:  url,
		Name: header.Filename,
	})
}

func NewImagesHandler(e *echo.Echo, userUsecase user.Usecase, imagesUsecase images.Usecase) imagesHandler {
	handler := imagesHandler{userUsecase: userUsecase, imagesUsecase: imagesUsecase}
	uploadUserAvatarsUrl := "/images/user/"
	uploadChatAvatarsUrl := "/images/chat/"
	uploadChatImagesUrl := "/images/chat/images/"

	api := e.Group("api/v1")

	uploadUserAvatars := api.Group(uploadUserAvatarsUrl)
	uploadChatAvatars := api.Group(uploadChatAvatarsUrl)
	uploadChatImages := api.Group(uploadChatImagesUrl)

	uploadUserAvatars.POST("", handler.UploadUserAvatarsHandler)
	uploadChatAvatars.POST("", handler.UploadChatAvatarsHandler)
	uploadChatImages.POST("", handler.UploadChatImagesHandler)

	return handler
}
