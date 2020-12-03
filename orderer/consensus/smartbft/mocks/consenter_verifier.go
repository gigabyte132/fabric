// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"
import protoutil "github.com/hyperledger/fabric/protoutil"

// ConsenterVerifier is an autogenerated mock type for the ConsenterVerifier type
type ConsenterVerifier struct {
	mock.Mock
}

// Evaluate provides a mock function with given fields: signatureSet
func (_m *ConsenterVerifier) Evaluate(signatureSet []*protoutil.SignedData) error {
	ret := _m.Called(signatureSet)

	var r0 error
	if rf, ok := ret.Get(0).(func([]*protoutil.SignedData) error); ok {
		r0 = rf(signatureSet)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
