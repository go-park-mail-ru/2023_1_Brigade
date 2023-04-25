package chat

import (
	"context"
	"google.golang.org/grpc"
	"project/internal/chat"
	"project/internal/generated"
	"project/internal/model"
	"project/internal/pkg/model_conversion"
)

type chatServiceGRPCClient struct {
	chatClient generated.ChatsClient
}

func NewChatServiceGRPSClient(con *grpc.ClientConn) chat.Usecase {
	return &chatServiceGRPCClient{
		chatClient: generated.NewChatsClient(con),
	}
}

func (c chatServiceGRPCClient) CheckExistUserInChat(ctx context.Context, chat model.Chat, userID uint64) error {
	_, err := c.chatClient.CheckExistUserInChat(ctx, &generated.ExistChatArguments{
		Chat:   model_conversion.FromChatToProtoChat(chat),
		UserID: model_conversion.FromUserIDToProtoUserID(userID),
	})

	return err
}

func (c chatServiceGRPCClient) GetChatById(ctx context.Context, chatID uint64) (model.Chat, error) {
	chat, err := c.chatClient.GetChatById(ctx, model_conversion.FromChatIDToProtoChatID(chatID))
	if err != nil {
		return model.Chat{}, err
	}

	return model_conversion.FromProtoChatToChat(chat), nil
}

func (c chatServiceGRPCClient) EditChat(ctx context.Context, editChat model.EditChat) (model.Chat, error) {
	chat, err := c.chatClient.EditChat(ctx, model_conversion.FromEditChatToProtoEditChat(editChat))

	if err != nil {
		return model.Chat{}, err
	}

	return model_conversion.FromProtoChatToChat(chat), err
}

func (c chatServiceGRPCClient) CreateChat(ctx context.Context, createChat model.CreateChat, userID uint64) (model.Chat, error) {
	chat, err := c.chatClient.CreateChat(ctx,
		&generated.CreateChatArguments{
			Chat:   model_conversion.FromCreateChatToProtoCreateChat(createChat),
			UserID: model_conversion.FromUserIDToProtoUserID(userID),
		})

	if err != nil {
		return model.Chat{}, err
	}

	return model_conversion.FromProtoChatToChat(chat), err
}

func (c chatServiceGRPCClient) DeleteChatById(ctx context.Context, chatID uint64) error {
	_, err := c.chatClient.DeleteChatById(ctx, model_conversion.FromChatIDToProtoChatID(chatID))
	return err
}

func (c chatServiceGRPCClient) GetListUserChats(ctx context.Context, userID uint64) ([]model.ChatInListUser, error) {
	chats, err := c.chatClient.GetListUserChats(ctx, model_conversion.FromUserIDToProtoUserID(userID))
	if err != nil {
		return nil, err
	}

	return model_conversion.FromProtoUserChatsToUserChats(chats), nil
}
