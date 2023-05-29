// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.22.3
// source: protobuf/chat.proto

package generated

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Chat struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       uint64     `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	MasterID uint64     `protobuf:"varint,2,opt,name=MasterID,proto3" json:"MasterID,omitempty"`
	Type     uint64     `protobuf:"varint,3,opt,name=Type,proto3" json:"Type,omitempty"`
	Title    string     `protobuf:"bytes,4,opt,name=Title,proto3" json:"Title,omitempty"`
	Avatar   string     `protobuf:"bytes,5,opt,name=Avatar,proto3" json:"Avatar,omitempty"`
	Members  []*User    `protobuf:"bytes,6,rep,name=Members,proto3" json:"Members,omitempty"`
	Messages []*Message `protobuf:"bytes,7,rep,name=Messages,proto3" json:"Messages,omitempty"`
}

func (x *Chat) Reset() {
	*x = Chat{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_chat_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Chat) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Chat) ProtoMessage() {}

func (x *Chat) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_chat_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Chat.ProtoReflect.Descriptor instead.
func (*Chat) Descriptor() ([]byte, []int) {
	return file_protobuf_chat_proto_rawDescGZIP(), []int{0}
}

func (x *Chat) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Chat) GetMasterID() uint64 {
	if x != nil {
		return x.MasterID
	}
	return 0
}

func (x *Chat) GetType() uint64 {
	if x != nil {
		return x.Type
	}
	return 0
}

func (x *Chat) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Chat) GetAvatar() string {
	if x != nil {
		return x.Avatar
	}
	return ""
}

func (x *Chat) GetMembers() []*User {
	if x != nil {
		return x.Members
	}
	return nil
}

func (x *Chat) GetMessages() []*Message {
	if x != nil {
		return x.Messages
	}
	return nil
}

type EditChatModel struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      uint64   `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	Type    uint64   `protobuf:"varint,2,opt,name=Type,proto3" json:"Type,omitempty"`
	Title   string   `protobuf:"bytes,3,opt,name=Title,proto3" json:"Title,omitempty"`
	Members []uint64 `protobuf:"varint,4,rep,packed,name=Members,proto3" json:"Members,omitempty"`
}

func (x *EditChatModel) Reset() {
	*x = EditChatModel{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_chat_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EditChatModel) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EditChatModel) ProtoMessage() {}

func (x *EditChatModel) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_chat_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EditChatModel.ProtoReflect.Descriptor instead.
func (*EditChatModel) Descriptor() ([]byte, []int) {
	return file_protobuf_chat_proto_rawDescGZIP(), []int{1}
}

func (x *EditChatModel) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *EditChatModel) GetType() uint64 {
	if x != nil {
		return x.Type
	}
	return 0
}

func (x *EditChatModel) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *EditChatModel) GetMembers() []uint64 {
	if x != nil {
		return x.Members
	}
	return nil
}

type CreateChat struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type    uint64   `protobuf:"varint,1,opt,name=Type,proto3" json:"Type,omitempty"`
	Title   string   `protobuf:"bytes,2,opt,name=Title,proto3" json:"Title,omitempty"`
	Avatar  string   `protobuf:"bytes,3,opt,name=Avatar,proto3" json:"Avatar,omitempty"`
	Members []uint64 `protobuf:"varint,4,rep,packed,name=Members,proto3" json:"Members,omitempty"`
}

func (x *CreateChat) Reset() {
	*x = CreateChat{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_chat_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateChat) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateChat) ProtoMessage() {}

func (x *CreateChat) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_chat_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateChat.ProtoReflect.Descriptor instead.
func (*CreateChat) Descriptor() ([]byte, []int) {
	return file_protobuf_chat_proto_rawDescGZIP(), []int{2}
}

func (x *CreateChat) GetType() uint64 {
	if x != nil {
		return x.Type
	}
	return 0
}

func (x *CreateChat) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *CreateChat) GetAvatar() string {
	if x != nil {
		return x.Avatar
	}
	return ""
}

func (x *CreateChat) GetMembers() []uint64 {
	if x != nil {
		return x.Members
	}
	return nil
}

type GetChatArguments struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ChatID uint64 `protobuf:"varint,1,opt,name=ChatID,proto3" json:"ChatID,omitempty"`
	UserID uint64 `protobuf:"varint,2,opt,name=UserID,proto3" json:"UserID,omitempty"`
}

func (x *GetChatArguments) Reset() {
	*x = GetChatArguments{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_chat_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetChatArguments) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetChatArguments) ProtoMessage() {}

func (x *GetChatArguments) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_chat_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetChatArguments.ProtoReflect.Descriptor instead.
func (*GetChatArguments) Descriptor() ([]byte, []int) {
	return file_protobuf_chat_proto_rawDescGZIP(), []int{3}
}

func (x *GetChatArguments) GetChatID() uint64 {
	if x != nil {
		return x.ChatID
	}
	return 0
}

func (x *GetChatArguments) GetUserID() uint64 {
	if x != nil {
		return x.UserID
	}
	return 0
}

type ChatID struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ChatID uint64 `protobuf:"varint,1,opt,name=ChatID,proto3" json:"ChatID,omitempty"`
}

func (x *ChatID) Reset() {
	*x = ChatID{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_chat_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChatID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChatID) ProtoMessage() {}

func (x *ChatID) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_chat_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChatID.ProtoReflect.Descriptor instead.
func (*ChatID) Descriptor() ([]byte, []int) {
	return file_protobuf_chat_proto_rawDescGZIP(), []int{4}
}

func (x *ChatID) GetChatID() uint64 {
	if x != nil {
		return x.ChatID
	}
	return 0
}

type CreateChatArguments struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Chat   *CreateChat `protobuf:"bytes,1,opt,name=Chat,proto3" json:"Chat,omitempty"`
	UserID *UserID     `protobuf:"bytes,2,opt,name=userID,proto3" json:"userID,omitempty"`
}

func (x *CreateChatArguments) Reset() {
	*x = CreateChatArguments{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_chat_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateChatArguments) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateChatArguments) ProtoMessage() {}

func (x *CreateChatArguments) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_chat_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateChatArguments.ProtoReflect.Descriptor instead.
func (*CreateChatArguments) Descriptor() ([]byte, []int) {
	return file_protobuf_chat_proto_rawDescGZIP(), []int{5}
}

func (x *CreateChatArguments) GetChat() *CreateChat {
	if x != nil {
		return x.Chat
	}
	return nil
}

func (x *CreateChatArguments) GetUserID() *UserID {
	if x != nil {
		return x.UserID
	}
	return nil
}

type ExistChatArguments struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Chat   *Chat   `protobuf:"bytes,1,opt,name=Chat,proto3" json:"Chat,omitempty"`
	UserID *UserID `protobuf:"bytes,2,opt,name=userID,proto3" json:"userID,omitempty"`
}

func (x *ExistChatArguments) Reset() {
	*x = ExistChatArguments{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_chat_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExistChatArguments) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExistChatArguments) ProtoMessage() {}

func (x *ExistChatArguments) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_chat_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExistChatArguments.ProtoReflect.Descriptor instead.
func (*ExistChatArguments) Descriptor() ([]byte, []int) {
	return file_protobuf_chat_proto_rawDescGZIP(), []int{6}
}

func (x *ExistChatArguments) GetChat() *Chat {
	if x != nil {
		return x.Chat
	}
	return nil
}

func (x *ExistChatArguments) GetUserID() *UserID {
	if x != nil {
		return x.UserID
	}
	return nil
}

type ChatInListUser struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                uint64   `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	Type              uint64   `protobuf:"varint,2,opt,name=Type,proto3" json:"Type,omitempty"`
	Title             string   `protobuf:"bytes,3,opt,name=Title,proto3" json:"Title,omitempty"`
	Avatar            string   `protobuf:"bytes,4,opt,name=Avatar,proto3" json:"Avatar,omitempty"`
	Members           []*User  `protobuf:"bytes,5,rep,name=Members,proto3" json:"Members,omitempty"`
	LastMessage       *Message `protobuf:"bytes,6,opt,name=LastMessage,proto3" json:"LastMessage,omitempty"`
	LastMessageAuthor *User    `protobuf:"bytes,7,opt,name=LastMessageAuthor,proto3" json:"LastMessageAuthor,omitempty"`
}

func (x *ChatInListUser) Reset() {
	*x = ChatInListUser{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_chat_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChatInListUser) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChatInListUser) ProtoMessage() {}

func (x *ChatInListUser) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_chat_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChatInListUser.ProtoReflect.Descriptor instead.
func (*ChatInListUser) Descriptor() ([]byte, []int) {
	return file_protobuf_chat_proto_rawDescGZIP(), []int{7}
}

func (x *ChatInListUser) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ChatInListUser) GetType() uint64 {
	if x != nil {
		return x.Type
	}
	return 0
}

func (x *ChatInListUser) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *ChatInListUser) GetAvatar() string {
	if x != nil {
		return x.Avatar
	}
	return ""
}

func (x *ChatInListUser) GetMembers() []*User {
	if x != nil {
		return x.Members
	}
	return nil
}

func (x *ChatInListUser) GetLastMessage() *Message {
	if x != nil {
		return x.LastMessage
	}
	return nil
}

func (x *ChatInListUser) GetLastMessageAuthor() *User {
	if x != nil {
		return x.LastMessageAuthor
	}
	return nil
}

type ArrayChatInListUser struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Chats []*ChatInListUser `protobuf:"bytes,1,rep,name=chats,proto3" json:"chats,omitempty"`
}

func (x *ArrayChatInListUser) Reset() {
	*x = ArrayChatInListUser{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_chat_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ArrayChatInListUser) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ArrayChatInListUser) ProtoMessage() {}

func (x *ArrayChatInListUser) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_chat_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ArrayChatInListUser.ProtoReflect.Descriptor instead.
func (*ArrayChatInListUser) Descriptor() ([]byte, []int) {
	return file_protobuf_chat_proto_rawDescGZIP(), []int{8}
}

func (x *ArrayChatInListUser) GetChats() []*ChatInListUser {
	if x != nil {
		return x.Chats
	}
	return nil
}

type FoundedChatsMessagesChannels struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FoundedChats    []*ChatInListUser `protobuf:"bytes,1,rep,name=FoundedChats,proto3" json:"FoundedChats,omitempty"`
	FoundedMessages []*ChatInListUser `protobuf:"bytes,2,rep,name=FoundedMessages,proto3" json:"FoundedMessages,omitempty"`
	FoundedChannels []*ChatInListUser `protobuf:"bytes,3,rep,name=FoundedChannels,proto3" json:"FoundedChannels,omitempty"`
	FoundedContacts *Contacts         `protobuf:"bytes,4,opt,name=FoundedContacts,proto3" json:"FoundedContacts,omitempty"`
}

func (x *FoundedChatsMessagesChannels) Reset() {
	*x = FoundedChatsMessagesChannels{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_chat_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FoundedChatsMessagesChannels) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FoundedChatsMessagesChannels) ProtoMessage() {}

func (x *FoundedChatsMessagesChannels) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_chat_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FoundedChatsMessagesChannels.ProtoReflect.Descriptor instead.
func (*FoundedChatsMessagesChannels) Descriptor() ([]byte, []int) {
	return file_protobuf_chat_proto_rawDescGZIP(), []int{9}
}

func (x *FoundedChatsMessagesChannels) GetFoundedChats() []*ChatInListUser {
	if x != nil {
		return x.FoundedChats
	}
	return nil
}

func (x *FoundedChatsMessagesChannels) GetFoundedMessages() []*ChatInListUser {
	if x != nil {
		return x.FoundedMessages
	}
	return nil
}

func (x *FoundedChatsMessagesChannels) GetFoundedChannels() []*ChatInListUser {
	if x != nil {
		return x.FoundedChannels
	}
	return nil
}

func (x *FoundedChatsMessagesChannels) GetFoundedContacts() *Contacts {
	if x != nil {
		return x.FoundedContacts
	}
	return nil
}

type SearchChatsArgumets struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserID  uint64 `protobuf:"varint,1,opt,name=UserID,proto3" json:"UserID,omitempty"`
	String_ string `protobuf:"bytes,2,opt,name=String,proto3" json:"String,omitempty"`
}

func (x *SearchChatsArgumets) Reset() {
	*x = SearchChatsArgumets{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_chat_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchChatsArgumets) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchChatsArgumets) ProtoMessage() {}

func (x *SearchChatsArgumets) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_chat_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchChatsArgumets.ProtoReflect.Descriptor instead.
func (*SearchChatsArgumets) Descriptor() ([]byte, []int) {
	return file_protobuf_chat_proto_rawDescGZIP(), []int{10}
}

func (x *SearchChatsArgumets) GetUserID() uint64 {
	if x != nil {
		return x.UserID
	}
	return 0
}

func (x *SearchChatsArgumets) GetString_() string {
	if x != nil {
		return x.String_
	}
	return ""
}

var File_protobuf_chat_proto protoreflect.FileDescriptor

var file_protobuf_chat_proto_rawDesc = []byte{
	0x0a, 0x13, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x63, 0x68, 0x61, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x1a,
	0x13, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xcd, 0x01,
	0x0a, 0x04, 0x43, 0x68, 0x61, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x02, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x4d, 0x61, 0x73, 0x74, 0x65, 0x72,
	0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x4d, 0x61, 0x73, 0x74, 0x65, 0x72,
	0x49, 0x44, 0x12, 0x12, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x16, 0x0a, 0x06,
	0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x41, 0x76,
	0x61, 0x74, 0x61, 0x72, 0x12, 0x28, 0x0a, 0x07, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x18,
	0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x07, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x12, 0x2d,
	0x0a, 0x08, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x11, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x52, 0x08, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x22, 0x63, 0x0a,
	0x0d, 0x45, 0x64, 0x69, 0x74, 0x43, 0x68, 0x61, 0x74, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x12, 0x0e,
	0x0a, 0x02, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x49, 0x64, 0x12, 0x12,
	0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x04, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x4d, 0x65, 0x6d, 0x62,
	0x65, 0x72, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x04, 0x52, 0x07, 0x4d, 0x65, 0x6d, 0x62, 0x65,
	0x72, 0x73, 0x22, 0x68, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x68, 0x61, 0x74,
	0x12, 0x12, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x04,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x41, 0x76,
	0x61, 0x74, 0x61, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x41, 0x76, 0x61, 0x74,
	0x61, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x18, 0x04, 0x20,
	0x03, 0x28, 0x04, 0x52, 0x07, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x22, 0x42, 0x0a, 0x10,
	0x47, 0x65, 0x74, 0x43, 0x68, 0x61, 0x74, 0x41, 0x72, 0x67, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x73,
	0x12, 0x16, 0x0a, 0x06, 0x43, 0x68, 0x61, 0x74, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x06, 0x43, 0x68, 0x61, 0x74, 0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72,
	0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44,
	0x22, 0x20, 0x0a, 0x06, 0x43, 0x68, 0x61, 0x74, 0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x43, 0x68,
	0x61, 0x74, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x43, 0x68, 0x61, 0x74,
	0x49, 0x44, 0x22, 0x69, 0x0a, 0x13, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x68, 0x61, 0x74,
	0x41, 0x72, 0x67, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x28, 0x0a, 0x04, 0x43, 0x68, 0x61,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x68, 0x61, 0x74, 0x52, 0x04, 0x43,
	0x68, 0x61, 0x74, 0x12, 0x28, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x55,
	0x73, 0x65, 0x72, 0x49, 0x44, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44, 0x22, 0x62, 0x0a,
	0x12, 0x45, 0x78, 0x69, 0x73, 0x74, 0x43, 0x68, 0x61, 0x74, 0x41, 0x72, 0x67, 0x75, 0x6d, 0x65,
	0x6e, 0x74, 0x73, 0x12, 0x22, 0x0a, 0x04, 0x43, 0x68, 0x61, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x43, 0x68, 0x61,
	0x74, 0x52, 0x04, 0x43, 0x68, 0x61, 0x74, 0x12, 0x28, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49,
	0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49,
	0x44, 0x22, 0xff, 0x01, 0x0a, 0x0e, 0x43, 0x68, 0x61, 0x74, 0x49, 0x6e, 0x4c, 0x69, 0x73, 0x74,
	0x55, 0x73, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x02, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x69, 0x74, 0x6c,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x16,
	0x0a, 0x06, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x12, 0x28, 0x0a, 0x07, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72,
	0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x07, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73,
	0x12, 0x33, 0x0a, 0x0b, 0x4c, 0x61, 0x73, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x0b, 0x4c, 0x61, 0x73, 0x74, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x3c, 0x0a, 0x11, 0x4c, 0x61, 0x73, 0x74, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x55, 0x73, 0x65, 0x72,
	0x52, 0x11, 0x4c, 0x61, 0x73, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x41, 0x75, 0x74,
	0x68, 0x6f, 0x72, 0x22, 0x45, 0x0a, 0x13, 0x41, 0x72, 0x72, 0x61, 0x79, 0x43, 0x68, 0x61, 0x74,
	0x49, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x55, 0x73, 0x65, 0x72, 0x12, 0x2e, 0x0a, 0x05, 0x63, 0x68,
	0x61, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x43, 0x68, 0x61, 0x74, 0x49, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x55,
	0x73, 0x65, 0x72, 0x52, 0x05, 0x63, 0x68, 0x61, 0x74, 0x73, 0x22, 0xa2, 0x02, 0x0a, 0x1c, 0x46,
	0x6f, 0x75, 0x6e, 0x64, 0x65, 0x64, 0x43, 0x68, 0x61, 0x74, 0x73, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x73, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x73, 0x12, 0x3c, 0x0a, 0x0c, 0x46,
	0x6f, 0x75, 0x6e, 0x64, 0x65, 0x64, 0x43, 0x68, 0x61, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x18, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x43, 0x68, 0x61,
	0x74, 0x49, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x0c, 0x46, 0x6f, 0x75,
	0x6e, 0x64, 0x65, 0x64, 0x43, 0x68, 0x61, 0x74, 0x73, 0x12, 0x42, 0x0a, 0x0f, 0x46, 0x6f, 0x75,
	0x6e, 0x64, 0x65, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x18, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x43, 0x68,
	0x61, 0x74, 0x49, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x0f, 0x46, 0x6f,
	0x75, 0x6e, 0x64, 0x65, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x12, 0x42, 0x0a,
	0x0f, 0x46, 0x6f, 0x75, 0x6e, 0x64, 0x65, 0x64, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x73,
	0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x43, 0x68, 0x61, 0x74, 0x49, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x55, 0x73, 0x65, 0x72,
	0x52, 0x0f, 0x46, 0x6f, 0x75, 0x6e, 0x64, 0x65, 0x64, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c,
	0x73, 0x12, 0x3c, 0x0a, 0x0f, 0x46, 0x6f, 0x75, 0x6e, 0x64, 0x65, 0x64, 0x43, 0x6f, 0x6e, 0x74,
	0x61, 0x63, 0x74, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x73, 0x52, 0x0f,
	0x46, 0x6f, 0x75, 0x6e, 0x64, 0x65, 0x64, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x73, 0x22,
	0x45, 0x0a, 0x13, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x43, 0x68, 0x61, 0x74, 0x73, 0x41, 0x72,
	0x67, 0x75, 0x6d, 0x65, 0x74, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x12, 0x16,
	0x0a, 0x06, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x42, 0x0d, 0x5a, 0x0b, 0x2e, 0x3b, 0x67, 0x65, 0x6e, 0x65,
	0x72, 0x61, 0x74, 0x65, 0x64, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protobuf_chat_proto_rawDescOnce sync.Once
	file_protobuf_chat_proto_rawDescData = file_protobuf_chat_proto_rawDesc
)

func file_protobuf_chat_proto_rawDescGZIP() []byte {
	file_protobuf_chat_proto_rawDescOnce.Do(func() {
		file_protobuf_chat_proto_rawDescData = protoimpl.X.CompressGZIP(file_protobuf_chat_proto_rawDescData)
	})
	return file_protobuf_chat_proto_rawDescData
}

var file_protobuf_chat_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_protobuf_chat_proto_goTypes = []interface{}{
	(*Chat)(nil),                         // 0: protobuf.Chat
	(*EditChatModel)(nil),                // 1: protobuf.EditChatModel
	(*CreateChat)(nil),                   // 2: protobuf.CreateChat
	(*GetChatArguments)(nil),             // 3: protobuf.GetChatArguments
	(*ChatID)(nil),                       // 4: protobuf.ChatID
	(*CreateChatArguments)(nil),          // 5: protobuf.CreateChatArguments
	(*ExistChatArguments)(nil),           // 6: protobuf.ExistChatArguments
	(*ChatInListUser)(nil),               // 7: protobuf.ChatInListUser
	(*ArrayChatInListUser)(nil),          // 8: protobuf.ArrayChatInListUser
	(*FoundedChatsMessagesChannels)(nil), // 9: protobuf.FoundedChatsMessagesChannels
	(*SearchChatsArgumets)(nil),          // 10: protobuf.SearchChatsArgumets
	(*User)(nil),                         // 11: protobuf.User
	(*Message)(nil),                      // 12: protobuf.Message
	(*UserID)(nil),                       // 13: protobuf.UserID
	(*Contacts)(nil),                     // 14: protobuf.Contacts
}
var file_protobuf_chat_proto_depIdxs = []int32{
	11, // 0: protobuf.Chat.Members:type_name -> protobuf.User
	12, // 1: protobuf.Chat.Messages:type_name -> protobuf.Message
	2,  // 2: protobuf.CreateChatArguments.Chat:type_name -> protobuf.CreateChat
	13, // 3: protobuf.CreateChatArguments.userID:type_name -> protobuf.UserID
	0,  // 4: protobuf.ExistChatArguments.Chat:type_name -> protobuf.Chat
	13, // 5: protobuf.ExistChatArguments.userID:type_name -> protobuf.UserID
	11, // 6: protobuf.ChatInListUser.Members:type_name -> protobuf.User
	12, // 7: protobuf.ChatInListUser.LastMessage:type_name -> protobuf.Message
	11, // 8: protobuf.ChatInListUser.LastMessageAuthor:type_name -> protobuf.User
	7,  // 9: protobuf.ArrayChatInListUser.chats:type_name -> protobuf.ChatInListUser
	7,  // 10: protobuf.FoundedChatsMessagesChannels.FoundedChats:type_name -> protobuf.ChatInListUser
	7,  // 11: protobuf.FoundedChatsMessagesChannels.FoundedMessages:type_name -> protobuf.ChatInListUser
	7,  // 12: protobuf.FoundedChatsMessagesChannels.FoundedChannels:type_name -> protobuf.ChatInListUser
	14, // 13: protobuf.FoundedChatsMessagesChannels.FoundedContacts:type_name -> protobuf.Contacts
	14, // [14:14] is the sub-list for method output_type
	14, // [14:14] is the sub-list for method input_type
	14, // [14:14] is the sub-list for extension type_name
	14, // [14:14] is the sub-list for extension extendee
	0,  // [0:14] is the sub-list for field type_name
}

func init() { file_protobuf_chat_proto_init() }
func file_protobuf_chat_proto_init() {
	if File_protobuf_chat_proto != nil {
		return
	}
	file_protobuf_user_proto_init()
	file_protobuf_messages_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_protobuf_chat_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Chat); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_protobuf_chat_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EditChatModel); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_protobuf_chat_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateChat); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_protobuf_chat_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetChatArguments); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_protobuf_chat_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChatID); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_protobuf_chat_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateChatArguments); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_protobuf_chat_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExistChatArguments); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_protobuf_chat_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChatInListUser); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_protobuf_chat_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ArrayChatInListUser); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_protobuf_chat_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FoundedChatsMessagesChannels); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_protobuf_chat_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchChatsArgumets); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_protobuf_chat_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_protobuf_chat_proto_goTypes,
		DependencyIndexes: file_protobuf_chat_proto_depIdxs,
		MessageInfos:      file_protobuf_chat_proto_msgTypes,
	}.Build()
	File_protobuf_chat_proto = out.File
	file_protobuf_chat_proto_rawDesc = nil
	file_protobuf_chat_proto_goTypes = nil
	file_protobuf_chat_proto_depIdxs = nil
}
