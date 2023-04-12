package http

import (
	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
	"project/internal/images"
)

type imagesHandler struct {
	imagesUsecase images.Usecase
}

func (h imagesHandler) LoadCurrentUserAvatarHandler(ctx echo.Context) error {
//	file, err := ctx.FormFile("image")
//	if err != nil {
//		log.Warn(err)
//	}
//	
//	log.Warn(file, 12131313)

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

	return nil

	//defer func() {
	//	err := file.
	//	if err != nil {
	//		log.Error(err)
	//	}
	//}()
	//session := ctx.Get("session").(model.Session)
	//userID := session.UserId
	//a, err := file.Open()
	//a.
	//url, err := h.imagesUsecase.LoadImage(ctx, file, header.Filename, userID)
	//if err != nil {
	//	return err
	//}

	//return ctx.JSON(http.StatusCreated, url)
	
	//err := ctx.Request().ParseMultipartForm(32 << 20) // максимальный размер файла 32 МБ
	//if err != nil {
	//	//http.Error(ctx.Response().Writer, err.Error(), http.StatusBadRequest)
	//	log.Error(err)
	//	return nil
	//}
	//
	//for _, fileHeaders := range ctx.Request().MultipartForm.File {
	//	for _, fileHeader := range fileHeaders {
	//		file, err := fileHeader.Open()
	//		if err != nil {
	//			//http.Error(ctx.Response().Writer, err.Error(), http.StatusBadRequest)
	//			//return
	//			log.Error(err)
	//		}
	//		defer file.Close()
	//
	//		out, err := os.Create(fileHeader.Filename)
	//		if err != nil {
	//			//http.Error(ctx.Response().Writer, err.Error(), http.StatusInternalServerError)
	//			//return
	//			log.Error(err)
	//		}
	//		defer out.Close()
	//
	//		_, err = io.Copy(out, file)
	//		if err != nil {
	//			//http.Error(ctx.Response().Writer, err.Error(), http.StatusInternalServerError)
	//			//return
	//			log.Error(err)
	//		}
	//	}
	//}

	//fmt.Fprintln(w, "Files uploaded successfully")
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
	//url, err := h.imagesUsecase.LoadImage(ctx, file, header.Filename, userID)
	//if err != nil {
	//	return err
	//}
	//
	//return ctx.JSON(http.StatusCreated, url)
}

func NewImagesHandler(e *echo.Echo, imagesUsecase images.Usecase) imagesHandler {
	handler := imagesHandler{imagesUsecase: imagesUsecase}
	loadImagesUrl := "/images/"

	api := e.Group("api/v1")

	loadImages := api.Group(loadImagesUrl)

	loadImages.POST("", handler.LoadCurrentUserAvatarHandler)

	return handler
}

