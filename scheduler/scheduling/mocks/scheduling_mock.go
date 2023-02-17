// Code generated by MockGen. DO NOT EDIT.
// Source: scheduling.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	set "d7y.io/dragonfly/v2/pkg/container/set"
	resource "d7y.io/dragonfly/v2/scheduler/resource"
	gomock "github.com/golang/mock/gomock"
)

// MockScheduling is a mock of Scheduling interface.
type MockScheduling struct {
	ctrl     *gomock.Controller
	recorder *MockSchedulingMockRecorder
}

// MockSchedulingMockRecorder is the mock recorder for MockScheduling.
type MockSchedulingMockRecorder struct {
	mock *MockScheduling
}

// NewMockScheduling creates a new mock instance.
func NewMockScheduling(ctrl *gomock.Controller) *MockScheduling {
	mock := &MockScheduling{ctrl: ctrl}
	mock.recorder = &MockSchedulingMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockScheduling) EXPECT() *MockSchedulingMockRecorder {
	return m.recorder
}

// FindParent mocks base method.
func (m *MockScheduling) FindParent(arg0 context.Context, arg1 *resource.Peer, arg2 set.SafeSet[string]) (*resource.Peer, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindParent", arg0, arg1, arg2)
	ret0, _ := ret[0].(*resource.Peer)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// FindParent indicates an expected call of FindParent.
func (mr *MockSchedulingMockRecorder) FindParent(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindParent", reflect.TypeOf((*MockScheduling)(nil).FindParent), arg0, arg1, arg2)
}

// NotifyAndFindParent mocks base method.
func (m *MockScheduling) NotifyAndFindParent(arg0 context.Context, arg1 *resource.Peer, arg2 set.SafeSet[string]) ([]*resource.Peer, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NotifyAndFindParent", arg0, arg1, arg2)
	ret0, _ := ret[0].([]*resource.Peer)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// NotifyAndFindParent indicates an expected call of NotifyAndFindParent.
func (mr *MockSchedulingMockRecorder) NotifyAndFindParent(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NotifyAndFindParent", reflect.TypeOf((*MockScheduling)(nil).NotifyAndFindParent), arg0, arg1, arg2)
}

// ScheduleParent mocks base method.
func (m *MockScheduling) ScheduleParent(arg0 context.Context, arg1 *resource.Peer, arg2 set.SafeSet[string]) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "ScheduleParent", arg0, arg1, arg2)
}

// ScheduleParent indicates an expected call of ScheduleParent.
func (mr *MockSchedulingMockRecorder) ScheduleParent(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ScheduleParent", reflect.TypeOf((*MockScheduling)(nil).ScheduleParent), arg0, arg1, arg2)
}
