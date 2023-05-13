package model_conversion

import (
	protobuf "project/internal/generated"
	"project/internal/model"
	"time"
)

func FromRegistrationUserToProtoRegistrationUser(registrationUser model.RegistrationUser) *protobuf.RegistrationUser {
	return &protobuf.RegistrationUser{
		Nickname: registrationUser.Nickname,
		Email:    registrationUser.Email,
		Password: registrationUser.Password,
	}
}

func FromProtoRegistrationUserToRegistrationUser(registrationUser *protobuf.RegistrationUser) model.RegistrationUser {
	return model.RegistrationUser{
		Nickname: registrationUser.Nickname,
		Email:    registrationUser.Email,
		Password: registrationUser.Password,
	}
}

func FromLoginUserToProtoLoginUser(loginUser model.LoginUser) *protobuf.LoginUser {
	return &protobuf.LoginUser{
		Email:    loginUser.Email,
		Password: loginUser.Password,
	}
}

func FromProtoLoginUserToLoginUser(loginUser *protobuf.LoginUser) model.LoginUser {
	return model.LoginUser{
		Email:    loginUser.Email,
		Password: loginUser.Password,
	}
}

func FromBytesToProtoBytes(bytes []byte) *protobuf.Bytes {
	return &protobuf.Bytes{
		Bytes: bytes,
	}
}

func FromProtoBytesToBytes(bytes *protobuf.Bytes) []byte {
	return bytes.Bytes
}

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

func FromProtoProducerMessageToProducerMessage(message *protobuf.ProducerMessage) model.ProducerMessage {
	return model.ProducerMessage{
		Id:         message.Id,
		Type:       message.Type,
		Body:       message.Body,
		AuthorId:   message.AuthorId,
		ChatID:     message.ChatId,
		ReceiverID: message.ReceiverID,
		CreatedAt:  time.Time{},
	}
}

func FromProducerMessageToProtoProducerMessage(message model.ProducerMessage) *protobuf.ProducerMessage {
	return &protobuf.ProducerMessage{
		Id:         message.Id,
		Type:       message.Type,
		Body:       message.Body,
		AuthorId:   message.AuthorId,
		ChatId:     message.ChatID,
		ReceiverID: message.ReceiverID,
		CreatedAt:  message.CreatedAt.String(),
	}
}

func FromProtoMessageToMessage(message *protobuf.Message) model.Message {
	return model.Message{
		Id:        message.Id,
		Body:      message.Body,
		AuthorId:  message.AuthorId,
		ChatId:    message.ChatId,
		CreatedAt: time.Time{},
	}
}

func FromMessageToProtoMessage(message model.Message) *protobuf.Message {
	return &protobuf.Message{
		Id:        message.Id,
		Body:      message.Body,
		AuthorId:  message.AuthorId,
		ChatId:    message.ChatId,
		CreatedAt: message.CreatedAt.String(),
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

func FromProtoSearchChatsToSearchChats(protoChats *protobuf.FoundedChatsMessagesChannels) model.FoundedChatsMessagesChannels {
	foundedChats := make([]model.ChatInListUser, len(protoChats.FoundedChats))
	for idx, value := range protoChats.FoundedChats {
		foundedChats[idx] = FromProtoUserChatToUserChat(value)
	}

	foundedMessages := make([]model.ChatInListUser, len(protoChats.FoundedMessages))
	for idx, value := range protoChats.FoundedMessages {
		foundedMessages[idx] = FromProtoUserChatToUserChat(value)
	}

	foundedChannels := make([]model.ChatInListUser, len(protoChats.FoundedChannels))
	for idx, value := range protoChats.FoundedChannels {
		foundedChannels[idx] = FromProtoUserChatToUserChat(value)
	}

	foundedContacts := make([]model.User, len(protoChats.FoundedContacts.Contacts))
	for idx, value := range protoChats.FoundedContacts.Contacts {
		foundedContacts[idx] = FromProtoUserToUser(value)
	}

	return model.FoundedChatsMessagesChannels{
		FoundedChats:    foundedChats,
		FoundedMessages: foundedMessages,
		FoundedChannels: foundedChannels,
		FoundedContacts: foundedContacts,
	}
}

func FromSearchChatsToProtoSearchChats(chats model.FoundedChatsMessagesChannels) *protobuf.FoundedChatsMessagesChannels {
	foundedChats := make([]*protobuf.ChatInListUser, len(chats.FoundedChats))
	for idx, value := range chats.FoundedChats {
		foundedChats[idx] = FromUserChatToProtoUserChat(value)
	}

	foundedMessages := make([]*protobuf.ChatInListUser, len(chats.FoundedMessages))
	for idx, value := range chats.FoundedMessages {
		foundedMessages[idx] = FromUserChatToProtoUserChat(value)
	}

	foundedChannels := make([]*protobuf.ChatInListUser, len(chats.FoundedChannels))
	for idx, value := range chats.FoundedChannels {
		foundedChannels[idx] = FromUserChatToProtoUserChat(value)
	}

	foundedContacts := make([]*protobuf.User, len(chats.FoundedContacts))
	for idx, value := range chats.FoundedContacts {
		foundedContacts[idx] = FromUserToProtoUser(value)
	}

	return &protobuf.FoundedChatsMessagesChannels{
		FoundedChats:    foundedChats,
		FoundedMessages: foundedMessages,
		FoundedChannels: foundedChannels,
		FoundedContacts: &protobuf.Contacts{Contacts: foundedContacts},
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
		MasterID: chat.MasterID,
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
		MasterID: chat.MasterID,
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
