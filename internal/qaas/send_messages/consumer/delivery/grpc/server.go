package grpc

import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	log "github.com/sirupsen/logrus"
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
	log.Warn("CONSUME SERVER")
	msg := c.consumerUsecase.ConsumeMessage(ctx)
	return &generated.Bytes{Bytes: msg}, nil
}

func (c *consumerServiceGRPCServer) StartConsumeMessages(ctx context.Context, _ *emptypb.Empty) (*emptypb.Empty, error) {
	log.Warn("CONSUME SERVER START MESSAGE")
	c.consumerUsecase.StartConsumeMessages(ctx)
	return new(empty.Empty), nil
}
