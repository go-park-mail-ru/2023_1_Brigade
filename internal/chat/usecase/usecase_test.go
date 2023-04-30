package usecase

import (
	"context"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	chatMock "project/internal/chat/repository/mocks"
	messageMock "project/internal/messages/repository/mocks"
	"project/internal/model"
	userMock "project/internal/user/repository/mocks"
	"testing"
)

func Test_CreateChat_OK(t *testing.T) {
	var members []model.User

	newChat := model.CreateChat{}

	createdChat := model.Chat{
		Avatar:   "",
		Messages: []model.Message{},
		Members:  members,
	}

	expectedChat := model.Chat{
		Id:       1,
		Avatar:   "",
		Messages: []model.Message{},
		Members:  members,
	}

	ctl := gomock.NewController(t)
	defer ctl.Finish()

	chatRepository := chatMock.NewMockRepository(ctl)
	userRepository := userMock.NewMockRepository(ctl)
	messagesRepository := messageMock.NewMockRepository(ctl)
	usecase := NewChatUsecase(chatRepository, userRepository, messagesRepository)

	chatRepository.EXPECT().CreateChat(context.Background(), createdChat).Return(expectedChat, nil).Times(1)

	chat, err := usecase.CreateChat(context.TODO(), newChat, 1)

	require.NoError(t, err)
	require.Equal(t, chat, expectedChat)
}

func Test_DeleteChat_OK(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	chatRepository := chatMock.NewMockRepository(ctl)
	userRepository := userMock.NewMockRepository(ctl)
	messagesRepository := messageMock.NewMockRepository(ctl)
	usecase := NewChatUsecase(chatRepository, userRepository, messagesRepository)

	chatRepository.EXPECT().DeleteChatById(context.Background(), uint64(1)).Return(nil).Times(1)

	err := usecase.DeleteChatById(context.TODO(), uint64(1))

	require.NoError(t, err)
}

func Test_GetChat_OK(t *testing.T) {
	expectedChat := model.Chat{
		Id:    1,
		Title: "",
	}

	ctl := gomock.NewController(t)
	defer ctl.Finish()

	chatRepository := chatMock.NewMockRepository(ctl)
	userRepository := userMock.NewMockRepository(ctl)
	messagesRepository := messageMock.NewMockRepository(ctl)
	usecase := NewChatUsecase(chatRepository, userRepository, messagesRepository)

	chatRepository.EXPECT().GetChatById(context.Background(), uint64(1)).Return(expectedChat, nil).Times(1)
	chatRepository.EXPECT().GetChatMembersByChatId(context.Background(), uint64(1)).Return([]model.ChatMembers{}, nil).Times(1)
	messagesRepository.EXPECT().GetChatMessages(context.Background(), uint64(1)).Times(1)

	chat, err := usecase.GetChatById(context.TODO(), uint64(1))

	require.NoError(t, err)
	require.Equal(t, chat, expectedChat)
}

func Test_GetListUserChats_OK(t *testing.T) {
	userChats := []model.ChatMembers{
		{
			ChatId:   1,
			MemberId: 1,
		},
	}
	chat := model.Chat{
		Id: 1,
	}
	chatMembers := []model.ChatMembers{
		{
			ChatId:   1,
			MemberId: 1,
		},
	}
	expectedChat := []model.ChatInListUser{
		{
			Id: 1,
			Members: []model.User{
				{
					Id: 1,
				},
			},
		},
	}

	ctl := gomock.NewController(t)
	defer ctl.Finish()

	chatRepository := chatMock.NewMockRepository(ctl)
	userRepository := userMock.NewMockRepository(ctl)
	messagesRepository := messageMock.NewMockRepository(ctl)
	usecase := NewChatUsecase(chatRepository, userRepository, messagesRepository)

	chatRepository.EXPECT().GetChatsByUserId(context.Background(), uint64(1)).Return(userChats, nil).Times(1)
	chatRepository.EXPECT().GetChatById(context.Background(), uint64(1)).Return(chat, nil).Times(1)
	chatRepository.EXPECT().GetChatMembersByChatId(context.Background(), uint64(1)).Return(chatMembers, nil).Times(1)
	userRepository.EXPECT().GetUserById(context.Background(), uint64(1)).Return(model.AuthorizedUser{Id: 1}, nil).Times(1)
	messagesRepository.EXPECT().GetLastChatMessage(context.Background(), uint64(1)).Return(model.Message{}, nil).Times(1)

	listChats, err := usecase.GetListUserChats(context.TODO(), uint64(1))

	require.NoError(t, err)
	require.Equal(t, expectedChat, listChats)
}
