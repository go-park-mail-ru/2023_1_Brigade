package grpc

import (
	"context"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	chatMock "project/internal/chat/usecase/mocks"
	"project/internal/configs"
	protobuf "project/internal/generated"
	"project/internal/model"
	"project/internal/pkg/model_conversion"
	"testing"
)

func TestServer_GetChatByID_OK(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	grpcServer := grpc.NewServer()
	chatUsecase := chatMock.NewMockUsecase(ctl)

	chatsService := NewChatsServiceGRPCServer(grpcServer, chatUsecase)

	chatID := uint64(1)
	userID := uint64(1)

	expectedChat := model.Chat{
		Id:    1,
		Type:  configs.Group,
		Title: "chat",
		Members: []model.User{
			{
				Id: 1,
			},
			{
				Id: 2,
			},
		},
		Messages: []model.Message{},
	}

	chatUsecase.EXPECT().GetChatById(context.TODO(), chatID, userID).Return(expectedChat, nil).Times(1)

	chat, err := chatsService.GetChatById(context.TODO(), &protobuf.GetChatArguments{ChatID: chatID, UserID: userID})

	require.NoError(t, err)
	require.Equal(t, expectedChat, model_conversion.FromProtoChatToChat(chat))
}

func TestServer_EditChat_OK(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	grpcServer := grpc.NewServer()
	chatUsecase := chatMock.NewMockUsecase(ctl)

	chatsService := NewChatsServiceGRPCServer(grpcServer, chatUsecase)

	editChat := model.EditChat{
		Id:      1,
		Type:    configs.Group,
		Title:   "chat",
		Members: []uint64{1, 2},
	}

	expectedChat := model.Chat{
		Id:    1,
		Type:  configs.Group,
		Title: "chat",
		Members: []model.User{
			{
				Id: 1,
			},
			{
				Id: 2,
			},
		},
		Messages: []model.Message{},
	}

	chatUsecase.EXPECT().EditChat(context.TODO(), editChat).Return(expectedChat, nil).Times(1)

	chat, err := chatsService.EditChat(context.TODO(), model_conversion.FromEditChatToProtoEditChat(editChat))

	require.NoError(t, err)
	require.Equal(t, expectedChat, model_conversion.FromProtoChatToChat(chat))
}

func TestServer_CreateChat_OK(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	grpcServer := grpc.NewServer()
	chatUsecase := chatMock.NewMockUsecase(ctl)

	chatsService := NewChatsServiceGRPCServer(grpcServer, chatUsecase)

	userID := uint64(1)

	createChat := model.CreateChat{
		Type:    configs.Group,
		Title:   "chat",
		Members: []uint64{1, 2},
	}

	expectedChat := model.Chat{
		Id:    1,
		Type:  configs.Group,
		Title: "chat",
		Members: []model.User{
			{
				Id: 1,
			},
			{
				Id: 2,
			},
		},
		Messages: []model.Message{},
	}

	chatUsecase.EXPECT().CreateChat(context.TODO(), createChat, userID).Return(expectedChat, nil).Times(1)

	chat, err := chatsService.CreateChat(context.TODO(), &protobuf.CreateChatArguments{
		Chat:   model_conversion.FromCreateChatToProtoCreateChat(createChat),
		UserID: model_conversion.FromUserIDToProtoUserID(userID),
	})

	require.NoError(t, err)
	require.Equal(t, expectedChat, model_conversion.FromProtoChatToChat(chat))
}

func TestServer_DeleteChatByID_OK(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	grpcServer := grpc.NewServer()
	chatUsecase := chatMock.NewMockUsecase(ctl)

	chatsService := NewChatsServiceGRPCServer(grpcServer, chatUsecase)

	chatID := uint64(1)

	chatUsecase.EXPECT().DeleteChatById(context.TODO(), chatID).Return(nil).Times(1)

	_, err := chatsService.DeleteChatById(context.TODO(), model_conversion.FromChatIDToProtoChatID(1))

	require.NoError(t, err)
}

func TestServer_CheckExistUserInChat_OK(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	grpcServer := grpc.NewServer()
	chatUsecase := chatMock.NewMockUsecase(ctl)

	chatsService := NewChatsServiceGRPCServer(grpcServer, chatUsecase)

	userID := uint64(1)

	chat := model.Chat{
		Id:    1,
		Type:  configs.Group,
		Title: "chat",
		Members: []model.User{
			{
				Id: 1,
			},
			{
				Id: 2,
			},
		},
		Messages: []model.Message{},
	}

	chatUsecase.EXPECT().CheckExistUserInChat(context.TODO(), chat, userID).Return(nil).Times(1)

	_, err := chatsService.CheckExistUserInChat(context.TODO(), &protobuf.ExistChatArguments{
		Chat:   model_conversion.FromChatToProtoChat(chat),
		UserID: model_conversion.FromUserIDToProtoUserID(userID),
	})

	require.NoError(t, err)
}

func TestServer_GetListUserChats_OK(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	grpcServer := grpc.NewServer()
	chatUsecase := chatMock.NewMockUsecase(ctl)

	chatsService := NewChatsServiceGRPCServer(grpcServer, chatUsecase)

	userID := uint64(1)

	expectedListUserChats := []model.ChatInListUser{
		{
			Id:    1,
			Type:  configs.Group,
			Title: "chat",
			Members: []model.User{
				{
					Id: 1,
				},
				{
					Id: 2,
				},
			},
			LastMessageAuthor: model.User{Id: userID},
			LastMessage:       model.Message{},
		},
	}

	chatUsecase.EXPECT().GetListUserChats(context.TODO(), userID).Return(expectedListUserChats, nil).Times(1)

	listUserChats, err := chatsService.GetListUserChats(context.TODO(), model_conversion.FromUserIDToProtoUserID(userID))

	require.NoError(t, err)
	require.Equal(t, expectedListUserChats[0], model_conversion.FromProtoUserChatToUserChat(listUserChats.Chats[0]))
}

func TestServer_GetSearchChatsMessagesChannels_OK(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	grpcServer := grpc.NewServer()
	chatUsecase := chatMock.NewMockUsecase(ctl)

	chatsService := NewChatsServiceGRPCServer(grpcServer, chatUsecase)

	chatUsecase.EXPECT().GetSearchChatsMessagesChannels(context.TODO(), uint64(1), "string").Return(model.FoundedChatsMessagesChannels{}, nil).Times(1)

	_, err := chatsService.GetSearchChatsMessagesChannels(context.TODO(), &protobuf.SearchChatsArgumets{UserID: uint64(1), String_: "string"})

	require.NoError(t, err)
}
