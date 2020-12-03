// Code generated by counterfeiter. DO NOT EDIT.
package mock

import (
	"context"
	"sync"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/hyperledger/fabric-protos-go/peer"
	"google.golang.org/grpc"
)

type SnapshotClient struct {
	CancelStub        func(context.Context, *peer.SignedSnapshotRequest, ...grpc.CallOption) (*empty.Empty, error)
	cancelMutex       sync.RWMutex
	cancelArgsForCall []struct {
		arg1 context.Context
		arg2 *peer.SignedSnapshotRequest
		arg3 []grpc.CallOption
	}
	cancelReturns struct {
		result1 *empty.Empty
		result2 error
	}
	cancelReturnsOnCall map[int]struct {
		result1 *empty.Empty
		result2 error
	}
	GenerateStub        func(context.Context, *peer.SignedSnapshotRequest, ...grpc.CallOption) (*empty.Empty, error)
	generateMutex       sync.RWMutex
	generateArgsForCall []struct {
		arg1 context.Context
		arg2 *peer.SignedSnapshotRequest
		arg3 []grpc.CallOption
	}
	generateReturns struct {
		result1 *empty.Empty
		result2 error
	}
	generateReturnsOnCall map[int]struct {
		result1 *empty.Empty
		result2 error
	}
	QueryPendingsStub        func(context.Context, *peer.SignedSnapshotRequest, ...grpc.CallOption) (*peer.QueryPendingSnapshotsResponse, error)
	queryPendingsMutex       sync.RWMutex
	queryPendingsArgsForCall []struct {
		arg1 context.Context
		arg2 *peer.SignedSnapshotRequest
		arg3 []grpc.CallOption
	}
	queryPendingsReturns struct {
		result1 *peer.QueryPendingSnapshotsResponse
		result2 error
	}
	queryPendingsReturnsOnCall map[int]struct {
		result1 *peer.QueryPendingSnapshotsResponse
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *SnapshotClient) Cancel(arg1 context.Context, arg2 *peer.SignedSnapshotRequest, arg3 ...grpc.CallOption) (*empty.Empty, error) {
	fake.cancelMutex.Lock()
	ret, specificReturn := fake.cancelReturnsOnCall[len(fake.cancelArgsForCall)]
	fake.cancelArgsForCall = append(fake.cancelArgsForCall, struct {
		arg1 context.Context
		arg2 *peer.SignedSnapshotRequest
		arg3 []grpc.CallOption
	}{arg1, arg2, arg3})
	stub := fake.CancelStub
	fakeReturns := fake.cancelReturns
	fake.recordInvocation("Cancel", []interface{}{arg1, arg2, arg3})
	fake.cancelMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3...)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *SnapshotClient) CancelCallCount() int {
	fake.cancelMutex.RLock()
	defer fake.cancelMutex.RUnlock()
	return len(fake.cancelArgsForCall)
}

func (fake *SnapshotClient) CancelCalls(stub func(context.Context, *peer.SignedSnapshotRequest, ...grpc.CallOption) (*empty.Empty, error)) {
	fake.cancelMutex.Lock()
	defer fake.cancelMutex.Unlock()
	fake.CancelStub = stub
}

func (fake *SnapshotClient) CancelArgsForCall(i int) (context.Context, *peer.SignedSnapshotRequest, []grpc.CallOption) {
	fake.cancelMutex.RLock()
	defer fake.cancelMutex.RUnlock()
	argsForCall := fake.cancelArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *SnapshotClient) CancelReturns(result1 *empty.Empty, result2 error) {
	fake.cancelMutex.Lock()
	defer fake.cancelMutex.Unlock()
	fake.CancelStub = nil
	fake.cancelReturns = struct {
		result1 *empty.Empty
		result2 error
	}{result1, result2}
}

func (fake *SnapshotClient) CancelReturnsOnCall(i int, result1 *empty.Empty, result2 error) {
	fake.cancelMutex.Lock()
	defer fake.cancelMutex.Unlock()
	fake.CancelStub = nil
	if fake.cancelReturnsOnCall == nil {
		fake.cancelReturnsOnCall = make(map[int]struct {
			result1 *empty.Empty
			result2 error
		})
	}
	fake.cancelReturnsOnCall[i] = struct {
		result1 *empty.Empty
		result2 error
	}{result1, result2}
}

func (fake *SnapshotClient) Generate(arg1 context.Context, arg2 *peer.SignedSnapshotRequest, arg3 ...grpc.CallOption) (*empty.Empty, error) {
	fake.generateMutex.Lock()
	ret, specificReturn := fake.generateReturnsOnCall[len(fake.generateArgsForCall)]
	fake.generateArgsForCall = append(fake.generateArgsForCall, struct {
		arg1 context.Context
		arg2 *peer.SignedSnapshotRequest
		arg3 []grpc.CallOption
	}{arg1, arg2, arg3})
	stub := fake.GenerateStub
	fakeReturns := fake.generateReturns
	fake.recordInvocation("Generate", []interface{}{arg1, arg2, arg3})
	fake.generateMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3...)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *SnapshotClient) GenerateCallCount() int {
	fake.generateMutex.RLock()
	defer fake.generateMutex.RUnlock()
	return len(fake.generateArgsForCall)
}

func (fake *SnapshotClient) GenerateCalls(stub func(context.Context, *peer.SignedSnapshotRequest, ...grpc.CallOption) (*empty.Empty, error)) {
	fake.generateMutex.Lock()
	defer fake.generateMutex.Unlock()
	fake.GenerateStub = stub
}

func (fake *SnapshotClient) GenerateArgsForCall(i int) (context.Context, *peer.SignedSnapshotRequest, []grpc.CallOption) {
	fake.generateMutex.RLock()
	defer fake.generateMutex.RUnlock()
	argsForCall := fake.generateArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *SnapshotClient) GenerateReturns(result1 *empty.Empty, result2 error) {
	fake.generateMutex.Lock()
	defer fake.generateMutex.Unlock()
	fake.GenerateStub = nil
	fake.generateReturns = struct {
		result1 *empty.Empty
		result2 error
	}{result1, result2}
}

func (fake *SnapshotClient) GenerateReturnsOnCall(i int, result1 *empty.Empty, result2 error) {
	fake.generateMutex.Lock()
	defer fake.generateMutex.Unlock()
	fake.GenerateStub = nil
	if fake.generateReturnsOnCall == nil {
		fake.generateReturnsOnCall = make(map[int]struct {
			result1 *empty.Empty
			result2 error
		})
	}
	fake.generateReturnsOnCall[i] = struct {
		result1 *empty.Empty
		result2 error
	}{result1, result2}
}

func (fake *SnapshotClient) QueryPendings(arg1 context.Context, arg2 *peer.SignedSnapshotRequest, arg3 ...grpc.CallOption) (*peer.QueryPendingSnapshotsResponse, error) {
	fake.queryPendingsMutex.Lock()
	ret, specificReturn := fake.queryPendingsReturnsOnCall[len(fake.queryPendingsArgsForCall)]
	fake.queryPendingsArgsForCall = append(fake.queryPendingsArgsForCall, struct {
		arg1 context.Context
		arg2 *peer.SignedSnapshotRequest
		arg3 []grpc.CallOption
	}{arg1, arg2, arg3})
	stub := fake.QueryPendingsStub
	fakeReturns := fake.queryPendingsReturns
	fake.recordInvocation("QueryPendings", []interface{}{arg1, arg2, arg3})
	fake.queryPendingsMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3...)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *SnapshotClient) QueryPendingsCallCount() int {
	fake.queryPendingsMutex.RLock()
	defer fake.queryPendingsMutex.RUnlock()
	return len(fake.queryPendingsArgsForCall)
}

func (fake *SnapshotClient) QueryPendingsCalls(stub func(context.Context, *peer.SignedSnapshotRequest, ...grpc.CallOption) (*peer.QueryPendingSnapshotsResponse, error)) {
	fake.queryPendingsMutex.Lock()
	defer fake.queryPendingsMutex.Unlock()
	fake.QueryPendingsStub = stub
}

func (fake *SnapshotClient) QueryPendingsArgsForCall(i int) (context.Context, *peer.SignedSnapshotRequest, []grpc.CallOption) {
	fake.queryPendingsMutex.RLock()
	defer fake.queryPendingsMutex.RUnlock()
	argsForCall := fake.queryPendingsArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *SnapshotClient) QueryPendingsReturns(result1 *peer.QueryPendingSnapshotsResponse, result2 error) {
	fake.queryPendingsMutex.Lock()
	defer fake.queryPendingsMutex.Unlock()
	fake.QueryPendingsStub = nil
	fake.queryPendingsReturns = struct {
		result1 *peer.QueryPendingSnapshotsResponse
		result2 error
	}{result1, result2}
}

func (fake *SnapshotClient) QueryPendingsReturnsOnCall(i int, result1 *peer.QueryPendingSnapshotsResponse, result2 error) {
	fake.queryPendingsMutex.Lock()
	defer fake.queryPendingsMutex.Unlock()
	fake.QueryPendingsStub = nil
	if fake.queryPendingsReturnsOnCall == nil {
		fake.queryPendingsReturnsOnCall = make(map[int]struct {
			result1 *peer.QueryPendingSnapshotsResponse
			result2 error
		})
	}
	fake.queryPendingsReturnsOnCall[i] = struct {
		result1 *peer.QueryPendingSnapshotsResponse
		result2 error
	}{result1, result2}
}

func (fake *SnapshotClient) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.cancelMutex.RLock()
	defer fake.cancelMutex.RUnlock()
	fake.generateMutex.RLock()
	defer fake.generateMutex.RUnlock()
	fake.queryPendingsMutex.RLock()
	defer fake.queryPendingsMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *SnapshotClient) recordInvocation(key string, args []interface{}) {
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
