package usecase

import (
	"context"
	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
	"project/internal/chat"
	"project/internal/configs"
	"project/internal/messages"
	"project/internal/model"
	myErrors "project/internal/pkg/errors"
	"project/internal/pkg/model_conversion"
	"project/internal/user"
)

type usecase struct {
	chatRepo     chat.Repository
	userRepo     user.Repository
	messagesRepo messages.Repository
}

func NewChatUsecase(chatRepo chat.Repository, userRepo user.Repository, messagesRepo messages.Repository) chat.Usecase {
	return &usecase{chatRepo: chatRepo, userRepo: userRepo, messagesRepo: messagesRepo}
}

func (u usecase) CheckExistUserInChat(ctx echo.Context, chat model.Chat, userID uint64) error {
	members := chat.Members
	for _, member := range members {
		if member.Id == userID {
			return myErrors.ErrUserIsAlreadyInChat
		}
	}

	return nil
}

func (u usecase) GetChatById(ctx echo.Context, chatID uint64) (model.Chat, error) {
	chat, err := u.chatRepo.GetChatById(context.Background(), chatID)
	if err != nil {
		return model.Chat{}, err
	}

	chatMembers, err := u.chatRepo.GetChatMembersByChatId(context.Background(), chatID)
	if err != nil {
		return model.Chat{}, err
	}

	var members []model.User
	for _, chatMember := range chatMembers {
		user, err := u.userRepo.GetUserById(context.Background(), chatMember.MemberId)
		if err != nil {
			return model.Chat{}, err
		}

		members = append(members, model_conversion.FromAuthorizedUserToUser(user))
	}

	chatMessages, err := u.messagesRepo.GetChatMessages(context.Background(), chatID)
	if err != nil {
		return model.Chat{}, err
	}

	var messages []model.Message
	for _, chatMessage := range chatMessages {
		message, err := u.messagesRepo.GetMessageById(context.Background(), chatMessage.MessageId)
		if err != nil {
			return model.Chat{}, err
		}

		messages = append(messages, message)
	}

	return model.Chat{
		Id:       chat.Id,
		Type:     chat.Type,
		Title:    chat.Title,
		Avatar:   chat.Avatar,
		Members:  members,
		Messages: messages,
	}, nil
}

func (u usecase) CreateChat(ctx echo.Context, chat model.CreateChat, userID uint64) (model.Chat, error) {
	var members []model.User
	for _, userID := range chat.Members {
		user, err := u.userRepo.GetUserById(context.Background(), userID)
		if err != nil {
			return model.Chat{}, err
		}

		members = append(members, model_conversion.FromAuthorizedUserToUser(user))
	}

	createdChat := model.Chat{
		Type:     chat.Type,
		Title:    chat.Title,
		Members:  members,
		Messages: []model.Message{},
	}

	if createdChat.Type == configs.Chat {
		if len(createdChat.Members) > 0 {
			if createdChat.Members[0].Id == userID {
				createdChat.Avatar = createdChat.Members[1].Avatar
			} else {
				createdChat.Avatar = createdChat.Members[0].Avatar
			}
		}
	}

	chatFromDB, err := u.chatRepo.CreateChat(context.Background(), createdChat)

	return chatFromDB, err
}

func (u usecase) DeleteChatById(ctx echo.Context, chatID uint64) error {
	err := u.chatRepo.DeleteChatById(context.Background(), chatID)
	return err
}

func (u usecase) GetListUserChats(ctx echo.Context, userID uint64) ([]model.ChatInListUser, error) {
	var chatsInListUser []model.ChatInListUser
	userChats, err := u.chatRepo.GetChatsByUserId(context.Background(), userID)

	if err != nil {
		return nil, err
	}

	for _, userChat := range userChats {
		chat, err := u.chatRepo.GetChatById(context.Background(), userChat.ChatId)
		if err != nil {
			return nil, err
		}

		chatMembers, err := u.chatRepo.GetChatMembersByChatId(context.Background(), chat.Id)
		if err != nil {
			return nil, err
		}

		var members []model.User
		for _, chatMember := range chatMembers {
			user, err := u.userRepo.GetUserById(context.Background(), chatMember.MemberId)
			if err != nil {
				return nil, err
			}

			members = append(members, model_conversion.FromAuthorizedUserToUser(user))
		}

		lastMessage, err := u.messagesRepo.GetLastChatMessage(context.Background(), chat.Id)
		if err != nil {
			return nil, err
		}

		var lastMessageAuthor model.AuthorizedUser
		if lastMessage.AuthorId != 0 {
			lastMessageAuthor, err = u.userRepo.GetUserById(context.Background(), lastMessage.AuthorId)
			if err != nil {
				return nil, err
			}
		}

		chatsInListUser = append(chatsInListUser, model.ChatInListUser{
			Id:                chat.Id,
			Type:              chat.Type,
			Title:             chat.Title,
			Avatar:            chat.Avatar,
			Members:           members,
			LastMessage:       lastMessage,
			LastMessageAuthor: model_conversion.FromAuthorizedUserToUser(lastMessageAuthor),
		})
	}

	return chatsInListUser, nil
}

func (u usecase) EditChat(ctx echo.Context, editChat model.EditChat) (model.Chat, error) {
	chat, err := u.chatRepo.UpdateChatById(context.Background(), editChat.Title, editChat.Id)
	if err != nil {
		return model.Chat{}, err
	}

	err = u.chatRepo.DeleteChatMembers(context.Background(), editChat.Id)
	if err != nil {
		return model.Chat{}, err
	}

	var members []model.User
	for _, memberID := range editChat.Members {
		err := u.userRepo.CheckExistUserById(context.Background(), memberID)
		if err != nil {
			log.Error(err)
		}

		log.Warn(editChat.Id)
		err = u.chatRepo.AddUserInChatDB(context.Background(), editChat.Id, memberID)
		if err != nil {
			log.Error(err)
		}

		user, err := u.userRepo.GetUserById(context.Background(), memberID)
		if err != nil {
			log.Error(err)
		}

		members = append(members, model_conversion.FromAuthorizedUserToUser(user))
	}
	chat.Members = members
	chat.Title = editChat.Title
	chat.Avatar = configs.DefaultAvatarUrl

	return chat, nil
}
