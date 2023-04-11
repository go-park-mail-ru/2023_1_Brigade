package usecase

import (
	"context"
	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/require"
	chatMock "project/internal/chat/repository/mocks"
	messageMock "project/internal/messages/repository/mocks"
	"project/internal/model"
	userMock "project/internal/user/repository/mocks"
	"testing"
)

//func Test_CreateChat_OK(t *testing.T) {
//	newChat := model.CreateChat{
//		Title:   "new_chat",
//		Members: []uint64{0},
//	}
//
//	dbChat := model.Chat{
//		Id:      0,
//		Title:   "new_chat",
//		Avatar:  configs.DefaultAvatarUrl,
//		Members: []model.User{{}},
//	}
//
//	expectedChat := model.Chat{
//		Id:      1,
//		Title:   "new_chat",
//		Avatar:  configs.DefaultAvatarUrl,
//		Members: []model.User{{}},
//	}
//
//	ctl := gomock.NewController(t)
//	defer ctl.Finish()
//
//	var ctx echo.Context
//	chatRepository := chatMock.NewMockRepository(ctl)
//	userRepository := userMock.NewMockRepository(ctl)
//	messagesRepository := messageMock.NewMockRepository(ctl)
//	usecase := NewChatUsecase(chatRepository, userRepository, messagesRepository)
//
//	userRepository.EXPECT().GetUserById(context.Background(), uint64(0)).Return(model.AuthorizedUser{}, nil).Times(1)
//	chatRepository.EXPECT().CreateChat(context.Background(), dbChat).Return(expectedChat, nil).Times(1)
//
//	chat, err := usecase.CreateChat(ctx, newChat)
//
//	require.NoError(t, err)
//	require.Equal(t, chat, expectedChat)
//}

func Test_DeleteChat_OK(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	var ctx echo.Context
	chatRepository := chatMock.NewMockRepository(ctl)
	userRepository := userMock.NewMockRepository(ctl)
	messagesRepository := messageMock.NewMockRepository(ctl)
	usecase := NewChatUsecase(chatRepository, userRepository, messagesRepository)

	chatRepository.EXPECT().DeleteChatById(context.Background(), uint64(1)).Return(nil).Times(1)

	err := usecase.DeleteChatById(ctx, uint64(1))

	require.NoError(t, err)
}

func Test_GetChat_OK(t *testing.T) {
	expectedChat := model.Chat{
		Id:    1,
		Title: "",
	}

	ctl := gomock.NewController(t)
	defer ctl.Finish()

	var ctx echo.Context
	chatRepository := chatMock.NewMockRepository(ctl)
	userRepository := userMock.NewMockRepository(ctl)
	messagesRepository := messageMock.NewMockRepository(ctl)
	usecase := NewChatUsecase(chatRepository, userRepository, messagesRepository)

	chatRepository.EXPECT().GetChatById(context.Background(), uint64(1)).Return(expectedChat, nil).Times(1)
	chatRepository.EXPECT().GetChatMembersByChatId(context.Background(), uint64(1)).Return([]model.ChatMembers{}, nil).Times(1)
	messagesRepository.EXPECT().GetChatMessages(context.Background(), uint64(1)).Times(1)

	chat, err := usecase.GetChatById(ctx, uint64(1))

	require.NoError(t, err)
	require.Equal(t, chat, expectedChat)
}

func Test_GetListUserChats_OK(t *testing.T) {
	var expectedChat []model.ChatInListUser

	ctl := gomock.NewController(t)
	defer ctl.Finish()

	var ctx echo.Context
	chatRepository := chatMock.NewMockRepository(ctl)
	userRepository := userMock.NewMockRepository(ctl)
	messagesRepository := messageMock.NewMockRepository(ctl)
	usecase := NewChatUsecase(chatRepository, userRepository, messagesRepository)

	chatRepository.EXPECT().GetChatsByUserId(context.Background(), uint64(1)).Return([]model.ChatMembers{}, nil).Times(1)

	chat, err := usecase.GetListUserChats(ctx, uint64(1))

	require.NoError(t, err)
	require.Equal(t, chat, expectedChat)
}
