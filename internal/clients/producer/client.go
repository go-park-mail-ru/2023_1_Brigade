package producer

import (
	"context"
	"project/internal/generated"
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

func (p producerServiceGRPCClient) ProduceMessage(ctx context.Context, message []byte) error {
	_, err := p.producerClient.ProduceMessage(ctx, &generated.Bytes{
		Bytes: message,
	})
	return err
}
