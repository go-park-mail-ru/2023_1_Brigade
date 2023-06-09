package usecase

import (
	"context"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"project/internal/config"
	imagesMock "project/internal/monolithic_services/images/repository/mocks"
	myErrors "project/internal/pkg/errors"
	"testing"
)

func Test_UploadImage_OK(t *testing.T) {
	filename := uuid.NewString()

	ctl := gomock.NewController(t)
	defer ctl.Finish()

	imagesRepository := imagesMock.NewMockRepository(ctl)
	usecase := NewImagesUsecase(imagesRepository)

	imagesRepository.EXPECT().UploadImage(context.TODO(), nil, config.UserAvatarsBucket, filename).Return(nil).Times(1)

	err := usecase.UploadImage(context.TODO(), nil, config.UserAvatarsBucket, filename)

	require.NoError(t, err)
}

func Test_UploadImage_Error(t *testing.T) {
	filename := uuid.NewString()

	ctl := gomock.NewController(t)
	defer ctl.Finish()

	imagesRepository := imagesMock.NewMockRepository(ctl)
	usecase := NewImagesUsecase(imagesRepository)

	imagesRepository.EXPECT().UploadImage(context.TODO(), nil, config.UserAvatarsBucket, filename).Return(myErrors.ErrInternal).Times(1)

	err := usecase.UploadImage(context.TODO(), nil, config.UserAvatarsBucket, filename)

	require.Error(t, err, myErrors.ErrInternal)
}

func Test_GetImage_OK(t *testing.T) {
	filename := uuid.NewString()
	expectedUrl := "vk.com"

	ctl := gomock.NewController(t)
	defer ctl.Finish()

	imagesRepository := imagesMock.NewMockRepository(ctl)
	usecase := NewImagesUsecase(imagesRepository)

	imagesRepository.EXPECT().GetImage(context.TODO(), config.UserAvatarsBucket, filename).Return(expectedUrl, nil).Times(1)

	url, err := usecase.GetImage(context.TODO(), config.UserAvatarsBucket, filename)
	require.NoError(t, err)
	require.Equal(t, expectedUrl, url)
}

func Test_GetImage_Error(t *testing.T) {
	filename := uuid.NewString()

	ctl := gomock.NewController(t)
	defer ctl.Finish()

	imagesRepository := imagesMock.NewMockRepository(ctl)
	usecase := NewImagesUsecase(imagesRepository)

	imagesRepository.EXPECT().GetImage(context.TODO(), config.UserAvatarsBucket, filename).Return("", myErrors.ErrInternal).Times(1)

	_, err := usecase.GetImage(context.TODO(), config.UserAvatarsBucket, filename)
	require.Error(t, err, myErrors.ErrInternal)
}
