package server

import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
	"net"
	"project/internal/generated"
	"project/internal/microservices/chat"
	"project/internal/pkg/model_conversion"
)

type chatsServiceGRPCServer struct {
	grpcServer  *grpc.Server
	chatUsecase chat.Usecase
}

func NewChatsServiceGRPCServer(grpcServer *grpc.Server, chatUsecase chat.Usecase) *chatsServiceGRPCServer {
	return &chatsServiceGRPCServer{
		grpcServer:  grpcServer,
		chatUsecase: chatUsecase,
	}
}

func (c *chatsServiceGRPCServer) StartGRPCServer(listenURL string) error {
	lis, err := net.Listen("tcp", listenURL)
	if err != nil {
		return err
	}

	generated.RegisterChatsServer(c.grpcServer, c)

	return c.grpcServer.Serve(lis)
}

func (c *chatsServiceGRPCServer) GetChatById(ctx context.Context, getChatArguments *generated.GetChatArguments) (*generated.Chat, error) {
	chat, err := c.chatUsecase.GetChatById(ctx, getChatArguments.ChatID, getChatArguments.UserID)
	if err != nil {
		return nil, err
	}

	return model_conversion.FromChatToProtoChat(chat), nil
}

func (c *chatsServiceGRPCServer) EditChat(ctx context.Context, editChat *generated.EditChatModel) (*generated.Chat, error) {
	chat, err := c.chatUsecase.EditChat(ctx, model_conversion.FromProtoEditChatToEditChat(editChat))
	if err != nil {
		return nil, err
	}

	return model_conversion.FromChatToProtoChat(chat), nil
}

func (c *chatsServiceGRPCServer) CreateChat(ctx context.Context, createChat *generated.CreateChatArguments) (*generated.Chat, error) {
	chat, err := c.chatUsecase.CreateChat(
		ctx,
		model_conversion.FromProtoCreateChatToCreateChat(createChat.Chat),
		model_conversion.FromProtoUserIDToUserID(createChat.UserID),
	)
	if err != nil {
		return nil, err
	}

	return model_conversion.FromChatToProtoChat(chat), nil
}

func (c *chatsServiceGRPCServer) DeleteChatById(ctx context.Context, chatID *generated.ChatID) (*empty.Empty, error) {
	err := c.chatUsecase.DeleteChatById(ctx, model_conversion.FromProtoChatIDToChatID(chatID))
	return &empty.Empty{}, err
}

func (c *chatsServiceGRPCServer) CheckExistUserInChat(ctx context.Context, existChat *generated.ExistChatArguments) (*empty.Empty, error) {
	err := c.chatUsecase.CheckExistUserInChat(
		ctx,
		model_conversion.FromProtoChatToChat(existChat.Chat),
		model_conversion.FromProtoUserIDToUserID(existChat.UserID),
	)
	return &empty.Empty{}, err
}

func (c *chatsServiceGRPCServer) GetListUserChats(ctx context.Context, userID *generated.UserID) (*generated.ArrayChatInListUser, error) {
	chats, err := c.chatUsecase.GetListUserChats(ctx, model_conversion.FromProtoUserIDToUserID(userID))
	if err != nil {
		return nil, err
	}

	return model_conversion.FromUserChatsToProtoUserChats(chats), nil
}

func (c *chatsServiceGRPCServer) GetSearchChatsMessagesChannels(ctx context.Context, argumets *generated.SearchChatsArgumets) (*generated.FoundedChatsMessagesChannels, error) {
	chats, err := c.chatUsecase.GetSearchChatsMessagesChannels(ctx, argumets.UserID, argumets.String_)
	if err != nil {
		return nil, err
	}

	return model_conversion.FromSearchChatsToProtoSearchChats(chats), nil
}
