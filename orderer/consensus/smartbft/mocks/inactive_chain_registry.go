// Code generated by mockery v2.10.4. DO NOT EDIT.

package mocks

import (
	common "github.com/hyperledger/fabric-protos-go/common"
	mock "github.com/stretchr/testify/mock"
)

// InactiveChainRegistry is an autogenerated mock type for the InactiveChainRegistry type
type InactiveChainRegistry struct {
	mock.Mock
}

// Stop provides a mock function with given fields:
func (_m *InactiveChainRegistry) Stop() {
	_m.Called()
}

// TrackChain provides a mock function with given fields: chainName, genesisBlock, createChain
func (_m *InactiveChainRegistry) TrackChain(chainName string, genesisBlock *common.Block, createChain func()) {
	_m.Called(chainName, genesisBlock, createChain)
}
