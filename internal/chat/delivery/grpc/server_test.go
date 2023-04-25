package grpc

import (
	"context"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	chatMock "project/internal/chat/usecase/mocks"
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

	chatUsecase.EXPECT().GetChatById(context.TODO(), uint64(1)).Return(model.Chat{}, nil).Times(1)

	chat, err := chatsService.GetChatById(context.TODO(), model_conversion.FromChatIDToProtoChatID(1))

	require.NoError(t, err)
	require.Equal(t, &protobuf.Chat{
		Members:  []*protobuf.User{},
		Messages: []*protobuf.Message{},
	}, chat)
}
