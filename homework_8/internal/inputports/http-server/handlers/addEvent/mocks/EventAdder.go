// Code generated by mockery v2.42.1. DO NOT EDIT.

package mocks

import (
	commands "github.com/andrei1998Front/go_course/homework_8/internal/app/event/commands"
	mock "github.com/stretchr/testify/mock"
)

// EventAdder is an autogenerated mock type for the EventAdder type
type EventAdder struct {
	mock.Mock
}

// Handle provides a mock function with given fields: query
func (_m *EventAdder) Handle(query *commands.AddEventRequest) error {
	ret := _m.Called(query)

	if len(ret) == 0 {
		panic("no return value specified for Handle")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(*commands.AddEventRequest) error); ok {
		r0 = rf(query)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewEventAdder creates a new instance of EventAdder. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewEventAdder(t interface {
	mock.TestingT
	Cleanup(func())
}) *EventAdder {
	mock := &EventAdder{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}