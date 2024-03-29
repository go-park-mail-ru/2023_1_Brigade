// Code generated by MockGen. DO NOT EDIT.
// Source: repository.go

// Package mock_messages is a generated GoMock package.
package mock_messages

import (
	context "context"
	model "project/internal/model"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockRepository is a mock of Repository interface.
type MockRepository struct {
	ctrl     *gomock.Controller
	recorder *MockRepositoryMockRecorder
}

// MockRepositoryMockRecorder is the mock recorder for MockRepository.
type MockRepositoryMockRecorder struct {
	mock *MockRepository
}

// NewMockRepository creates a new mock instance.
func NewMockRepository(ctrl *gomock.Controller) *MockRepository {
	mock := &MockRepository{ctrl: ctrl}
	mock.recorder = &MockRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRepository) EXPECT() *MockRepositoryMockRecorder {
	return m.recorder
}

// DeleteMessageById mocks base method.
func (m *MockRepository) DeleteMessageById(ctx context.Context, messageID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteMessageById", ctx, messageID)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteMessageById indicates an expected call of DeleteMessageById.
func (mr *MockRepositoryMockRecorder) DeleteMessageById(ctx, messageID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteMessageById", reflect.TypeOf((*MockRepository)(nil).DeleteMessageById), ctx, messageID)
}

// EditMessageById mocks base method.
func (m *MockRepository) EditMessageById(ctx context.Context, producerMessage model.ProducerMessage) (model.Message, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EditMessageById", ctx, producerMessage)
	ret0, _ := ret[0].(model.Message)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// EditMessageById indicates an expected call of EditMessageById.
func (mr *MockRepositoryMockRecorder) EditMessageById(ctx, producerMessage interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EditMessageById", reflect.TypeOf((*MockRepository)(nil).EditMessageById), ctx, producerMessage)
}

// GetChatMessages mocks base method.
func (m *MockRepository) GetChatMessages(ctx context.Context, chatID uint64) ([]model.ChatMessages, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetChatMessages", ctx, chatID)
	ret0, _ := ret[0].([]model.ChatMessages)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetChatMessages indicates an expected call of GetChatMessages.
func (mr *MockRepositoryMockRecorder) GetChatMessages(ctx, chatID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetChatMessages", reflect.TypeOf((*MockRepository)(nil).GetChatMessages), ctx, chatID)
}

// GetLastChatMessage mocks base method.
func (m *MockRepository) GetLastChatMessage(ctx context.Context, chatID uint64) (model.Message, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLastChatMessage", ctx, chatID)
	ret0, _ := ret[0].(model.Message)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetLastChatMessage indicates an expected call of GetLastChatMessage.
func (mr *MockRepositoryMockRecorder) GetLastChatMessage(ctx, chatID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLastChatMessage", reflect.TypeOf((*MockRepository)(nil).GetLastChatMessage), ctx, chatID)
}

// GetMessageById mocks base method.
func (m *MockRepository) GetMessageById(ctx context.Context, messageID string) (model.Message, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMessageById", ctx, messageID)
	ret0, _ := ret[0].(model.Message)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMessageById indicates an expected call of GetMessageById.
func (mr *MockRepositoryMockRecorder) GetMessageById(ctx, messageID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMessageById", reflect.TypeOf((*MockRepository)(nil).GetMessageById), ctx, messageID)
}

// GetSearchMessages mocks base method.
func (m *MockRepository) GetSearchMessages(ctx context.Context, userID uint64, string string) ([]model.Message, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSearchMessages", ctx, userID, string)
	ret0, _ := ret[0].([]model.Message)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSearchMessages indicates an expected call of GetSearchMessages.
func (mr *MockRepositoryMockRecorder) GetSearchMessages(ctx, userID, string interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSearchMessages", reflect.TypeOf((*MockRepository)(nil).GetSearchMessages), ctx, userID, string)
}

// InsertMessageInDB mocks base method.
func (m *MockRepository) InsertMessageInDB(ctx context.Context, message model.Message) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertMessageInDB", ctx, message)
	ret0, _ := ret[0].(error)
	return ret0
}

// InsertMessageInDB indicates an expected call of InsertMessageInDB.
func (mr *MockRepositoryMockRecorder) InsertMessageInDB(ctx, message interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertMessageInDB", reflect.TypeOf((*MockRepository)(nil).InsertMessageInDB), ctx, message)
}
