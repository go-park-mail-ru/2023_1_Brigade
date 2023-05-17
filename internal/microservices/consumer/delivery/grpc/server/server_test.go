package server

import (
	"context"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	mockConsumer "project/internal/microservices/consumer/usecase/mocks"
	"testing"
)

func TestServer_StartConsumeMessages_OK(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	grpcServer := grpc.NewServer()

	consumerUsecase := mockConsumer.NewMockUsecase(ctl)

	consumerService := NewConsumerServiceGRPCServer(grpcServer, consumerUsecase)

	consumerUsecase.EXPECT().StartConsumeMessages(context.TODO()).AnyTimes()

	_, err := consumerService.StartConsumeMessages(context.TODO(), nil)

	require.NoError(t, err)
}
