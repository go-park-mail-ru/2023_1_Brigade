package usecase

import (
	"context"
	"project/internal/config"
	"project/internal/microservices/chat"
	"project/internal/microservices/messages"
	"project/internal/microservices/user"
	"project/internal/model"
	"project/internal/monolithic_services/images"
	myErrors "project/internal/pkg/errors"
	"project/internal/pkg/model_conversion"
	"sort"
)

type usecase struct {
	chatRepo      chat.Repository
	userRepo      user.Repository
	messagesRepo  messages.Repository
	imagesUsecase images.Usecase
}

func NewChatUsecase(chatRepo chat.Repository, userRepo user.Repository, messagesRepo messages.Repository, imagesUsecase images.Usecase) chat.Usecase {
	return &usecase{chatRepo: chatRepo, userRepo: userRepo, messagesRepo: messagesRepo, imagesUsecase: imagesUsecase}
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

func (u usecase) GetChatById(ctx context.Context, chatID uint64, userID uint64) (model.Chat, error) {
	chat, err := u.chatRepo.GetChatById(ctx, chatID)
	if err != nil {
		return model.Chat{}, err
	}

	chatMembers, err := u.chatRepo.GetChatMembersByChatId(ctx, chatID)
	if err != nil {
		return model.Chat{}, err
	}

	if chat.Type != config.Channel {
		userInChat := false
		for _, chatMember := range chatMembers {
			if chatMember.MemberId == userID {
				userInChat = true
				continue
			}
		}

		if !userInChat {
			return model.Chat{}, myErrors.ErrNotChatAccess
		}
	}

	var members []model.User
	for _, chatMember := range chatMembers {
		user, err := u.userRepo.GetUserById(ctx, chatMember.MemberId)
		if err != nil {
			return model.Chat{}, err
		}

		members = append(members, model_conversion.FromAuthorizedUserToUser(user))
	}

	chatMessages, err := u.messagesRepo.GetChatMessages(ctx, chatID)
	if err != nil {
		return model.Chat{}, err
	}

	var messages []model.Message
	for _, chatMessage := range chatMessages {
		message, err := u.messagesRepo.GetMessageById(ctx, chatMessage.MessageId)
		if err != nil {
			return model.Chat{}, err
		}

		messages = append(messages, message)
	}

	returnedChat := model.Chat{
		Id:          chat.Id,
		MasterID:    chat.MasterID,
		Type:        chat.Type,
		Description: chat.Description,
		Title:       chat.Title,
		Avatar:      chat.Avatar,
		Members:     members,
		Messages:    messages,
	}

	if returnedChat.Type == config.Chat {
		if len(returnedChat.Members) > 0 {
			if returnedChat.Members[0].Id == userID {
				returnedChat.Title = returnedChat.Members[1].Nickname
				returnedChat.Avatar = returnedChat.Members[1].Avatar
				returnedChat.Description = returnedChat.Members[1].Status
			} else {
				returnedChat.Title = returnedChat.Members[0].Nickname
				returnedChat.Avatar = returnedChat.Members[0].Avatar
				returnedChat.Description = returnedChat.Members[0].Status
			}
		}
	}

	return returnedChat, nil
}

func (u usecase) GetChatInfoById(ctx context.Context, chatID uint64, userID uint64) (model.ChatInListUser, error) {
	chat, err := u.chatRepo.GetChatById(ctx, chatID)
	if err != nil {
		return model.ChatInListUser{}, err
	}

	chatMembers, err := u.chatRepo.GetChatMembersByChatId(ctx, chatID)
	if err != nil {
		return model.ChatInListUser{}, err
	}

	if chat.Type != config.Channel {
		userInChat := false
		for _, chatMember := range chatMembers {
			if chatMember.MemberId == userID {
				userInChat = true
				continue
			}
		}

		if !userInChat {
			return model.ChatInListUser{}, myErrors.ErrNotChatAccess
		}
	}

	var members []model.User
	for _, chatMember := range chatMembers {
		user, err := u.userRepo.GetUserById(ctx, chatMember.MemberId)
		if err != nil {
			return model.ChatInListUser{}, err
		}

		members = append(members, model_conversion.FromAuthorizedUserToUser(user))
	}

	lastMessage, err := u.messagesRepo.GetLastChatMessage(ctx, chatID)
	if err != nil {
		return model.ChatInListUser{}, err
	}

	lastMessageAuthor, err := u.userRepo.GetUserById(ctx, userID)
	if err != nil {
		return model.ChatInListUser{}, err
	}

	returnedChat := model.ChatInListUser{
		Id:                chat.Id,
		Type:              chat.Type,
		Title:             chat.Title,
		Avatar:            chat.Avatar,
		Members:           members,
		LastMessage:       lastMessage,
		LastMessageAuthor: model_conversion.FromAuthorizedUserToUser(lastMessageAuthor),
	}

	if returnedChat.Type == config.Chat {
		if len(returnedChat.Members) > 0 {
			if returnedChat.Members[0].Id == userID {
				returnedChat.Title = returnedChat.Members[1].Nickname
				returnedChat.Avatar = returnedChat.Members[1].Avatar
			} else {
				returnedChat.Title = returnedChat.Members[0].Nickname
				returnedChat.Avatar = returnedChat.Members[0].Avatar
			}
		}
	}

	return returnedChat, nil
}

func (u usecase) CreateChat(ctx context.Context, chat model.CreateChat, userID uint64) (model.Chat, error) {
	var members []model.User
	for _, memberID := range chat.Members {
		user, err := u.userRepo.GetUserById(ctx, memberID)
		if err != nil {
			return model.Chat{}, err
		}

		members = append(members, model_conversion.FromAuthorizedUserToUser(user))
	}

	createdChat := model.Chat{
		MasterID:    userID,
		Type:        chat.Type,
		Description: chat.Description,
		Title:       chat.Title,
		Avatar:      chat.Avatar,
		Members:     members,
		Messages:    []model.Message{},
	}

	chatFromDB, err := u.chatRepo.CreateChat(ctx, createdChat)
	if err != nil {
		return model.Chat{}, err
	}

	if createdChat.Type != config.Chat && createdChat.Avatar == "" {
		//filename := strconv.FormatUint(chatFromDB.Id, 10)
		//firstCharacterName := string([]rune(chat.Title)[0])

		//err = u.imagesUsecase.UploadGeneratedImage(ctx, config.ChatAvatarsBucket, filename, firstCharacterName)
		//if err != nil {
		//	return model.Chat{}, err
		//}
		//
		//url, err := u.imagesUsecase.GetImage(ctx, config.ChatAvatarsBucket, filename)
		//if err != nil {
		//	return model.Chat{}, err
		//}

		url := "https://brigade_chat_avatars.hb.bizmrg.com/logo.png"

		chatFromDB, err = u.chatRepo.UpdateChatAvatar(ctx, url, chatFromDB.Id)
		if err != nil {
			return model.Chat{}, err
		}
	}

	if chat.Type == config.Chat {
		if len(chatFromDB.Members) > 0 {
			if chatFromDB.Members[0].Id == userID {
				chatFromDB.Title = chatFromDB.Members[1].Nickname
				chatFromDB.Avatar = chatFromDB.Members[1].Avatar
				chatFromDB.Description = chatFromDB.Members[1].Status
			} else {
				chatFromDB.Title = chatFromDB.Members[0].Nickname
				chatFromDB.Avatar = chatFromDB.Members[0].Avatar
				chatFromDB.Description = chatFromDB.Members[0].Status
			}
		}
	}

	chatFromDB.MasterID = userID
	return chatFromDB, nil
}

func (u usecase) DeleteChatById(ctx context.Context, chatID uint64) error {
	err := u.chatRepo.DeleteChatById(ctx, chatID)
	return err
}

func (u usecase) GetListUserChats(ctx context.Context, userID uint64) ([]model.ChatInListUser, error) {
	var chatsInListUser []model.ChatInListUser
	userChats, err := u.chatRepo.GetChatsByUserId(ctx, userID)

	if err != nil {
		return nil, err
	}

	for _, userChat := range userChats {
		chat, err := u.chatRepo.GetChatById(ctx, userChat.ChatId)
		if err != nil {
			return nil, err
		}

		chatMembers, err := u.chatRepo.GetChatMembersByChatId(ctx, chat.Id)
		if err != nil {
			return nil, err
		}

		var members []model.User
		for _, chatMember := range chatMembers {
			user, err := u.userRepo.GetUserById(ctx, chatMember.MemberId)
			if err != nil {
				return nil, err
			}

			members = append(members, model_conversion.FromAuthorizedUserToUser(user))
		}

		lastMessage, err := u.messagesRepo.GetLastChatMessage(ctx, chat.Id)
		if err != nil {
			return nil, err
		}

		var lastMessageAuthor model.AuthorizedUser
		if lastMessage.AuthorId != 0 {
			lastMessageAuthor, err = u.userRepo.GetUserById(ctx, lastMessage.AuthorId)
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

	for ind := range chatsInListUser {
		if chatsInListUser[ind].Type == config.Chat {
			if chatsInListUser[ind].Members[0].Id == userID && len(chatsInListUser[ind].Members) > 1 {
				chatsInListUser[ind].Title = chatsInListUser[ind].Members[1].Nickname
				chatsInListUser[ind].Avatar = chatsInListUser[ind].Members[1].Avatar
			} else {
				chatsInListUser[ind].Title = chatsInListUser[ind].Members[0].Nickname
				chatsInListUser[ind].Avatar = chatsInListUser[ind].Members[0].Avatar
			}
		}
	}

	sort.Slice(chatsInListUser, func(i, j int) bool {
		return chatsInListUser[i].LastMessage.CreatedAt > chatsInListUser[j].LastMessage.CreatedAt
	})

	return chatsInListUser, nil
}

func (u usecase) EditChat(ctx context.Context, editChat model.EditChat) (model.Chat, error) {
	chatFromDB, err := u.chatRepo.UpdateChatById(ctx, editChat)
	if err != nil {
		return model.Chat{}, err
	}

	chat := model.Chat{
		Id:          chatFromDB.Id,
		MasterID:    chatFromDB.MasterID,
		Type:        chatFromDB.Type,
		Description: chatFromDB.Description,
		Title:       chatFromDB.Title,
		Avatar:      chatFromDB.Avatar,
	}

	err = u.chatRepo.DeleteChatMembers(context.TODO(), editChat.Id)
	if err != nil {
		return model.Chat{}, err
	}

	var members []model.User
	for _, memberID := range editChat.Members {
		err = u.userRepo.CheckExistUserById(context.TODO(), memberID)
		if err != nil {
			return model.Chat{}, err
		}

		err = u.chatRepo.AddUserInChatDB(context.TODO(), editChat.Id, memberID)
		if err != nil {
			return model.Chat{}, err
		}

		user, err := u.userRepo.GetUserById(context.TODO(), memberID)
		if err != nil {
			return model.Chat{}, err
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

	chats, err := u.chatRepo.GetSearchChats(ctx, userID, string)
	if err != nil {
		return model.FoundedChatsMessagesChannels{}, err
	}

	messages, err := u.messagesRepo.GetSearchMessages(ctx, userID, string)
	if err != nil {
		return model.FoundedChatsMessagesChannels{}, err
	}

	foundedContacts, err := u.userRepo.GetSearchUsers(ctx, string)
	if err != nil {
		return model.FoundedChatsMessagesChannels{}, err
	}

	var foundedChannels []model.ChatInListUser
	var foundedChats []model.ChatInListUser
	var foundedMessages []model.ChatInListUser

	for _, channel := range channels {
		foundedChannel := model.ChatInListUser{
			Id:     channel.Id,
			Type:   channel.Type,
			Title:  channel.Title,
			Avatar: channel.Avatar,
		}

		lastMessage, err := u.messagesRepo.GetLastChatMessage(ctx, channel.Id)
		if err != nil {
			return model.FoundedChatsMessagesChannels{}, err
		}
		foundedChannel.LastMessage = lastMessage

		foundedChannels = append(foundedChannels, foundedChannel)
	}

	for _, chat := range chats {
		foundedChat := model.ChatInListUser{
			Id:     chat.Id,
			Type:   chat.Type,
			Title:  chat.Title,
			Avatar: chat.Avatar,
		}

		lastMessage, err := u.messagesRepo.GetLastChatMessage(ctx, chat.Id)
		if err != nil {
			return model.FoundedChatsMessagesChannels{}, err
		}
		foundedChat.LastMessage = lastMessage

		foundedChats = append(foundedChats, foundedChat)
	}

	for _, message := range messages {
		chat, err := u.chatRepo.GetChatById(ctx, message.ChatId)
		if err != nil {
			return model.FoundedChatsMessagesChannels{}, err
		}

		foundedChat := model.ChatInListUser{
			Id:     chat.Id,
			Type:   chat.Type,
			Title:  chat.Title,
			Avatar: chat.Avatar,
		}
		if chat.Type == config.Chat {
			chatMembers, err := u.chatRepo.GetChatMembersByChatId(ctx, message.ChatId)
			if err != nil {
				return model.FoundedChatsMessagesChannels{}, err
			}

			var members []model.User
			for _, chatMember := range chatMembers {
				member, err := u.userRepo.GetUserById(ctx, chatMember.MemberId)
				if err != nil {
					return model.FoundedChatsMessagesChannels{}, err
				}

				members = append(members, model_conversion.FromAuthorizedUserToUser(member))
			}
			if len(members) > 1 {
				if members[0].Id == userID {
					foundedChat.Title = members[1].Nickname
					foundedChat.Avatar = members[1].Avatar
				} else {
					foundedChat.Title = members[0].Nickname
					foundedChat.Avatar = members[0].Avatar
				}
			}
		}
		foundedChat.LastMessage = message

		foundedMessages = append(foundedMessages, foundedChat)
	}

	return model.FoundedChatsMessagesChannels{
		FoundedChats:    foundedChats,
		FoundedMessages: foundedMessages,
		FoundedChannels: foundedChannels,
		FoundedContacts: model_conversion.FromAuthorizedUserArrayToUserArray(foundedContacts),
	}, nil
}
