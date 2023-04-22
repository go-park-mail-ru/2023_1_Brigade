package grpc

import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
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
	log.Warn("Server send messages", string(bytes.Bytes))
	var echoCtx echo.Context
	err := c.messagesUsecase.SendMessage(echoCtx, bytes.Bytes)
	log.Warn("Server send messages error", err)
	return nil, err
}

func (c *messagesServiceGRPCServer) ReceiveMessage(ctx context.Context, empty *empty.Empty) (*generated.Bytes, error) {
	var echoCtx echo.Context
	bytes, err := c.messagesUsecase.ReceiveMessage(echoCtx)
	log.Warn("Server receive messages", string(bytes))
	if err != nil {
		log.Warn("Server receive messages error", err)
		return nil, err
	}

	return &generated.Bytes{Bytes: bytes}, err
}
