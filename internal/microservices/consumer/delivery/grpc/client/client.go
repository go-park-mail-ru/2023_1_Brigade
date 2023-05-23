package client

import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	"project/internal/generated"
	consumer "project/internal/microservices/consumer/usecase"

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

func (c consumerServiceGRPCClient) StartConsumeMessages(ctx context.Context) {
	// nolint: errcheck
	c.consumerClient.StartConsumeMessages(ctx, new(empty.Empty))
}
