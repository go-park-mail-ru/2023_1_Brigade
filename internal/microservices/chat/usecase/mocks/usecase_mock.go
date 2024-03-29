// Code generated by MockGen. DO NOT EDIT.
// Source: usecase.go

// Package mock_chat is a generated GoMock package.
package mock_chat

import (
	context "context"
	model "project/internal/model"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockUsecase is a mock of Usecase interface.
type MockUsecase struct {
	ctrl     *gomock.Controller
	recorder *MockUsecaseMockRecorder
}

// MockUsecaseMockRecorder is the mock recorder for MockUsecase.
type MockUsecaseMockRecorder struct {
	mock *MockUsecase
}

// NewMockUsecase creates a new mock instance.
func NewMockUsecase(ctrl *gomock.Controller) *MockUsecase {
	mock := &MockUsecase{ctrl: ctrl}
	mock.recorder = &MockUsecaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUsecase) EXPECT() *MockUsecaseMockRecorder {
	return m.recorder
}

// CheckExistUserInChat mocks base method.
func (m *MockUsecase) CheckExistUserInChat(ctx context.Context, chat model.Chat, userID uint64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckExistUserInChat", ctx, chat, userID)
	ret0, _ := ret[0].(error)
	return ret0
}

// CheckExistUserInChat indicates an expected call of CheckExistUserInChat.
func (mr *MockUsecaseMockRecorder) CheckExistUserInChat(ctx, chat, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckExistUserInChat", reflect.TypeOf((*MockUsecase)(nil).CheckExistUserInChat), ctx, chat, userID)
}

// CreateChat mocks base method.
func (m *MockUsecase) CreateChat(ctx context.Context, chat model.CreateChat, userID uint64) (model.Chat, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateChat", ctx, chat, userID)
	ret0, _ := ret[0].(model.Chat)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateChat indicates an expected call of CreateChat.
func (mr *MockUsecaseMockRecorder) CreateChat(ctx, chat, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateChat", reflect.TypeOf((*MockUsecase)(nil).CreateChat), ctx, chat, userID)
}

// DeleteChatById mocks base method.
func (m *MockUsecase) DeleteChatById(ctx context.Context, chatID uint64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteChatById", ctx, chatID)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteChatById indicates an expected call of DeleteChatById.
func (mr *MockUsecaseMockRecorder) DeleteChatById(ctx, chatID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteChatById", reflect.TypeOf((*MockUsecase)(nil).DeleteChatById), ctx, chatID)
}

// EditChat mocks base method.
func (m *MockUsecase) EditChat(ctx context.Context, editChat model.EditChat) (model.Chat, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EditChat", ctx, editChat)
	ret0, _ := ret[0].(model.Chat)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// EditChat indicates an expected call of EditChat.
func (mr *MockUsecaseMockRecorder) EditChat(ctx, editChat interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EditChat", reflect.TypeOf((*MockUsecase)(nil).EditChat), ctx, editChat)
}

// GetChatById mocks base method.
func (m *MockUsecase) GetChatById(ctx context.Context, chatID, userID uint64) (model.Chat, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetChatById", ctx, chatID, userID)
	ret0, _ := ret[0].(model.Chat)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetChatById indicates an expected call of GetChatById.
func (mr *MockUsecaseMockRecorder) GetChatById(ctx, chatID, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetChatById", reflect.TypeOf((*MockUsecase)(nil).GetChatById), ctx, chatID, userID)
}

// GetListUserChats mocks base method.
func (m *MockUsecase) GetListUserChats(ctx context.Context, userID uint64) ([]model.ChatInListUser, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetListUserChats", ctx, userID)
	ret0, _ := ret[0].([]model.ChatInListUser)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetListUserChats indicates an expected call of GetListUserChats.
func (mr *MockUsecaseMockRecorder) GetListUserChats(ctx, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetListUserChats", reflect.TypeOf((*MockUsecase)(nil).GetListUserChats), ctx, userID)
}

// GetSearchChatsMessagesChannels mocks base method.
func (m *MockUsecase) GetSearchChatsMessagesChannels(ctx context.Context, userID uint64, string string) (model.FoundedChatsMessagesChannels, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSearchChatsMessagesChannels", ctx, userID, string)
	ret0, _ := ret[0].(model.FoundedChatsMessagesChannels)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSearchChatsMessagesChannels indicates an expected call of GetSearchChatsMessagesChannels.
func (mr *MockUsecaseMockRecorder) GetSearchChatsMessagesChannels(ctx, userID, string interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSearchChatsMessagesChannels", reflect.TypeOf((*MockUsecase)(nil).GetSearchChatsMessagesChannels), ctx, userID, string)
}
