// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/ssargent/bbq/bbq-apiserver/security (interfaces: AuthenticationService)

// Package mock_security is a generated GoMock package.
package mock_security

import (
	gomock "github.com/golang/mock/gomock"
	security "github.com/ssargent/bbq/bbq-apiserver/security"
	http "net/http"
	reflect "reflect"
)

// MockAuthenticationService is a mock of AuthenticationService interface
type MockAuthenticationService struct {
	ctrl     *gomock.Controller
	recorder *MockAuthenticationServiceMockRecorder
}

// MockAuthenticationServiceMockRecorder is the mock recorder for MockAuthenticationService
type MockAuthenticationServiceMockRecorder struct {
	mock *MockAuthenticationService
}

// NewMockAuthenticationService creates a new mock instance
func NewMockAuthenticationService(ctrl *gomock.Controller) *MockAuthenticationService {
	mock := &MockAuthenticationService{ctrl: ctrl}
	mock.recorder = &MockAuthenticationServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockAuthenticationService) EXPECT() *MockAuthenticationServiceMockRecorder {
	return m.recorder
}

// GetLoginSession mocks base method
func (m *MockAuthenticationService) GetLoginSession(arg0 *http.Request) (security.LoginSession, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLoginSession", arg0)
	ret0, _ := ret[0].(security.LoginSession)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetLoginSession indicates an expected call of GetLoginSession
func (mr *MockAuthenticationServiceMockRecorder) GetLoginSession(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLoginSession", reflect.TypeOf((*MockAuthenticationService)(nil).GetLoginSession), arg0)
}
