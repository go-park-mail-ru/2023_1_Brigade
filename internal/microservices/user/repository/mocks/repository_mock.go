// Code generated by MockGen. DO NOT EDIT.
// Source: repository.go

// Package mock_user is a generated GoMock package.
package mock_user

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

// AddUserInContact mocks base method.
func (m *MockRepository) AddUserInContact(ctx context.Context, contact model.UserContact) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddUserInContact", ctx, contact)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddUserInContact indicates an expected call of AddUserInContact.
func (mr *MockRepositoryMockRecorder) AddUserInContact(ctx, contact interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddUserInContact", reflect.TypeOf((*MockRepository)(nil).AddUserInContact), ctx, contact)
}

// CheckExistUserById mocks base method.
func (m *MockRepository) CheckExistUserById(ctx context.Context, userID uint64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckExistUserById", ctx, userID)
	ret0, _ := ret[0].(error)
	return ret0
}

// CheckExistUserById indicates an expected call of CheckExistUserById.
func (mr *MockRepositoryMockRecorder) CheckExistUserById(ctx, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckExistUserById", reflect.TypeOf((*MockRepository)(nil).CheckExistUserById), ctx, userID)
}

// CheckUserIsContact mocks base method.
func (m *MockRepository) CheckUserIsContact(ctx context.Context, contact model.UserContact) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckUserIsContact", ctx, contact)
	ret0, _ := ret[0].(error)
	return ret0
}

// CheckUserIsContact indicates an expected call of CheckUserIsContact.
func (mr *MockRepositoryMockRecorder) CheckUserIsContact(ctx, contact interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckUserIsContact", reflect.TypeOf((*MockRepository)(nil).CheckUserIsContact), ctx, contact)
}

// DeleteUserById mocks base method.
func (m *MockRepository) DeleteUserById(ctx context.Context, userID uint64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteUserById", ctx, userID)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteUserById indicates an expected call of DeleteUserById.
func (mr *MockRepositoryMockRecorder) DeleteUserById(ctx, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteUserById", reflect.TypeOf((*MockRepository)(nil).DeleteUserById), ctx, userID)
}

// GetAllUsersExceptCurrentUser mocks base method.
func (m *MockRepository) GetAllUsersExceptCurrentUser(ctx context.Context, userID uint64) ([]model.AuthorizedUser, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllUsersExceptCurrentUser", ctx, userID)
	ret0, _ := ret[0].([]model.AuthorizedUser)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllUsersExceptCurrentUser indicates an expected call of GetAllUsersExceptCurrentUser.
func (mr *MockRepositoryMockRecorder) GetAllUsersExceptCurrentUser(ctx, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllUsersExceptCurrentUser", reflect.TypeOf((*MockRepository)(nil).GetAllUsersExceptCurrentUser), ctx, userID)
}

// GetSearchUsers mocks base method.
func (m *MockRepository) GetSearchUsers(ctx context.Context, string string) ([]model.AuthorizedUser, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSearchUsers", ctx, string)
	ret0, _ := ret[0].([]model.AuthorizedUser)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSearchUsers indicates an expected call of GetSearchUsers.
func (mr *MockRepositoryMockRecorder) GetSearchUsers(ctx, string interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSearchUsers", reflect.TypeOf((*MockRepository)(nil).GetSearchUsers), ctx, string)
}

// GetUserByEmail mocks base method.
func (m *MockRepository) GetUserByEmail(ctx context.Context, email string) (model.AuthorizedUser, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserByEmail", ctx, email)
	ret0, _ := ret[0].(model.AuthorizedUser)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserByEmail indicates an expected call of GetUserByEmail.
func (mr *MockRepositoryMockRecorder) GetUserByEmail(ctx, email interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserByEmail", reflect.TypeOf((*MockRepository)(nil).GetUserByEmail), ctx, email)
}

// GetUserById mocks base method.
func (m *MockRepository) GetUserById(ctx context.Context, userID uint64) (model.AuthorizedUser, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserById", ctx, userID)
	ret0, _ := ret[0].(model.AuthorizedUser)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserById indicates an expected call of GetUserById.
func (mr *MockRepositoryMockRecorder) GetUserById(ctx, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserById", reflect.TypeOf((*MockRepository)(nil).GetUserById), ctx, userID)
}

// GetUserContacts mocks base method.
func (m *MockRepository) GetUserContacts(ctx context.Context, userID uint64) ([]model.AuthorizedUser, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserContacts", ctx, userID)
	ret0, _ := ret[0].([]model.AuthorizedUser)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserContacts indicates an expected call of GetUserContacts.
func (mr *MockRepositoryMockRecorder) GetUserContacts(ctx, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserContacts", reflect.TypeOf((*MockRepository)(nil).GetUserContacts), ctx, userID)
}

// UpdateUserAvatarNicknameById mocks base method.
func (m *MockRepository) UpdateUserAvatarNicknameById(ctx context.Context, user model.AuthorizedUser) (model.AuthorizedUser, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUserAvatarNicknameById", ctx, user)
	ret0, _ := ret[0].(model.AuthorizedUser)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateUserAvatarNicknameById indicates an expected call of UpdateUserAvatarNicknameById.
func (mr *MockRepositoryMockRecorder) UpdateUserAvatarNicknameById(ctx, user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUserAvatarNicknameById", reflect.TypeOf((*MockRepository)(nil).UpdateUserAvatarNicknameById), ctx, user)
}

// UpdateUserEmailStatusById mocks base method.
func (m *MockRepository) UpdateUserEmailStatusById(ctx context.Context, user model.AuthorizedUser) (model.AuthorizedUser, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUserEmailStatusById", ctx, user)
	ret0, _ := ret[0].(model.AuthorizedUser)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateUserEmailStatusById indicates an expected call of UpdateUserEmailStatusById.
func (mr *MockRepositoryMockRecorder) UpdateUserEmailStatusById(ctx, user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUserEmailStatusById", reflect.TypeOf((*MockRepository)(nil).UpdateUserEmailStatusById), ctx, user)
}

// UpdateUserPasswordById mocks base method.
func (m *MockRepository) UpdateUserPasswordById(ctx context.Context, user model.AuthorizedUser) (model.AuthorizedUser, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUserPasswordById", ctx, user)
	ret0, _ := ret[0].(model.AuthorizedUser)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateUserPasswordById indicates an expected call of UpdateUserPasswordById.
func (mr *MockRepositoryMockRecorder) UpdateUserPasswordById(ctx, user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUserPasswordById", reflect.TypeOf((*MockRepository)(nil).UpdateUserPasswordById), ctx, user)
}
