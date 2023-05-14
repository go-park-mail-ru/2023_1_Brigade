package server

import (
	"context"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	mockMessages "project/internal/microservices/messages/usecase/mocks"
	"project/internal/pkg/model_conversion"
	"testing"
)

//func TestServer_SwitchMessageType_OK(t *testing.T) {
//	ctl := gomock.NewController(t)
//	defer ctl.Finish()
//
//	bytes := []byte(`"msg":"hello world"}`)
//
//	grpcServer := grpc.NewServer()
//
//	messagesUsecase := mockMessages.NewMockUsecase(ctl)
//
//	messagesSerivce := NewMessagesServiceGRPCServer(grpcServer, messagesUsecase)
//
//	messagesUsecase.EXPECT().SwitchMessageType(context.TODO(), bytes).Return(nil).Times(1)
//
//	_, err := messagesSerivce.SwitchMessageType(context.TODO(), model_conversion.FromBytesToProtoBytes(bytes))
//
//	require.NoError(t, err)
//}

func TestServer_PutInProducer_OK(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	bytes := []byte(`{
		Id:         "sfdst3523rrfgdxxf0",
		Type:       config.Create,
		Body:       "Hello world!",
		AuthorId:   1,
		ChatID:     1,
		ReceiverID: 2,
	}`)

	grpcServer := grpc.NewServer()

	messagesUsecase := mockMessages.NewMockUsecase(ctl)

	messagesSerivce := NewMessagesServiceGRPCServer(grpcServer, messagesUsecase)

	messagesUsecase.EXPECT().PutInProducer(context.TODO(), bytes).Return(nil).Times(1)

	_, err := messagesSerivce.PutInProducer(context.TODO(), model_conversion.FromBytesToProtoBytes(bytes))

	require.NoError(t, err)
}

//func TestServer_PullFromConsumer_OK(t *testing.T) {
//	ctl := gomock.NewController(t)
//	defer ctl.Finish()
//
//	expectedBytes := []byte(`"msg":"hello world"}`)
//
//	grpcServer := grpc.NewServer()
//
//	messagesUsecase := mockMessages.NewMockUsecase(ctl)
//
//	messagesSerivce := NewMessagesServiceGRPCServer(grpcServer, messagesUsecase)
//
//	messagesUsecase.EXPECT().PullFromConsumer(context.TODO()).Return(expectedBytes, nil).Times(1)
//
//	bytes, err := messagesSerivce.PullFromConsumer(context.TODO(), nil)
//
//	require.NoError(t, err)
//	require.Equal(t, expectedBytes, model_conversion.FromProtoBytesToBytes(bytes))
//}
