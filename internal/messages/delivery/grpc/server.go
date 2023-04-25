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

func (c *messagesServiceGRPCServer) SwitchMesssageType(ctx context.Context, bytes *generated.Bytes) (*empty.Empty, error) {
	err := c.messagesUsecase.SwitchMesssageType(ctx, bytes.Bytes)
	return nil, err
}

func (c *messagesServiceGRPCServer) SendMessage(ctx context.Context, message *generated.WebSocketMessage) (*empty.Empty, error) {
	err := c.messagesUsecase.SendMessage(ctx, model_conversion.FromProtoWebSocketMessageToWebSocketMessage(message))
	return nil, err
}

func (c *messagesServiceGRPCServer) EditMessage(ctx context.Context, message *generated.WebSocketMessage) (*empty.Empty, error) {
	//TODO implement me
	panic("implement me")
}

func (c *messagesServiceGRPCServer) DeleteMessage(ctx context.Context, message *generated.WebSocketMessage) (*empty.Empty, error) {
	//TODO implement me
	panic("implement me")
}

func (c *messagesServiceGRPCServer) ReceiveMessage(ctx context.Context, empty *empty.Empty) (*generated.Bytes, error) {
	bytes, err := c.messagesUsecase.ReceiveMessage(ctx)
	if err != nil {
		return nil, err
	}

	return &generated.Bytes{Bytes: bytes}, err
}
