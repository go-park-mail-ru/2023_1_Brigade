package usecase

import (
	"context"
	"io"
	"os"
	"project/internal/monolithic_services/images"
	"project/internal/pkg/image_generation"
)

type usecase struct {
	imagesRepo images.Repository
}

func NewImagesUsecase(imagesRepo images.Repository) images.Usecase {
	return &usecase{imagesRepo: imagesRepo}
}

func (u usecase) UploadGeneratedImage(ctx context.Context, bucketName string, filename string, firstCharacterName string) error {
	err := image_generation.GenerateAvatar(firstCharacterName)
	if err != nil {
		return err
	}

	file, err := os.Open("../background.png")
	if err != nil {
		return err
	}

	err = u.UploadImage(ctx, file, bucketName, filename)
	if err != nil {
		return err
	}

	return nil
}

func (u usecase) UploadImage(ctx context.Context, file io.Reader, bucketName string, filename string) error {
	err := u.imagesRepo.UploadImage(ctx, file, bucketName, filename)
	if err != nil {
		return err
	}

	return err
}

func (u usecase) GetImage(ctx context.Context, bucketName string, filename string) (string, error) {
	url, err := u.imagesRepo.GetImage(ctx, bucketName, filename)
	if err != nil {
		return "", err
	}

	return url, nil
}
