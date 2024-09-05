// Code generated by counterfeiter. DO NOT EDIT.
package mocks

import (
	"context"
	"sync"

	"github.com/bborbe/k8s"
	v1 "k8s.io/api/core/v1"
)

type K8sServiceBuilder struct {
	BuildStub        func(context.Context) (*v1.Service, error)
	buildMutex       sync.RWMutex
	buildArgsForCall []struct {
		arg1 context.Context
	}
	buildReturns struct {
		result1 *v1.Service
		result2 error
	}
	buildReturnsOnCall map[int]struct {
		result1 *v1.Service
		result2 error
	}
	SetNameStub        func(k8s.Name) k8s.ServiceBuilder
	setNameMutex       sync.RWMutex
	setNameArgsForCall []struct {
		arg1 k8s.Name
	}
	setNameReturns struct {
		result1 k8s.ServiceBuilder
	}
	setNameReturnsOnCall map[int]struct {
		result1 k8s.ServiceBuilder
	}
	SetObjectMetaBuilderStub        func(k8s.ObjectMetaBuilder) k8s.ServiceBuilder
	setObjectMetaBuilderMutex       sync.RWMutex
	setObjectMetaBuilderArgsForCall []struct {
		arg1 k8s.ObjectMetaBuilder
	}
	setObjectMetaBuilderReturns struct {
		result1 k8s.ServiceBuilder
	}
	setObjectMetaBuilderReturnsOnCall map[int]struct {
		result1 k8s.ServiceBuilder
	}
	SetServicePortNameStub        func(string) k8s.ServiceBuilder
	setServicePortNameMutex       sync.RWMutex
	setServicePortNameArgsForCall []struct {
		arg1 string
	}
	setServicePortNameReturns struct {
		result1 k8s.ServiceBuilder
	}
	setServicePortNameReturnsOnCall map[int]struct {
		result1 k8s.ServiceBuilder
	}
	SetServicePortNumberStub        func(int32) k8s.ServiceBuilder
	setServicePortNumberMutex       sync.RWMutex
	setServicePortNumberArgsForCall []struct {
		arg1 int32
	}
	setServicePortNumberReturns struct {
		result1 k8s.ServiceBuilder
	}
	setServicePortNumberReturnsOnCall map[int]struct {
		result1 k8s.ServiceBuilder
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *K8sServiceBuilder) Build(arg1 context.Context) (*v1.Service, error) {
	fake.buildMutex.Lock()
	ret, specificReturn := fake.buildReturnsOnCall[len(fake.buildArgsForCall)]
	fake.buildArgsForCall = append(fake.buildArgsForCall, struct {
		arg1 context.Context
	}{arg1})
	stub := fake.BuildStub
	fakeReturns := fake.buildReturns
	fake.recordInvocation("Build", []interface{}{arg1})
	fake.buildMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *K8sServiceBuilder) BuildCallCount() int {
	fake.buildMutex.RLock()
	defer fake.buildMutex.RUnlock()
	return len(fake.buildArgsForCall)
}

func (fake *K8sServiceBuilder) BuildCalls(stub func(context.Context) (*v1.Service, error)) {
	fake.buildMutex.Lock()
	defer fake.buildMutex.Unlock()
	fake.BuildStub = stub
}

func (fake *K8sServiceBuilder) BuildArgsForCall(i int) context.Context {
	fake.buildMutex.RLock()
	defer fake.buildMutex.RUnlock()
	argsForCall := fake.buildArgsForCall[i]
	return argsForCall.arg1
}

func (fake *K8sServiceBuilder) BuildReturns(result1 *v1.Service, result2 error) {
	fake.buildMutex.Lock()
	defer fake.buildMutex.Unlock()
	fake.BuildStub = nil
	fake.buildReturns = struct {
		result1 *v1.Service
		result2 error
	}{result1, result2}
}

func (fake *K8sServiceBuilder) BuildReturnsOnCall(i int, result1 *v1.Service, result2 error) {
	fake.buildMutex.Lock()
	defer fake.buildMutex.Unlock()
	fake.BuildStub = nil
	if fake.buildReturnsOnCall == nil {
		fake.buildReturnsOnCall = make(map[int]struct {
			result1 *v1.Service
			result2 error
		})
	}
	fake.buildReturnsOnCall[i] = struct {
		result1 *v1.Service
		result2 error
	}{result1, result2}
}

func (fake *K8sServiceBuilder) SetName(arg1 k8s.Name) k8s.ServiceBuilder {
	fake.setNameMutex.Lock()
	ret, specificReturn := fake.setNameReturnsOnCall[len(fake.setNameArgsForCall)]
	fake.setNameArgsForCall = append(fake.setNameArgsForCall, struct {
		arg1 k8s.Name
	}{arg1})
	stub := fake.SetNameStub
	fakeReturns := fake.setNameReturns
	fake.recordInvocation("SetName", []interface{}{arg1})
	fake.setNameMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *K8sServiceBuilder) SetNameCallCount() int {
	fake.setNameMutex.RLock()
	defer fake.setNameMutex.RUnlock()
	return len(fake.setNameArgsForCall)
}

func (fake *K8sServiceBuilder) SetNameCalls(stub func(k8s.Name) k8s.ServiceBuilder) {
	fake.setNameMutex.Lock()
	defer fake.setNameMutex.Unlock()
	fake.SetNameStub = stub
}

func (fake *K8sServiceBuilder) SetNameArgsForCall(i int) k8s.Name {
	fake.setNameMutex.RLock()
	defer fake.setNameMutex.RUnlock()
	argsForCall := fake.setNameArgsForCall[i]
	return argsForCall.arg1
}

func (fake *K8sServiceBuilder) SetNameReturns(result1 k8s.ServiceBuilder) {
	fake.setNameMutex.Lock()
	defer fake.setNameMutex.Unlock()
	fake.SetNameStub = nil
	fake.setNameReturns = struct {
		result1 k8s.ServiceBuilder
	}{result1}
}

func (fake *K8sServiceBuilder) SetNameReturnsOnCall(i int, result1 k8s.ServiceBuilder) {
	fake.setNameMutex.Lock()
	defer fake.setNameMutex.Unlock()
	fake.SetNameStub = nil
	if fake.setNameReturnsOnCall == nil {
		fake.setNameReturnsOnCall = make(map[int]struct {
			result1 k8s.ServiceBuilder
		})
	}
	fake.setNameReturnsOnCall[i] = struct {
		result1 k8s.ServiceBuilder
	}{result1}
}

func (fake *K8sServiceBuilder) SetObjectMetaBuilder(arg1 k8s.ObjectMetaBuilder) k8s.ServiceBuilder {
	fake.setObjectMetaBuilderMutex.Lock()
	ret, specificReturn := fake.setObjectMetaBuilderReturnsOnCall[len(fake.setObjectMetaBuilderArgsForCall)]
	fake.setObjectMetaBuilderArgsForCall = append(fake.setObjectMetaBuilderArgsForCall, struct {
		arg1 k8s.ObjectMetaBuilder
	}{arg1})
	stub := fake.SetObjectMetaBuilderStub
	fakeReturns := fake.setObjectMetaBuilderReturns
	fake.recordInvocation("SetObjectMetaBuilder", []interface{}{arg1})
	fake.setObjectMetaBuilderMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *K8sServiceBuilder) SetObjectMetaBuilderCallCount() int {
	fake.setObjectMetaBuilderMutex.RLock()
	defer fake.setObjectMetaBuilderMutex.RUnlock()
	return len(fake.setObjectMetaBuilderArgsForCall)
}

func (fake *K8sServiceBuilder) SetObjectMetaBuilderCalls(stub func(k8s.ObjectMetaBuilder) k8s.ServiceBuilder) {
	fake.setObjectMetaBuilderMutex.Lock()
	defer fake.setObjectMetaBuilderMutex.Unlock()
	fake.SetObjectMetaBuilderStub = stub
}

func (fake *K8sServiceBuilder) SetObjectMetaBuilderArgsForCall(i int) k8s.ObjectMetaBuilder {
	fake.setObjectMetaBuilderMutex.RLock()
	defer fake.setObjectMetaBuilderMutex.RUnlock()
	argsForCall := fake.setObjectMetaBuilderArgsForCall[i]
	return argsForCall.arg1
}

func (fake *K8sServiceBuilder) SetObjectMetaBuilderReturns(result1 k8s.ServiceBuilder) {
	fake.setObjectMetaBuilderMutex.Lock()
	defer fake.setObjectMetaBuilderMutex.Unlock()
	fake.SetObjectMetaBuilderStub = nil
	fake.setObjectMetaBuilderReturns = struct {
		result1 k8s.ServiceBuilder
	}{result1}
}

func (fake *K8sServiceBuilder) SetObjectMetaBuilderReturnsOnCall(i int, result1 k8s.ServiceBuilder) {
	fake.setObjectMetaBuilderMutex.Lock()
	defer fake.setObjectMetaBuilderMutex.Unlock()
	fake.SetObjectMetaBuilderStub = nil
	if fake.setObjectMetaBuilderReturnsOnCall == nil {
		fake.setObjectMetaBuilderReturnsOnCall = make(map[int]struct {
			result1 k8s.ServiceBuilder
		})
	}
	fake.setObjectMetaBuilderReturnsOnCall[i] = struct {
		result1 k8s.ServiceBuilder
	}{result1}
}

func (fake *K8sServiceBuilder) SetServicePortName(arg1 string) k8s.ServiceBuilder {
	fake.setServicePortNameMutex.Lock()
	ret, specificReturn := fake.setServicePortNameReturnsOnCall[len(fake.setServicePortNameArgsForCall)]
	fake.setServicePortNameArgsForCall = append(fake.setServicePortNameArgsForCall, struct {
		arg1 string
	}{arg1})
	stub := fake.SetServicePortNameStub
	fakeReturns := fake.setServicePortNameReturns
	fake.recordInvocation("SetServicePortName", []interface{}{arg1})
	fake.setServicePortNameMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *K8sServiceBuilder) SetServicePortNameCallCount() int {
	fake.setServicePortNameMutex.RLock()
	defer fake.setServicePortNameMutex.RUnlock()
	return len(fake.setServicePortNameArgsForCall)
}

func (fake *K8sServiceBuilder) SetServicePortNameCalls(stub func(string) k8s.ServiceBuilder) {
	fake.setServicePortNameMutex.Lock()
	defer fake.setServicePortNameMutex.Unlock()
	fake.SetServicePortNameStub = stub
}

func (fake *K8sServiceBuilder) SetServicePortNameArgsForCall(i int) string {
	fake.setServicePortNameMutex.RLock()
	defer fake.setServicePortNameMutex.RUnlock()
	argsForCall := fake.setServicePortNameArgsForCall[i]
	return argsForCall.arg1
}

func (fake *K8sServiceBuilder) SetServicePortNameReturns(result1 k8s.ServiceBuilder) {
	fake.setServicePortNameMutex.Lock()
	defer fake.setServicePortNameMutex.Unlock()
	fake.SetServicePortNameStub = nil
	fake.setServicePortNameReturns = struct {
		result1 k8s.ServiceBuilder
	}{result1}
}

func (fake *K8sServiceBuilder) SetServicePortNameReturnsOnCall(i int, result1 k8s.ServiceBuilder) {
	fake.setServicePortNameMutex.Lock()
	defer fake.setServicePortNameMutex.Unlock()
	fake.SetServicePortNameStub = nil
	if fake.setServicePortNameReturnsOnCall == nil {
		fake.setServicePortNameReturnsOnCall = make(map[int]struct {
			result1 k8s.ServiceBuilder
		})
	}
	fake.setServicePortNameReturnsOnCall[i] = struct {
		result1 k8s.ServiceBuilder
	}{result1}
}

func (fake *K8sServiceBuilder) SetServicePortNumber(arg1 int32) k8s.ServiceBuilder {
	fake.setServicePortNumberMutex.Lock()
	ret, specificReturn := fake.setServicePortNumberReturnsOnCall[len(fake.setServicePortNumberArgsForCall)]
	fake.setServicePortNumberArgsForCall = append(fake.setServicePortNumberArgsForCall, struct {
		arg1 int32
	}{arg1})
	stub := fake.SetServicePortNumberStub
	fakeReturns := fake.setServicePortNumberReturns
	fake.recordInvocation("SetServicePortNumber", []interface{}{arg1})
	fake.setServicePortNumberMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *K8sServiceBuilder) SetServicePortNumberCallCount() int {
	fake.setServicePortNumberMutex.RLock()
	defer fake.setServicePortNumberMutex.RUnlock()
	return len(fake.setServicePortNumberArgsForCall)
}

func (fake *K8sServiceBuilder) SetServicePortNumberCalls(stub func(int32) k8s.ServiceBuilder) {
	fake.setServicePortNumberMutex.Lock()
	defer fake.setServicePortNumberMutex.Unlock()
	fake.SetServicePortNumberStub = stub
}

func (fake *K8sServiceBuilder) SetServicePortNumberArgsForCall(i int) int32 {
	fake.setServicePortNumberMutex.RLock()
	defer fake.setServicePortNumberMutex.RUnlock()
	argsForCall := fake.setServicePortNumberArgsForCall[i]
	return argsForCall.arg1
}

func (fake *K8sServiceBuilder) SetServicePortNumberReturns(result1 k8s.ServiceBuilder) {
	fake.setServicePortNumberMutex.Lock()
	defer fake.setServicePortNumberMutex.Unlock()
	fake.SetServicePortNumberStub = nil
	fake.setServicePortNumberReturns = struct {
		result1 k8s.ServiceBuilder
	}{result1}
}

func (fake *K8sServiceBuilder) SetServicePortNumberReturnsOnCall(i int, result1 k8s.ServiceBuilder) {
	fake.setServicePortNumberMutex.Lock()
	defer fake.setServicePortNumberMutex.Unlock()
	fake.SetServicePortNumberStub = nil
	if fake.setServicePortNumberReturnsOnCall == nil {
		fake.setServicePortNumberReturnsOnCall = make(map[int]struct {
			result1 k8s.ServiceBuilder
		})
	}
	fake.setServicePortNumberReturnsOnCall[i] = struct {
		result1 k8s.ServiceBuilder
	}{result1}
}

func (fake *K8sServiceBuilder) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.buildMutex.RLock()
	defer fake.buildMutex.RUnlock()
	fake.setNameMutex.RLock()
	defer fake.setNameMutex.RUnlock()
	fake.setObjectMetaBuilderMutex.RLock()
	defer fake.setObjectMetaBuilderMutex.RUnlock()
	fake.setServicePortNameMutex.RLock()
	defer fake.setServicePortNameMutex.RUnlock()
	fake.setServicePortNumberMutex.RLock()
	defer fake.setServicePortNumberMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *K8sServiceBuilder) recordInvocation(key string, args []interface{}) {
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

var _ k8s.ServiceBuilder = new(K8sServiceBuilder)
