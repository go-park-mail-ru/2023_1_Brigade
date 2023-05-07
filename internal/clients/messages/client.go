package messages

import (
	"context"
	"google.golang.org/grpc"
	"project/internal/generated"
	"project/internal/messages"
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

func (m messagesServiceGRPCClient) PutInProducer(ctx context.Context, jsonWebSocketMessage []byte) error {
	_, err := m.messagesClient.PutInProducer(ctx, model_conversion.FromBytesToProtoBytes(jsonWebSocketMessage))
	if err != nil {
		return err
	}

	return nil
}
