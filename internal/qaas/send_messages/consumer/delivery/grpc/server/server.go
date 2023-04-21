package server

import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
	"net"
	protobuf "project/internal/model/generated"
	"project/internal/qaas/send_messages/consumer"
	"project/internal/qaas/send_messages/consumer/delivery/grpc/service"
)

type consumersServiceGRPCServer struct {
	grpcServer      *grpc.Server
	consumerUsecase consumer.Usecase
}

func NewConsumersServiceGRPCServer(grpcServer *grpc.Server, consumerUsecase consumer.Usecase) *consumersServiceGRPCServer {
	return &consumersServiceGRPCServer{
		grpcServer:      grpcServer,
		consumerUsecase: consumerUsecase,
	}
}

func (c *consumersServiceGRPCServer) StartGRPCServer(listenURL string) error {
	lis, err := net.Listen("tcp", listenURL)
	if err != nil {
		return err
	}

	service.RegisterConsumersServer(c.grpcServer, c)

	return c.grpcServer.Serve(lis)
}

func (c *consumersServiceGRPCServer) ConsumeMessage(ctx context.Context, empty *empty.Empty) (*protobuf.Bytes, error) {
	msg := c.consumerUsecase.ConsumeMessage()
	return &protobuf.Bytes{Bytes: msg}, nil
}

func (c *consumersServiceGRPCServer) StartConsumeMessages(ctx context.Context, empty *empty.Empty) (*empty.Empty, error) {
	c.consumerUsecase.StartConsumeMessages()
	return empty, nil
}

//func (c *messagesServiceGRPCServer) SendMessage(ctx context.Context, bytes *protobuf.Bytes) (*empty.Empty, error) {
//	var echoCtx echo.Context
//	err := c.messagesUsecase.SendMessage(echoCtx, bytes.Bytes)
//	return nil, err
//}
//
//func (c *messagesServiceGRPCServer) ReceiveMessage(ctx context.Context, empty *empty.Empty) (*protobuf.Bytes, error) {
//	var echoCtx echo.Context
//	bytes, err := c.messagesUsecase.ReceiveMessage(echoCtx)
//	if err != nil {
//		return nil, err
//	}
//
//	return &protobuf.Bytes{Bytes: bytes}, err
//}
