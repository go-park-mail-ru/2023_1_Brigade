package client

import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	log "github.com/sirupsen/logrus"
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
	_, err := c.consumerClient.StartConsumeMessages(ctx, new(empty.Empty))
	if err != nil {
		log.Error(err)
	}
}
