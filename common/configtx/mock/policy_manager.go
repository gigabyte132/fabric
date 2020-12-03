// Code generated by counterfeiter. DO NOT EDIT.
package mock

import (
	"sync"

	"github.com/hyperledger/fabric/common/policies"
)

type PolicyManager struct {
	GetPolicyStub        func(string) (policies.Policy, bool)
	getPolicyMutex       sync.RWMutex
	getPolicyArgsForCall []struct {
		arg1 string
	}
	getPolicyReturns struct {
		result1 policies.Policy
		result2 bool
	}
	getPolicyReturnsOnCall map[int]struct {
		result1 policies.Policy
		result2 bool
	}
	ManagerStub        func([]string) (policies.Manager, bool)
	managerMutex       sync.RWMutex
	managerArgsForCall []struct {
		arg1 []string
	}
	managerReturns struct {
		result1 policies.Manager
		result2 bool
	}
	managerReturnsOnCall map[int]struct {
		result1 policies.Manager
		result2 bool
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *PolicyManager) GetPolicy(arg1 string) (policies.Policy, bool) {
	fake.getPolicyMutex.Lock()
	ret, specificReturn := fake.getPolicyReturnsOnCall[len(fake.getPolicyArgsForCall)]
	fake.getPolicyArgsForCall = append(fake.getPolicyArgsForCall, struct {
		arg1 string
	}{arg1})
	stub := fake.GetPolicyStub
	fakeReturns := fake.getPolicyReturns
	fake.recordInvocation("GetPolicy", []interface{}{arg1})
	fake.getPolicyMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *PolicyManager) GetPolicyCallCount() int {
	fake.getPolicyMutex.RLock()
	defer fake.getPolicyMutex.RUnlock()
	return len(fake.getPolicyArgsForCall)
}

func (fake *PolicyManager) GetPolicyCalls(stub func(string) (policies.Policy, bool)) {
	fake.getPolicyMutex.Lock()
	defer fake.getPolicyMutex.Unlock()
	fake.GetPolicyStub = stub
}

func (fake *PolicyManager) GetPolicyArgsForCall(i int) string {
	fake.getPolicyMutex.RLock()
	defer fake.getPolicyMutex.RUnlock()
	argsForCall := fake.getPolicyArgsForCall[i]
	return argsForCall.arg1
}

func (fake *PolicyManager) GetPolicyReturns(result1 policies.Policy, result2 bool) {
	fake.getPolicyMutex.Lock()
	defer fake.getPolicyMutex.Unlock()
	fake.GetPolicyStub = nil
	fake.getPolicyReturns = struct {
		result1 policies.Policy
		result2 bool
	}{result1, result2}
}

func (fake *PolicyManager) GetPolicyReturnsOnCall(i int, result1 policies.Policy, result2 bool) {
	fake.getPolicyMutex.Lock()
	defer fake.getPolicyMutex.Unlock()
	fake.GetPolicyStub = nil
	if fake.getPolicyReturnsOnCall == nil {
		fake.getPolicyReturnsOnCall = make(map[int]struct {
			result1 policies.Policy
			result2 bool
		})
	}
	fake.getPolicyReturnsOnCall[i] = struct {
		result1 policies.Policy
		result2 bool
	}{result1, result2}
}

func (fake *PolicyManager) Manager(arg1 []string) (policies.Manager, bool) {
	var arg1Copy []string
	if arg1 != nil {
		arg1Copy = make([]string, len(arg1))
		copy(arg1Copy, arg1)
	}
	fake.managerMutex.Lock()
	ret, specificReturn := fake.managerReturnsOnCall[len(fake.managerArgsForCall)]
	fake.managerArgsForCall = append(fake.managerArgsForCall, struct {
		arg1 []string
	}{arg1Copy})
	stub := fake.ManagerStub
	fakeReturns := fake.managerReturns
	fake.recordInvocation("Manager", []interface{}{arg1Copy})
	fake.managerMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *PolicyManager) ManagerCallCount() int {
	fake.managerMutex.RLock()
	defer fake.managerMutex.RUnlock()
	return len(fake.managerArgsForCall)
}

func (fake *PolicyManager) ManagerCalls(stub func([]string) (policies.Manager, bool)) {
	fake.managerMutex.Lock()
	defer fake.managerMutex.Unlock()
	fake.ManagerStub = stub
}

func (fake *PolicyManager) ManagerArgsForCall(i int) []string {
	fake.managerMutex.RLock()
	defer fake.managerMutex.RUnlock()
	argsForCall := fake.managerArgsForCall[i]
	return argsForCall.arg1
}

func (fake *PolicyManager) ManagerReturns(result1 policies.Manager, result2 bool) {
	fake.managerMutex.Lock()
	defer fake.managerMutex.Unlock()
	fake.ManagerStub = nil
	fake.managerReturns = struct {
		result1 policies.Manager
		result2 bool
	}{result1, result2}
}

func (fake *PolicyManager) ManagerReturnsOnCall(i int, result1 policies.Manager, result2 bool) {
	fake.managerMutex.Lock()
	defer fake.managerMutex.Unlock()
	fake.ManagerStub = nil
	if fake.managerReturnsOnCall == nil {
		fake.managerReturnsOnCall = make(map[int]struct {
			result1 policies.Manager
			result2 bool
		})
	}
	fake.managerReturnsOnCall[i] = struct {
		result1 policies.Manager
		result2 bool
	}{result1, result2}
}

func (fake *PolicyManager) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.getPolicyMutex.RLock()
	defer fake.getPolicyMutex.RUnlock()
	fake.managerMutex.RLock()
	defer fake.managerMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *PolicyManager) recordInvocation(key string, args []interface{}) {
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
