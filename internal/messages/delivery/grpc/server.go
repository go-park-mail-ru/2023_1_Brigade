package grpc

import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
	"net"
	"project/internal/generated"
	"project/internal/messages"
	"project/internal/pkg/model_conversion"
)

type messagesServiceGRPCServer struct {
	grpcServer      *grpc.Server
	messagesUsecase messages.Usecase
}

func NewMessagesServiceGRPCServer(grpcServer *grpc.Server, messagesUsecase messages.Usecase) *messagesServiceGRPCServer {
	return &messagesServiceGRPCServer{
		grpcServer:      grpcServer,
		messagesUsecase: messagesUsecase,
	}
}

func (c *messagesServiceGRPCServer) StartGRPCServer(listenURL string) error {
	lis, err := net.Listen("tcp", listenURL)
	if err != nil {
		return err
	}

	generated.RegisterMessagesServer(c.grpcServer, c)

	return c.grpcServer.Serve(lis)
}

func (c *messagesServiceGRPCServer) PutInProducer(ctx context.Context, bytes *generated.Bytes) (*empty.Empty, error) {
	err := c.messagesUsecase.PutInProducer(ctx, model_conversion.FromProtoBytesToBytes(bytes))
	return new(empty.Empty), err
}
