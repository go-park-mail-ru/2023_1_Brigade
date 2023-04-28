package producer

import (
	"context"
	log "github.com/sirupsen/logrus"
	"project/internal/generated"
	"project/internal/qaas/send_messages/producer"

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
	log.Warn("PRODUCER CLIENT")
	_, err := p.producerClient.ProduceMessage(ctx, &generated.Bytes{
		Bytes: message,
	})
	return err
}
