package usecase

import (
	"context"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	imagesMock "project/internal/images/repository/mocks"
	"testing"
)

func Test_GetChat_OK(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	filename := "photo.jpg"
	userID := uint64(1)
	expectedUrl := "https://techogramm.ru/avatars/ab4242fdfadc.jpg"

	imagesRepository := imagesMock.NewMockRepostiory(ctl)

	usecase := NewImagesUsecase(imagesRepository)

	imagesRepository.EXPECT().LoadImage(context.TODO(), nil, filename, userID).Return(expectedUrl, nil).Times(1)

	url, err := usecase.LoadImage(context.TODO(), nil, filename, userID)

	require.NoError(t, err)
	require.Equal(t, expectedUrl, url)
}
