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

func (m messagesServiceGRPCClient) SwitchMessageType(ctx context.Context, jsonWebSocketMessage []byte) error {
	_, err := m.messagesClient.SwitchMessageType(ctx, &generated.Bytes{Bytes: jsonWebSocketMessage})
	return err
}

func (m messagesServiceGRPCClient) PutInProducer(ctx context.Context, webSocketMessage model.WebSocketMessage) error {
	_, err := m.messagesClient.PutInProducer(ctx, model_conversion.FromWebSocketMessageToProtoWebSocketMessage(webSocketMessage))
	return err
}

func (m messagesServiceGRPCClient) PullFromConsumer(ctx context.Context) ([]byte, error) {
	bytes, err := m.messagesClient.PullFromConsumer(ctx, &empty.Empty{})
	if err != nil {
		return nil, err
	}

	return bytes.Bytes, nil
}
