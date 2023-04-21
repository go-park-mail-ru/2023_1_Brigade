package client

import (
	"context"
	"google.golang.org/grpc"
	protobuf "project/internal/model/generated"
	"project/internal/qaas/send_messages/producer"
	"project/internal/qaas/send_messages/producer/delivery/grpc/service"
)

type producersServiceGRPCClient struct {
	producersClient service.ProducersClient
}

func NewConsumersServiceGRPSClient(con *grpc.ClientConn) producer.Usecase {
	return &producersServiceGRPCClient{
		producersClient: service.NewProducersClient(con),
	}
}

func (p producersServiceGRPCClient) ProduceMessage(message []byte) error {
	_, err := p.producersClient.ProduceMessage(context.TODO(), &protobuf.Bytes{Bytes: message})
	return err
}
