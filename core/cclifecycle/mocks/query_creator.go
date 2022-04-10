// Code generated by mockery v2.10.4. DO NOT EDIT.

package mocks

import (
	cclifecycle "github.com/hyperledger/fabric/core/cclifecycle"
	mock "github.com/stretchr/testify/mock"
)

// QueryCreator is an autogenerated mock type for the QueryCreator type
type QueryCreator struct {
	mock.Mock
}

// NewQuery provides a mock function with given fields:
func (_m *QueryCreator) NewQuery() (cclifecycle.Query, error) {
	ret := _m.Called()

	var r0 cclifecycle.Query
	if rf, ok := ret.Get(0).(func() cclifecycle.Query); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(cclifecycle.Query)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
