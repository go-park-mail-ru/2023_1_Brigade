package producer

import (
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

func (p producerServiceGRPCClient) ProduceMessage(message []byte) error {
	p.producerClient.ProduceMessage()
}
