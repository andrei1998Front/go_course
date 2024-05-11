// Code generated by mockery v2.42.1. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// FlagGetter is an autogenerated mock type for the FlagGetter type
type FlagGetter struct {
	mock.Mock
}

// Args provides a mock function with given fields:
func (_m *FlagGetter) Args() []string {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Args")
	}

	var r0 []string
	if rf, ok := ret.Get(0).(func() []string); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	return r0
}

// String provides a mock function with given fields: name, value, usage
func (_m *FlagGetter) String(name string, value string, usage string) *string {
	ret := _m.Called(name, value, usage)

	if len(ret) == 0 {
		panic("no return value specified for String")
	}

	var r0 *string
	if rf, ok := ret.Get(0).(func(string, string, string) *string); ok {
		r0 = rf(name, value, usage)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*string)
		}
	}

	return r0
}

// NewFlagGetter creates a new instance of FlagGetter. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewFlagGetter(t interface {
	mock.TestingT
	Cleanup(func())
}) *FlagGetter {
	mock := &FlagGetter{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
