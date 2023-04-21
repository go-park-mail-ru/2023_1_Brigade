package model_conversion

import (
	protobuf "project/internal/generated"
	"project/internal/model"
)

func FromUserIDToProtoUserID(userID uint64) *protobuf.UserID {
	return &protobuf.UserID{UserID: userID}
}

func FromProtoUserIDToUserID(userID *protobuf.UserID) uint64 {
	return userID.UserID
}

func FromChatIDToProtoChatID(chatID uint64) *protobuf.ChatID {
	return &protobuf.ChatID{ChatID: chatID}
}

func FromProtoChatIDToChatID(chatID *protobuf.ChatID) uint64 {
	return chatID.ChatID
}

func FromProtoUserToUser(user *protobuf.User) model.User {
	return model.User{
		Id:       user.Id,
		Username: user.Username,
		Nickname: user.Nickname,
		Email:    user.Email,
		Status:   user.Status,
		Avatar:   user.Avatar,
	}
}

func FromProtoMessageToMessage(message *protobuf.Message) model.Message {
	return model.Message{
		Id:       message.Id,
		Body:     message.Body,
		AuthorId: message.AuthorId,
		ChatId:   message.ChatId,
	}
}

func FromUserToProtoUser(user model.User) *protobuf.User {
	return &protobuf.User{
		Id:       user.Id,
		Username: user.Username,
		Nickname: user.Nickname,
		Email:    user.Email,
		Status:   user.Status,
		Avatar:   user.Avatar,
	}
}

func FromMessageToProtoMessage(message model.Message) *protobuf.Message {
	return &protobuf.Message{
		Id:       message.Id,
		Body:     message.Body,
		AuthorId: message.AuthorId,
		ChatId:   message.ChatId,
	}
}

func FromUserChatToProtoUserChat(chat model.ChatInListUser) *protobuf.ChatInListUser {
	return &protobuf.ChatInListUser{
		Id:                chat.Id,
		Type:              chat.Type,
		Title:             chat.Title,
		Avatar:            chat.Avatar,
		Members:           FromMembersToProtoMembers(chat.Members),
		LastMessage:       FromMessageToProtoMessage(chat.LastMessage),
		LastMessageAuthor: FromUserToProtoUser(chat.LastMessageAuthor),
	}
}

func FromProtoUserChatToUserChat(chat *protobuf.ChatInListUser) model.ChatInListUser {
	return model.ChatInListUser{
		Id:                chat.Id,
		Type:              chat.Type,
		Title:             chat.Title,
		Avatar:            chat.Avatar,
		Members:           FromProtoMembersToMembers(chat.Members),
		LastMessage:       FromProtoMessageToMessage(chat.LastMessage),
		LastMessageAuthor: FromProtoUserToUser(chat.LastMessageAuthor),
	}
}

func FromProtoEditChatToEditChat(chat *protobuf.EditChatModel) model.EditChat {
	return model.EditChat{
		Id:      chat.Id,
		Type:    chat.Type,
		Title:   chat.Title,
		Members: chat.Members,
	}
}

func FromEditChatToProtoEditChat(chat model.EditChat) *protobuf.EditChatModel {
	return &protobuf.EditChatModel{
		Id:      chat.Id,
		Type:    chat.Type,
		Title:   chat.Title,
		Members: chat.Members,
	}
}

func FromProtoUpdateUserToUpdateUser(user *protobuf.UpdateUser) model.UpdateUser {
	return model.UpdateUser{
		Username:        user.Username,
		Nickname:        user.Nickname,
		Status:          user.Status,
		CurrentPassword: user.CurrentPassword,
		NewPassword:     user.NewPassword,
	}
}

func FromUpdateUserToProtoUpdateUser(user model.UpdateUser) *protobuf.UpdateUser {
	return &protobuf.UpdateUser{
		Username:        user.Username,
		Nickname:        user.Nickname,
		Status:          user.Status,
		CurrentPassword: user.CurrentPassword,
		NewPassword:     user.NewPassword,
	}
}

func FromProtoCreateChatToCreateChat(chat *protobuf.CreateChat) model.CreateChat {
	return model.CreateChat{
		Type:    chat.Type,
		Title:   chat.Title,
		Members: chat.Members,
	}
}

func FromCreateChatToProtoCreateChat(chat model.CreateChat) *protobuf.CreateChat {
	return &protobuf.CreateChat{
		Type:    chat.Type,
		Title:   chat.Title,
		Members: chat.Members,
	}
}

func FromProtoMembersToMembers(members []*protobuf.User) []model.User {
	res := make([]model.User, len(members))

	for idx, value := range members {
		res[idx] = FromProtoUserToUser(value)
	}

	return res
}

func FromMembersToProtoMembers(members []model.User) []*protobuf.User {
	res := make([]*protobuf.User, len(members))

	for idx, value := range members {
		res[idx] = FromUserToProtoUser(value)
	}

	return res
}

func FromProtoMessagesToMessages(messages []*protobuf.Message) []model.Message {
	res := make([]model.Message, len(messages))

	for idx, value := range messages {
		res[idx] = FromProtoMessageToMessage(value)
	}

	return res
}

func FromMessagesToProtoMessages(messages []model.Message) []*protobuf.Message {
	res := make([]*protobuf.Message, len(messages))

	for idx, value := range messages {
		res[idx] = FromMessageToProtoMessage(value)
	}

	return res
}

func FromProtoUserChatsToUserChats(chats *protobuf.ArrayChatInListUser) []model.ChatInListUser {
	res := make([]model.ChatInListUser, len(chats.Chats))

	for idx, value := range chats.Chats {
		res[idx] = FromProtoUserChatToUserChat(value)
	}

	return res
}

func FromUserChatsToProtoUserChats(chats []model.ChatInListUser) *protobuf.ArrayChatInListUser {
	var chatsArr protobuf.ArrayChatInListUser

	for _, value := range chats {
		chatsArr.Chats = append(chatsArr.Chats, FromUserChatToProtoUserChat(value))
	}

	return &chatsArr
}

func FromProtoChatToChat(chat *protobuf.Chat) model.Chat {
	return model.Chat{
		Id:       chat.Id,
		Type:     chat.Type,
		Title:    chat.Title,
		Avatar:   chat.Avatar,
		Members:  FromProtoMembersToMembers(chat.Members),
		Messages: FromProtoMessagesToMessages(chat.Messages),
	}
}

func FromChatToProtoChat(chat model.Chat) *protobuf.Chat {
	return &protobuf.Chat{
		Id:       chat.Id,
		Type:     chat.Type,
		Title:    chat.Title,
		Avatar:   chat.Avatar,
		Members:  FromMembersToProtoMembers(chat.Members),
		Messages: FromMessagesToProtoMessages(chat.Messages),
	}
}

func FromAuthorizedUserArrayToUserArray(authorizedUsers []model.AuthorizedUser) []model.User {
	var users []model.User
	for _, user := range authorizedUsers {
		users = append(users, model.User{
			Id:       user.Id,
			Avatar:   user.Avatar,
			Username: user.Username,
			Nickname: user.Nickname,
			Email:    user.Email,
			Status:   user.Status,
		})
	}

	return users
}

func FromAuthorizedUserToUser(user model.AuthorizedUser) model.User {
	return model.User{
		Id:       user.Id,
		Avatar:   user.Avatar,
		Username: user.Username,
		Nickname: user.Nickname,
		Email:    user.Email,
		Status:   user.Status,
	}
}
