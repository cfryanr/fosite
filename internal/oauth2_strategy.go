// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/ory/fosite/handler/oauth2 (interfaces: CoreStrategy)

// Package internal is a generated GoMock package.
package internal

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"

	fosite "github.com/ory/fosite"
)

// MockCoreStrategy is a mock of CoreStrategy interface
type MockCoreStrategy struct {
	ctrl     *gomock.Controller
	recorder *MockCoreStrategyMockRecorder
}

// MockCoreStrategyMockRecorder is the mock recorder for MockCoreStrategy
type MockCoreStrategyMockRecorder struct {
	mock *MockCoreStrategy
}

// NewMockCoreStrategy creates a new mock instance
func NewMockCoreStrategy(ctrl *gomock.Controller) *MockCoreStrategy {
	mock := &MockCoreStrategy{ctrl: ctrl}
	mock.recorder = &MockCoreStrategyMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockCoreStrategy) EXPECT() *MockCoreStrategyMockRecorder {
	return m.recorder
}

// AccessTokenSignature mocks base method
func (m *MockCoreStrategy) AccessTokenSignature(arg0 string) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AccessTokenSignature", arg0)
	ret0, _ := ret[0].(string)
	return ret0
}

// AccessTokenSignature indicates an expected call of AccessTokenSignature
func (mr *MockCoreStrategyMockRecorder) AccessTokenSignature(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AccessTokenSignature", reflect.TypeOf((*MockCoreStrategy)(nil).AccessTokenSignature), arg0)
}

// AuthorizeCodeSignature mocks base method
func (m *MockCoreStrategy) AuthorizeCodeSignature(arg0 string) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AuthorizeCodeSignature", arg0)
	ret0, _ := ret[0].(string)
	return ret0
}

// AuthorizeCodeSignature indicates an expected call of AuthorizeCodeSignature
func (mr *MockCoreStrategyMockRecorder) AuthorizeCodeSignature(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AuthorizeCodeSignature", reflect.TypeOf((*MockCoreStrategy)(nil).AuthorizeCodeSignature), arg0)
}

// GenerateAccessToken mocks base method
func (m *MockCoreStrategy) GenerateAccessToken(arg0 context.Context, arg1 fosite.Requester) (string, string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenerateAccessToken", arg0, arg1)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(string)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GenerateAccessToken indicates an expected call of GenerateAccessToken
func (mr *MockCoreStrategyMockRecorder) GenerateAccessToken(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateAccessToken", reflect.TypeOf((*MockCoreStrategy)(nil).GenerateAccessToken), arg0, arg1)
}

// GenerateAuthorizeCode mocks base method
func (m *MockCoreStrategy) GenerateAuthorizeCode(arg0 context.Context, arg1 fosite.Requester) (string, string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenerateAuthorizeCode", arg0, arg1)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(string)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GenerateAuthorizeCode indicates an expected call of GenerateAuthorizeCode
func (mr *MockCoreStrategyMockRecorder) GenerateAuthorizeCode(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateAuthorizeCode", reflect.TypeOf((*MockCoreStrategy)(nil).GenerateAuthorizeCode), arg0, arg1)
}

// GenerateRefreshToken mocks base method
func (m *MockCoreStrategy) GenerateRefreshToken(arg0 context.Context, arg1 fosite.Requester) (string, string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenerateRefreshToken", arg0, arg1)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(string)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GenerateRefreshToken indicates an expected call of GenerateRefreshToken
func (mr *MockCoreStrategyMockRecorder) GenerateRefreshToken(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateRefreshToken", reflect.TypeOf((*MockCoreStrategy)(nil).GenerateRefreshToken), arg0, arg1)
}

// RefreshTokenSignature mocks base method
func (m *MockCoreStrategy) RefreshTokenSignature(arg0 string) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RefreshTokenSignature", arg0)
	ret0, _ := ret[0].(string)
	return ret0
}

// RefreshTokenSignature indicates an expected call of RefreshTokenSignature
func (mr *MockCoreStrategyMockRecorder) RefreshTokenSignature(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RefreshTokenSignature", reflect.TypeOf((*MockCoreStrategy)(nil).RefreshTokenSignature), arg0)
}

// ValidateAccessToken mocks base method
func (m *MockCoreStrategy) ValidateAccessToken(arg0 context.Context, arg1 fosite.Requester, arg2 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ValidateAccessToken", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// ValidateAccessToken indicates an expected call of ValidateAccessToken
func (mr *MockCoreStrategyMockRecorder) ValidateAccessToken(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ValidateAccessToken", reflect.TypeOf((*MockCoreStrategy)(nil).ValidateAccessToken), arg0, arg1, arg2)
}

// ValidateAuthorizeCode mocks base method
func (m *MockCoreStrategy) ValidateAuthorizeCode(arg0 context.Context, arg1 fosite.Requester, arg2 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ValidateAuthorizeCode", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// ValidateAuthorizeCode indicates an expected call of ValidateAuthorizeCode
func (mr *MockCoreStrategyMockRecorder) ValidateAuthorizeCode(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ValidateAuthorizeCode", reflect.TypeOf((*MockCoreStrategy)(nil).ValidateAuthorizeCode), arg0, arg1, arg2)
}

// ValidateRefreshToken mocks base method
func (m *MockCoreStrategy) ValidateRefreshToken(arg0 context.Context, arg1 fosite.Requester, arg2 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ValidateRefreshToken", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// ValidateRefreshToken indicates an expected call of ValidateRefreshToken
func (mr *MockCoreStrategyMockRecorder) ValidateRefreshToken(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ValidateRefreshToken", reflect.TypeOf((*MockCoreStrategy)(nil).ValidateRefreshToken), arg0, arg1, arg2)
}
