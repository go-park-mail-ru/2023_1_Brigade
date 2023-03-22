// Code generated by MockGen. DO NOT EDIT.
// Source: authservice.go

// Package mockAuthService is a generated GoMock package.
package mocks

import (
	"github.com/labstack/echo/v4"
	"project/internal/model"
	reflect "reflect"
	gomock "github.com/golang/mock/gomock"
)

// MockAuthService is a mock of AuthService interface.
type MockRepository struct {
	ctrl     *gomock.Controller
	recorder *MockRepositoryMockRecorder
}

// MockAuthServiceMockRecorder is the mock recorder for MockAuthService.
type MockRepositoryMockRecorder struct {
	mock *MockRepository
}

// NewMockAuthService creates a new mock instance.
func NewMockRepository(ctrl *gomock.Controller) *MockRepository {
	mock := &MockRepository{ctrl: ctrl}
	mock.recorder = &MockRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRepository) EXPECT() *MockRepositoryMockRecorder {
	return m.recorder
}

// CreateUser mocks base method.
func (m *MockRepository) CreateUser(ctx echo.Context, user model.User) (model.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", ctx, user)
	ret0, _ := ret[0].(model.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUser indicates an expected call of Auth.
func (mr *MockRepositoryMockRecorder) CreateUser(ctx echo.Context, user model.User) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockRepository)(nil).CreateUser), ctx, user)
}

// CheckCorrectPassword mocks base method.
func (m *MockRepository) CheckCorrectPassword(ctx echo.Context, user model.User) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckCorrectPassword", ctx, user)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckCorrectPassword indicates an expected call of Auth.
func (mr *MockRepositoryMockRecorder) CheckCorrectPassword(ctx echo.Context, user model.User) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckCorrectPassword", reflect.TypeOf((*MockRepository)(nil).CheckCorrectPassword), ctx, user)
}

// CheckExistEmail mocks base method.
func (m *MockRepository) CheckExistEmail(ctx echo.Context, email string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckExistEmail", ctx, email)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckExistEmail indicates an expected call of Auth.
func (mr *MockRepositoryMockRecorder) CheckExistEmail(ctx echo.Context, email string) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckExistEmail", reflect.TypeOf((*MockRepository)(nil).CheckExistEmail), ctx, email)
}

// CheckExistUsername mocks base method.
func (m *MockRepository) CheckExistUsername(ctx echo.Context, username string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckExistUsername", ctx, username)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckExistUsername indicates an expected call of Auth.
func (mr *MockRepositoryMockRecorder) CheckExistUsername(ctx echo.Context, username string) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckExistUsername", reflect.TypeOf((*MockRepository)(nil).CheckExistUsername), ctx, username)
}

// GetSessionByCookie mocks base method.
func (m *MockRepository) GetSessionByCookie(ctx echo.Context, cookie string) (model.Session, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSessionByCookie", ctx, cookie)
	ret0, _ := ret[0].(model.Session)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSessionByCookie indicates an expected call of Auth.
func (mr *MockRepositoryMockRecorder) GetSessionByCookie(ctx echo.Context, cookie string) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSessionByCookie", reflect.TypeOf((*MockRepository)(nil).GetSessionByCookie), ctx, cookie)
}

// CreateSession mocks base method.
func (m *MockRepository) CreateSession(ctx echo.Context, session model.Session) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateSession", ctx, session)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateSession indicates an expected call of Auth.
func (mr *MockRepositoryMockRecorder) CreateSession(ctx echo.Context, session model.Session) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateSession", reflect.TypeOf((*MockRepository)(nil).CreateSession), ctx, session)
}

// DeleteSession mocks base method.
func (m *MockRepository) DeleteSession(ctx echo.Context, cookie string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteSession", ctx, cookie)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteSession indicates an expected call of Auth.
func (mr *MockRepositoryMockRecorder) DeleteSession(ctx echo.Context, cookie string) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteSession", reflect.TypeOf((*MockRepository)(nil).DeleteSession), ctx, cookie)
}
