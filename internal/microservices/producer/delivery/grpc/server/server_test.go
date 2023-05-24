package server

import (
	"context"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	mockProducer "project/internal/microservices/producer/usecase/mocks"
	"project/internal/model"
	"project/internal/pkg/model_conversion"
	"testing"
)

func TestServer_ProduceMessage_OK(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	msg := model.ProducerMessage{
		Attachments: []model.File{},
	}

	grpcServer := grpc.NewServer()

	producerUsecase := mockProducer.NewMockUsecase(ctl)

	producerService := NewProducerServiceGRPCServer(grpcServer, producerUsecase)

	producerUsecase.EXPECT().ProduceMessage(context.TODO(), msg).Return(nil).Times(1)

	_, err := producerService.ProduceMessage(context.TODO(), model_conversion.FromProducerMessageToProtoProducerMessage(msg))

	require.NoError(t, err)
}
