package messages

import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/labstack/echo/v4"
	"google.golang.org/grpc"
	"project/internal/generated"
	"project/internal/messages"
)

type messagesServiceGRPCClient struct {
	messagesClient generated.MessagesClient
}

func NewMessagesServiceGRPSClient(con *grpc.ClientConn) messages.Usecase {
	return &messagesServiceGRPCClient{
		messagesClient: generated.NewMessagesClient(con),
	}
}

func (m messagesServiceGRPCClient) SendMessage(ctx echo.Context, jsonWebSocketMessage []byte) error {
	_, err := m.messagesClient.SendMessage(context.TODO(), &generated.Bytes{Bytes: jsonWebSocketMessage})
	return err
}

func (m messagesServiceGRPCClient) ReceiveMessage(ctx echo.Context) ([]byte, error) {
	bytes, err := m.messagesClient.ReceiveMessage(context.TODO(), &empty.Empty{})
	if err != nil {
		return nil, err
	}

	return bytes.Bytes, nil
}
