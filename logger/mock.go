package logger

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockLogger is a mock of Logger interface
type MockLogger struct {
	ctrl     *gomock.Controller
	recorder *MockLoggerMockRecorder
}

// MockLoggerMockRecorder is the mock recorder for MockLogger
type MockLoggerMockRecorder struct {
	mock *MockLogger
}

// NewMockLogger creates a new mock instance
func NewMockLogger(ctrl *gomock.Controller) *MockLogger {
	mock := &MockLogger{ctrl: ctrl}
	mock.recorder = &MockLoggerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockLogger) EXPECT() *MockLoggerMockRecorder {
	return m.recorder
}

// Info mocks base method
func (m *MockLogger) Info(arg0 Context, arg1 string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Info", arg0, arg1)
}

// Info indicates an expected call of Info
func (mr *MockLoggerMockRecorder) Info(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Info", reflect.TypeOf((*MockLogger)(nil).Info), arg0, arg1)
}

// Error mocks base method
func (m *MockLogger) Error(arg0 Context, arg1 string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Error", arg0, arg1)
}

// Error indicates an expected call of Error
func (mr *MockLoggerMockRecorder) Error(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Error", reflect.TypeOf((*MockLogger)(nil).Error), arg0, arg1)
}

// Fatal mocks base method
func (m *MockLogger) Fatal(arg0 Context, arg1 string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Fatal", arg0, arg1)
}

// Fatal indicates an expected call of Fatal
func (mr *MockLoggerMockRecorder) Fatal(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Fatal", reflect.TypeOf((*MockLogger)(nil).Fatal), arg0, arg1)
}

// Warning mocks base method
func (m *MockLogger) Warning(arg0 Context, arg1 string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Warning", arg0, arg1)
}

// Warning indicates an expected call of Warning
func (mr *MockLoggerMockRecorder) Warning(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Warning", reflect.TypeOf((*MockLogger)(nil).Warning), arg0, arg1)
}
