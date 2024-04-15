// Code generated by mockery v2.42.1. DO NOT EDIT.

package mocks

import (
	queries "github.com/andrei1998Front/go_course/homework_8/internal/app/event/queries"
	mock "github.com/stretchr/testify/mock"
)

// EventGetter is an autogenerated mock type for the EventGetter type
type EventGetter struct {
	mock.Mock
}

// Handle provides a mock function with given fields: query
func (_m *EventGetter) Handle(query queries.GetEventRequest) (*queries.GetEventResponce, error) {
	ret := _m.Called(query)

	if len(ret) == 0 {
		panic("no return value specified for Handle")
	}

	var r0 *queries.GetEventResponce
	var r1 error
	if rf, ok := ret.Get(0).(func(queries.GetEventRequest) (*queries.GetEventResponce, error)); ok {
		return rf(query)
	}
	if rf, ok := ret.Get(0).(func(queries.GetEventRequest) *queries.GetEventResponce); ok {
		r0 = rf(query)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*queries.GetEventResponce)
		}
	}

	if rf, ok := ret.Get(1).(func(queries.GetEventRequest) error); ok {
		r1 = rf(query)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewEventGetter creates a new instance of EventGetter. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewEventGetter(t interface {
	mock.TestingT
	Cleanup(func())
}) *EventGetter {
	mock := &EventGetter{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}