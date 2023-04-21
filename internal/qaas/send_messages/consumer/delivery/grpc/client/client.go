package client

import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"project/internal/qaas/send_messages/consumer"
	"project/internal/qaas/send_messages/consumer/delivery/grpc/service"
)

type consumersServiceGRPCClient struct {
	consumersClient service.ConsumersClient
}

func NewConsumersServiceGRPSClient(con *grpc.ClientConn) consumer.Usecase {
	return &consumersServiceGRPCClient{
		consumersClient: service.NewConsumersClient(con),
	}
}

func (c consumersServiceGRPCClient) ConsumeMessage() []byte {
	msg, err := c.consumersClient.ConsumeMessage(context.TODO(), &empty.Empty{})
	if err != nil {
		return nil
	}

	return msg.Bytes
}

func (c consumersServiceGRPCClient) StartConsumeMessages() {
	_, err := c.consumersClient.StartConsumeMessages(context.TODO(), &empty.Empty{})
	log.Error(err)
}
