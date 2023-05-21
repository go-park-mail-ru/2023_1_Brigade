package http

import (
	"context"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
	"net/http"
	"project/internal/config"
	"project/internal/microservices/user"
	"project/internal/model"
	"project/internal/monolithic_services/images"
	"strconv"
)

type imagesHandler struct {
	userUsecase   user.Usecase
	imagesUsecase images.Usecase
}

func (h imagesHandler) UploadUserAvatarsHandler(ctx echo.Context) error {
	session := ctx.Get("session").(model.Session)
	userID := strconv.FormatUint(session.UserId, 10)

	maxSize := int64(64 << 20)
	err := ctx.Request().ParseMultipartForm(maxSize)
	if err != nil {
		return err
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

	err = h.imagesUsecase.UploadImage(context.TODO(), file, config.UserAvatarsBucket, string(userID))
	if err != nil {
		return err
	}

	user, err := h.userUsecase.GetUserById(context.TODO(), session.UserId)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusCreated, user)
}

func (h imagesHandler) UploadChatAvatarsHandler(ctx echo.Context) error {
	session := ctx.Get("session").(model.Session)
	userID := strconv.FormatUint(session.UserId, 10)

	maxSize := int64(64 << 20)
	err := ctx.Request().ParseMultipartForm(maxSize)
	if err != nil {
		return err
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

	err = h.imagesUsecase.UploadImage(context.TODO(), file, config.UserAvatarsBucket, string(userID))
	if err != nil {
		return err
	}

	user, err := h.userUsecase.GetUserById(context.TODO(), session.UserId)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusCreated, user)

	//return nil
	//session := ctx.Get("session").(model.Session)
	//	userID := session.UserId
	//
	//	maxSize := int64(64 << 20)
	//	err := ctx.Request().ParseMultipartForm(maxSize)
	//	if err != nil {
	//		return err
	//	}
	//
	//	file, header, err := ctx.Request().FormFile("image")
	//	if err != nil {
	//		return err
	//	}
	//
	//	defer func() {
	//		err := file.Close()
	//		if err != nil {
	//			log.Error(err)
	//		}
	//	}()
	//
	//	url, err := h.imagesUsecase.UploadImage(context.TODO(), file, header.Filename, userID)
	//	if err != nil {
	//		return err
	//	}
	//
	//	user, err := h.userUsecase.GetUserById(context.TODO(), session.UserId)
	//	if err != nil {
	//		return err
	//	}
	//	user.Avatar = url
	//
	//	return ctx.JSON(http.StatusCreated, user)
}

func (h imagesHandler) UploadChatImagesHandler(ctx echo.Context) error {
	//session := ctx.Get("session").(model.Session)
	//userID := strconv.FormatUint(session.UserId, 10)

	maxSize := int64(64 << 20)
	err := ctx.Request().ParseMultipartForm(maxSize)
	if err != nil {
		return err
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
	err = h.imagesUsecase.UploadImage(context.TODO(), file, config.ChatImagesBucket, filename)
	if err != nil {
		return err
	}

	url, err := h.imagesUsecase.GetImage(context.TODO(), config.ChatImagesBucket, filename)
	if err != nil {
		return err
	}

	//user, err := h.userUsecase.GetUserById(context.TODO(), session.UserId)
	//if err != nil {
	//	return err
	//}

	return ctx.JSON(http.StatusCreated, url)
	//return nil
	//session := ctx.Get("session").(model.Session)
	//userID := session.UserId
	//
	//maxSize := int64(64 << 20)
	//err := ctx.Request().ParseMultipartForm(maxSize)
	//if err != nil {
	//	return err
	//}
	//
	//file, header, err := ctx.Request().FormFile("image")
	//if err != nil {
	//	return err
	//}
	//
	//defer func() {
	//	err := file.Close()
	//	if err != nil {
	//		log.Error(err)
	//	}
	//}()
	//
	//url, err := h.imagesUsecase.UploadImage(context.TODO(), file, header.Filename, userID)
	//if err != nil {
	//	return err
	//}
	//
	//user, err := h.userUsecase.GetUserById(context.TODO(), session.UserId)
	//if err != nil {
	//	return err
	//}
	//user.Avatar = url
	//
	//return ctx.JSON(http.StatusCreated, user)
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
