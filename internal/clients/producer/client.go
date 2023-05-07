package producer

import (
	"context"
	"project/internal/generated"
	"project/internal/model"
	"project/internal/pkg/model_conversion"
	producer "project/internal/qaas/send_messages/producer/usecase"

	"google.golang.org/grpc"
)

type producerServiceGRPCClient struct {
	producerClient generated.ProducerClient
}

func NewProducerServiceGRPCClient(con *grpc.ClientConn) producer.Usecase {
	return &producerServiceGRPCClient{
		producerClient: generated.NewProducerClient(con),
	}
}

func (p producerServiceGRPCClient) ProduceMessage(ctx context.Context, message model.ProducerMessage) error {
	_, err := p.producerClient.ProduceMessage(ctx, model_conversion.FromProducerMessageToProtoProducerMessage(message))
	if err != nil {
		return err
	}

	return nil
}
