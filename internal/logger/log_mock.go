// Code generated by MockGen. DO NOT EDIT.
// Source: internal/logger/log.go

// Package logger is a generated GoMock package.
package logger

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockFdWriter is a mock of FdWriter interface
type MockFdWriter struct {
	ctrl     *gomock.Controller
	recorder *MockFdWriterMockRecorder
}

// MockFdWriterMockRecorder is the mock recorder for MockFdWriter
type MockFdWriterMockRecorder struct {
	mock *MockFdWriter
}

// NewMockFdWriter creates a new mock instance
func NewMockFdWriter(ctrl *gomock.Controller) *MockFdWriter {
	mock := &MockFdWriter{ctrl: ctrl}
	mock.recorder = &MockFdWriterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockFdWriter) EXPECT() *MockFdWriterMockRecorder {
	return m.recorder
}

// Write mocks base method
func (m *MockFdWriter) Write(p []byte) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Write", p)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Write indicates an expected call of Write
func (mr *MockFdWriterMockRecorder) Write(p interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Write", reflect.TypeOf((*MockFdWriter)(nil).Write), p)
}

// Fd mocks base method
func (m *MockFdWriter) Fd() uintptr {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Fd")
	ret0, _ := ret[0].(uintptr)
	return ret0
}

// Fd indicates an expected call of Fd
func (mr *MockFdWriterMockRecorder) Fd() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Fd", reflect.TypeOf((*MockFdWriter)(nil).Fd))
}

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

// WithColor mocks base method
func (m *MockLogger) WithColor() Logger {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WithColor")
	ret0, _ := ret[0].(Logger)
	return ret0
}

// WithColor indicates an expected call of WithColor
func (mr *MockLoggerMockRecorder) WithColor() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WithColor", reflect.TypeOf((*MockLogger)(nil).WithColor))
}

// WithoutColor mocks base method
func (m *MockLogger) WithoutColor() Logger {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WithoutColor")
	ret0, _ := ret[0].(Logger)
	return ret0
}

// WithoutColor indicates an expected call of WithoutColor
func (mr *MockLoggerMockRecorder) WithoutColor() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WithoutColor", reflect.TypeOf((*MockLogger)(nil).WithoutColor))
}

// WithDebug mocks base method
func (m *MockLogger) WithDebug() Logger {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WithDebug")
	ret0, _ := ret[0].(Logger)
	return ret0
}

// WithDebug indicates an expected call of WithDebug
func (mr *MockLoggerMockRecorder) WithDebug() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WithDebug", reflect.TypeOf((*MockLogger)(nil).WithDebug))
}

// WithoutDebug mocks base method
func (m *MockLogger) WithoutDebug() Logger {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WithoutDebug")
	ret0, _ := ret[0].(Logger)
	return ret0
}

// WithoutDebug indicates an expected call of WithoutDebug
func (mr *MockLoggerMockRecorder) WithoutDebug() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WithoutDebug", reflect.TypeOf((*MockLogger)(nil).WithoutDebug))
}

// IsDebug mocks base method
func (m *MockLogger) IsDebug() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsDebug")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsDebug indicates an expected call of IsDebug
func (mr *MockLoggerMockRecorder) IsDebug() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsDebug", reflect.TypeOf((*MockLogger)(nil).IsDebug))
}

// WithTimestamp mocks base method
func (m *MockLogger) WithTimestamp() Logger {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WithTimestamp")
	ret0, _ := ret[0].(Logger)
	return ret0
}

// WithTimestamp indicates an expected call of WithTimestamp
func (mr *MockLoggerMockRecorder) WithTimestamp() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WithTimestamp", reflect.TypeOf((*MockLogger)(nil).WithTimestamp))
}

// WithoutTimestamp mocks base method
func (m *MockLogger) WithoutTimestamp() Logger {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WithoutTimestamp")
	ret0, _ := ret[0].(Logger)
	return ret0
}

// WithoutTimestamp indicates an expected call of WithoutTimestamp
func (mr *MockLoggerMockRecorder) WithoutTimestamp() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WithoutTimestamp", reflect.TypeOf((*MockLogger)(nil).WithoutTimestamp))
}

// Quiet mocks base method
func (m *MockLogger) Quiet() Logger {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Quiet")
	ret0, _ := ret[0].(Logger)
	return ret0
}

// Quiet indicates an expected call of Quiet
func (mr *MockLoggerMockRecorder) Quiet() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Quiet", reflect.TypeOf((*MockLogger)(nil).Quiet))
}

// NoQuiet mocks base method
func (m *MockLogger) NoQuiet() Logger {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NoQuiet")
	ret0, _ := ret[0].(Logger)
	return ret0
}

// NoQuiet indicates an expected call of NoQuiet
func (mr *MockLoggerMockRecorder) NoQuiet() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NoQuiet", reflect.TypeOf((*MockLogger)(nil).NoQuiet))
}

// IsQuiet mocks base method
func (m *MockLogger) IsQuiet() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsQuiet")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsQuiet indicates an expected call of IsQuiet
func (mr *MockLoggerMockRecorder) IsQuiet() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsQuiet", reflect.TypeOf((*MockLogger)(nil).IsQuiet))
}

// Output mocks base method
func (m *MockLogger) Output(depth int, prefix Prefix, data string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Output", depth, prefix, data)
	ret0, _ := ret[0].(error)
	return ret0
}

// Output indicates an expected call of Output
func (mr *MockLoggerMockRecorder) Output(depth, prefix, data interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Output", reflect.TypeOf((*MockLogger)(nil).Output), depth, prefix, data)
}

// Fatal mocks base method
func (m *MockLogger) Fatal(v ...interface{}) {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range v {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Fatal", varargs...)
}

// Fatal indicates an expected call of Fatal
func (mr *MockLoggerMockRecorder) Fatal(v ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Fatal", reflect.TypeOf((*MockLogger)(nil).Fatal), v...)
}

// Fatalf mocks base method
func (m *MockLogger) Fatalf(format string, v ...interface{}) {
	m.ctrl.T.Helper()
	varargs := []interface{}{format}
	for _, a := range v {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Fatalf", varargs...)
}

// Fatalf indicates an expected call of Fatalf
func (mr *MockLoggerMockRecorder) Fatalf(format interface{}, v ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{format}, v...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Fatalf", reflect.TypeOf((*MockLogger)(nil).Fatalf), varargs...)
}

// Error mocks base method
func (m *MockLogger) Error(v ...interface{}) {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range v {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Error", varargs...)
}

// Error indicates an expected call of Error
func (mr *MockLoggerMockRecorder) Error(v ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Error", reflect.TypeOf((*MockLogger)(nil).Error), v...)
}

// Errorf mocks base method
func (m *MockLogger) Errorf(format string, v ...interface{}) {
	m.ctrl.T.Helper()
	varargs := []interface{}{format}
	for _, a := range v {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Errorf", varargs...)
}

// Errorf indicates an expected call of Errorf
func (mr *MockLoggerMockRecorder) Errorf(format interface{}, v ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{format}, v...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Errorf", reflect.TypeOf((*MockLogger)(nil).Errorf), varargs...)
}

// Warn mocks base method
func (m *MockLogger) Warn(v ...interface{}) {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range v {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Warn", varargs...)
}

// Warn indicates an expected call of Warn
func (mr *MockLoggerMockRecorder) Warn(v ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Warn", reflect.TypeOf((*MockLogger)(nil).Warn), v...)
}

// Warnf mocks base method
func (m *MockLogger) Warnf(format string, v ...interface{}) {
	m.ctrl.T.Helper()
	varargs := []interface{}{format}
	for _, a := range v {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Warnf", varargs...)
}

// Warnf indicates an expected call of Warnf
func (mr *MockLoggerMockRecorder) Warnf(format interface{}, v ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{format}, v...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Warnf", reflect.TypeOf((*MockLogger)(nil).Warnf), varargs...)
}

// Info mocks base method
func (m *MockLogger) Info(v ...interface{}) {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range v {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Info", varargs...)
}

// Info indicates an expected call of Info
func (mr *MockLoggerMockRecorder) Info(v ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Info", reflect.TypeOf((*MockLogger)(nil).Info), v...)
}

// Infof mocks base method
func (m *MockLogger) Infof(format string, v ...interface{}) {
	m.ctrl.T.Helper()
	varargs := []interface{}{format}
	for _, a := range v {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Infof", varargs...)
}

// Infof indicates an expected call of Infof
func (mr *MockLoggerMockRecorder) Infof(format interface{}, v ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{format}, v...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Infof", reflect.TypeOf((*MockLogger)(nil).Infof), varargs...)
}

// Debug mocks base method
func (m *MockLogger) Debug(v ...interface{}) {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range v {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Debug", varargs...)
}

// Debug indicates an expected call of Debug
func (mr *MockLoggerMockRecorder) Debug(v ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Debug", reflect.TypeOf((*MockLogger)(nil).Debug), v...)
}

// Debugf mocks base method
func (m *MockLogger) Debugf(format string, v ...interface{}) {
	m.ctrl.T.Helper()
	varargs := []interface{}{format}
	for _, a := range v {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Debugf", varargs...)
}

// Debugf indicates an expected call of Debugf
func (mr *MockLoggerMockRecorder) Debugf(format interface{}, v ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{format}, v...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Debugf", reflect.TypeOf((*MockLogger)(nil).Debugf), varargs...)
}

// Trace mocks base method
func (m *MockLogger) Trace(v ...interface{}) {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range v {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Trace", varargs...)
}

// Trace indicates an expected call of Trace
func (mr *MockLoggerMockRecorder) Trace(v ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Trace", reflect.TypeOf((*MockLogger)(nil).Trace), v...)
}

// Tracef mocks base method
func (m *MockLogger) Tracef(format string, v ...interface{}) {
	m.ctrl.T.Helper()
	varargs := []interface{}{format}
	for _, a := range v {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Tracef", varargs...)
}

// Tracef indicates an expected call of Tracef
func (mr *MockLoggerMockRecorder) Tracef(format interface{}, v ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{format}, v...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Tracef", reflect.TypeOf((*MockLogger)(nil).Tracef), varargs...)
}
