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

func (u *usecase) LoadImage(ctx echo.Context, file multipart.File, filename string) (string, error) {
	err := u.imagesRepo.LoadImage(ctx, file, filename)
	if err != nil {
		return "", err
	}

	imageUrl, err := u.imagesRepo.GetImage(ctx, filename)

	return imageUrl.Path, err
}
