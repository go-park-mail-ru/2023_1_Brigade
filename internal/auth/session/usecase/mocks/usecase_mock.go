package mocks

import (
	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"project/internal/model"
	"reflect"
)

// MockAuthService is a mock of AuthService interface.
type MockUsecase struct {
	ctrl     *gomock.Controller
	recorder *MockUsecaseMockRecorder
}

// MockAuthServiceMockRecorder is the mock recorder for MockAuthService.
type MockUsecaseMockRecorder struct {
	mock *MockUsecase
}

// NewMockAuthService creates a new mock instance.
func NewMockUsecase(ctrl *gomock.Controller) *MockUsecase {
	mock := &MockUsecase{ctrl: ctrl}
	mock.recorder = &MockUsecaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUsecase) EXPECT() *MockUsecaseMockRecorder {
	return m.recorder
}

// GetSessionByCookie mocks base method.
func (m *MockUsecase) GetSessionByCookie(ctx echo.Context, cookie string) (model.Session, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSessionByCookie", ctx, cookie)
	ret0, _ := ret[0].(model.Session)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSessionByCookie indicates an expected call of Auth.
func (mr *MockUsecaseMockRecorder) GetSessionByCookie(ctx echo.Context, cookie string) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSessionByCookie", reflect.TypeOf((*MockUsecase)(nil).GetSessionByCookie), ctx, cookie)
}

// CreateSessionById mocks base method.
func (m *MockUsecase) CreateSessionById(ctx echo.Context, userID uint64) (model.Session, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateSessionById", ctx, userID)
	ret0, _ := ret[0].(model.Session)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateSessionById indicates an expected call of Auth.
func (mr *MockUsecaseMockRecorder) CreateSessionById(ctx echo.Context, userID uint64) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateSessionById", reflect.TypeOf((*MockUsecase)(nil).CreateSessionById), ctx, userID)
}

// DeleteSessionByCookie mocks base method.
func (m *MockUsecase) DeleteSessionByCookie(ctx echo.Context, cookie string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteSessionByCookie", ctx, cookie)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteSessionByCookie indicates an expected call of Auth.
func (mr *MockUsecaseMockRecorder) DeleteSessionByCookie(ctx echo.Context, cookie string) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteSessionByCookie", reflect.TypeOf((*MockUsecase)(nil).DeleteSessionByCookie), ctx, cookie)
}
