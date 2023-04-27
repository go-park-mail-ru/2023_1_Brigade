package consumer

import (
	"context"
	"github.com/labstack/gommon/log"
	"project/internal/generated"
	"project/internal/qaas/send_messages/consumer"

	"google.golang.org/grpc"
)

type consumerServiceGRPCClient struct {
	consumerClient generated.ConsumerClient
}

func NewConsumerServiceGRPCClient(con *grpc.ClientConn) consumer.Usecase {
	return &consumerServiceGRPCClient{
		consumerClient: generated.NewConsumerClient(con),
	}
}

func (c consumerServiceGRPCClient) ConsumeMessage(ctx context.Context) []byte {
	bytes, err := c.consumerClient.ConsumeMessage(ctx, nil)
	if err != nil {
		log.Error("client consumer error: ", err)
	}
	if err != nil {
		return []byte{}
	}

	return bytes.Bytes
}

func (c consumerServiceGRPCClient) StartConsumeMessages(ctx context.Context) {
	c.consumerClient.StartConsumeMessages(ctx, nil)
}
