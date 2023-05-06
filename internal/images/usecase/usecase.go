package usecase

import (
	"context"
	"io"
	"os"
	"project/internal/images"
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

	file, err := os.Open("../../avatars/background.png")
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
	filename = filename + ".png"
	err := u.imagesRepo.UploadImage(ctx, file, bucketName, filename)
	return err
}

func (u usecase) GetImage(ctx context.Context, bucketName string, filename string) (string, error) {
	filename = filename + ".png"
	url, err := u.imagesRepo.GetImage(ctx, bucketName, filename)
	if err != nil {
		return "", err
	}

	return url, nil
}
