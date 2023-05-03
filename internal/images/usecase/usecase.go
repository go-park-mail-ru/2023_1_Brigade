package usecase

import (
	"context"
	"mime/multipart"
	"project/internal/images"
)

type usecase struct {
	imagesRepo images.Repostiory
}

func NewImagesUsecase(imagesRepo images.Repostiory) images.Usecase {
	return &usecase{imagesRepo: imagesRepo}
}

func (u usecase) LoadImage(ctx context.Context, file multipart.File, filename string, userID uint64) (string, error) {
	imageUrl, err := u.imagesRepo.LoadImage(context.Background(), file, filename, userID)
	return imageUrl, err
}
