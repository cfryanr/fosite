// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/ory/fosite/handler/oauth2 (interfaces: ResourceOwnerPasswordCredentialsGrantStorage)

// Package internal is a generated GoMock package.
package internal

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"

	fosite "github.com/ory/fosite"
)

// MockResourceOwnerPasswordCredentialsGrantStorage is a mock of ResourceOwnerPasswordCredentialsGrantStorage interface
type MockResourceOwnerPasswordCredentialsGrantStorage struct {
	ctrl     *gomock.Controller
	recorder *MockResourceOwnerPasswordCredentialsGrantStorageMockRecorder
}

// MockResourceOwnerPasswordCredentialsGrantStorageMockRecorder is the mock recorder for MockResourceOwnerPasswordCredentialsGrantStorage
type MockResourceOwnerPasswordCredentialsGrantStorageMockRecorder struct {
	mock *MockResourceOwnerPasswordCredentialsGrantStorage
}

// NewMockResourceOwnerPasswordCredentialsGrantStorage creates a new mock instance
func NewMockResourceOwnerPasswordCredentialsGrantStorage(ctrl *gomock.Controller) *MockResourceOwnerPasswordCredentialsGrantStorage {
	mock := &MockResourceOwnerPasswordCredentialsGrantStorage{ctrl: ctrl}
	mock.recorder = &MockResourceOwnerPasswordCredentialsGrantStorageMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockResourceOwnerPasswordCredentialsGrantStorage) EXPECT() *MockResourceOwnerPasswordCredentialsGrantStorageMockRecorder {
	return m.recorder
}

// Authenticate mocks base method
func (m *MockResourceOwnerPasswordCredentialsGrantStorage) Authenticate(arg0 context.Context, arg1, arg2 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Authenticate", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// Authenticate indicates an expected call of Authenticate
func (mr *MockResourceOwnerPasswordCredentialsGrantStorageMockRecorder) Authenticate(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Authenticate", reflect.TypeOf((*MockResourceOwnerPasswordCredentialsGrantStorage)(nil).Authenticate), arg0, arg1, arg2)
}

// CreateAccessTokenSession mocks base method
func (m *MockResourceOwnerPasswordCredentialsGrantStorage) CreateAccessTokenSession(arg0 context.Context, arg1 string, arg2 fosite.Requester) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateAccessTokenSession", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateAccessTokenSession indicates an expected call of CreateAccessTokenSession
func (mr *MockResourceOwnerPasswordCredentialsGrantStorageMockRecorder) CreateAccessTokenSession(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateAccessTokenSession", reflect.TypeOf((*MockResourceOwnerPasswordCredentialsGrantStorage)(nil).CreateAccessTokenSession), arg0, arg1, arg2)
}

// CreateRefreshTokenSession mocks base method
func (m *MockResourceOwnerPasswordCredentialsGrantStorage) CreateRefreshTokenSession(arg0 context.Context, arg1 string, arg2 fosite.Requester) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateRefreshTokenSession", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateRefreshTokenSession indicates an expected call of CreateRefreshTokenSession
func (mr *MockResourceOwnerPasswordCredentialsGrantStorageMockRecorder) CreateRefreshTokenSession(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateRefreshTokenSession", reflect.TypeOf((*MockResourceOwnerPasswordCredentialsGrantStorage)(nil).CreateRefreshTokenSession), arg0, arg1, arg2)
}

// DeleteAccessTokenSession mocks base method
func (m *MockResourceOwnerPasswordCredentialsGrantStorage) DeleteAccessTokenSession(arg0 context.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAccessTokenSession", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteAccessTokenSession indicates an expected call of DeleteAccessTokenSession
func (mr *MockResourceOwnerPasswordCredentialsGrantStorageMockRecorder) DeleteAccessTokenSession(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAccessTokenSession", reflect.TypeOf((*MockResourceOwnerPasswordCredentialsGrantStorage)(nil).DeleteAccessTokenSession), arg0, arg1)
}

// DeleteRefreshTokenSession mocks base method
func (m *MockResourceOwnerPasswordCredentialsGrantStorage) DeleteRefreshTokenSession(arg0 context.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteRefreshTokenSession", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteRefreshTokenSession indicates an expected call of DeleteRefreshTokenSession
func (mr *MockResourceOwnerPasswordCredentialsGrantStorageMockRecorder) DeleteRefreshTokenSession(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteRefreshTokenSession", reflect.TypeOf((*MockResourceOwnerPasswordCredentialsGrantStorage)(nil).DeleteRefreshTokenSession), arg0, arg1)
}

// GetAccessTokenSession mocks base method
func (m *MockResourceOwnerPasswordCredentialsGrantStorage) GetAccessTokenSession(arg0 context.Context, arg1 string, arg2 fosite.Session) (fosite.Requester, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAccessTokenSession", arg0, arg1, arg2)
	ret0, _ := ret[0].(fosite.Requester)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAccessTokenSession indicates an expected call of GetAccessTokenSession
func (mr *MockResourceOwnerPasswordCredentialsGrantStorageMockRecorder) GetAccessTokenSession(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAccessTokenSession", reflect.TypeOf((*MockResourceOwnerPasswordCredentialsGrantStorage)(nil).GetAccessTokenSession), arg0, arg1, arg2)
}

// GetRefreshTokenSession mocks base method
func (m *MockResourceOwnerPasswordCredentialsGrantStorage) GetRefreshTokenSession(arg0 context.Context, arg1 string, arg2 fosite.Session) (fosite.Requester, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRefreshTokenSession", arg0, arg1, arg2)
	ret0, _ := ret[0].(fosite.Requester)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRefreshTokenSession indicates an expected call of GetRefreshTokenSession
func (mr *MockResourceOwnerPasswordCredentialsGrantStorageMockRecorder) GetRefreshTokenSession(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRefreshTokenSession", reflect.TypeOf((*MockResourceOwnerPasswordCredentialsGrantStorage)(nil).GetRefreshTokenSession), arg0, arg1, arg2)
}
