package server

import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
	"net"
	protobuf "project/internal/model/generated"
	"project/internal/qaas/send_messages/producer"
	"project/internal/qaas/send_messages/producer/delivery/grpc/service"
)

type producersServiceGRPCServer struct {
	grpcServer      *grpc.Server
	producerUsecase producer.Usecase
}

func NewProducersServiceGRPCServer(grpcServer *grpc.Server, producerUsecase producer.Usecase) *producersServiceGRPCServer {
	return &producersServiceGRPCServer{
		grpcServer:      grpcServer,
		producerUsecase: producerUsecase,
	}
}

func (c *producersServiceGRPCServer) StartGRPCServer(listenURL string) error {
	lis, err := net.Listen("tcp", listenURL)
	if err != nil {
		return err
	}

	service.RegisterProducersServer(c.grpcServer, c)

	return c.grpcServer.Serve(lis)
}

func (c *producersServiceGRPCServer) ProduceMessage(ctx context.Context, bytes *protobuf.Bytes) (*empty.Empty, error) {
	err := c.producerUsecase.ProduceMessage(bytes.Bytes)
	return nil, err
}
