package consumer

import (
	"context"
	"project/internal/generated"
	"project/internal/qaas/send_messages/consumer"

	"github.com/golang/protobuf/ptypes/empty"

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
	bytes, err := c.consumerClient.ConsumeMessage(ctx, &empty.Empty{})
	if err != nil {
		return []byte{}
	}
	return bytes.Bytes
}

func (c consumerServiceGRPCClient) StartConsumeMessages(ctx context.Context) {
	c.consumerClient.StartConsumeMessages(ctx, &empty.Empty{})
}
