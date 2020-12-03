// Code generated by counterfeiter. DO NOT EDIT.
package mock

import (
	"sync"

	"github.com/hyperledger/fabric/core/ledger"
	"github.com/hyperledger/fabric/msp"
	"github.com/hyperledger/fabric/protoutil"
)

type CollectionPolicyChecker struct {
	CheckCollectionPolicyStub        func(uint64, string, string, ledger.ConfigHistoryRetriever, msp.IdentityDeserializer, *protoutil.SignedData) (bool, error)
	checkCollectionPolicyMutex       sync.RWMutex
	checkCollectionPolicyArgsForCall []struct {
		arg1 uint64
		arg2 string
		arg3 string
		arg4 ledger.ConfigHistoryRetriever
		arg5 msp.IdentityDeserializer
		arg6 *protoutil.SignedData
	}
	checkCollectionPolicyReturns struct {
		result1 bool
		result2 error
	}
	checkCollectionPolicyReturnsOnCall map[int]struct {
		result1 bool
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *CollectionPolicyChecker) CheckCollectionPolicy(arg1 uint64, arg2 string, arg3 string, arg4 ledger.ConfigHistoryRetriever, arg5 msp.IdentityDeserializer, arg6 *protoutil.SignedData) (bool, error) {
	fake.checkCollectionPolicyMutex.Lock()
	ret, specificReturn := fake.checkCollectionPolicyReturnsOnCall[len(fake.checkCollectionPolicyArgsForCall)]
	fake.checkCollectionPolicyArgsForCall = append(fake.checkCollectionPolicyArgsForCall, struct {
		arg1 uint64
		arg2 string
		arg3 string
		arg4 ledger.ConfigHistoryRetriever
		arg5 msp.IdentityDeserializer
		arg6 *protoutil.SignedData
	}{arg1, arg2, arg3, arg4, arg5, arg6})
	stub := fake.CheckCollectionPolicyStub
	fakeReturns := fake.checkCollectionPolicyReturns
	fake.recordInvocation("CheckCollectionPolicy", []interface{}{arg1, arg2, arg3, arg4, arg5, arg6})
	fake.checkCollectionPolicyMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3, arg4, arg5, arg6)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *CollectionPolicyChecker) CheckCollectionPolicyCallCount() int {
	fake.checkCollectionPolicyMutex.RLock()
	defer fake.checkCollectionPolicyMutex.RUnlock()
	return len(fake.checkCollectionPolicyArgsForCall)
}

func (fake *CollectionPolicyChecker) CheckCollectionPolicyCalls(stub func(uint64, string, string, ledger.ConfigHistoryRetriever, msp.IdentityDeserializer, *protoutil.SignedData) (bool, error)) {
	fake.checkCollectionPolicyMutex.Lock()
	defer fake.checkCollectionPolicyMutex.Unlock()
	fake.CheckCollectionPolicyStub = stub
}

func (fake *CollectionPolicyChecker) CheckCollectionPolicyArgsForCall(i int) (uint64, string, string, ledger.ConfigHistoryRetriever, msp.IdentityDeserializer, *protoutil.SignedData) {
	fake.checkCollectionPolicyMutex.RLock()
	defer fake.checkCollectionPolicyMutex.RUnlock()
	argsForCall := fake.checkCollectionPolicyArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4, argsForCall.arg5, argsForCall.arg6
}

func (fake *CollectionPolicyChecker) CheckCollectionPolicyReturns(result1 bool, result2 error) {
	fake.checkCollectionPolicyMutex.Lock()
	defer fake.checkCollectionPolicyMutex.Unlock()
	fake.CheckCollectionPolicyStub = nil
	fake.checkCollectionPolicyReturns = struct {
		result1 bool
		result2 error
	}{result1, result2}
}

func (fake *CollectionPolicyChecker) CheckCollectionPolicyReturnsOnCall(i int, result1 bool, result2 error) {
	fake.checkCollectionPolicyMutex.Lock()
	defer fake.checkCollectionPolicyMutex.Unlock()
	fake.CheckCollectionPolicyStub = nil
	if fake.checkCollectionPolicyReturnsOnCall == nil {
		fake.checkCollectionPolicyReturnsOnCall = make(map[int]struct {
			result1 bool
			result2 error
		})
	}
	fake.checkCollectionPolicyReturnsOnCall[i] = struct {
		result1 bool
		result2 error
	}{result1, result2}
}

func (fake *CollectionPolicyChecker) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.checkCollectionPolicyMutex.RLock()
	defer fake.checkCollectionPolicyMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *CollectionPolicyChecker) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}
