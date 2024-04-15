// Code generated by counterfeiter. DO NOT EDIT.
package mocks

import (
	"context"
	"sync"

	"github.com/bborbe/k8s"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type K8sObjectMetaBuilder struct {
	AddAnnotationStub        func(string, string) k8s.ObjectMetaBuilder
	addAnnotationMutex       sync.RWMutex
	addAnnotationArgsForCall []struct {
		arg1 string
		arg2 string
	}
	addAnnotationReturns struct {
		result1 k8s.ObjectMetaBuilder
	}
	addAnnotationReturnsOnCall map[int]struct {
		result1 k8s.ObjectMetaBuilder
	}
	AddLabelStub        func(string, string) k8s.ObjectMetaBuilder
	addLabelMutex       sync.RWMutex
	addLabelArgsForCall []struct {
		arg1 string
		arg2 string
	}
	addLabelReturns struct {
		result1 k8s.ObjectMetaBuilder
	}
	addLabelReturnsOnCall map[int]struct {
		result1 k8s.ObjectMetaBuilder
	}
	BuildStub        func(context.Context) (*v1.ObjectMeta, error)
	buildMutex       sync.RWMutex
	buildArgsForCall []struct {
		arg1 context.Context
	}
	buildReturns struct {
		result1 *v1.ObjectMeta
		result2 error
	}
	buildReturnsOnCall map[int]struct {
		result1 *v1.ObjectMeta
		result2 error
	}
	SetComponentStub        func(string) k8s.ObjectMetaBuilder
	setComponentMutex       sync.RWMutex
	setComponentArgsForCall []struct {
		arg1 string
	}
	setComponentReturns struct {
		result1 k8s.ObjectMetaBuilder
	}
	setComponentReturnsOnCall map[int]struct {
		result1 k8s.ObjectMetaBuilder
	}
	SetFinalizersStub        func([]string) k8s.ObjectMetaBuilder
	setFinalizersMutex       sync.RWMutex
	setFinalizersArgsForCall []struct {
		arg1 []string
	}
	setFinalizersReturns struct {
		result1 k8s.ObjectMetaBuilder
	}
	setFinalizersReturnsOnCall map[int]struct {
		result1 k8s.ObjectMetaBuilder
	}
	SetGenerateNameStub        func(string) k8s.ObjectMetaBuilder
	setGenerateNameMutex       sync.RWMutex
	setGenerateNameArgsForCall []struct {
		arg1 string
	}
	setGenerateNameReturns struct {
		result1 k8s.ObjectMetaBuilder
	}
	setGenerateNameReturnsOnCall map[int]struct {
		result1 k8s.ObjectMetaBuilder
	}
	SetNameStub        func(string) k8s.ObjectMetaBuilder
	setNameMutex       sync.RWMutex
	setNameArgsForCall []struct {
		arg1 string
	}
	setNameReturns struct {
		result1 k8s.ObjectMetaBuilder
	}
	setNameReturnsOnCall map[int]struct {
		result1 k8s.ObjectMetaBuilder
	}
	SetNamespaceStub        func(k8s.Namespace) k8s.ObjectMetaBuilder
	setNamespaceMutex       sync.RWMutex
	setNamespaceArgsForCall []struct {
		arg1 k8s.Namespace
	}
	setNamespaceReturns struct {
		result1 k8s.ObjectMetaBuilder
	}
	setNamespaceReturnsOnCall map[int]struct {
		result1 k8s.ObjectMetaBuilder
	}
	ValidateStub        func(context.Context) error
	validateMutex       sync.RWMutex
	validateArgsForCall []struct {
		arg1 context.Context
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

func (fake *K8sObjectMetaBuilder) AddAnnotation(arg1 string, arg2 string) k8s.ObjectMetaBuilder {
	fake.addAnnotationMutex.Lock()
	ret, specificReturn := fake.addAnnotationReturnsOnCall[len(fake.addAnnotationArgsForCall)]
	fake.addAnnotationArgsForCall = append(fake.addAnnotationArgsForCall, struct {
		arg1 string
		arg2 string
	}{arg1, arg2})
	stub := fake.AddAnnotationStub
	fakeReturns := fake.addAnnotationReturns
	fake.recordInvocation("AddAnnotation", []interface{}{arg1, arg2})
	fake.addAnnotationMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *K8sObjectMetaBuilder) AddAnnotationCallCount() int {
	fake.addAnnotationMutex.RLock()
	defer fake.addAnnotationMutex.RUnlock()
	return len(fake.addAnnotationArgsForCall)
}

func (fake *K8sObjectMetaBuilder) AddAnnotationCalls(stub func(string, string) k8s.ObjectMetaBuilder) {
	fake.addAnnotationMutex.Lock()
	defer fake.addAnnotationMutex.Unlock()
	fake.AddAnnotationStub = stub
}

func (fake *K8sObjectMetaBuilder) AddAnnotationArgsForCall(i int) (string, string) {
	fake.addAnnotationMutex.RLock()
	defer fake.addAnnotationMutex.RUnlock()
	argsForCall := fake.addAnnotationArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *K8sObjectMetaBuilder) AddAnnotationReturns(result1 k8s.ObjectMetaBuilder) {
	fake.addAnnotationMutex.Lock()
	defer fake.addAnnotationMutex.Unlock()
	fake.AddAnnotationStub = nil
	fake.addAnnotationReturns = struct {
		result1 k8s.ObjectMetaBuilder
	}{result1}
}

func (fake *K8sObjectMetaBuilder) AddAnnotationReturnsOnCall(i int, result1 k8s.ObjectMetaBuilder) {
	fake.addAnnotationMutex.Lock()
	defer fake.addAnnotationMutex.Unlock()
	fake.AddAnnotationStub = nil
	if fake.addAnnotationReturnsOnCall == nil {
		fake.addAnnotationReturnsOnCall = make(map[int]struct {
			result1 k8s.ObjectMetaBuilder
		})
	}
	fake.addAnnotationReturnsOnCall[i] = struct {
		result1 k8s.ObjectMetaBuilder
	}{result1}
}

func (fake *K8sObjectMetaBuilder) AddLabel(arg1 string, arg2 string) k8s.ObjectMetaBuilder {
	fake.addLabelMutex.Lock()
	ret, specificReturn := fake.addLabelReturnsOnCall[len(fake.addLabelArgsForCall)]
	fake.addLabelArgsForCall = append(fake.addLabelArgsForCall, struct {
		arg1 string
		arg2 string
	}{arg1, arg2})
	stub := fake.AddLabelStub
	fakeReturns := fake.addLabelReturns
	fake.recordInvocation("AddLabel", []interface{}{arg1, arg2})
	fake.addLabelMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *K8sObjectMetaBuilder) AddLabelCallCount() int {
	fake.addLabelMutex.RLock()
	defer fake.addLabelMutex.RUnlock()
	return len(fake.addLabelArgsForCall)
}

func (fake *K8sObjectMetaBuilder) AddLabelCalls(stub func(string, string) k8s.ObjectMetaBuilder) {
	fake.addLabelMutex.Lock()
	defer fake.addLabelMutex.Unlock()
	fake.AddLabelStub = stub
}

func (fake *K8sObjectMetaBuilder) AddLabelArgsForCall(i int) (string, string) {
	fake.addLabelMutex.RLock()
	defer fake.addLabelMutex.RUnlock()
	argsForCall := fake.addLabelArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *K8sObjectMetaBuilder) AddLabelReturns(result1 k8s.ObjectMetaBuilder) {
	fake.addLabelMutex.Lock()
	defer fake.addLabelMutex.Unlock()
	fake.AddLabelStub = nil
	fake.addLabelReturns = struct {
		result1 k8s.ObjectMetaBuilder
	}{result1}
}

func (fake *K8sObjectMetaBuilder) AddLabelReturnsOnCall(i int, result1 k8s.ObjectMetaBuilder) {
	fake.addLabelMutex.Lock()
	defer fake.addLabelMutex.Unlock()
	fake.AddLabelStub = nil
	if fake.addLabelReturnsOnCall == nil {
		fake.addLabelReturnsOnCall = make(map[int]struct {
			result1 k8s.ObjectMetaBuilder
		})
	}
	fake.addLabelReturnsOnCall[i] = struct {
		result1 k8s.ObjectMetaBuilder
	}{result1}
}

func (fake *K8sObjectMetaBuilder) Build(arg1 context.Context) (*v1.ObjectMeta, error) {
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

func (fake *K8sObjectMetaBuilder) BuildCallCount() int {
	fake.buildMutex.RLock()
	defer fake.buildMutex.RUnlock()
	return len(fake.buildArgsForCall)
}

func (fake *K8sObjectMetaBuilder) BuildCalls(stub func(context.Context) (*v1.ObjectMeta, error)) {
	fake.buildMutex.Lock()
	defer fake.buildMutex.Unlock()
	fake.BuildStub = stub
}

func (fake *K8sObjectMetaBuilder) BuildArgsForCall(i int) context.Context {
	fake.buildMutex.RLock()
	defer fake.buildMutex.RUnlock()
	argsForCall := fake.buildArgsForCall[i]
	return argsForCall.arg1
}

func (fake *K8sObjectMetaBuilder) BuildReturns(result1 *v1.ObjectMeta, result2 error) {
	fake.buildMutex.Lock()
	defer fake.buildMutex.Unlock()
	fake.BuildStub = nil
	fake.buildReturns = struct {
		result1 *v1.ObjectMeta
		result2 error
	}{result1, result2}
}

func (fake *K8sObjectMetaBuilder) BuildReturnsOnCall(i int, result1 *v1.ObjectMeta, result2 error) {
	fake.buildMutex.Lock()
	defer fake.buildMutex.Unlock()
	fake.BuildStub = nil
	if fake.buildReturnsOnCall == nil {
		fake.buildReturnsOnCall = make(map[int]struct {
			result1 *v1.ObjectMeta
			result2 error
		})
	}
	fake.buildReturnsOnCall[i] = struct {
		result1 *v1.ObjectMeta
		result2 error
	}{result1, result2}
}

func (fake *K8sObjectMetaBuilder) SetComponent(arg1 string) k8s.ObjectMetaBuilder {
	fake.setComponentMutex.Lock()
	ret, specificReturn := fake.setComponentReturnsOnCall[len(fake.setComponentArgsForCall)]
	fake.setComponentArgsForCall = append(fake.setComponentArgsForCall, struct {
		arg1 string
	}{arg1})
	stub := fake.SetComponentStub
	fakeReturns := fake.setComponentReturns
	fake.recordInvocation("SetComponent", []interface{}{arg1})
	fake.setComponentMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *K8sObjectMetaBuilder) SetComponentCallCount() int {
	fake.setComponentMutex.RLock()
	defer fake.setComponentMutex.RUnlock()
	return len(fake.setComponentArgsForCall)
}

func (fake *K8sObjectMetaBuilder) SetComponentCalls(stub func(string) k8s.ObjectMetaBuilder) {
	fake.setComponentMutex.Lock()
	defer fake.setComponentMutex.Unlock()
	fake.SetComponentStub = stub
}

func (fake *K8sObjectMetaBuilder) SetComponentArgsForCall(i int) string {
	fake.setComponentMutex.RLock()
	defer fake.setComponentMutex.RUnlock()
	argsForCall := fake.setComponentArgsForCall[i]
	return argsForCall.arg1
}

func (fake *K8sObjectMetaBuilder) SetComponentReturns(result1 k8s.ObjectMetaBuilder) {
	fake.setComponentMutex.Lock()
	defer fake.setComponentMutex.Unlock()
	fake.SetComponentStub = nil
	fake.setComponentReturns = struct {
		result1 k8s.ObjectMetaBuilder
	}{result1}
}

func (fake *K8sObjectMetaBuilder) SetComponentReturnsOnCall(i int, result1 k8s.ObjectMetaBuilder) {
	fake.setComponentMutex.Lock()
	defer fake.setComponentMutex.Unlock()
	fake.SetComponentStub = nil
	if fake.setComponentReturnsOnCall == nil {
		fake.setComponentReturnsOnCall = make(map[int]struct {
			result1 k8s.ObjectMetaBuilder
		})
	}
	fake.setComponentReturnsOnCall[i] = struct {
		result1 k8s.ObjectMetaBuilder
	}{result1}
}

func (fake *K8sObjectMetaBuilder) SetFinalizers(arg1 []string) k8s.ObjectMetaBuilder {
	var arg1Copy []string
	if arg1 != nil {
		arg1Copy = make([]string, len(arg1))
		copy(arg1Copy, arg1)
	}
	fake.setFinalizersMutex.Lock()
	ret, specificReturn := fake.setFinalizersReturnsOnCall[len(fake.setFinalizersArgsForCall)]
	fake.setFinalizersArgsForCall = append(fake.setFinalizersArgsForCall, struct {
		arg1 []string
	}{arg1Copy})
	stub := fake.SetFinalizersStub
	fakeReturns := fake.setFinalizersReturns
	fake.recordInvocation("SetFinalizers", []interface{}{arg1Copy})
	fake.setFinalizersMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *K8sObjectMetaBuilder) SetFinalizersCallCount() int {
	fake.setFinalizersMutex.RLock()
	defer fake.setFinalizersMutex.RUnlock()
	return len(fake.setFinalizersArgsForCall)
}

func (fake *K8sObjectMetaBuilder) SetFinalizersCalls(stub func([]string) k8s.ObjectMetaBuilder) {
	fake.setFinalizersMutex.Lock()
	defer fake.setFinalizersMutex.Unlock()
	fake.SetFinalizersStub = stub
}

func (fake *K8sObjectMetaBuilder) SetFinalizersArgsForCall(i int) []string {
	fake.setFinalizersMutex.RLock()
	defer fake.setFinalizersMutex.RUnlock()
	argsForCall := fake.setFinalizersArgsForCall[i]
	return argsForCall.arg1
}

func (fake *K8sObjectMetaBuilder) SetFinalizersReturns(result1 k8s.ObjectMetaBuilder) {
	fake.setFinalizersMutex.Lock()
	defer fake.setFinalizersMutex.Unlock()
	fake.SetFinalizersStub = nil
	fake.setFinalizersReturns = struct {
		result1 k8s.ObjectMetaBuilder
	}{result1}
}

func (fake *K8sObjectMetaBuilder) SetFinalizersReturnsOnCall(i int, result1 k8s.ObjectMetaBuilder) {
	fake.setFinalizersMutex.Lock()
	defer fake.setFinalizersMutex.Unlock()
	fake.SetFinalizersStub = nil
	if fake.setFinalizersReturnsOnCall == nil {
		fake.setFinalizersReturnsOnCall = make(map[int]struct {
			result1 k8s.ObjectMetaBuilder
		})
	}
	fake.setFinalizersReturnsOnCall[i] = struct {
		result1 k8s.ObjectMetaBuilder
	}{result1}
}

func (fake *K8sObjectMetaBuilder) SetGenerateName(arg1 string) k8s.ObjectMetaBuilder {
	fake.setGenerateNameMutex.Lock()
	ret, specificReturn := fake.setGenerateNameReturnsOnCall[len(fake.setGenerateNameArgsForCall)]
	fake.setGenerateNameArgsForCall = append(fake.setGenerateNameArgsForCall, struct {
		arg1 string
	}{arg1})
	stub := fake.SetGenerateNameStub
	fakeReturns := fake.setGenerateNameReturns
	fake.recordInvocation("SetGenerateName", []interface{}{arg1})
	fake.setGenerateNameMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *K8sObjectMetaBuilder) SetGenerateNameCallCount() int {
	fake.setGenerateNameMutex.RLock()
	defer fake.setGenerateNameMutex.RUnlock()
	return len(fake.setGenerateNameArgsForCall)
}

func (fake *K8sObjectMetaBuilder) SetGenerateNameCalls(stub func(string) k8s.ObjectMetaBuilder) {
	fake.setGenerateNameMutex.Lock()
	defer fake.setGenerateNameMutex.Unlock()
	fake.SetGenerateNameStub = stub
}

func (fake *K8sObjectMetaBuilder) SetGenerateNameArgsForCall(i int) string {
	fake.setGenerateNameMutex.RLock()
	defer fake.setGenerateNameMutex.RUnlock()
	argsForCall := fake.setGenerateNameArgsForCall[i]
	return argsForCall.arg1
}

func (fake *K8sObjectMetaBuilder) SetGenerateNameReturns(result1 k8s.ObjectMetaBuilder) {
	fake.setGenerateNameMutex.Lock()
	defer fake.setGenerateNameMutex.Unlock()
	fake.SetGenerateNameStub = nil
	fake.setGenerateNameReturns = struct {
		result1 k8s.ObjectMetaBuilder
	}{result1}
}

func (fake *K8sObjectMetaBuilder) SetGenerateNameReturnsOnCall(i int, result1 k8s.ObjectMetaBuilder) {
	fake.setGenerateNameMutex.Lock()
	defer fake.setGenerateNameMutex.Unlock()
	fake.SetGenerateNameStub = nil
	if fake.setGenerateNameReturnsOnCall == nil {
		fake.setGenerateNameReturnsOnCall = make(map[int]struct {
			result1 k8s.ObjectMetaBuilder
		})
	}
	fake.setGenerateNameReturnsOnCall[i] = struct {
		result1 k8s.ObjectMetaBuilder
	}{result1}
}

func (fake *K8sObjectMetaBuilder) SetName(arg1 string) k8s.ObjectMetaBuilder {
	fake.setNameMutex.Lock()
	ret, specificReturn := fake.setNameReturnsOnCall[len(fake.setNameArgsForCall)]
	fake.setNameArgsForCall = append(fake.setNameArgsForCall, struct {
		arg1 string
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

func (fake *K8sObjectMetaBuilder) SetNameCallCount() int {
	fake.setNameMutex.RLock()
	defer fake.setNameMutex.RUnlock()
	return len(fake.setNameArgsForCall)
}

func (fake *K8sObjectMetaBuilder) SetNameCalls(stub func(string) k8s.ObjectMetaBuilder) {
	fake.setNameMutex.Lock()
	defer fake.setNameMutex.Unlock()
	fake.SetNameStub = stub
}

func (fake *K8sObjectMetaBuilder) SetNameArgsForCall(i int) string {
	fake.setNameMutex.RLock()
	defer fake.setNameMutex.RUnlock()
	argsForCall := fake.setNameArgsForCall[i]
	return argsForCall.arg1
}

func (fake *K8sObjectMetaBuilder) SetNameReturns(result1 k8s.ObjectMetaBuilder) {
	fake.setNameMutex.Lock()
	defer fake.setNameMutex.Unlock()
	fake.SetNameStub = nil
	fake.setNameReturns = struct {
		result1 k8s.ObjectMetaBuilder
	}{result1}
}

func (fake *K8sObjectMetaBuilder) SetNameReturnsOnCall(i int, result1 k8s.ObjectMetaBuilder) {
	fake.setNameMutex.Lock()
	defer fake.setNameMutex.Unlock()
	fake.SetNameStub = nil
	if fake.setNameReturnsOnCall == nil {
		fake.setNameReturnsOnCall = make(map[int]struct {
			result1 k8s.ObjectMetaBuilder
		})
	}
	fake.setNameReturnsOnCall[i] = struct {
		result1 k8s.ObjectMetaBuilder
	}{result1}
}

func (fake *K8sObjectMetaBuilder) SetNamespace(arg1 k8s.Namespace) k8s.ObjectMetaBuilder {
	fake.setNamespaceMutex.Lock()
	ret, specificReturn := fake.setNamespaceReturnsOnCall[len(fake.setNamespaceArgsForCall)]
	fake.setNamespaceArgsForCall = append(fake.setNamespaceArgsForCall, struct {
		arg1 k8s.Namespace
	}{arg1})
	stub := fake.SetNamespaceStub
	fakeReturns := fake.setNamespaceReturns
	fake.recordInvocation("SetNamespace", []interface{}{arg1})
	fake.setNamespaceMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *K8sObjectMetaBuilder) SetNamespaceCallCount() int {
	fake.setNamespaceMutex.RLock()
	defer fake.setNamespaceMutex.RUnlock()
	return len(fake.setNamespaceArgsForCall)
}

func (fake *K8sObjectMetaBuilder) SetNamespaceCalls(stub func(k8s.Namespace) k8s.ObjectMetaBuilder) {
	fake.setNamespaceMutex.Lock()
	defer fake.setNamespaceMutex.Unlock()
	fake.SetNamespaceStub = stub
}

func (fake *K8sObjectMetaBuilder) SetNamespaceArgsForCall(i int) k8s.Namespace {
	fake.setNamespaceMutex.RLock()
	defer fake.setNamespaceMutex.RUnlock()
	argsForCall := fake.setNamespaceArgsForCall[i]
	return argsForCall.arg1
}

func (fake *K8sObjectMetaBuilder) SetNamespaceReturns(result1 k8s.ObjectMetaBuilder) {
	fake.setNamespaceMutex.Lock()
	defer fake.setNamespaceMutex.Unlock()
	fake.SetNamespaceStub = nil
	fake.setNamespaceReturns = struct {
		result1 k8s.ObjectMetaBuilder
	}{result1}
}

func (fake *K8sObjectMetaBuilder) SetNamespaceReturnsOnCall(i int, result1 k8s.ObjectMetaBuilder) {
	fake.setNamespaceMutex.Lock()
	defer fake.setNamespaceMutex.Unlock()
	fake.SetNamespaceStub = nil
	if fake.setNamespaceReturnsOnCall == nil {
		fake.setNamespaceReturnsOnCall = make(map[int]struct {
			result1 k8s.ObjectMetaBuilder
		})
	}
	fake.setNamespaceReturnsOnCall[i] = struct {
		result1 k8s.ObjectMetaBuilder
	}{result1}
}

func (fake *K8sObjectMetaBuilder) Validate(arg1 context.Context) error {
	fake.validateMutex.Lock()
	ret, specificReturn := fake.validateReturnsOnCall[len(fake.validateArgsForCall)]
	fake.validateArgsForCall = append(fake.validateArgsForCall, struct {
		arg1 context.Context
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

func (fake *K8sObjectMetaBuilder) ValidateCallCount() int {
	fake.validateMutex.RLock()
	defer fake.validateMutex.RUnlock()
	return len(fake.validateArgsForCall)
}

func (fake *K8sObjectMetaBuilder) ValidateCalls(stub func(context.Context) error) {
	fake.validateMutex.Lock()
	defer fake.validateMutex.Unlock()
	fake.ValidateStub = stub
}

func (fake *K8sObjectMetaBuilder) ValidateArgsForCall(i int) context.Context {
	fake.validateMutex.RLock()
	defer fake.validateMutex.RUnlock()
	argsForCall := fake.validateArgsForCall[i]
	return argsForCall.arg1
}

func (fake *K8sObjectMetaBuilder) ValidateReturns(result1 error) {
	fake.validateMutex.Lock()
	defer fake.validateMutex.Unlock()
	fake.ValidateStub = nil
	fake.validateReturns = struct {
		result1 error
	}{result1}
}

func (fake *K8sObjectMetaBuilder) ValidateReturnsOnCall(i int, result1 error) {
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

func (fake *K8sObjectMetaBuilder) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.addAnnotationMutex.RLock()
	defer fake.addAnnotationMutex.RUnlock()
	fake.addLabelMutex.RLock()
	defer fake.addLabelMutex.RUnlock()
	fake.buildMutex.RLock()
	defer fake.buildMutex.RUnlock()
	fake.setComponentMutex.RLock()
	defer fake.setComponentMutex.RUnlock()
	fake.setFinalizersMutex.RLock()
	defer fake.setFinalizersMutex.RUnlock()
	fake.setGenerateNameMutex.RLock()
	defer fake.setGenerateNameMutex.RUnlock()
	fake.setNameMutex.RLock()
	defer fake.setNameMutex.RUnlock()
	fake.setNamespaceMutex.RLock()
	defer fake.setNamespaceMutex.RUnlock()
	fake.validateMutex.RLock()
	defer fake.validateMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *K8sObjectMetaBuilder) recordInvocation(key string, args []interface{}) {
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

var _ k8s.ObjectMetaBuilder = new(K8sObjectMetaBuilder)
