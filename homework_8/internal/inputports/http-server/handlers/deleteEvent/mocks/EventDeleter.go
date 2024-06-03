// Code generated by mockery v2.42.1. DO NOT EDIT.

package mocks

import (
	commands "github.com/andrei1998Front/go_course/homework_8/internal/app/event/commands"

	mock "github.com/stretchr/testify/mock"
)

// EventDeleter is an autogenerated mock type for the EventDeleter type
type EventDeleter struct {
	mock.Mock
}

// Handle provides a mock function with given fields: query
func (_m *EventDeleter) Handle(query commands.DeleteEventRequest) error {
	ret := _m.Called(query)

	if len(ret) == 0 {
		panic("no return value specified for Handle")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(commands.DeleteEventRequest) error); ok {
		r0 = rf(query)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewEventDeleter creates a new instance of EventDeleter. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewEventDeleter(t interface {
	mock.TestingT
	Cleanup(func())
}) *EventDeleter {
	mock := &EventDeleter{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}