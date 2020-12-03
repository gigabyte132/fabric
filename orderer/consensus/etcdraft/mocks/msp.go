// Code generated by counterfeiter. DO NOT EDIT.
package mocks

import (
	"sync"

	mspa "github.com/hyperledger/fabric-protos-go/msp"
	"github.com/hyperledger/fabric/msp"
)

type MSP struct {
	DeserializeIdentityStub        func([]byte) (msp.Identity, error)
	deserializeIdentityMutex       sync.RWMutex
	deserializeIdentityArgsForCall []struct {
		arg1 []byte
	}
	deserializeIdentityReturns struct {
		result1 msp.Identity
		result2 error
	}
	deserializeIdentityReturnsOnCall map[int]struct {
		result1 msp.Identity
		result2 error
	}
	GetDefaultSigningIdentityStub        func() (msp.SigningIdentity, error)
	getDefaultSigningIdentityMutex       sync.RWMutex
	getDefaultSigningIdentityArgsForCall []struct {
	}
	getDefaultSigningIdentityReturns struct {
		result1 msp.SigningIdentity
		result2 error
	}
	getDefaultSigningIdentityReturnsOnCall map[int]struct {
		result1 msp.SigningIdentity
		result2 error
	}
	GetIdentifierStub        func() (string, error)
	getIdentifierMutex       sync.RWMutex
	getIdentifierArgsForCall []struct {
	}
	getIdentifierReturns struct {
		result1 string
		result2 error
	}
	getIdentifierReturnsOnCall map[int]struct {
		result1 string
		result2 error
	}
	GetSigningIdentityStub        func(*msp.IdentityIdentifier) (msp.SigningIdentity, error)
	getSigningIdentityMutex       sync.RWMutex
	getSigningIdentityArgsForCall []struct {
		arg1 *msp.IdentityIdentifier
	}
	getSigningIdentityReturns struct {
		result1 msp.SigningIdentity
		result2 error
	}
	getSigningIdentityReturnsOnCall map[int]struct {
		result1 msp.SigningIdentity
		result2 error
	}
	GetTLSIntermediateCertsStub        func() [][]byte
	getTLSIntermediateCertsMutex       sync.RWMutex
	getTLSIntermediateCertsArgsForCall []struct {
	}
	getTLSIntermediateCertsReturns struct {
		result1 [][]byte
	}
	getTLSIntermediateCertsReturnsOnCall map[int]struct {
		result1 [][]byte
	}
	GetTLSRootCertsStub        func() [][]byte
	getTLSRootCertsMutex       sync.RWMutex
	getTLSRootCertsArgsForCall []struct {
	}
	getTLSRootCertsReturns struct {
		result1 [][]byte
	}
	getTLSRootCertsReturnsOnCall map[int]struct {
		result1 [][]byte
	}
	GetTypeStub        func() msp.ProviderType
	getTypeMutex       sync.RWMutex
	getTypeArgsForCall []struct {
	}
	getTypeReturns struct {
		result1 msp.ProviderType
	}
	getTypeReturnsOnCall map[int]struct {
		result1 msp.ProviderType
	}
	GetVersionStub        func() msp.MSPVersion
	getVersionMutex       sync.RWMutex
	getVersionArgsForCall []struct {
	}
	getVersionReturns struct {
		result1 msp.MSPVersion
	}
	getVersionReturnsOnCall map[int]struct {
		result1 msp.MSPVersion
	}
	IsWellFormedStub        func(*mspa.SerializedIdentity) error
	isWellFormedMutex       sync.RWMutex
	isWellFormedArgsForCall []struct {
		arg1 *mspa.SerializedIdentity
	}
	isWellFormedReturns struct {
		result1 error
	}
	isWellFormedReturnsOnCall map[int]struct {
		result1 error
	}
	SatisfiesPrincipalStub        func(msp.Identity, *mspa.MSPPrincipal) error
	satisfiesPrincipalMutex       sync.RWMutex
	satisfiesPrincipalArgsForCall []struct {
		arg1 msp.Identity
		arg2 *mspa.MSPPrincipal
	}
	satisfiesPrincipalReturns struct {
		result1 error
	}
	satisfiesPrincipalReturnsOnCall map[int]struct {
		result1 error
	}
	SetupStub        func(*mspa.MSPConfig) error
	setupMutex       sync.RWMutex
	setupArgsForCall []struct {
		arg1 *mspa.MSPConfig
	}
	setupReturns struct {
		result1 error
	}
	setupReturnsOnCall map[int]struct {
		result1 error
	}
	ValidateStub        func(msp.Identity) error
	validateMutex       sync.RWMutex
	validateArgsForCall []struct {
		arg1 msp.Identity
	}
	validateReturns struct {
		result1 error
	}
	validateReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *MSP) DeserializeIdentity(arg1 []byte) (msp.Identity, error) {
	var arg1Copy []byte
	if arg1 != nil {
		arg1Copy = make([]byte, len(arg1))
		copy(arg1Copy, arg1)
	}
	fake.deserializeIdentityMutex.Lock()
	ret, specificReturn := fake.deserializeIdentityReturnsOnCall[len(fake.deserializeIdentityArgsForCall)]
	fake.deserializeIdentityArgsForCall = append(fake.deserializeIdentityArgsForCall, struct {
		arg1 []byte
	}{arg1Copy})
	stub := fake.DeserializeIdentityStub
	fakeReturns := fake.deserializeIdentityReturns
	fake.recordInvocation("DeserializeIdentity", []interface{}{arg1Copy})
	fake.deserializeIdentityMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *MSP) DeserializeIdentityCallCount() int {
	fake.deserializeIdentityMutex.RLock()
	defer fake.deserializeIdentityMutex.RUnlock()
	return len(fake.deserializeIdentityArgsForCall)
}

func (fake *MSP) DeserializeIdentityCalls(stub func([]byte) (msp.Identity, error)) {
	fake.deserializeIdentityMutex.Lock()
	defer fake.deserializeIdentityMutex.Unlock()
	fake.DeserializeIdentityStub = stub
}

func (fake *MSP) DeserializeIdentityArgsForCall(i int) []byte {
	fake.deserializeIdentityMutex.RLock()
	defer fake.deserializeIdentityMutex.RUnlock()
	argsForCall := fake.deserializeIdentityArgsForCall[i]
	return argsForCall.arg1
}

func (fake *MSP) DeserializeIdentityReturns(result1 msp.Identity, result2 error) {
	fake.deserializeIdentityMutex.Lock()
	defer fake.deserializeIdentityMutex.Unlock()
	fake.DeserializeIdentityStub = nil
	fake.deserializeIdentityReturns = struct {
		result1 msp.Identity
		result2 error
	}{result1, result2}
}

func (fake *MSP) DeserializeIdentityReturnsOnCall(i int, result1 msp.Identity, result2 error) {
	fake.deserializeIdentityMutex.Lock()
	defer fake.deserializeIdentityMutex.Unlock()
	fake.DeserializeIdentityStub = nil
	if fake.deserializeIdentityReturnsOnCall == nil {
		fake.deserializeIdentityReturnsOnCall = make(map[int]struct {
			result1 msp.Identity
			result2 error
		})
	}
	fake.deserializeIdentityReturnsOnCall[i] = struct {
		result1 msp.Identity
		result2 error
	}{result1, result2}
}

func (fake *MSP) GetDefaultSigningIdentity() (msp.SigningIdentity, error) {
	fake.getDefaultSigningIdentityMutex.Lock()
	ret, specificReturn := fake.getDefaultSigningIdentityReturnsOnCall[len(fake.getDefaultSigningIdentityArgsForCall)]
	fake.getDefaultSigningIdentityArgsForCall = append(fake.getDefaultSigningIdentityArgsForCall, struct {
	}{})
	stub := fake.GetDefaultSigningIdentityStub
	fakeReturns := fake.getDefaultSigningIdentityReturns
	fake.recordInvocation("GetDefaultSigningIdentity", []interface{}{})
	fake.getDefaultSigningIdentityMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *MSP) GetDefaultSigningIdentityCallCount() int {
	fake.getDefaultSigningIdentityMutex.RLock()
	defer fake.getDefaultSigningIdentityMutex.RUnlock()
	return len(fake.getDefaultSigningIdentityArgsForCall)
}

func (fake *MSP) GetDefaultSigningIdentityCalls(stub func() (msp.SigningIdentity, error)) {
	fake.getDefaultSigningIdentityMutex.Lock()
	defer fake.getDefaultSigningIdentityMutex.Unlock()
	fake.GetDefaultSigningIdentityStub = stub
}

func (fake *MSP) GetDefaultSigningIdentityReturns(result1 msp.SigningIdentity, result2 error) {
	fake.getDefaultSigningIdentityMutex.Lock()
	defer fake.getDefaultSigningIdentityMutex.Unlock()
	fake.GetDefaultSigningIdentityStub = nil
	fake.getDefaultSigningIdentityReturns = struct {
		result1 msp.SigningIdentity
		result2 error
	}{result1, result2}
}

func (fake *MSP) GetDefaultSigningIdentityReturnsOnCall(i int, result1 msp.SigningIdentity, result2 error) {
	fake.getDefaultSigningIdentityMutex.Lock()
	defer fake.getDefaultSigningIdentityMutex.Unlock()
	fake.GetDefaultSigningIdentityStub = nil
	if fake.getDefaultSigningIdentityReturnsOnCall == nil {
		fake.getDefaultSigningIdentityReturnsOnCall = make(map[int]struct {
			result1 msp.SigningIdentity
			result2 error
		})
	}
	fake.getDefaultSigningIdentityReturnsOnCall[i] = struct {
		result1 msp.SigningIdentity
		result2 error
	}{result1, result2}
}

func (fake *MSP) GetIdentifier() (string, error) {
	fake.getIdentifierMutex.Lock()
	ret, specificReturn := fake.getIdentifierReturnsOnCall[len(fake.getIdentifierArgsForCall)]
	fake.getIdentifierArgsForCall = append(fake.getIdentifierArgsForCall, struct {
	}{})
	stub := fake.GetIdentifierStub
	fakeReturns := fake.getIdentifierReturns
	fake.recordInvocation("GetIdentifier", []interface{}{})
	fake.getIdentifierMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *MSP) GetIdentifierCallCount() int {
	fake.getIdentifierMutex.RLock()
	defer fake.getIdentifierMutex.RUnlock()
	return len(fake.getIdentifierArgsForCall)
}

func (fake *MSP) GetIdentifierCalls(stub func() (string, error)) {
	fake.getIdentifierMutex.Lock()
	defer fake.getIdentifierMutex.Unlock()
	fake.GetIdentifierStub = stub
}

func (fake *MSP) GetIdentifierReturns(result1 string, result2 error) {
	fake.getIdentifierMutex.Lock()
	defer fake.getIdentifierMutex.Unlock()
	fake.GetIdentifierStub = nil
	fake.getIdentifierReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *MSP) GetIdentifierReturnsOnCall(i int, result1 string, result2 error) {
	fake.getIdentifierMutex.Lock()
	defer fake.getIdentifierMutex.Unlock()
	fake.GetIdentifierStub = nil
	if fake.getIdentifierReturnsOnCall == nil {
		fake.getIdentifierReturnsOnCall = make(map[int]struct {
			result1 string
			result2 error
		})
	}
	fake.getIdentifierReturnsOnCall[i] = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *MSP) GetSigningIdentity(arg1 *msp.IdentityIdentifier) (msp.SigningIdentity, error) {
	fake.getSigningIdentityMutex.Lock()
	ret, specificReturn := fake.getSigningIdentityReturnsOnCall[len(fake.getSigningIdentityArgsForCall)]
	fake.getSigningIdentityArgsForCall = append(fake.getSigningIdentityArgsForCall, struct {
		arg1 *msp.IdentityIdentifier
	}{arg1})
	stub := fake.GetSigningIdentityStub
	fakeReturns := fake.getSigningIdentityReturns
	fake.recordInvocation("GetSigningIdentity", []interface{}{arg1})
	fake.getSigningIdentityMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *MSP) GetSigningIdentityCallCount() int {
	fake.getSigningIdentityMutex.RLock()
	defer fake.getSigningIdentityMutex.RUnlock()
	return len(fake.getSigningIdentityArgsForCall)
}

func (fake *MSP) GetSigningIdentityCalls(stub func(*msp.IdentityIdentifier) (msp.SigningIdentity, error)) {
	fake.getSigningIdentityMutex.Lock()
	defer fake.getSigningIdentityMutex.Unlock()
	fake.GetSigningIdentityStub = stub
}

func (fake *MSP) GetSigningIdentityArgsForCall(i int) *msp.IdentityIdentifier {
	fake.getSigningIdentityMutex.RLock()
	defer fake.getSigningIdentityMutex.RUnlock()
	argsForCall := fake.getSigningIdentityArgsForCall[i]
	return argsForCall.arg1
}

func (fake *MSP) GetSigningIdentityReturns(result1 msp.SigningIdentity, result2 error) {
	fake.getSigningIdentityMutex.Lock()
	defer fake.getSigningIdentityMutex.Unlock()
	fake.GetSigningIdentityStub = nil
	fake.getSigningIdentityReturns = struct {
		result1 msp.SigningIdentity
		result2 error
	}{result1, result2}
}

func (fake *MSP) GetSigningIdentityReturnsOnCall(i int, result1 msp.SigningIdentity, result2 error) {
	fake.getSigningIdentityMutex.Lock()
	defer fake.getSigningIdentityMutex.Unlock()
	fake.GetSigningIdentityStub = nil
	if fake.getSigningIdentityReturnsOnCall == nil {
		fake.getSigningIdentityReturnsOnCall = make(map[int]struct {
			result1 msp.SigningIdentity
			result2 error
		})
	}
	fake.getSigningIdentityReturnsOnCall[i] = struct {
		result1 msp.SigningIdentity
		result2 error
	}{result1, result2}
}

func (fake *MSP) GetTLSIntermediateCerts() [][]byte {
	fake.getTLSIntermediateCertsMutex.Lock()
	ret, specificReturn := fake.getTLSIntermediateCertsReturnsOnCall[len(fake.getTLSIntermediateCertsArgsForCall)]
	fake.getTLSIntermediateCertsArgsForCall = append(fake.getTLSIntermediateCertsArgsForCall, struct {
	}{})
	stub := fake.GetTLSIntermediateCertsStub
	fakeReturns := fake.getTLSIntermediateCertsReturns
	fake.recordInvocation("GetTLSIntermediateCerts", []interface{}{})
	fake.getTLSIntermediateCertsMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *MSP) GetTLSIntermediateCertsCallCount() int {
	fake.getTLSIntermediateCertsMutex.RLock()
	defer fake.getTLSIntermediateCertsMutex.RUnlock()
	return len(fake.getTLSIntermediateCertsArgsForCall)
}

func (fake *MSP) GetTLSIntermediateCertsCalls(stub func() [][]byte) {
	fake.getTLSIntermediateCertsMutex.Lock()
	defer fake.getTLSIntermediateCertsMutex.Unlock()
	fake.GetTLSIntermediateCertsStub = stub
}

func (fake *MSP) GetTLSIntermediateCertsReturns(result1 [][]byte) {
	fake.getTLSIntermediateCertsMutex.Lock()
	defer fake.getTLSIntermediateCertsMutex.Unlock()
	fake.GetTLSIntermediateCertsStub = nil
	fake.getTLSIntermediateCertsReturns = struct {
		result1 [][]byte
	}{result1}
}

func (fake *MSP) GetTLSIntermediateCertsReturnsOnCall(i int, result1 [][]byte) {
	fake.getTLSIntermediateCertsMutex.Lock()
	defer fake.getTLSIntermediateCertsMutex.Unlock()
	fake.GetTLSIntermediateCertsStub = nil
	if fake.getTLSIntermediateCertsReturnsOnCall == nil {
		fake.getTLSIntermediateCertsReturnsOnCall = make(map[int]struct {
			result1 [][]byte
		})
	}
	fake.getTLSIntermediateCertsReturnsOnCall[i] = struct {
		result1 [][]byte
	}{result1}
}

func (fake *MSP) GetTLSRootCerts() [][]byte {
	fake.getTLSRootCertsMutex.Lock()
	ret, specificReturn := fake.getTLSRootCertsReturnsOnCall[len(fake.getTLSRootCertsArgsForCall)]
	fake.getTLSRootCertsArgsForCall = append(fake.getTLSRootCertsArgsForCall, struct {
	}{})
	stub := fake.GetTLSRootCertsStub
	fakeReturns := fake.getTLSRootCertsReturns
	fake.recordInvocation("GetTLSRootCerts", []interface{}{})
	fake.getTLSRootCertsMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *MSP) GetTLSRootCertsCallCount() int {
	fake.getTLSRootCertsMutex.RLock()
	defer fake.getTLSRootCertsMutex.RUnlock()
	return len(fake.getTLSRootCertsArgsForCall)
}

func (fake *MSP) GetTLSRootCertsCalls(stub func() [][]byte) {
	fake.getTLSRootCertsMutex.Lock()
	defer fake.getTLSRootCertsMutex.Unlock()
	fake.GetTLSRootCertsStub = stub
}

func (fake *MSP) GetTLSRootCertsReturns(result1 [][]byte) {
	fake.getTLSRootCertsMutex.Lock()
	defer fake.getTLSRootCertsMutex.Unlock()
	fake.GetTLSRootCertsStub = nil
	fake.getTLSRootCertsReturns = struct {
		result1 [][]byte
	}{result1}
}

func (fake *MSP) GetTLSRootCertsReturnsOnCall(i int, result1 [][]byte) {
	fake.getTLSRootCertsMutex.Lock()
	defer fake.getTLSRootCertsMutex.Unlock()
	fake.GetTLSRootCertsStub = nil
	if fake.getTLSRootCertsReturnsOnCall == nil {
		fake.getTLSRootCertsReturnsOnCall = make(map[int]struct {
			result1 [][]byte
		})
	}
	fake.getTLSRootCertsReturnsOnCall[i] = struct {
		result1 [][]byte
	}{result1}
}

func (fake *MSP) GetType() msp.ProviderType {
	fake.getTypeMutex.Lock()
	ret, specificReturn := fake.getTypeReturnsOnCall[len(fake.getTypeArgsForCall)]
	fake.getTypeArgsForCall = append(fake.getTypeArgsForCall, struct {
	}{})
	stub := fake.GetTypeStub
	fakeReturns := fake.getTypeReturns
	fake.recordInvocation("GetType", []interface{}{})
	fake.getTypeMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *MSP) GetTypeCallCount() int {
	fake.getTypeMutex.RLock()
	defer fake.getTypeMutex.RUnlock()
	return len(fake.getTypeArgsForCall)
}

func (fake *MSP) GetTypeCalls(stub func() msp.ProviderType) {
	fake.getTypeMutex.Lock()
	defer fake.getTypeMutex.Unlock()
	fake.GetTypeStub = stub
}

func (fake *MSP) GetTypeReturns(result1 msp.ProviderType) {
	fake.getTypeMutex.Lock()
	defer fake.getTypeMutex.Unlock()
	fake.GetTypeStub = nil
	fake.getTypeReturns = struct {
		result1 msp.ProviderType
	}{result1}
}

func (fake *MSP) GetTypeReturnsOnCall(i int, result1 msp.ProviderType) {
	fake.getTypeMutex.Lock()
	defer fake.getTypeMutex.Unlock()
	fake.GetTypeStub = nil
	if fake.getTypeReturnsOnCall == nil {
		fake.getTypeReturnsOnCall = make(map[int]struct {
			result1 msp.ProviderType
		})
	}
	fake.getTypeReturnsOnCall[i] = struct {
		result1 msp.ProviderType
	}{result1}
}

func (fake *MSP) GetVersion() msp.MSPVersion {
	fake.getVersionMutex.Lock()
	ret, specificReturn := fake.getVersionReturnsOnCall[len(fake.getVersionArgsForCall)]
	fake.getVersionArgsForCall = append(fake.getVersionArgsForCall, struct {
	}{})
	stub := fake.GetVersionStub
	fakeReturns := fake.getVersionReturns
	fake.recordInvocation("GetVersion", []interface{}{})
	fake.getVersionMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *MSP) GetVersionCallCount() int {
	fake.getVersionMutex.RLock()
	defer fake.getVersionMutex.RUnlock()
	return len(fake.getVersionArgsForCall)
}

func (fake *MSP) GetVersionCalls(stub func() msp.MSPVersion) {
	fake.getVersionMutex.Lock()
	defer fake.getVersionMutex.Unlock()
	fake.GetVersionStub = stub
}

func (fake *MSP) GetVersionReturns(result1 msp.MSPVersion) {
	fake.getVersionMutex.Lock()
	defer fake.getVersionMutex.Unlock()
	fake.GetVersionStub = nil
	fake.getVersionReturns = struct {
		result1 msp.MSPVersion
	}{result1}
}

func (fake *MSP) GetVersionReturnsOnCall(i int, result1 msp.MSPVersion) {
	fake.getVersionMutex.Lock()
	defer fake.getVersionMutex.Unlock()
	fake.GetVersionStub = nil
	if fake.getVersionReturnsOnCall == nil {
		fake.getVersionReturnsOnCall = make(map[int]struct {
			result1 msp.MSPVersion
		})
	}
	fake.getVersionReturnsOnCall[i] = struct {
		result1 msp.MSPVersion
	}{result1}
}

func (fake *MSP) IsWellFormed(arg1 *mspa.SerializedIdentity) error {
	fake.isWellFormedMutex.Lock()
	ret, specificReturn := fake.isWellFormedReturnsOnCall[len(fake.isWellFormedArgsForCall)]
	fake.isWellFormedArgsForCall = append(fake.isWellFormedArgsForCall, struct {
		arg1 *mspa.SerializedIdentity
	}{arg1})
	stub := fake.IsWellFormedStub
	fakeReturns := fake.isWellFormedReturns
	fake.recordInvocation("IsWellFormed", []interface{}{arg1})
	fake.isWellFormedMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *MSP) IsWellFormedCallCount() int {
	fake.isWellFormedMutex.RLock()
	defer fake.isWellFormedMutex.RUnlock()
	return len(fake.isWellFormedArgsForCall)
}

func (fake *MSP) IsWellFormedCalls(stub func(*mspa.SerializedIdentity) error) {
	fake.isWellFormedMutex.Lock()
	defer fake.isWellFormedMutex.Unlock()
	fake.IsWellFormedStub = stub
}

func (fake *MSP) IsWellFormedArgsForCall(i int) *mspa.SerializedIdentity {
	fake.isWellFormedMutex.RLock()
	defer fake.isWellFormedMutex.RUnlock()
	argsForCall := fake.isWellFormedArgsForCall[i]
	return argsForCall.arg1
}

func (fake *MSP) IsWellFormedReturns(result1 error) {
	fake.isWellFormedMutex.Lock()
	defer fake.isWellFormedMutex.Unlock()
	fake.IsWellFormedStub = nil
	fake.isWellFormedReturns = struct {
		result1 error
	}{result1}
}

func (fake *MSP) IsWellFormedReturnsOnCall(i int, result1 error) {
	fake.isWellFormedMutex.Lock()
	defer fake.isWellFormedMutex.Unlock()
	fake.IsWellFormedStub = nil
	if fake.isWellFormedReturnsOnCall == nil {
		fake.isWellFormedReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.isWellFormedReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *MSP) SatisfiesPrincipal(arg1 msp.Identity, arg2 *mspa.MSPPrincipal) error {
	fake.satisfiesPrincipalMutex.Lock()
	ret, specificReturn := fake.satisfiesPrincipalReturnsOnCall[len(fake.satisfiesPrincipalArgsForCall)]
	fake.satisfiesPrincipalArgsForCall = append(fake.satisfiesPrincipalArgsForCall, struct {
		arg1 msp.Identity
		arg2 *mspa.MSPPrincipal
	}{arg1, arg2})
	stub := fake.SatisfiesPrincipalStub
	fakeReturns := fake.satisfiesPrincipalReturns
	fake.recordInvocation("SatisfiesPrincipal", []interface{}{arg1, arg2})
	fake.satisfiesPrincipalMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *MSP) SatisfiesPrincipalCallCount() int {
	fake.satisfiesPrincipalMutex.RLock()
	defer fake.satisfiesPrincipalMutex.RUnlock()
	return len(fake.satisfiesPrincipalArgsForCall)
}

func (fake *MSP) SatisfiesPrincipalCalls(stub func(msp.Identity, *mspa.MSPPrincipal) error) {
	fake.satisfiesPrincipalMutex.Lock()
	defer fake.satisfiesPrincipalMutex.Unlock()
	fake.SatisfiesPrincipalStub = stub
}

func (fake *MSP) SatisfiesPrincipalArgsForCall(i int) (msp.Identity, *mspa.MSPPrincipal) {
	fake.satisfiesPrincipalMutex.RLock()
	defer fake.satisfiesPrincipalMutex.RUnlock()
	argsForCall := fake.satisfiesPrincipalArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *MSP) SatisfiesPrincipalReturns(result1 error) {
	fake.satisfiesPrincipalMutex.Lock()
	defer fake.satisfiesPrincipalMutex.Unlock()
	fake.SatisfiesPrincipalStub = nil
	fake.satisfiesPrincipalReturns = struct {
		result1 error
	}{result1}
}

func (fake *MSP) SatisfiesPrincipalReturnsOnCall(i int, result1 error) {
	fake.satisfiesPrincipalMutex.Lock()
	defer fake.satisfiesPrincipalMutex.Unlock()
	fake.SatisfiesPrincipalStub = nil
	if fake.satisfiesPrincipalReturnsOnCall == nil {
		fake.satisfiesPrincipalReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.satisfiesPrincipalReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *MSP) Setup(arg1 *mspa.MSPConfig) error {
	fake.setupMutex.Lock()
	ret, specificReturn := fake.setupReturnsOnCall[len(fake.setupArgsForCall)]
	fake.setupArgsForCall = append(fake.setupArgsForCall, struct {
		arg1 *mspa.MSPConfig
	}{arg1})
	stub := fake.SetupStub
	fakeReturns := fake.setupReturns
	fake.recordInvocation("Setup", []interface{}{arg1})
	fake.setupMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *MSP) SetupCallCount() int {
	fake.setupMutex.RLock()
	defer fake.setupMutex.RUnlock()
	return len(fake.setupArgsForCall)
}

func (fake *MSP) SetupCalls(stub func(*mspa.MSPConfig) error) {
	fake.setupMutex.Lock()
	defer fake.setupMutex.Unlock()
	fake.SetupStub = stub
}

func (fake *MSP) SetupArgsForCall(i int) *mspa.MSPConfig {
	fake.setupMutex.RLock()
	defer fake.setupMutex.RUnlock()
	argsForCall := fake.setupArgsForCall[i]
	return argsForCall.arg1
}

func (fake *MSP) SetupReturns(result1 error) {
	fake.setupMutex.Lock()
	defer fake.setupMutex.Unlock()
	fake.SetupStub = nil
	fake.setupReturns = struct {
		result1 error
	}{result1}
}

func (fake *MSP) SetupReturnsOnCall(i int, result1 error) {
	fake.setupMutex.Lock()
	defer fake.setupMutex.Unlock()
	fake.SetupStub = nil
	if fake.setupReturnsOnCall == nil {
		fake.setupReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.setupReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *MSP) Validate(arg1 msp.Identity) error {
	fake.validateMutex.Lock()
	ret, specificReturn := fake.validateReturnsOnCall[len(fake.validateArgsForCall)]
	fake.validateArgsForCall = append(fake.validateArgsForCall, struct {
		arg1 msp.Identity
	}{arg1})
	stub := fake.ValidateStub
	fakeReturns := fake.validateReturns
	fake.recordInvocation("Validate", []interface{}{arg1})
	fake.validateMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *MSP) ValidateCallCount() int {
	fake.validateMutex.RLock()
	defer fake.validateMutex.RUnlock()
	return len(fake.validateArgsForCall)
}

func (fake *MSP) ValidateCalls(stub func(msp.Identity) error) {
	fake.validateMutex.Lock()
	defer fake.validateMutex.Unlock()
	fake.ValidateStub = stub
}

func (fake *MSP) ValidateArgsForCall(i int) msp.Identity {
	fake.validateMutex.RLock()
	defer fake.validateMutex.RUnlock()
	argsForCall := fake.validateArgsForCall[i]
	return argsForCall.arg1
}

func (fake *MSP) ValidateReturns(result1 error) {
	fake.validateMutex.Lock()
	defer fake.validateMutex.Unlock()
	fake.ValidateStub = nil
	fake.validateReturns = struct {
		result1 error
	}{result1}
}

func (fake *MSP) ValidateReturnsOnCall(i int, result1 error) {
	fake.validateMutex.Lock()
	defer fake.validateMutex.Unlock()
	fake.ValidateStub = nil
	if fake.validateReturnsOnCall == nil {
		fake.validateReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.validateReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *MSP) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.deserializeIdentityMutex.RLock()
	defer fake.deserializeIdentityMutex.RUnlock()
	fake.getDefaultSigningIdentityMutex.RLock()
	defer fake.getDefaultSigningIdentityMutex.RUnlock()
	fake.getIdentifierMutex.RLock()
	defer fake.getIdentifierMutex.RUnlock()
	fake.getSigningIdentityMutex.RLock()
	defer fake.getSigningIdentityMutex.RUnlock()
	fake.getTLSIntermediateCertsMutex.RLock()
	defer fake.getTLSIntermediateCertsMutex.RUnlock()
	fake.getTLSRootCertsMutex.RLock()
	defer fake.getTLSRootCertsMutex.RUnlock()
	fake.getTypeMutex.RLock()
	defer fake.getTypeMutex.RUnlock()
	fake.getVersionMutex.RLock()
	defer fake.getVersionMutex.RUnlock()
	fake.isWellFormedMutex.RLock()
	defer fake.isWellFormedMutex.RUnlock()
	fake.satisfiesPrincipalMutex.RLock()
	defer fake.satisfiesPrincipalMutex.RUnlock()
	fake.setupMutex.RLock()
	defer fake.setupMutex.RUnlock()
	fake.validateMutex.RLock()
	defer fake.validateMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *MSP) recordInvocation(key string, args []interface{}) {
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
