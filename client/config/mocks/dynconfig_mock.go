// Code generated by MockGen. DO NOT EDIT.
// Source: dynconfig.go

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	config "d7y.io/dragonfly/v2/client/config"
	manager "d7y.io/dragonfly/v2/pkg/rpc/manager"
	gomock "github.com/golang/mock/gomock"
)

// MockDynconfig is a mock of Dynconfig interface.
type MockDynconfig struct {
	ctrl     *gomock.Controller
	recorder *MockDynconfigMockRecorder
}

// MockDynconfigMockRecorder is the mock recorder for MockDynconfig.
type MockDynconfigMockRecorder struct {
	mock *MockDynconfig
}

// NewMockDynconfig creates a new mock instance.
func NewMockDynconfig(ctrl *gomock.Controller) *MockDynconfig {
	mock := &MockDynconfig{ctrl: ctrl}
	mock.recorder = &MockDynconfigMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDynconfig) EXPECT() *MockDynconfigMockRecorder {
	return m.recorder
}

// Deregister mocks base method.
func (m *MockDynconfig) Deregister(arg0 config.Observer) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Deregister", arg0)
}

// Deregister indicates an expected call of Deregister.
func (mr *MockDynconfigMockRecorder) Deregister(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Deregister", reflect.TypeOf((*MockDynconfig)(nil).Deregister), arg0)
}

// Get mocks base method.
func (m *MockDynconfig) Get() (*config.DynconfigData, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get")
	ret0, _ := ret[0].(*config.DynconfigData)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockDynconfigMockRecorder) Get() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockDynconfig)(nil).Get))
}

// GetObjectStorage mocks base method.
func (m *MockDynconfig) GetObjectStorage() (*manager.ObjectStorage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetObjectStorage")
	ret0, _ := ret[0].(*manager.ObjectStorage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetObjectStorage indicates an expected call of GetObjectStorage.
func (mr *MockDynconfigMockRecorder) GetObjectStorage() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetObjectStorage", reflect.TypeOf((*MockDynconfig)(nil).GetObjectStorage))
}

// GetSchedulers mocks base method.
func (m *MockDynconfig) GetSchedulers() ([]*manager.Scheduler, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSchedulers")
	ret0, _ := ret[0].([]*manager.Scheduler)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSchedulers indicates an expected call of GetSchedulers.
func (mr *MockDynconfigMockRecorder) GetSchedulers() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSchedulers", reflect.TypeOf((*MockDynconfig)(nil).GetSchedulers))
}

// Notify mocks base method.
func (m *MockDynconfig) Notify() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Notify")
	ret0, _ := ret[0].(error)
	return ret0
}

// Notify indicates an expected call of Notify.
func (mr *MockDynconfigMockRecorder) Notify() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Notify", reflect.TypeOf((*MockDynconfig)(nil).Notify))
}

// Register mocks base method.
func (m *MockDynconfig) Register(arg0 config.Observer) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Register", arg0)
}

// Register indicates an expected call of Register.
func (mr *MockDynconfigMockRecorder) Register(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Register", reflect.TypeOf((*MockDynconfig)(nil).Register), arg0)
}

// Serve mocks base method.
func (m *MockDynconfig) Serve() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Serve")
	ret0, _ := ret[0].(error)
	return ret0
}

// Serve indicates an expected call of Serve.
func (mr *MockDynconfigMockRecorder) Serve() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Serve", reflect.TypeOf((*MockDynconfig)(nil).Serve))
}

// Stop mocks base method.
func (m *MockDynconfig) Stop() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Stop")
	ret0, _ := ret[0].(error)
	return ret0
}

// Stop indicates an expected call of Stop.
func (mr *MockDynconfigMockRecorder) Stop() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stop", reflect.TypeOf((*MockDynconfig)(nil).Stop))
}

// MockObserver is a mock of Observer interface.
type MockObserver struct {
	ctrl     *gomock.Controller
	recorder *MockObserverMockRecorder
}

// MockObserverMockRecorder is the mock recorder for MockObserver.
type MockObserverMockRecorder struct {
	mock *MockObserver
}

// NewMockObserver creates a new mock instance.
func NewMockObserver(ctrl *gomock.Controller) *MockObserver {
	mock := &MockObserver{ctrl: ctrl}
	mock.recorder = &MockObserverMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockObserver) EXPECT() *MockObserverMockRecorder {
	return m.recorder
}

// OnNotify mocks base method.
func (m *MockObserver) OnNotify(arg0 *config.DynconfigData) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "OnNotify", arg0)
}

// OnNotify indicates an expected call of OnNotify.
func (mr *MockObserverMockRecorder) OnNotify(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OnNotify", reflect.TypeOf((*MockObserver)(nil).OnNotify), arg0)
}