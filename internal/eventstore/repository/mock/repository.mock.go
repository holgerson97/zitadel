// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/zitadel/zitadel/internal/eventstore (interfaces: Querier,Pusher)

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"
	time "time"

	gomock "github.com/golang/mock/gomock"
	repository "github.com/zitadel/zitadel/internal/eventstore/repository"
	eventstore "github.com/zitadel/zitadel/internal/eventstore/v3"
)

// MockQuerier is a mock of Querier interface.
type MockQuerier struct {
	ctrl     *gomock.Controller
	recorder *MockQuerierMockRecorder
}

// MockQuerierMockRecorder is the mock recorder for MockQuerier.
type MockQuerierMockRecorder struct {
	mock *MockQuerier
}

// NewMockQuerier creates a new mock instance.
func NewMockQuerier(ctrl *gomock.Controller) *MockQuerier {
	mock := &MockQuerier{ctrl: ctrl}
	mock.recorder = &MockQuerierMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockQuerier) EXPECT() *MockQuerierMockRecorder {
	return m.recorder
}

// CreateInstance mocks base method.
func (m *MockQuerier) CreateInstance(arg0 context.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateInstance", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateInstance indicates an expected call of CreateInstance.
func (mr *MockQuerierMockRecorder) CreateInstance(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateInstance", reflect.TypeOf((*MockQuerier)(nil).CreateInstance), arg0, arg1)
}

// Filter mocks base method.
func (m *MockQuerier) Filter(arg0 context.Context, arg1 *repository.SearchQuery) ([]*repository.Event, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Filter", arg0, arg1)
	ret0, _ := ret[0].([]*repository.Event)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Filter indicates an expected call of Filter.
func (mr *MockQuerierMockRecorder) Filter(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Filter", reflect.TypeOf((*MockQuerier)(nil).Filter), arg0, arg1)
}

// Health mocks base method.
func (m *MockQuerier) Health(arg0 context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Health", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Health indicates an expected call of Health.
func (mr *MockQuerierMockRecorder) Health(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Health", reflect.TypeOf((*MockQuerier)(nil).Health), arg0)
}

// InstanceIDs mocks base method.
func (m *MockQuerier) InstanceIDs(arg0 context.Context, arg1 *repository.SearchQuery) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InstanceIDs", arg0, arg1)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InstanceIDs indicates an expected call of InstanceIDs.
func (mr *MockQuerierMockRecorder) InstanceIDs(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InstanceIDs", reflect.TypeOf((*MockQuerier)(nil).InstanceIDs), arg0, arg1)
}

// LatestSequence mocks base method.
func (m *MockQuerier) LatestSequence(arg0 context.Context, arg1 *repository.SearchQuery) (time.Time, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LatestSequence", arg0, arg1)
	ret0, _ := ret[0].(time.Time)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LatestSequence indicates an expected call of LatestSequence.
func (mr *MockQuerierMockRecorder) LatestSequence(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LatestSequence", reflect.TypeOf((*MockQuerier)(nil).LatestSequence), arg0, arg1)
}

// MockPusher is a mock of Pusher interface.
type MockPusher struct {
	ctrl     *gomock.Controller
	recorder *MockPusherMockRecorder
}

// MockPusherMockRecorder is the mock recorder for MockPusher.
type MockPusherMockRecorder struct {
	mock *MockPusher
}

// NewMockPusher creates a new mock instance.
func NewMockPusher(ctrl *gomock.Controller) *MockPusher {
	mock := &MockPusher{ctrl: ctrl}
	mock.recorder = &MockPusherMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPusher) EXPECT() *MockPusherMockRecorder {
	return m.recorder
}

// Health mocks base method.
func (m *MockPusher) Health(arg0 context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Health", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Health indicates an expected call of Health.
func (mr *MockPusherMockRecorder) Health(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Health", reflect.TypeOf((*MockPusher)(nil).Health), arg0)
}

// Push mocks base method.
func (m *MockPusher) Push(arg0 context.Context, arg1 ...eventstore.Command) ([]eventstore.Event, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Push", varargs...)
	ret0, _ := ret[0].([]eventstore.Event)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Push indicates an expected call of Push.
func (mr *MockPusherMockRecorder) Push(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Push", reflect.TypeOf((*MockPusher)(nil).Push), varargs...)
}
