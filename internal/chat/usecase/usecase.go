package usecase

import (
	"context"
	log "github.com/sirupsen/logrus"
	"project/internal/chat"
	"project/internal/configs"
	"project/internal/messages"
	"project/internal/model"
	myErrors "project/internal/pkg/errors"
	"project/internal/pkg/image_generation"
	"project/internal/pkg/model_conversion"
	"project/internal/user"
	"strings"
)

type usecase struct {
	chatRepo     chat.Repository
	userRepo     user.Repository
	messagesRepo messages.Repository
}

func NewChatUsecase(chatRepo chat.Repository, userRepo user.Repository, messagesRepo messages.Repository) chat.Usecase {
	return &usecase{chatRepo: chatRepo, userRepo: userRepo, messagesRepo: messagesRepo}
}

func (u usecase) CheckExistUserInChat(ctx context.Context, chat model.Chat, userID uint64) error {
	members := chat.Members
	for _, member := range members {
		if member.Id == userID {
			return myErrors.ErrUserIsAlreadyInChat
		}
	}

	return nil
}

func (u usecase) GetChatById(ctx context.Context, chatID uint64) (model.Chat, error) {
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

func (u usecase) CreateChat(ctx context.Context, chat model.CreateChat, userID uint64) (model.Chat, error) {
	var members []model.User
	for _, userID := range chat.Members {
		user, err := u.userRepo.GetUserById(context.Background(), userID)
		if err != nil {
			return model.Chat{}, err
		}

		members = append(members, model_conversion.FromAuthorizedUserToUser(user))
	}

	createdChat := model.Chat{
		MasterID: userID,
		Type:     chat.Type,
		Title:    chat.Title,
		Members:  members,
		Messages: []model.Message{},
	}

	log.Info(createdChat)

	if createdChat.Type != configs.Chat {
		avatar, err := image_generation.GenerateAvatar(string(chat.Title[0]))
		if err != nil {
			log.Error(err)
		}

		createdChat.Avatar = avatar
	}

	chatFromDB, err := u.chatRepo.CreateChat(context.Background(), createdChat)
	chatFromDB.MasterID = createdChat.MasterID
	log.Info(chatFromDB)

	return chatFromDB, err
}

func (u usecase) DeleteChatById(ctx context.Context, chatID uint64) error {
	err := u.chatRepo.DeleteChatById(context.Background(), chatID)
	return err
}

func (u usecase) GetListUserChats(ctx context.Context, userID uint64) ([]model.ChatInListUser, error) {
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

func (u usecase) EditChat(ctx context.Context, editChat model.EditChat) (model.Chat, error) {
	chatFromDB, err := u.chatRepo.UpdateChatById(ctx, editChat.Title, editChat.Id)
	if err != nil {
		return model.Chat{}, err
	}
	chat := model.Chat{
		Id:     chatFromDB.Id,
		Type:   chatFromDB.Type,
		Title:  chatFromDB.Title,
		Avatar: chatFromDB.Avatar,
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

	return chat, nil
}

func (u usecase) GetSearchChatsMessagesChannels(ctx context.Context, userID uint64, string string) (model.FoundedChatsMessagesChannels, error) {
	channels, err := u.chatRepo.GetSearchChannels(ctx, string, userID)
	if err != nil {
		return model.FoundedChatsMessagesChannels{}, err
	}

	chatMembers, err := u.chatRepo.GetChatsByUserId(ctx, userID)
	if err != nil {
		return model.FoundedChatsMessagesChannels{}, err
	}

	contacts, err := u.userRepo.GetAllUsersExceptCurrentUser(ctx, userID)
	if err != nil {
		return model.FoundedChatsMessagesChannels{}, err
	}

	var lastMessages []model.Message
	var chats []model.Chat
	for _, chatMember := range chatMembers {
		lastMessage, err := u.messagesRepo.GetLastChatMessage(ctx, chatMember.ChatId)
		if err != nil {
			return model.FoundedChatsMessagesChannels{}, err
		}

		chat, err := u.chatRepo.GetChatById(ctx, chatMember.ChatId)
		if err != nil {
			return model.FoundedChatsMessagesChannels{}, err
		}
		if chat.Type == configs.Chat {
			continue
		}

		lastMessages = append(lastMessages, lastMessage)
		chats = append(chats, chat)
	}

	var correctLastMessages []model.ChatInListUser
	var correctChats []model.ChatInListUser
	var correctChannels []model.ChatInListUser
	for _, message := range lastMessages {
		if strings.Contains(strings.ToLower(message.Body), strings.ToLower(string)) {
			chat, err := u.chatRepo.GetChatById(ctx, message.ChatId)
			if err != nil {
				return model.FoundedChatsMessagesChannels{}, err
			}

			messageToArray := model.ChatInListUser{
				Id:          chat.Id,
				Type:        chat.Type,
				Title:       chat.Title,
				Avatar:      chat.Avatar,
				LastMessage: message,
			}

			correctLastMessages = append(correctLastMessages, messageToArray)
		}
	}

	for _, chat := range chats {
		if strings.Contains(strings.ToLower(chat.Title), strings.ToLower(string)) {
			chatToArray := model.ChatInListUser{
				Id:     chat.Id,
				Type:   chat.Type,
				Title:  chat.Title,
				Avatar: chat.Avatar,
			}

			lastMessage, err := u.messagesRepo.GetLastChatMessage(ctx, chat.Id)
			if err != nil {
				return model.FoundedChatsMessagesChannels{}, err
			}
			chatToArray.LastMessage = lastMessage

			correctChats = append(correctChats, chatToArray)
		}
	}

	for _, channel := range channels {
		channelToArray := model.ChatInListUser{
			Id:     channel.Id,
			Type:   channel.Type,
			Title:  channel.Title,
			Avatar: channel.Avatar,
		}

		lastMessage, err := u.messagesRepo.GetLastChatMessage(ctx, channel.Id)
		if err != nil {
			return model.FoundedChatsMessagesChannels{}, err
		}
		channelToArray.LastMessage = lastMessage

		correctChannels = append(correctChannels, channelToArray)
	}

	return model.FoundedChatsMessagesChannels{
		FoundedChats:    correctChats,
		FoundedMessages: correctLastMessages,
		FoundedChannels: correctChannels,
		FoundedContacts: model_conversion.FromAuthorizedUserArrayToUserArray(contacts),
	}, nil
}
