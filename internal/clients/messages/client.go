package messages

import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
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
	log.Warn("Client send messages", string(jsonWebSocketMessage))
	_, err := m.messagesClient.SendMessage(context.TODO(), &generated.Bytes{Bytes: jsonWebSocketMessage})
	log.Warn("Client send messages error", err)
	return err
}

func (m messagesServiceGRPCClient) ReceiveMessage(ctx echo.Context) ([]byte, error) {
	bytes, err := m.messagesClient.ReceiveMessage(context.TODO(), &empty.Empty{})
	log.Warn("Client receive messages", string(bytes.Bytes))
	if err != nil {
		log.Warn("Client receive messages err", err)
		return nil, err
	}

	return bytes.Bytes, nil
}
