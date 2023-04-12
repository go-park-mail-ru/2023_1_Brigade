// Code generated by MockGen. DO NOT EDIT.
// Source: usecase.go

// Package mock_user is a generated GoMock package.
package mock_user

import (
model "project/internal/model"
reflect "reflect"

gomock "github.com/golang/mock/gomock"
echo "github.com/labstack/echo/v4"
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

// AddUserContact mocks base method.
func (m *MockUsecase) AddUserContact(ctx echo.Context, userID, contactID uint64) ([]model.User, error) {
m.ctrl.T.Helper()
ret := m.ctrl.Call(m, "AddUserContact", ctx, userID, contactID)
ret0, _ := ret[0].([]model.User)
ret1, _ := ret[1].(error)
return ret0, ret1
}

// AddUserContact indicates an expected call of AddUserContact.
func (mr *MockUsecaseMockRecorder) AddUserContact(ctx, userID, contactID interface{}) *gomock.Call {
mr.mock.ctrl.T.Helper()
return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddUserContact", reflect.TypeOf((*MockUsecase)(nil).AddUserContact), ctx, userID, contactID)
}

// CheckExistUserById mocks base method.
func (m *MockUsecase) CheckExistUserById(ctx echo.Context, userID uint64) error {
m.ctrl.T.Helper()
ret := m.ctrl.Call(m, "CheckExistUserById", ctx, userID)
ret0, _ := ret[0].(error)
return ret0
}

// CheckExistUserById indicates an expected call of CheckExistUserById.
func (mr *MockUsecaseMockRecorder) CheckExistUserById(ctx, userID interface{}) *gomock.Call {
mr.mock.ctrl.T.Helper()
return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckExistUserById", reflect.TypeOf((*MockUsecase)(nil).CheckExistUserById), ctx, userID)
}

// DeleteUserById mocks base method.
func (m *MockUsecase) DeleteUserById(ctx echo.Context, userID uint64) error {
m.ctrl.T.Helper()
ret := m.ctrl.Call(m, "DeleteUserById", ctx, userID)
ret0, _ := ret[0].(error)
return ret0
}

// DeleteUserById indicates an expected call of DeleteUserById.
func (mr *MockUsecaseMockRecorder) DeleteUserById(ctx, userID interface{}) *gomock.Call {
mr.mock.ctrl.T.Helper()
return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteUserById", reflect.TypeOf((*MockUsecase)(nil).DeleteUserById), ctx, userID)
}

// GetAllUsersExceptCurrentUser mocks base method.
func (m *MockUsecase) GetAllUsersExceptCurrentUser(ctx echo.Context, userID uint64) ([]model.User, error) {
m.ctrl.T.Helper()
ret := m.ctrl.Call(m, "GetAllUsersExceptCurrentUser", ctx, userID)
ret0, _ := ret[0].([]model.User)
ret1, _ := ret[1].(error)
return ret0, ret1
}

// GetAllUsersExceptCurrentUser indicates an expected call of GetAllUsersExceptCurrentUser.
func (mr *MockUsecaseMockRecorder) GetAllUsersExceptCurrentUser(ctx, userID interface{}) *gomock.Call {
mr.mock.ctrl.T.Helper()
return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllUsersExceptCurrentUser", reflect.TypeOf((*MockUsecase)(nil).GetAllUsersExceptCurrentUser), ctx, userID)
}

// GetUserById mocks base method.
func (m *MockUsecase) GetUserById(ctx echo.Context, userID uint64) (model.User, error) {
m.ctrl.T.Helper()
ret := m.ctrl.Call(m, "GetUserById", ctx, userID)
ret0, _ := ret[0].(model.User)
ret1, _ := ret[1].(error)
return ret0, ret1
}

// GetUserById indicates an expected call of GetUserById.
func (mr *MockUsecaseMockRecorder) GetUserById(ctx, userID interface{}) *gomock.Call {
mr.mock.ctrl.T.Helper()
return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserById", reflect.TypeOf((*MockUsecase)(nil).GetUserById), ctx, userID)
}

// GetUserContacts mocks base method.
func (m *MockUsecase) GetUserContacts(ctx echo.Context, userID uint64) ([]model.User, error) {
m.ctrl.T.Helper()
ret := m.ctrl.Call(m, "GetUserContacts", ctx, userID)
ret0, _ := ret[0].([]model.User)
ret1, _ := ret[1].(error)
return ret0, ret1
}

// GetUserContacts indicates an expected call of GetUserContacts.
func (mr *MockUsecaseMockRecorder) GetUserContacts(ctx, userID interface{}) *gomock.Call {
mr.mock.ctrl.T.Helper()
return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserContacts", reflect.TypeOf((*MockUsecase)(nil).GetUserContacts), ctx, userID)
}

// PutUserById mocks base method.
func (m *MockUsecase) PutUserById(ctx echo.Context, user model.UpdateUser, userID uint64) (model.User, error) {
m.ctrl.T.Helper()
ret := m.ctrl.Call(m, "PutUserById", ctx, user, userID)
ret0, _ := ret[0].(model.User)
ret1, _ := ret[1].(error)
return ret0, ret1
}

// PutUserById indicates an expected call of PutUserById.
func (mr *MockUsecaseMockRecorder) PutUserById(ctx, user, userID interface{}) *gomock.Call {
mr.mock.ctrl.T.Helper()
return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PutUserById", reflect.TypeOf((*MockUsecase)(nil).PutUserById), ctx, user, userID)
}
