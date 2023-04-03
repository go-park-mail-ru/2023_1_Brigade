package usecase

import (
	"context"
	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/require"
	chatMock "project/internal/chat/repository/mocks"
	"project/internal/model"
	userMock "project/internal/user/repository/mocks"
	"testing"
)

func Test_CreateChat_OK(t *testing.T) {
	newChat := model.CreateChat{
		Title:   "new_chat",
		Members: []uint64{0},
	}

	dbChat := model.Chat{
		Id:      0,
		Title:   "new_chat",
		Members: []model.User{{}},
	}

	expectedChat := model.Chat{
		Id:      1,
		Title:   "new_chat",
		Members: []model.User{{}},
	}

	ctl := gomock.NewController(t)
	defer ctl.Finish()

	var ctx echo.Context
	chatRepository := chatMock.NewMockRepository(ctl)
	userRepository := userMock.NewMockRepository(ctl)
	usecase := NewChatUsecase(chatRepository, userRepository)

	userRepository.EXPECT().GetUserById(context.Background(), uint64(0)).Return(model.User{}, nil).Times(1)
	chatRepository.EXPECT().CreateChat(context.Background(), dbChat).Return(expectedChat, nil).Times(1)

	chat, err := usecase.CreateChat(ctx, newChat)

	require.NoError(t, err)
	require.Equal(t, chat, expectedChat)
}

func Test_DeleteChat_OK(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	var ctx echo.Context
	chatRepository := chatMock.NewMockRepository(ctl)
	userRepository := userMock.NewMockRepository(ctl)
	usecase := NewChatUsecase(chatRepository, userRepository)

	chatRepository.EXPECT().DeleteChatById(context.Background(), uint64(1)).Return(nil).Times(1)

	err := usecase.DeleteChatById(ctx, uint64(1))

	require.NoError(t, err)
}

func Test_GetChat_OK(t *testing.T) {
	expectedChat := model.Chat{
		Id:      1,
		Title:   "",
		Members: []model.User{{}},
	}

	ctl := gomock.NewController(t)
	defer ctl.Finish()

	var ctx echo.Context
	chatRepository := chatMock.NewMockRepository(ctl)
	userRepository := userMock.NewMockRepository(ctl)
	usecase := NewChatUsecase(chatRepository, userRepository)

	chatRepository.EXPECT().GetChatById(context.Background(), uint64(1)).Return(expectedChat, nil).Times(1)

	chat, err := usecase.GetChatById(ctx, uint64(1))

	require.NoError(t, err)
	require.Equal(t, chat, expectedChat)
}
