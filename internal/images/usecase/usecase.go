package usecase

import (
	"context"
	"io"
	"project/internal/images"
)

type usecase struct {
	imagesRepo images.Repostiory
}

func NewImagesUsecase(imagesRepo images.Repostiory) images.Usecase {
	return &usecase{imagesRepo: imagesRepo}
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

//func (u usecase) LoadImage(ctx context.Context, file multipart.File, filename string, userID uint64) (string, error) {
//	imageUrl, err := u.imagesRepo.LoadImage(context.Background(), file, filename, userID)
//	return imageUrl, err
//}
