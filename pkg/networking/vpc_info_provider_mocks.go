// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/sonal-chauhan/aws-load-balancer-controller/pkg/networking (interfaces: VPCInfoProvider)

// Package networking is a generated GoMock package.
package networking

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockVPCInfoProvider is a mock of VPCInfoProvider interface.
type MockVPCInfoProvider struct {
	ctrl     *gomock.Controller
	recorder *MockVPCInfoProviderMockRecorder
}

// MockVPCInfoProviderMockRecorder is the mock recorder for MockVPCInfoProvider.
type MockVPCInfoProviderMockRecorder struct {
	mock *MockVPCInfoProvider
}

// NewMockVPCInfoProvider creates a new mock instance.
func NewMockVPCInfoProvider(ctrl *gomock.Controller) *MockVPCInfoProvider {
	mock := &MockVPCInfoProvider{ctrl: ctrl}
	mock.recorder = &MockVPCInfoProviderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockVPCInfoProvider) EXPECT() *MockVPCInfoProviderMockRecorder {
	return m.recorder
}

// FetchVPCInfo mocks base method.
func (m *MockVPCInfoProvider) FetchVPCInfo(arg0 context.Context, arg1 string, arg2 ...FetchVPCInfoOption) (VPCInfo, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "FetchVPCInfo", varargs...)
	ret0, _ := ret[0].(VPCInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchVPCInfo indicates an expected call of FetchVPCInfo.
func (mr *MockVPCInfoProviderMockRecorder) FetchVPCInfo(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchVPCInfo", reflect.TypeOf((*MockVPCInfoProvider)(nil).FetchVPCInfo), varargs...)
}
