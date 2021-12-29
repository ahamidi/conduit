// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/conduitio/conduit/pkg/plugins/kafka (interfaces: Consumer)

// Package mock is a generated GoMock package.
package mock

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	kafka "github.com/segmentio/kafka-go"
)

// Consumer is a mock of Consumer interface.
type Consumer struct {
	ctrl     *gomock.Controller
	recorder *ConsumerMockRecorder
}

// ConsumerMockRecorder is the mock recorder for Consumer.
type ConsumerMockRecorder struct {
	mock *Consumer
}

// NewConsumer creates a new mock instance.
func NewConsumer(ctrl *gomock.Controller) *Consumer {
	mock := &Consumer{ctrl: ctrl}
	mock.recorder = &ConsumerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *Consumer) EXPECT() *ConsumerMockRecorder {
	return m.recorder
}

// Close mocks base method.
func (m *Consumer) Close() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Close")
}

// Close indicates an expected call of Close.
func (mr *ConsumerMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*Consumer)(nil).Close))
}

// Get mocks base method.
func (m *Consumer) Get() (*kafka.Message, map[int]int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get")
	ret0, _ := ret[0].(*kafka.Message)
	ret1, _ := ret[1].(map[int]int64)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Get indicates an expected call of Get.
func (mr *ConsumerMockRecorder) Get() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*Consumer)(nil).Get))
}

// StartFrom mocks base method.
func (m *Consumer) StartFrom(arg0 string, arg1 map[int]int64, arg2 bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StartFrom", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// StartFrom indicates an expected call of StartFrom.
func (mr *ConsumerMockRecorder) StartFrom(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StartFrom", reflect.TypeOf((*Consumer)(nil).StartFrom), arg0, arg1, arg2)
}
