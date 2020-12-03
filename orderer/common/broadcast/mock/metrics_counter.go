// Code generated by counterfeiter. DO NOT EDIT.
package mock

import (
	"sync"

	"github.com/hyperledger/fabric/common/metrics"
)

type MetricsCounter struct {
	AddStub        func(float64)
	addMutex       sync.RWMutex
	addArgsForCall []struct {
		arg1 float64
	}
	WithStub        func(...string) metrics.Counter
	withMutex       sync.RWMutex
	withArgsForCall []struct {
		arg1 []string
	}
	withReturns struct {
		result1 metrics.Counter
	}
	withReturnsOnCall map[int]struct {
		result1 metrics.Counter
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *MetricsCounter) Add(arg1 float64) {
	fake.addMutex.Lock()
	fake.addArgsForCall = append(fake.addArgsForCall, struct {
		arg1 float64
	}{arg1})
	stub := fake.AddStub
	fake.recordInvocation("Add", []interface{}{arg1})
	fake.addMutex.Unlock()
	if stub != nil {
		fake.AddStub(arg1)
	}
}

func (fake *MetricsCounter) AddCallCount() int {
	fake.addMutex.RLock()
	defer fake.addMutex.RUnlock()
	return len(fake.addArgsForCall)
}

func (fake *MetricsCounter) AddCalls(stub func(float64)) {
	fake.addMutex.Lock()
	defer fake.addMutex.Unlock()
	fake.AddStub = stub
}

func (fake *MetricsCounter) AddArgsForCall(i int) float64 {
	fake.addMutex.RLock()
	defer fake.addMutex.RUnlock()
	argsForCall := fake.addArgsForCall[i]
	return argsForCall.arg1
}

func (fake *MetricsCounter) With(arg1 ...string) metrics.Counter {
	fake.withMutex.Lock()
	ret, specificReturn := fake.withReturnsOnCall[len(fake.withArgsForCall)]
	fake.withArgsForCall = append(fake.withArgsForCall, struct {
		arg1 []string
	}{arg1})
	stub := fake.WithStub
	fakeReturns := fake.withReturns
	fake.recordInvocation("With", []interface{}{arg1})
	fake.withMutex.Unlock()
	if stub != nil {
		return stub(arg1...)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *MetricsCounter) WithCallCount() int {
	fake.withMutex.RLock()
	defer fake.withMutex.RUnlock()
	return len(fake.withArgsForCall)
}

func (fake *MetricsCounter) WithCalls(stub func(...string) metrics.Counter) {
	fake.withMutex.Lock()
	defer fake.withMutex.Unlock()
	fake.WithStub = stub
}

func (fake *MetricsCounter) WithArgsForCall(i int) []string {
	fake.withMutex.RLock()
	defer fake.withMutex.RUnlock()
	argsForCall := fake.withArgsForCall[i]
	return argsForCall.arg1
}

func (fake *MetricsCounter) WithReturns(result1 metrics.Counter) {
	fake.withMutex.Lock()
	defer fake.withMutex.Unlock()
	fake.WithStub = nil
	fake.withReturns = struct {
		result1 metrics.Counter
	}{result1}
}

func (fake *MetricsCounter) WithReturnsOnCall(i int, result1 metrics.Counter) {
	fake.withMutex.Lock()
	defer fake.withMutex.Unlock()
	fake.WithStub = nil
	if fake.withReturnsOnCall == nil {
		fake.withReturnsOnCall = make(map[int]struct {
			result1 metrics.Counter
		})
	}
	fake.withReturnsOnCall[i] = struct {
		result1 metrics.Counter
	}{result1}
}

func (fake *MetricsCounter) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.addMutex.RLock()
	defer fake.addMutex.RUnlock()
	fake.withMutex.RLock()
	defer fake.withMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *MetricsCounter) recordInvocation(key string, args []interface{}) {
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
