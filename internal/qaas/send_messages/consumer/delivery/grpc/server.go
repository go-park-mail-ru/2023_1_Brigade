package grpc

import (
	"context"
	"net"
	"project/internal/generated"
	"project/internal/qaas/send_messages/consumer"

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

func (c *consumerServiceGRPCServer) ConsumeMessage(ctx context.Context, _ *emptypb.Empty) (*generated.Bytes, error) {
	return &generated.Bytes{
		Bytes: c.consumerUsecase.ConsumeMessage(ctx),
	}, nil
}

func (c *consumerServiceGRPCServer) StartConsumeMessages(ctx context.Context, _ *emptypb.Empty) (*emptypb.Empty, error) {
	c.consumerUsecase.StartConsumeMessages(ctx)
}
