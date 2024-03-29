package server

import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	"net"
	"project/internal/generated"
	consumer "project/internal/microservices/consumer/usecase"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

type consumerServiceGRPCServer struct {
	grpcServer      *grpc.Server
	consumerUsecase consumer.Usecase
}

func NewConsumerServiceGRPCServer(grpcServer *grpc.Server, consumerUsecase consumer.Usecase) *consumerServiceGRPCServer {
	return &consumerServiceGRPCServer{
		grpcServer:      grpcServer,
		consumerUsecase: consumerUsecase,
	}
}

func (c *consumerServiceGRPCServer) StartGRPCServer(listenURL string) error {
	lis, err := net.Listen("tcp", listenURL)
	if err != nil {
		return err
	}

	generated.RegisterConsumerServer(c.grpcServer, c)

	return c.grpcServer.Serve(lis)
}

func (c *consumerServiceGRPCServer) StartConsumeMessages(ctx context.Context, _ *emptypb.Empty) (*emptypb.Empty, error) {
	go func() {
		c.consumerUsecase.StartConsumeMessages(ctx)
	}()
	return new(empty.Empty), nil
}
