// Code generated by counterfeiter. DO NOT EDIT.
package mocks

import (
	"context"
	"sync"

	"github.com/bborbe/k8s"
	v1 "k8s.io/api/core/v1"
)

type K8sContainerBuilder struct {
	AddVolumeMountsStub        func(...v1.VolumeMount) k8s.ContainerBuilder
	addVolumeMountsMutex       sync.RWMutex
	addVolumeMountsArgsForCall []struct {
		arg1 []v1.VolumeMount
	}
	addVolumeMountsReturns struct {
		result1 k8s.ContainerBuilder
	}
	addVolumeMountsReturnsOnCall map[int]struct {
		result1 k8s.ContainerBuilder
	}
	BuildStub        func(context.Context) (*v1.Container, error)
	buildMutex       sync.RWMutex
	buildArgsForCall []struct {
		arg1 context.Context
	}
	buildReturns struct {
		result1 *v1.Container
		result2 error
	}
	buildReturnsOnCall map[int]struct {
		result1 *v1.Container
		result2 error
	}
	SetArgsStub        func([]string) k8s.ContainerBuilder
	setArgsMutex       sync.RWMutex
	setArgsArgsForCall []struct {
		arg1 []string
	}
	setArgsReturns struct {
		result1 k8s.ContainerBuilder
	}
	setArgsReturnsOnCall map[int]struct {
		result1 k8s.ContainerBuilder
	}
	SetCommandStub        func([]string) k8s.ContainerBuilder
	setCommandMutex       sync.RWMutex
	setCommandArgsForCall []struct {
		arg1 []string
	}
	setCommandReturns struct {
		result1 k8s.ContainerBuilder
	}
	setCommandReturnsOnCall map[int]struct {
		result1 k8s.ContainerBuilder
	}
	SetCpuLimitStub        func(string) k8s.ContainerBuilder
	setCpuLimitMutex       sync.RWMutex
	setCpuLimitArgsForCall []struct {
		arg1 string
	}
	setCpuLimitReturns struct {
		result1 k8s.ContainerBuilder
	}
	setCpuLimitReturnsOnCall map[int]struct {
		result1 k8s.ContainerBuilder
	}
	SetCpuRequestStub        func(string) k8s.ContainerBuilder
	setCpuRequestMutex       sync.RWMutex
	setCpuRequestArgsForCall []struct {
		arg1 string
	}
	setCpuRequestReturns struct {
		result1 k8s.ContainerBuilder
	}
	setCpuRequestReturnsOnCall map[int]struct {
		result1 k8s.ContainerBuilder
	}
	SetEnvBuilderStub        func(k8s.EnvBuilder) k8s.ContainerBuilder
	setEnvBuilderMutex       sync.RWMutex
	setEnvBuilderArgsForCall []struct {
		arg1 k8s.EnvBuilder
	}
	setEnvBuilderReturns struct {
		result1 k8s.ContainerBuilder
	}
	setEnvBuilderReturnsOnCall map[int]struct {
		result1 k8s.ContainerBuilder
	}
	SetImageStub        func(string) k8s.ContainerBuilder
	setImageMutex       sync.RWMutex
	setImageArgsForCall []struct {
		arg1 string
	}
	setImageReturns struct {
		result1 k8s.ContainerBuilder
	}
	setImageReturnsOnCall map[int]struct {
		result1 k8s.ContainerBuilder
	}
	SetLivenessProbeStub        func(v1.Probe) k8s.ContainerBuilder
	setLivenessProbeMutex       sync.RWMutex
	setLivenessProbeArgsForCall []struct {
		arg1 v1.Probe
	}
	setLivenessProbeReturns struct {
		result1 k8s.ContainerBuilder
	}
	setLivenessProbeReturnsOnCall map[int]struct {
		result1 k8s.ContainerBuilder
	}
	SetMemoryLimitStub        func(string) k8s.ContainerBuilder
	setMemoryLimitMutex       sync.RWMutex
	setMemoryLimitArgsForCall []struct {
		arg1 string
	}
	setMemoryLimitReturns struct {
		result1 k8s.ContainerBuilder
	}
	setMemoryLimitReturnsOnCall map[int]struct {
		result1 k8s.ContainerBuilder
	}
	SetMemoryRequestStub        func(string) k8s.ContainerBuilder
	setMemoryRequestMutex       sync.RWMutex
	setMemoryRequestArgsForCall []struct {
		arg1 string
	}
	setMemoryRequestReturns struct {
		result1 k8s.ContainerBuilder
	}
	setMemoryRequestReturnsOnCall map[int]struct {
		result1 k8s.ContainerBuilder
	}
	SetNameStub        func(string) k8s.ContainerBuilder
	setNameMutex       sync.RWMutex
	setNameArgsForCall []struct {
		arg1 string
	}
	setNameReturns struct {
		result1 k8s.ContainerBuilder
	}
	setNameReturnsOnCall map[int]struct {
		result1 k8s.ContainerBuilder
	}
	SetPortsStub        func([]v1.ContainerPort) k8s.ContainerBuilder
	setPortsMutex       sync.RWMutex
	setPortsArgsForCall []struct {
		arg1 []v1.ContainerPort
	}
	setPortsReturns struct {
		result1 k8s.ContainerBuilder
	}
	setPortsReturnsOnCall map[int]struct {
		result1 k8s.ContainerBuilder
	}
	SetReadinessProbeStub        func(v1.Probe) k8s.ContainerBuilder
	setReadinessProbeMutex       sync.RWMutex
	setReadinessProbeArgsForCall []struct {
		arg1 v1.Probe
	}
	setReadinessProbeReturns struct {
		result1 k8s.ContainerBuilder
	}
	setReadinessProbeReturnsOnCall map[int]struct {
		result1 k8s.ContainerBuilder
	}
	SetVolumeMountsStub        func([]v1.VolumeMount) k8s.ContainerBuilder
	setVolumeMountsMutex       sync.RWMutex
	setVolumeMountsArgsForCall []struct {
		arg1 []v1.VolumeMount
	}
	setVolumeMountsReturns struct {
		result1 k8s.ContainerBuilder
	}
	setVolumeMountsReturnsOnCall map[int]struct {
		result1 k8s.ContainerBuilder
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

func (fake *K8sContainerBuilder) AddVolumeMounts(arg1 ...v1.VolumeMount) k8s.ContainerBuilder {
	fake.addVolumeMountsMutex.Lock()
	ret, specificReturn := fake.addVolumeMountsReturnsOnCall[len(fake.addVolumeMountsArgsForCall)]
	fake.addVolumeMountsArgsForCall = append(fake.addVolumeMountsArgsForCall, struct {
		arg1 []v1.VolumeMount
	}{arg1})
	stub := fake.AddVolumeMountsStub
	fakeReturns := fake.addVolumeMountsReturns
	fake.recordInvocation("AddVolumeMounts", []interface{}{arg1})
	fake.addVolumeMountsMutex.Unlock()
	if stub != nil {
		return stub(arg1...)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *K8sContainerBuilder) AddVolumeMountsCallCount() int {
	fake.addVolumeMountsMutex.RLock()
	defer fake.addVolumeMountsMutex.RUnlock()
	return len(fake.addVolumeMountsArgsForCall)
}

func (fake *K8sContainerBuilder) AddVolumeMountsCalls(stub func(...v1.VolumeMount) k8s.ContainerBuilder) {
	fake.addVolumeMountsMutex.Lock()
	defer fake.addVolumeMountsMutex.Unlock()
	fake.AddVolumeMountsStub = stub
}

func (fake *K8sContainerBuilder) AddVolumeMountsArgsForCall(i int) []v1.VolumeMount {
	fake.addVolumeMountsMutex.RLock()
	defer fake.addVolumeMountsMutex.RUnlock()
	argsForCall := fake.addVolumeMountsArgsForCall[i]
	return argsForCall.arg1
}

func (fake *K8sContainerBuilder) AddVolumeMountsReturns(result1 k8s.ContainerBuilder) {
	fake.addVolumeMountsMutex.Lock()
	defer fake.addVolumeMountsMutex.Unlock()
	fake.AddVolumeMountsStub = nil
	fake.addVolumeMountsReturns = struct {
		result1 k8s.ContainerBuilder
	}{result1}
}

func (fake *K8sContainerBuilder) AddVolumeMountsReturnsOnCall(i int, result1 k8s.ContainerBuilder) {
	fake.addVolumeMountsMutex.Lock()
	defer fake.addVolumeMountsMutex.Unlock()
	fake.AddVolumeMountsStub = nil
	if fake.addVolumeMountsReturnsOnCall == nil {
		fake.addVolumeMountsReturnsOnCall = make(map[int]struct {
			result1 k8s.ContainerBuilder
		})
	}
	fake.addVolumeMountsReturnsOnCall[i] = struct {
		result1 k8s.ContainerBuilder
	}{result1}
}

func (fake *K8sContainerBuilder) Build(arg1 context.Context) (*v1.Container, error) {
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

func (fake *K8sContainerBuilder) BuildCallCount() int {
	fake.buildMutex.RLock()
	defer fake.buildMutex.RUnlock()
	return len(fake.buildArgsForCall)
}

func (fake *K8sContainerBuilder) BuildCalls(stub func(context.Context) (*v1.Container, error)) {
	fake.buildMutex.Lock()
	defer fake.buildMutex.Unlock()
	fake.BuildStub = stub
}

func (fake *K8sContainerBuilder) BuildArgsForCall(i int) context.Context {
	fake.buildMutex.RLock()
	defer fake.buildMutex.RUnlock()
	argsForCall := fake.buildArgsForCall[i]
	return argsForCall.arg1
}

func (fake *K8sContainerBuilder) BuildReturns(result1 *v1.Container, result2 error) {
	fake.buildMutex.Lock()
	defer fake.buildMutex.Unlock()
	fake.BuildStub = nil
	fake.buildReturns = struct {
		result1 *v1.Container
		result2 error
	}{result1, result2}
}

func (fake *K8sContainerBuilder) BuildReturnsOnCall(i int, result1 *v1.Container, result2 error) {
	fake.buildMutex.Lock()
	defer fake.buildMutex.Unlock()
	fake.BuildStub = nil
	if fake.buildReturnsOnCall == nil {
		fake.buildReturnsOnCall = make(map[int]struct {
			result1 *v1.Container
			result2 error
		})
	}
	fake.buildReturnsOnCall[i] = struct {
		result1 *v1.Container
		result2 error
	}{result1, result2}
}

func (fake *K8sContainerBuilder) SetArgs(arg1 []string) k8s.ContainerBuilder {
	var arg1Copy []string
	if arg1 != nil {
		arg1Copy = make([]string, len(arg1))
		copy(arg1Copy, arg1)
	}
	fake.setArgsMutex.Lock()
	ret, specificReturn := fake.setArgsReturnsOnCall[len(fake.setArgsArgsForCall)]
	fake.setArgsArgsForCall = append(fake.setArgsArgsForCall, struct {
		arg1 []string
	}{arg1Copy})
	stub := fake.SetArgsStub
	fakeReturns := fake.setArgsReturns
	fake.recordInvocation("SetArgs", []interface{}{arg1Copy})
	fake.setArgsMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *K8sContainerBuilder) SetArgsCallCount() int {
	fake.setArgsMutex.RLock()
	defer fake.setArgsMutex.RUnlock()
	return len(fake.setArgsArgsForCall)
}

func (fake *K8sContainerBuilder) SetArgsCalls(stub func([]string) k8s.ContainerBuilder) {
	fake.setArgsMutex.Lock()
	defer fake.setArgsMutex.Unlock()
	fake.SetArgsStub = stub
}

func (fake *K8sContainerBuilder) SetArgsArgsForCall(i int) []string {
	fake.setArgsMutex.RLock()
	defer fake.setArgsMutex.RUnlock()
	argsForCall := fake.setArgsArgsForCall[i]
	return argsForCall.arg1
}

func (fake *K8sContainerBuilder) SetArgsReturns(result1 k8s.ContainerBuilder) {
	fake.setArgsMutex.Lock()
	defer fake.setArgsMutex.Unlock()
	fake.SetArgsStub = nil
	fake.setArgsReturns = struct {
		result1 k8s.ContainerBuilder
	}{result1}
}

func (fake *K8sContainerBuilder) SetArgsReturnsOnCall(i int, result1 k8s.ContainerBuilder) {
	fake.setArgsMutex.Lock()
	defer fake.setArgsMutex.Unlock()
	fake.SetArgsStub = nil
	if fake.setArgsReturnsOnCall == nil {
		fake.setArgsReturnsOnCall = make(map[int]struct {
			result1 k8s.ContainerBuilder
		})
	}
	fake.setArgsReturnsOnCall[i] = struct {
		result1 k8s.ContainerBuilder
	}{result1}
}

func (fake *K8sContainerBuilder) SetCommand(arg1 []string) k8s.ContainerBuilder {
	var arg1Copy []string
	if arg1 != nil {
		arg1Copy = make([]string, len(arg1))
		copy(arg1Copy, arg1)
	}
	fake.setCommandMutex.Lock()
	ret, specificReturn := fake.setCommandReturnsOnCall[len(fake.setCommandArgsForCall)]
	fake.setCommandArgsForCall = append(fake.setCommandArgsForCall, struct {
		arg1 []string
	}{arg1Copy})
	stub := fake.SetCommandStub
	fakeReturns := fake.setCommandReturns
	fake.recordInvocation("SetCommand", []interface{}{arg1Copy})
	fake.setCommandMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *K8sContainerBuilder) SetCommandCallCount() int {
	fake.setCommandMutex.RLock()
	defer fake.setCommandMutex.RUnlock()
	return len(fake.setCommandArgsForCall)
}

func (fake *K8sContainerBuilder) SetCommandCalls(stub func([]string) k8s.ContainerBuilder) {
	fake.setCommandMutex.Lock()
	defer fake.setCommandMutex.Unlock()
	fake.SetCommandStub = stub
}

func (fake *K8sContainerBuilder) SetCommandArgsForCall(i int) []string {
	fake.setCommandMutex.RLock()
	defer fake.setCommandMutex.RUnlock()
	argsForCall := fake.setCommandArgsForCall[i]
	return argsForCall.arg1
}

func (fake *K8sContainerBuilder) SetCommandReturns(result1 k8s.ContainerBuilder) {
	fake.setCommandMutex.Lock()
	defer fake.setCommandMutex.Unlock()
	fake.SetCommandStub = nil
	fake.setCommandReturns = struct {
		result1 k8s.ContainerBuilder
	}{result1}
}

func (fake *K8sContainerBuilder) SetCommandReturnsOnCall(i int, result1 k8s.ContainerBuilder) {
	fake.setCommandMutex.Lock()
	defer fake.setCommandMutex.Unlock()
	fake.SetCommandStub = nil
	if fake.setCommandReturnsOnCall == nil {
		fake.setCommandReturnsOnCall = make(map[int]struct {
			result1 k8s.ContainerBuilder
		})
	}
	fake.setCommandReturnsOnCall[i] = struct {
		result1 k8s.ContainerBuilder
	}{result1}
}

func (fake *K8sContainerBuilder) SetCpuLimit(arg1 string) k8s.ContainerBuilder {
	fake.setCpuLimitMutex.Lock()
	ret, specificReturn := fake.setCpuLimitReturnsOnCall[len(fake.setCpuLimitArgsForCall)]
	fake.setCpuLimitArgsForCall = append(fake.setCpuLimitArgsForCall, struct {
		arg1 string
	}{arg1})
	stub := fake.SetCpuLimitStub
	fakeReturns := fake.setCpuLimitReturns
	fake.recordInvocation("SetCpuLimit", []interface{}{arg1})
	fake.setCpuLimitMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *K8sContainerBuilder) SetCpuLimitCallCount() int {
	fake.setCpuLimitMutex.RLock()
	defer fake.setCpuLimitMutex.RUnlock()
	return len(fake.setCpuLimitArgsForCall)
}

func (fake *K8sContainerBuilder) SetCpuLimitCalls(stub func(string) k8s.ContainerBuilder) {
	fake.setCpuLimitMutex.Lock()
	defer fake.setCpuLimitMutex.Unlock()
	fake.SetCpuLimitStub = stub
}

func (fake *K8sContainerBuilder) SetCpuLimitArgsForCall(i int) string {
	fake.setCpuLimitMutex.RLock()
	defer fake.setCpuLimitMutex.RUnlock()
	argsForCall := fake.setCpuLimitArgsForCall[i]
	return argsForCall.arg1
}

func (fake *K8sContainerBuilder) SetCpuLimitReturns(result1 k8s.ContainerBuilder) {
	fake.setCpuLimitMutex.Lock()
	defer fake.setCpuLimitMutex.Unlock()
	fake.SetCpuLimitStub = nil
	fake.setCpuLimitReturns = struct {
		result1 k8s.ContainerBuilder
	}{result1}
}

func (fake *K8sContainerBuilder) SetCpuLimitReturnsOnCall(i int, result1 k8s.ContainerBuilder) {
	fake.setCpuLimitMutex.Lock()
	defer fake.setCpuLimitMutex.Unlock()
	fake.SetCpuLimitStub = nil
	if fake.setCpuLimitReturnsOnCall == nil {
		fake.setCpuLimitReturnsOnCall = make(map[int]struct {
			result1 k8s.ContainerBuilder
		})
	}
	fake.setCpuLimitReturnsOnCall[i] = struct {
		result1 k8s.ContainerBuilder
	}{result1}
}

func (fake *K8sContainerBuilder) SetCpuRequest(arg1 string) k8s.ContainerBuilder {
	fake.setCpuRequestMutex.Lock()
	ret, specificReturn := fake.setCpuRequestReturnsOnCall[len(fake.setCpuRequestArgsForCall)]
	fake.setCpuRequestArgsForCall = append(fake.setCpuRequestArgsForCall, struct {
		arg1 string
	}{arg1})
	stub := fake.SetCpuRequestStub
	fakeReturns := fake.setCpuRequestReturns
	fake.recordInvocation("SetCpuRequest", []interface{}{arg1})
	fake.setCpuRequestMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *K8sContainerBuilder) SetCpuRequestCallCount() int {
	fake.setCpuRequestMutex.RLock()
	defer fake.setCpuRequestMutex.RUnlock()
	return len(fake.setCpuRequestArgsForCall)
}

func (fake *K8sContainerBuilder) SetCpuRequestCalls(stub func(string) k8s.ContainerBuilder) {
	fake.setCpuRequestMutex.Lock()
	defer fake.setCpuRequestMutex.Unlock()
	fake.SetCpuRequestStub = stub
}

func (fake *K8sContainerBuilder) SetCpuRequestArgsForCall(i int) string {
	fake.setCpuRequestMutex.RLock()
	defer fake.setCpuRequestMutex.RUnlock()
	argsForCall := fake.setCpuRequestArgsForCall[i]
	return argsForCall.arg1
}

func (fake *K8sContainerBuilder) SetCpuRequestReturns(result1 k8s.ContainerBuilder) {
	fake.setCpuRequestMutex.Lock()
	defer fake.setCpuRequestMutex.Unlock()
	fake.SetCpuRequestStub = nil
	fake.setCpuRequestReturns = struct {
		result1 k8s.ContainerBuilder
	}{result1}
}

func (fake *K8sContainerBuilder) SetCpuRequestReturnsOnCall(i int, result1 k8s.ContainerBuilder) {
	fake.setCpuRequestMutex.Lock()
	defer fake.setCpuRequestMutex.Unlock()
	fake.SetCpuRequestStub = nil
	if fake.setCpuRequestReturnsOnCall == nil {
		fake.setCpuRequestReturnsOnCall = make(map[int]struct {
			result1 k8s.ContainerBuilder
		})
	}
	fake.setCpuRequestReturnsOnCall[i] = struct {
		result1 k8s.ContainerBuilder
	}{result1}
}

func (fake *K8sContainerBuilder) SetEnvBuilder(arg1 k8s.EnvBuilder) k8s.ContainerBuilder {
	fake.setEnvBuilderMutex.Lock()
	ret, specificReturn := fake.setEnvBuilderReturnsOnCall[len(fake.setEnvBuilderArgsForCall)]
	fake.setEnvBuilderArgsForCall = append(fake.setEnvBuilderArgsForCall, struct {
		arg1 k8s.EnvBuilder
	}{arg1})
	stub := fake.SetEnvBuilderStub
	fakeReturns := fake.setEnvBuilderReturns
	fake.recordInvocation("SetEnvBuilder", []interface{}{arg1})
	fake.setEnvBuilderMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *K8sContainerBuilder) SetEnvBuilderCallCount() int {
	fake.setEnvBuilderMutex.RLock()
	defer fake.setEnvBuilderMutex.RUnlock()
	return len(fake.setEnvBuilderArgsForCall)
}

func (fake *K8sContainerBuilder) SetEnvBuilderCalls(stub func(k8s.EnvBuilder) k8s.ContainerBuilder) {
	fake.setEnvBuilderMutex.Lock()
	defer fake.setEnvBuilderMutex.Unlock()
	fake.SetEnvBuilderStub = stub
}

func (fake *K8sContainerBuilder) SetEnvBuilderArgsForCall(i int) k8s.EnvBuilder {
	fake.setEnvBuilderMutex.RLock()
	defer fake.setEnvBuilderMutex.RUnlock()
	argsForCall := fake.setEnvBuilderArgsForCall[i]
	return argsForCall.arg1
}

func (fake *K8sContainerBuilder) SetEnvBuilderReturns(result1 k8s.ContainerBuilder) {
	fake.setEnvBuilderMutex.Lock()
	defer fake.setEnvBuilderMutex.Unlock()
	fake.SetEnvBuilderStub = nil
	fake.setEnvBuilderReturns = struct {
		result1 k8s.ContainerBuilder
	}{result1}
}

func (fake *K8sContainerBuilder) SetEnvBuilderReturnsOnCall(i int, result1 k8s.ContainerBuilder) {
	fake.setEnvBuilderMutex.Lock()
	defer fake.setEnvBuilderMutex.Unlock()
	fake.SetEnvBuilderStub = nil
	if fake.setEnvBuilderReturnsOnCall == nil {
		fake.setEnvBuilderReturnsOnCall = make(map[int]struct {
			result1 k8s.ContainerBuilder
		})
	}
	fake.setEnvBuilderReturnsOnCall[i] = struct {
		result1 k8s.ContainerBuilder
	}{result1}
}

func (fake *K8sContainerBuilder) SetImage(arg1 string) k8s.ContainerBuilder {
	fake.setImageMutex.Lock()
	ret, specificReturn := fake.setImageReturnsOnCall[len(fake.setImageArgsForCall)]
	fake.setImageArgsForCall = append(fake.setImageArgsForCall, struct {
		arg1 string
	}{arg1})
	stub := fake.SetImageStub
	fakeReturns := fake.setImageReturns
	fake.recordInvocation("SetImage", []interface{}{arg1})
	fake.setImageMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *K8sContainerBuilder) SetImageCallCount() int {
	fake.setImageMutex.RLock()
	defer fake.setImageMutex.RUnlock()
	return len(fake.setImageArgsForCall)
}

func (fake *K8sContainerBuilder) SetImageCalls(stub func(string) k8s.ContainerBuilder) {
	fake.setImageMutex.Lock()
	defer fake.setImageMutex.Unlock()
	fake.SetImageStub = stub
}

func (fake *K8sContainerBuilder) SetImageArgsForCall(i int) string {
	fake.setImageMutex.RLock()
	defer fake.setImageMutex.RUnlock()
	argsForCall := fake.setImageArgsForCall[i]
	return argsForCall.arg1
}

func (fake *K8sContainerBuilder) SetImageReturns(result1 k8s.ContainerBuilder) {
	fake.setImageMutex.Lock()
	defer fake.setImageMutex.Unlock()
	fake.SetImageStub = nil
	fake.setImageReturns = struct {
		result1 k8s.ContainerBuilder
	}{result1}
}

func (fake *K8sContainerBuilder) SetImageReturnsOnCall(i int, result1 k8s.ContainerBuilder) {
	fake.setImageMutex.Lock()
	defer fake.setImageMutex.Unlock()
	fake.SetImageStub = nil
	if fake.setImageReturnsOnCall == nil {
		fake.setImageReturnsOnCall = make(map[int]struct {
			result1 k8s.ContainerBuilder
		})
	}
	fake.setImageReturnsOnCall[i] = struct {
		result1 k8s.ContainerBuilder
	}{result1}
}

func (fake *K8sContainerBuilder) SetLivenessProbe(arg1 v1.Probe) k8s.ContainerBuilder {
	fake.setLivenessProbeMutex.Lock()
	ret, specificReturn := fake.setLivenessProbeReturnsOnCall[len(fake.setLivenessProbeArgsForCall)]
	fake.setLivenessProbeArgsForCall = append(fake.setLivenessProbeArgsForCall, struct {
		arg1 v1.Probe
	}{arg1})
	stub := fake.SetLivenessProbeStub
	fakeReturns := fake.setLivenessProbeReturns
	fake.recordInvocation("SetLivenessProbe", []interface{}{arg1})
	fake.setLivenessProbeMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *K8sContainerBuilder) SetLivenessProbeCallCount() int {
	fake.setLivenessProbeMutex.RLock()
	defer fake.setLivenessProbeMutex.RUnlock()
	return len(fake.setLivenessProbeArgsForCall)
}

func (fake *K8sContainerBuilder) SetLivenessProbeCalls(stub func(v1.Probe) k8s.ContainerBuilder) {
	fake.setLivenessProbeMutex.Lock()
	defer fake.setLivenessProbeMutex.Unlock()
	fake.SetLivenessProbeStub = stub
}

func (fake *K8sContainerBuilder) SetLivenessProbeArgsForCall(i int) v1.Probe {
	fake.setLivenessProbeMutex.RLock()
	defer fake.setLivenessProbeMutex.RUnlock()
	argsForCall := fake.setLivenessProbeArgsForCall[i]
	return argsForCall.arg1
}

func (fake *K8sContainerBuilder) SetLivenessProbeReturns(result1 k8s.ContainerBuilder) {
	fake.setLivenessProbeMutex.Lock()
	defer fake.setLivenessProbeMutex.Unlock()
	fake.SetLivenessProbeStub = nil
	fake.setLivenessProbeReturns = struct {
		result1 k8s.ContainerBuilder
	}{result1}
}

func (fake *K8sContainerBuilder) SetLivenessProbeReturnsOnCall(i int, result1 k8s.ContainerBuilder) {
	fake.setLivenessProbeMutex.Lock()
	defer fake.setLivenessProbeMutex.Unlock()
	fake.SetLivenessProbeStub = nil
	if fake.setLivenessProbeReturnsOnCall == nil {
		fake.setLivenessProbeReturnsOnCall = make(map[int]struct {
			result1 k8s.ContainerBuilder
		})
	}
	fake.setLivenessProbeReturnsOnCall[i] = struct {
		result1 k8s.ContainerBuilder
	}{result1}
}

func (fake *K8sContainerBuilder) SetMemoryLimit(arg1 string) k8s.ContainerBuilder {
	fake.setMemoryLimitMutex.Lock()
	ret, specificReturn := fake.setMemoryLimitReturnsOnCall[len(fake.setMemoryLimitArgsForCall)]
	fake.setMemoryLimitArgsForCall = append(fake.setMemoryLimitArgsForCall, struct {
		arg1 string
	}{arg1})
	stub := fake.SetMemoryLimitStub
	fakeReturns := fake.setMemoryLimitReturns
	fake.recordInvocation("SetMemoryLimit", []interface{}{arg1})
	fake.setMemoryLimitMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *K8sContainerBuilder) SetMemoryLimitCallCount() int {
	fake.setMemoryLimitMutex.RLock()
	defer fake.setMemoryLimitMutex.RUnlock()
	return len(fake.setMemoryLimitArgsForCall)
}

func (fake *K8sContainerBuilder) SetMemoryLimitCalls(stub func(string) k8s.ContainerBuilder) {
	fake.setMemoryLimitMutex.Lock()
	defer fake.setMemoryLimitMutex.Unlock()
	fake.SetMemoryLimitStub = stub
}

func (fake *K8sContainerBuilder) SetMemoryLimitArgsForCall(i int) string {
	fake.setMemoryLimitMutex.RLock()
	defer fake.setMemoryLimitMutex.RUnlock()
	argsForCall := fake.setMemoryLimitArgsForCall[i]
	return argsForCall.arg1
}

func (fake *K8sContainerBuilder) SetMemoryLimitReturns(result1 k8s.ContainerBuilder) {
	fake.setMemoryLimitMutex.Lock()
	defer fake.setMemoryLimitMutex.Unlock()
	fake.SetMemoryLimitStub = nil
	fake.setMemoryLimitReturns = struct {
		result1 k8s.ContainerBuilder
	}{result1}
}

func (fake *K8sContainerBuilder) SetMemoryLimitReturnsOnCall(i int, result1 k8s.ContainerBuilder) {
	fake.setMemoryLimitMutex.Lock()
	defer fake.setMemoryLimitMutex.Unlock()
	fake.SetMemoryLimitStub = nil
	if fake.setMemoryLimitReturnsOnCall == nil {
		fake.setMemoryLimitReturnsOnCall = make(map[int]struct {
			result1 k8s.ContainerBuilder
		})
	}
	fake.setMemoryLimitReturnsOnCall[i] = struct {
		result1 k8s.ContainerBuilder
	}{result1}
}

func (fake *K8sContainerBuilder) SetMemoryRequest(arg1 string) k8s.ContainerBuilder {
	fake.setMemoryRequestMutex.Lock()
	ret, specificReturn := fake.setMemoryRequestReturnsOnCall[len(fake.setMemoryRequestArgsForCall)]
	fake.setMemoryRequestArgsForCall = append(fake.setMemoryRequestArgsForCall, struct {
		arg1 string
	}{arg1})
	stub := fake.SetMemoryRequestStub
	fakeReturns := fake.setMemoryRequestReturns
	fake.recordInvocation("SetMemoryRequest", []interface{}{arg1})
	fake.setMemoryRequestMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *K8sContainerBuilder) SetMemoryRequestCallCount() int {
	fake.setMemoryRequestMutex.RLock()
	defer fake.setMemoryRequestMutex.RUnlock()
	return len(fake.setMemoryRequestArgsForCall)
}

func (fake *K8sContainerBuilder) SetMemoryRequestCalls(stub func(string) k8s.ContainerBuilder) {
	fake.setMemoryRequestMutex.Lock()
	defer fake.setMemoryRequestMutex.Unlock()
	fake.SetMemoryRequestStub = stub
}

func (fake *K8sContainerBuilder) SetMemoryRequestArgsForCall(i int) string {
	fake.setMemoryRequestMutex.RLock()
	defer fake.setMemoryRequestMutex.RUnlock()
	argsForCall := fake.setMemoryRequestArgsForCall[i]
	return argsForCall.arg1
}

func (fake *K8sContainerBuilder) SetMemoryRequestReturns(result1 k8s.ContainerBuilder) {
	fake.setMemoryRequestMutex.Lock()
	defer fake.setMemoryRequestMutex.Unlock()
	fake.SetMemoryRequestStub = nil
	fake.setMemoryRequestReturns = struct {
		result1 k8s.ContainerBuilder
	}{result1}
}

func (fake *K8sContainerBuilder) SetMemoryRequestReturnsOnCall(i int, result1 k8s.ContainerBuilder) {
	fake.setMemoryRequestMutex.Lock()
	defer fake.setMemoryRequestMutex.Unlock()
	fake.SetMemoryRequestStub = nil
	if fake.setMemoryRequestReturnsOnCall == nil {
		fake.setMemoryRequestReturnsOnCall = make(map[int]struct {
			result1 k8s.ContainerBuilder
		})
	}
	fake.setMemoryRequestReturnsOnCall[i] = struct {
		result1 k8s.ContainerBuilder
	}{result1}
}

func (fake *K8sContainerBuilder) SetName(arg1 string) k8s.ContainerBuilder {
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

func (fake *K8sContainerBuilder) SetNameCallCount() int {
	fake.setNameMutex.RLock()
	defer fake.setNameMutex.RUnlock()
	return len(fake.setNameArgsForCall)
}

func (fake *K8sContainerBuilder) SetNameCalls(stub func(string) k8s.ContainerBuilder) {
	fake.setNameMutex.Lock()
	defer fake.setNameMutex.Unlock()
	fake.SetNameStub = stub
}

func (fake *K8sContainerBuilder) SetNameArgsForCall(i int) string {
	fake.setNameMutex.RLock()
	defer fake.setNameMutex.RUnlock()
	argsForCall := fake.setNameArgsForCall[i]
	return argsForCall.arg1
}

func (fake *K8sContainerBuilder) SetNameReturns(result1 k8s.ContainerBuilder) {
	fake.setNameMutex.Lock()
	defer fake.setNameMutex.Unlock()
	fake.SetNameStub = nil
	fake.setNameReturns = struct {
		result1 k8s.ContainerBuilder
	}{result1}
}

func (fake *K8sContainerBuilder) SetNameReturnsOnCall(i int, result1 k8s.ContainerBuilder) {
	fake.setNameMutex.Lock()
	defer fake.setNameMutex.Unlock()
	fake.SetNameStub = nil
	if fake.setNameReturnsOnCall == nil {
		fake.setNameReturnsOnCall = make(map[int]struct {
			result1 k8s.ContainerBuilder
		})
	}
	fake.setNameReturnsOnCall[i] = struct {
		result1 k8s.ContainerBuilder
	}{result1}
}

func (fake *K8sContainerBuilder) SetPorts(arg1 []v1.ContainerPort) k8s.ContainerBuilder {
	var arg1Copy []v1.ContainerPort
	if arg1 != nil {
		arg1Copy = make([]v1.ContainerPort, len(arg1))
		copy(arg1Copy, arg1)
	}
	fake.setPortsMutex.Lock()
	ret, specificReturn := fake.setPortsReturnsOnCall[len(fake.setPortsArgsForCall)]
	fake.setPortsArgsForCall = append(fake.setPortsArgsForCall, struct {
		arg1 []v1.ContainerPort
	}{arg1Copy})
	stub := fake.SetPortsStub
	fakeReturns := fake.setPortsReturns
	fake.recordInvocation("SetPorts", []interface{}{arg1Copy})
	fake.setPortsMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *K8sContainerBuilder) SetPortsCallCount() int {
	fake.setPortsMutex.RLock()
	defer fake.setPortsMutex.RUnlock()
	return len(fake.setPortsArgsForCall)
}

func (fake *K8sContainerBuilder) SetPortsCalls(stub func([]v1.ContainerPort) k8s.ContainerBuilder) {
	fake.setPortsMutex.Lock()
	defer fake.setPortsMutex.Unlock()
	fake.SetPortsStub = stub
}

func (fake *K8sContainerBuilder) SetPortsArgsForCall(i int) []v1.ContainerPort {
	fake.setPortsMutex.RLock()
	defer fake.setPortsMutex.RUnlock()
	argsForCall := fake.setPortsArgsForCall[i]
	return argsForCall.arg1
}

func (fake *K8sContainerBuilder) SetPortsReturns(result1 k8s.ContainerBuilder) {
	fake.setPortsMutex.Lock()
	defer fake.setPortsMutex.Unlock()
	fake.SetPortsStub = nil
	fake.setPortsReturns = struct {
		result1 k8s.ContainerBuilder
	}{result1}
}

func (fake *K8sContainerBuilder) SetPortsReturnsOnCall(i int, result1 k8s.ContainerBuilder) {
	fake.setPortsMutex.Lock()
	defer fake.setPortsMutex.Unlock()
	fake.SetPortsStub = nil
	if fake.setPortsReturnsOnCall == nil {
		fake.setPortsReturnsOnCall = make(map[int]struct {
			result1 k8s.ContainerBuilder
		})
	}
	fake.setPortsReturnsOnCall[i] = struct {
		result1 k8s.ContainerBuilder
	}{result1}
}

func (fake *K8sContainerBuilder) SetReadinessProbe(arg1 v1.Probe) k8s.ContainerBuilder {
	fake.setReadinessProbeMutex.Lock()
	ret, specificReturn := fake.setReadinessProbeReturnsOnCall[len(fake.setReadinessProbeArgsForCall)]
	fake.setReadinessProbeArgsForCall = append(fake.setReadinessProbeArgsForCall, struct {
		arg1 v1.Probe
	}{arg1})
	stub := fake.SetReadinessProbeStub
	fakeReturns := fake.setReadinessProbeReturns
	fake.recordInvocation("SetReadinessProbe", []interface{}{arg1})
	fake.setReadinessProbeMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *K8sContainerBuilder) SetReadinessProbeCallCount() int {
	fake.setReadinessProbeMutex.RLock()
	defer fake.setReadinessProbeMutex.RUnlock()
	return len(fake.setReadinessProbeArgsForCall)
}

func (fake *K8sContainerBuilder) SetReadinessProbeCalls(stub func(v1.Probe) k8s.ContainerBuilder) {
	fake.setReadinessProbeMutex.Lock()
	defer fake.setReadinessProbeMutex.Unlock()
	fake.SetReadinessProbeStub = stub
}

func (fake *K8sContainerBuilder) SetReadinessProbeArgsForCall(i int) v1.Probe {
	fake.setReadinessProbeMutex.RLock()
	defer fake.setReadinessProbeMutex.RUnlock()
	argsForCall := fake.setReadinessProbeArgsForCall[i]
	return argsForCall.arg1
}

func (fake *K8sContainerBuilder) SetReadinessProbeReturns(result1 k8s.ContainerBuilder) {
	fake.setReadinessProbeMutex.Lock()
	defer fake.setReadinessProbeMutex.Unlock()
	fake.SetReadinessProbeStub = nil
	fake.setReadinessProbeReturns = struct {
		result1 k8s.ContainerBuilder
	}{result1}
}

func (fake *K8sContainerBuilder) SetReadinessProbeReturnsOnCall(i int, result1 k8s.ContainerBuilder) {
	fake.setReadinessProbeMutex.Lock()
	defer fake.setReadinessProbeMutex.Unlock()
	fake.SetReadinessProbeStub = nil
	if fake.setReadinessProbeReturnsOnCall == nil {
		fake.setReadinessProbeReturnsOnCall = make(map[int]struct {
			result1 k8s.ContainerBuilder
		})
	}
	fake.setReadinessProbeReturnsOnCall[i] = struct {
		result1 k8s.ContainerBuilder
	}{result1}
}

func (fake *K8sContainerBuilder) SetVolumeMounts(arg1 []v1.VolumeMount) k8s.ContainerBuilder {
	var arg1Copy []v1.VolumeMount
	if arg1 != nil {
		arg1Copy = make([]v1.VolumeMount, len(arg1))
		copy(arg1Copy, arg1)
	}
	fake.setVolumeMountsMutex.Lock()
	ret, specificReturn := fake.setVolumeMountsReturnsOnCall[len(fake.setVolumeMountsArgsForCall)]
	fake.setVolumeMountsArgsForCall = append(fake.setVolumeMountsArgsForCall, struct {
		arg1 []v1.VolumeMount
	}{arg1Copy})
	stub := fake.SetVolumeMountsStub
	fakeReturns := fake.setVolumeMountsReturns
	fake.recordInvocation("SetVolumeMounts", []interface{}{arg1Copy})
	fake.setVolumeMountsMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *K8sContainerBuilder) SetVolumeMountsCallCount() int {
	fake.setVolumeMountsMutex.RLock()
	defer fake.setVolumeMountsMutex.RUnlock()
	return len(fake.setVolumeMountsArgsForCall)
}

func (fake *K8sContainerBuilder) SetVolumeMountsCalls(stub func([]v1.VolumeMount) k8s.ContainerBuilder) {
	fake.setVolumeMountsMutex.Lock()
	defer fake.setVolumeMountsMutex.Unlock()
	fake.SetVolumeMountsStub = stub
}

func (fake *K8sContainerBuilder) SetVolumeMountsArgsForCall(i int) []v1.VolumeMount {
	fake.setVolumeMountsMutex.RLock()
	defer fake.setVolumeMountsMutex.RUnlock()
	argsForCall := fake.setVolumeMountsArgsForCall[i]
	return argsForCall.arg1
}

func (fake *K8sContainerBuilder) SetVolumeMountsReturns(result1 k8s.ContainerBuilder) {
	fake.setVolumeMountsMutex.Lock()
	defer fake.setVolumeMountsMutex.Unlock()
	fake.SetVolumeMountsStub = nil
	fake.setVolumeMountsReturns = struct {
		result1 k8s.ContainerBuilder
	}{result1}
}

func (fake *K8sContainerBuilder) SetVolumeMountsReturnsOnCall(i int, result1 k8s.ContainerBuilder) {
	fake.setVolumeMountsMutex.Lock()
	defer fake.setVolumeMountsMutex.Unlock()
	fake.SetVolumeMountsStub = nil
	if fake.setVolumeMountsReturnsOnCall == nil {
		fake.setVolumeMountsReturnsOnCall = make(map[int]struct {
			result1 k8s.ContainerBuilder
		})
	}
	fake.setVolumeMountsReturnsOnCall[i] = struct {
		result1 k8s.ContainerBuilder
	}{result1}
}

func (fake *K8sContainerBuilder) Validate(arg1 context.Context) error {
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

func (fake *K8sContainerBuilder) ValidateCallCount() int {
	fake.validateMutex.RLock()
	defer fake.validateMutex.RUnlock()
	return len(fake.validateArgsForCall)
}

func (fake *K8sContainerBuilder) ValidateCalls(stub func(context.Context) error) {
	fake.validateMutex.Lock()
	defer fake.validateMutex.Unlock()
	fake.ValidateStub = stub
}

func (fake *K8sContainerBuilder) ValidateArgsForCall(i int) context.Context {
	fake.validateMutex.RLock()
	defer fake.validateMutex.RUnlock()
	argsForCall := fake.validateArgsForCall[i]
	return argsForCall.arg1
}

func (fake *K8sContainerBuilder) ValidateReturns(result1 error) {
	fake.validateMutex.Lock()
	defer fake.validateMutex.Unlock()
	fake.ValidateStub = nil
	fake.validateReturns = struct {
		result1 error
	}{result1}
}

func (fake *K8sContainerBuilder) ValidateReturnsOnCall(i int, result1 error) {
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

func (fake *K8sContainerBuilder) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.addVolumeMountsMutex.RLock()
	defer fake.addVolumeMountsMutex.RUnlock()
	fake.buildMutex.RLock()
	defer fake.buildMutex.RUnlock()
	fake.setArgsMutex.RLock()
	defer fake.setArgsMutex.RUnlock()
	fake.setCommandMutex.RLock()
	defer fake.setCommandMutex.RUnlock()
	fake.setCpuLimitMutex.RLock()
	defer fake.setCpuLimitMutex.RUnlock()
	fake.setCpuRequestMutex.RLock()
	defer fake.setCpuRequestMutex.RUnlock()
	fake.setEnvBuilderMutex.RLock()
	defer fake.setEnvBuilderMutex.RUnlock()
	fake.setImageMutex.RLock()
	defer fake.setImageMutex.RUnlock()
	fake.setLivenessProbeMutex.RLock()
	defer fake.setLivenessProbeMutex.RUnlock()
	fake.setMemoryLimitMutex.RLock()
	defer fake.setMemoryLimitMutex.RUnlock()
	fake.setMemoryRequestMutex.RLock()
	defer fake.setMemoryRequestMutex.RUnlock()
	fake.setNameMutex.RLock()
	defer fake.setNameMutex.RUnlock()
	fake.setPortsMutex.RLock()
	defer fake.setPortsMutex.RUnlock()
	fake.setReadinessProbeMutex.RLock()
	defer fake.setReadinessProbeMutex.RUnlock()
	fake.setVolumeMountsMutex.RLock()
	defer fake.setVolumeMountsMutex.RUnlock()
	fake.validateMutex.RLock()
	defer fake.validateMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *K8sContainerBuilder) recordInvocation(key string, args []interface{}) {
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

var _ k8s.ContainerBuilder = new(K8sContainerBuilder)