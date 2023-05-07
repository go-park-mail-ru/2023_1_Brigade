package consumer

import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	"project/internal/generated"
	consumer "project/internal/qaas/send_messages/consumer/usecase"

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
	bytes, err := c.consumerClient.ConsumeMessage(ctx, new(empty.Empty))
	if err != nil {
		return []byte{}
	}

	return bytes.Bytes
}

func (c consumerServiceGRPCClient) StartConsumeMessages(ctx context.Context) {
	c.consumerClient.StartConsumeMessages(ctx, new(empty.Empty))
}
