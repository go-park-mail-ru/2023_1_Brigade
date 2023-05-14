package usecase

import (
	"context"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"project/internal/config"
	imagesMock "project/internal/monolithic_services/images/repository/mocks"
	"testing"
)

func Test_GetChat_OK(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	filename := "photo"

	imagesRepository := imagesMock.NewMockRepository(ctl)

	usecase := NewImagesUsecase(imagesRepository)

	imagesRepository.EXPECT().UploadImage(context.TODO(), nil, config.UserAvatarsBucket, filename+`.png`).Return(nil).Times(1)

	err := usecase.UploadImage(context.TODO(), nil, config.UserAvatarsBucket, filename)

	require.NoError(t, err)
}
