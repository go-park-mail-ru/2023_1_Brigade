package messages

import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
	"project/internal/generated"
	"project/internal/messages"
	"project/internal/model"
	"project/internal/pkg/model_conversion"
)

type messagesServiceGRPCClient struct {
	messagesClient generated.MessagesClient
}

func NewMessagesServiceGRPSClient(con *grpc.ClientConn) messages.Usecase {
	return &messagesServiceGRPCClient{
		messagesClient: generated.NewMessagesClient(con),
	}
}

func (m messagesServiceGRPCClient) SwitchMesssageType(ctx context.Context, jsonWebSocketMessage []byte) error {
	_, err := m.messagesClient.SwitchMesssageType(ctx, &generated.Bytes{Bytes: jsonWebSocketMessage})
	return err
}

func (m messagesServiceGRPCClient) SendMessage(ctx context.Context, webSocketMessage model.WebSocketMessage) error {
	_, err := m.messagesClient.SendMessage(ctx, model_conversion.FromWebSocketMessageToProtoWebSocketMessage(webSocketMessage))
	return err
}

func (m messagesServiceGRPCClient) EditMessage(ctx context.Context, webSocketMessage model.WebSocketMessage) error {
	//TODO implement me
	panic("implement me")
}

func (m messagesServiceGRPCClient) DeleteMessage(ctx context.Context, webSocketMessage model.WebSocketMessage) error {
	//TODO implement me
	panic("implement me")
}

func (m messagesServiceGRPCClient) ReceiveMessage(ctx context.Context) ([]byte, error) {
	bytes, err := m.messagesClient.ReceiveMessage(ctx, &empty.Empty{})
	if err != nil {
		return nil, err
	}

	return bytes.Bytes, nil
}
