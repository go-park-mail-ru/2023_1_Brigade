package usecase

import (
	"github.com/labstack/echo/v4"
	"mime/multipart"
	"project/internal/images"
)

type usecase struct {
	imagesRepo images.Repostiory
}

func NewChatUsecase(imagesRepo images.Repostiory) images.Usecase {
	return &usecase{imagesRepo: imagesRepo}
}

func (u *usecase) LoadImage(ctx echo.Context, file multipart.File, filename string, userID uint64) (string, error) {
	imageUrl, err := u.imagesRepo.LoadImage(ctx, file, filename, userID)
	if imageUrl == nil {
		return "", err
	}

	return imageUrl.String(), err
}
