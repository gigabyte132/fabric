// Code generated by mockery v2.10.4. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// Sequencer is an autogenerated mock type for the Sequencer type
type Sequencer struct {
	mock.Mock
}

// Sequence provides a mock function with given fields:
func (_m *Sequencer) Sequence() uint64 {
	ret := _m.Called()

	var r0 uint64
	if rf, ok := ret.Get(0).(func() uint64); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(uint64)
	}

	return r0
}
