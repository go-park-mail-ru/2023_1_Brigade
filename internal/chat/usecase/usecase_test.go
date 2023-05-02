package usecase

import (
	"context"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	chatMock "project/internal/chat/repository/mocks"
	"project/internal/configs"
	messageMock "project/internal/messages/repository/mocks"
	"project/internal/model"
	"project/internal/pkg/model_conversion"
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

func Test_EditChat_OK(t *testing.T) {
	user := model.AuthorizedUser{
		Id:       0,
		Avatar:   "",
		Nickname: "marcussss",
		Email:    "marcussss@gmail.com",
		Status:   "Привет, я использую технограм!",
		Password: "hashedPassword",
	}

	editChat := model.EditChat{
		Id:      1,
		Type:    configs.Chat,
		Title:   "title",
		Members: []uint64{1},
	}

	dbChat := model.DBChat{
		Id:     editChat.Id,
		Type:   editChat.Type,
		Title:  editChat.Title,
		Avatar: "",
	}
	expectedChat := model.Chat{
		Id:     dbChat.Id,
		Type:   dbChat.Type,
		Title:  dbChat.Title,
		Avatar: dbChat.Avatar,
		Members: []model.User{
			model_conversion.FromAuthorizedUserToUser(user),
		},
	}

	ctl := gomock.NewController(t)
	defer ctl.Finish()

	chatRepository := chatMock.NewMockRepository(ctl)
	userRepository := userMock.NewMockRepository(ctl)
	messagesRepository := messageMock.NewMockRepository(ctl)
	usecase := NewChatUsecase(chatRepository, userRepository, messagesRepository)

	chatRepository.EXPECT().UpdateChatById(context.TODO(), editChat.Title, editChat.Id).Return(dbChat, nil).Times(1)
	chatRepository.EXPECT().DeleteChatMembers(context.TODO(), editChat.Id).Return(nil).Times(1)
	userRepository.EXPECT().CheckExistUserById(context.TODO(), uint64(1)).Return(nil).Times(1)
	chatRepository.EXPECT().AddUserInChatDB(context.TODO(), editChat.Id, uint64(1)).Return(nil).Times(1)
	userRepository.EXPECT().GetUserById(context.Background(), uint64(1)).Return(user, nil).Times(1)

	chat, err := usecase.EditChat(context.TODO(), editChat)

	require.NoError(t, err)
	require.Equal(t, expectedChat, chat)
}

func Test_GetSearchChatsMessagesChannels_OK(t *testing.T) {
	userID := uint64(1)
	string := "ba"
	expectedChats := model.FoundedChatsMessagesChannels{}

	ctl := gomock.NewController(t)
	defer ctl.Finish()

	chatRepository := chatMock.NewMockRepository(ctl)
	userRepository := userMock.NewMockRepository(ctl)
	messagesRepository := messageMock.NewMockRepository(ctl)
	usecase := NewChatUsecase(chatRepository, userRepository, messagesRepository)

	chatRepository.EXPECT().GetSearchChannels(context.TODO(), string, userID).Return([]model.Chat{}, nil).Times(1)
	chatRepository.EXPECT().GetChatMembersByChatId(context.TODO(), userID).Return([]model.ChatMembers{}, nil).Times(1)

	chats, err := usecase.GetSearchChatsMessagesChannels(context.TODO(), userID, string)

	require.NoError(t, err)
	require.Equal(t, expectedChats, chats)
}
