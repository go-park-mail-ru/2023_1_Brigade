package grpc

import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/labstack/echo/v4"
	"google.golang.org/grpc"
	"net"
	"project/internal/generated"
	"project/internal/messages"
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

func (c *messagesServiceGRPCServer) SendMessage(ctx context.Context, bytes *generated.Bytes) (*empty.Empty, error) {
	var echoCtx echo.Context
	err := c.messagesUsecase.SendMessage(echoCtx, bytes.Bytes)
	return nil, err
}

func (c *messagesServiceGRPCServer) ReceiveMessage(ctx context.Context, empty *empty.Empty) (*generated.Bytes, error) {
	var echoCtx echo.Context
	bytes, err := c.messagesUsecase.ReceiveMessage(echoCtx)
	if err != nil {
		return nil, err
	}

	return &generated.Bytes{Bytes: bytes}, err
}
