// Code generated by mockery v2.10.4. DO NOT EDIT.

package mocks

import (
	common "github.com/hyperledger/fabric-protos-go/common"
	mock "github.com/stretchr/testify/mock"
)

// Filters is an autogenerated mock type for the Filters type
type Filters struct {
	mock.Mock
}

// ApplyFilters provides a mock function with given fields: channel, env
func (_m *Filters) ApplyFilters(channel string, env *common.Envelope) error {
	ret := _m.Called(channel, env)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, *common.Envelope) error); ok {
		r0 = rf(channel, env)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
