// Code generated by MockGen. DO NOT EDIT.
// Source: usecase.go

// Package mock_usecase is a generated GoMock package.
package mock_usecase

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

// ProduceMessage mocks base method.
func (m *MockUsecase) ProduceMessage(ctx context.Context, message model.ProducerMessage) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ProduceMessage", ctx, message)
	ret0, _ := ret[0].(error)
	return ret0
}

// ProduceMessage indicates an expected call of ProduceMessage.
func (mr *MockUsecaseMockRecorder) ProduceMessage(ctx, message interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProduceMessage", reflect.TypeOf((*MockUsecase)(nil).ProduceMessage), ctx, message)
}
