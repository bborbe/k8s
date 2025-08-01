// Code generated by counterfeiter. DO NOT EDIT.
package mocks

import (
	"context"
	"sync"

	"github.com/bborbe/k8s"
	v1a "k8s.io/api/apps/v1"
	v1 "k8s.io/api/core/v1"
	v1b "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type K8sDeploymentBuilder struct {
	AddImagePullSecretsStub        func(...string) k8s.DeploymentBuilder
	addImagePullSecretsMutex       sync.RWMutex
	addImagePullSecretsArgsForCall []struct {
		arg1 []string
	}
	addImagePullSecretsReturns struct {
		result1 k8s.DeploymentBuilder
	}
	addImagePullSecretsReturnsOnCall map[int]struct {
		result1 k8s.DeploymentBuilder
	}
	AddVolumesStub        func(...v1.Volume) k8s.DeploymentBuilder
	addVolumesMutex       sync.RWMutex
	addVolumesArgsForCall []struct {
		arg1 []v1.Volume
	}
	addVolumesReturns struct {
		result1 k8s.DeploymentBuilder
	}
	addVolumesReturnsOnCall map[int]struct {
		result1 k8s.DeploymentBuilder
	}
	BuildStub        func(context.Context) (*v1a.Deployment, error)
	buildMutex       sync.RWMutex
	buildArgsForCall []struct {
		arg1 context.Context
	}
	buildReturns struct {
		result1 *v1a.Deployment
		result2 error
	}
	buildReturnsOnCall map[int]struct {
		result1 *v1a.Deployment
		result2 error
	}
	SetAffinityStub        func(v1.Affinity) k8s.DeploymentBuilder
	setAffinityMutex       sync.RWMutex
	setAffinityArgsForCall []struct {
		arg1 v1.Affinity
	}
	setAffinityReturns struct {
		result1 k8s.DeploymentBuilder
	}
	setAffinityReturnsOnCall map[int]struct {
		result1 k8s.DeploymentBuilder
	}
	SetComponentStub        func(string) k8s.DeploymentBuilder
	setComponentMutex       sync.RWMutex
	setComponentArgsForCall []struct {
		arg1 string
	}
	setComponentReturns struct {
		result1 k8s.DeploymentBuilder
	}
	setComponentReturnsOnCall map[int]struct {
		result1 k8s.DeploymentBuilder
	}
	SetContainersStub        func([]v1.Container) k8s.DeploymentBuilder
	setContainersMutex       sync.RWMutex
	setContainersArgsForCall []struct {
		arg1 []v1.Container
	}
	setContainersReturns struct {
		result1 k8s.DeploymentBuilder
	}
	setContainersReturnsOnCall map[int]struct {
		result1 k8s.DeploymentBuilder
	}
	SetContainersBuilderStub        func(k8s.HasBuildContainers) k8s.DeploymentBuilder
	setContainersBuilderMutex       sync.RWMutex
	setContainersBuilderArgsForCall []struct {
		arg1 k8s.HasBuildContainers
	}
	setContainersBuilderReturns struct {
		result1 k8s.DeploymentBuilder
	}
	setContainersBuilderReturnsOnCall map[int]struct {
		result1 k8s.DeploymentBuilder
	}
	SetImagePullSecretsStub        func([]string) k8s.DeploymentBuilder
	setImagePullSecretsMutex       sync.RWMutex
	setImagePullSecretsArgsForCall []struct {
		arg1 []string
	}
	setImagePullSecretsReturns struct {
		result1 k8s.DeploymentBuilder
	}
	setImagePullSecretsReturnsOnCall map[int]struct {
		result1 k8s.DeploymentBuilder
	}
	SetNameStub        func(k8s.Name) k8s.DeploymentBuilder
	setNameMutex       sync.RWMutex
	setNameArgsForCall []struct {
		arg1 k8s.Name
	}
	setNameReturns struct {
		result1 k8s.DeploymentBuilder
	}
	setNameReturnsOnCall map[int]struct {
		result1 k8s.DeploymentBuilder
	}
	SetObjectMetaStub        func(v1b.ObjectMeta) k8s.DeploymentBuilder
	setObjectMetaMutex       sync.RWMutex
	setObjectMetaArgsForCall []struct {
		arg1 v1b.ObjectMeta
	}
	setObjectMetaReturns struct {
		result1 k8s.DeploymentBuilder
	}
	setObjectMetaReturnsOnCall map[int]struct {
		result1 k8s.DeploymentBuilder
	}
	SetObjectMetaBuilderStub        func(k8s.HasBuildObjectMeta) k8s.DeploymentBuilder
	setObjectMetaBuilderMutex       sync.RWMutex
	setObjectMetaBuilderArgsForCall []struct {
		arg1 k8s.HasBuildObjectMeta
	}
	setObjectMetaBuilderReturns struct {
		result1 k8s.DeploymentBuilder
	}
	setObjectMetaBuilderReturnsOnCall map[int]struct {
		result1 k8s.DeploymentBuilder
	}
	SetReplicasStub        func(int32) k8s.DeploymentBuilder
	setReplicasMutex       sync.RWMutex
	setReplicasArgsForCall []struct {
		arg1 int32
	}
	setReplicasReturns struct {
		result1 k8s.DeploymentBuilder
	}
	setReplicasReturnsOnCall map[int]struct {
		result1 k8s.DeploymentBuilder
	}
	SetServiceAccountNameStub        func(string) k8s.DeploymentBuilder
	setServiceAccountNameMutex       sync.RWMutex
	setServiceAccountNameArgsForCall []struct {
		arg1 string
	}
	setServiceAccountNameReturns struct {
		result1 k8s.DeploymentBuilder
	}
	setServiceAccountNameReturnsOnCall map[int]struct {
		result1 k8s.DeploymentBuilder
	}
	SetVolumesStub        func([]v1.Volume) k8s.DeploymentBuilder
	setVolumesMutex       sync.RWMutex
	setVolumesArgsForCall []struct {
		arg1 []v1.Volume
	}
	setVolumesReturns struct {
		result1 k8s.DeploymentBuilder
	}
	setVolumesReturnsOnCall map[int]struct {
		result1 k8s.DeploymentBuilder
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

func (fake *K8sDeploymentBuilder) AddImagePullSecrets(arg1 ...string) k8s.DeploymentBuilder {
	fake.addImagePullSecretsMutex.Lock()
	ret, specificReturn := fake.addImagePullSecretsReturnsOnCall[len(fake.addImagePullSecretsArgsForCall)]
	fake.addImagePullSecretsArgsForCall = append(fake.addImagePullSecretsArgsForCall, struct {
		arg1 []string
	}{arg1})
	stub := fake.AddImagePullSecretsStub
	fakeReturns := fake.addImagePullSecretsReturns
	fake.recordInvocation("AddImagePullSecrets", []interface{}{arg1})
	fake.addImagePullSecretsMutex.Unlock()
	if stub != nil {
		return stub(arg1...)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *K8sDeploymentBuilder) AddImagePullSecretsCallCount() int {
	fake.addImagePullSecretsMutex.RLock()
	defer fake.addImagePullSecretsMutex.RUnlock()
	return len(fake.addImagePullSecretsArgsForCall)
}

func (fake *K8sDeploymentBuilder) AddImagePullSecretsCalls(stub func(...string) k8s.DeploymentBuilder) {
	fake.addImagePullSecretsMutex.Lock()
	defer fake.addImagePullSecretsMutex.Unlock()
	fake.AddImagePullSecretsStub = stub
}

func (fake *K8sDeploymentBuilder) AddImagePullSecretsArgsForCall(i int) []string {
	fake.addImagePullSecretsMutex.RLock()
	defer fake.addImagePullSecretsMutex.RUnlock()
	argsForCall := fake.addImagePullSecretsArgsForCall[i]
	return argsForCall.arg1
}

func (fake *K8sDeploymentBuilder) AddImagePullSecretsReturns(result1 k8s.DeploymentBuilder) {
	fake.addImagePullSecretsMutex.Lock()
	defer fake.addImagePullSecretsMutex.Unlock()
	fake.AddImagePullSecretsStub = nil
	fake.addImagePullSecretsReturns = struct {
		result1 k8s.DeploymentBuilder
	}{result1}
}

func (fake *K8sDeploymentBuilder) AddImagePullSecretsReturnsOnCall(i int, result1 k8s.DeploymentBuilder) {
	fake.addImagePullSecretsMutex.Lock()
	defer fake.addImagePullSecretsMutex.Unlock()
	fake.AddImagePullSecretsStub = nil
	if fake.addImagePullSecretsReturnsOnCall == nil {
		fake.addImagePullSecretsReturnsOnCall = make(map[int]struct {
			result1 k8s.DeploymentBuilder
		})
	}
	fake.addImagePullSecretsReturnsOnCall[i] = struct {
		result1 k8s.DeploymentBuilder
	}{result1}
}

func (fake *K8sDeploymentBuilder) AddVolumes(arg1 ...v1.Volume) k8s.DeploymentBuilder {
	fake.addVolumesMutex.Lock()
	ret, specificReturn := fake.addVolumesReturnsOnCall[len(fake.addVolumesArgsForCall)]
	fake.addVolumesArgsForCall = append(fake.addVolumesArgsForCall, struct {
		arg1 []v1.Volume
	}{arg1})
	stub := fake.AddVolumesStub
	fakeReturns := fake.addVolumesReturns
	fake.recordInvocation("AddVolumes", []interface{}{arg1})
	fake.addVolumesMutex.Unlock()
	if stub != nil {
		return stub(arg1...)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *K8sDeploymentBuilder) AddVolumesCallCount() int {
	fake.addVolumesMutex.RLock()
	defer fake.addVolumesMutex.RUnlock()
	return len(fake.addVolumesArgsForCall)
}

func (fake *K8sDeploymentBuilder) AddVolumesCalls(stub func(...v1.Volume) k8s.DeploymentBuilder) {
	fake.addVolumesMutex.Lock()
	defer fake.addVolumesMutex.Unlock()
	fake.AddVolumesStub = stub
}

func (fake *K8sDeploymentBuilder) AddVolumesArgsForCall(i int) []v1.Volume {
	fake.addVolumesMutex.RLock()
	defer fake.addVolumesMutex.RUnlock()
	argsForCall := fake.addVolumesArgsForCall[i]
	return argsForCall.arg1
}

func (fake *K8sDeploymentBuilder) AddVolumesReturns(result1 k8s.DeploymentBuilder) {
	fake.addVolumesMutex.Lock()
	defer fake.addVolumesMutex.Unlock()
	fake.AddVolumesStub = nil
	fake.addVolumesReturns = struct {
		result1 k8s.DeploymentBuilder
	}{result1}
}

func (fake *K8sDeploymentBuilder) AddVolumesReturnsOnCall(i int, result1 k8s.DeploymentBuilder) {
	fake.addVolumesMutex.Lock()
	defer fake.addVolumesMutex.Unlock()
	fake.AddVolumesStub = nil
	if fake.addVolumesReturnsOnCall == nil {
		fake.addVolumesReturnsOnCall = make(map[int]struct {
			result1 k8s.DeploymentBuilder
		})
	}
	fake.addVolumesReturnsOnCall[i] = struct {
		result1 k8s.DeploymentBuilder
	}{result1}
}

func (fake *K8sDeploymentBuilder) Build(arg1 context.Context) (*v1a.Deployment, error) {
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

func (fake *K8sDeploymentBuilder) BuildCallCount() int {
	fake.buildMutex.RLock()
	defer fake.buildMutex.RUnlock()
	return len(fake.buildArgsForCall)
}

func (fake *K8sDeploymentBuilder) BuildCalls(stub func(context.Context) (*v1a.Deployment, error)) {
	fake.buildMutex.Lock()
	defer fake.buildMutex.Unlock()
	fake.BuildStub = stub
}

func (fake *K8sDeploymentBuilder) BuildArgsForCall(i int) context.Context {
	fake.buildMutex.RLock()
	defer fake.buildMutex.RUnlock()
	argsForCall := fake.buildArgsForCall[i]
	return argsForCall.arg1
}

func (fake *K8sDeploymentBuilder) BuildReturns(result1 *v1a.Deployment, result2 error) {
	fake.buildMutex.Lock()
	defer fake.buildMutex.Unlock()
	fake.BuildStub = nil
	fake.buildReturns = struct {
		result1 *v1a.Deployment
		result2 error
	}{result1, result2}
}

func (fake *K8sDeploymentBuilder) BuildReturnsOnCall(i int, result1 *v1a.Deployment, result2 error) {
	fake.buildMutex.Lock()
	defer fake.buildMutex.Unlock()
	fake.BuildStub = nil
	if fake.buildReturnsOnCall == nil {
		fake.buildReturnsOnCall = make(map[int]struct {
			result1 *v1a.Deployment
			result2 error
		})
	}
	fake.buildReturnsOnCall[i] = struct {
		result1 *v1a.Deployment
		result2 error
	}{result1, result2}
}

func (fake *K8sDeploymentBuilder) SetAffinity(arg1 v1.Affinity) k8s.DeploymentBuilder {
	fake.setAffinityMutex.Lock()
	ret, specificReturn := fake.setAffinityReturnsOnCall[len(fake.setAffinityArgsForCall)]
	fake.setAffinityArgsForCall = append(fake.setAffinityArgsForCall, struct {
		arg1 v1.Affinity
	}{arg1})
	stub := fake.SetAffinityStub
	fakeReturns := fake.setAffinityReturns
	fake.recordInvocation("SetAffinity", []interface{}{arg1})
	fake.setAffinityMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *K8sDeploymentBuilder) SetAffinityCallCount() int {
	fake.setAffinityMutex.RLock()
	defer fake.setAffinityMutex.RUnlock()
	return len(fake.setAffinityArgsForCall)
}

func (fake *K8sDeploymentBuilder) SetAffinityCalls(stub func(v1.Affinity) k8s.DeploymentBuilder) {
	fake.setAffinityMutex.Lock()
	defer fake.setAffinityMutex.Unlock()
	fake.SetAffinityStub = stub
}

func (fake *K8sDeploymentBuilder) SetAffinityArgsForCall(i int) v1.Affinity {
	fake.setAffinityMutex.RLock()
	defer fake.setAffinityMutex.RUnlock()
	argsForCall := fake.setAffinityArgsForCall[i]
	return argsForCall.arg1
}

func (fake *K8sDeploymentBuilder) SetAffinityReturns(result1 k8s.DeploymentBuilder) {
	fake.setAffinityMutex.Lock()
	defer fake.setAffinityMutex.Unlock()
	fake.SetAffinityStub = nil
	fake.setAffinityReturns = struct {
		result1 k8s.DeploymentBuilder
	}{result1}
}

func (fake *K8sDeploymentBuilder) SetAffinityReturnsOnCall(i int, result1 k8s.DeploymentBuilder) {
	fake.setAffinityMutex.Lock()
	defer fake.setAffinityMutex.Unlock()
	fake.SetAffinityStub = nil
	if fake.setAffinityReturnsOnCall == nil {
		fake.setAffinityReturnsOnCall = make(map[int]struct {
			result1 k8s.DeploymentBuilder
		})
	}
	fake.setAffinityReturnsOnCall[i] = struct {
		result1 k8s.DeploymentBuilder
	}{result1}
}

func (fake *K8sDeploymentBuilder) SetComponent(arg1 string) k8s.DeploymentBuilder {
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

func (fake *K8sDeploymentBuilder) SetComponentCallCount() int {
	fake.setComponentMutex.RLock()
	defer fake.setComponentMutex.RUnlock()
	return len(fake.setComponentArgsForCall)
}

func (fake *K8sDeploymentBuilder) SetComponentCalls(stub func(string) k8s.DeploymentBuilder) {
	fake.setComponentMutex.Lock()
	defer fake.setComponentMutex.Unlock()
	fake.SetComponentStub = stub
}

func (fake *K8sDeploymentBuilder) SetComponentArgsForCall(i int) string {
	fake.setComponentMutex.RLock()
	defer fake.setComponentMutex.RUnlock()
	argsForCall := fake.setComponentArgsForCall[i]
	return argsForCall.arg1
}

func (fake *K8sDeploymentBuilder) SetComponentReturns(result1 k8s.DeploymentBuilder) {
	fake.setComponentMutex.Lock()
	defer fake.setComponentMutex.Unlock()
	fake.SetComponentStub = nil
	fake.setComponentReturns = struct {
		result1 k8s.DeploymentBuilder
	}{result1}
}

func (fake *K8sDeploymentBuilder) SetComponentReturnsOnCall(i int, result1 k8s.DeploymentBuilder) {
	fake.setComponentMutex.Lock()
	defer fake.setComponentMutex.Unlock()
	fake.SetComponentStub = nil
	if fake.setComponentReturnsOnCall == nil {
		fake.setComponentReturnsOnCall = make(map[int]struct {
			result1 k8s.DeploymentBuilder
		})
	}
	fake.setComponentReturnsOnCall[i] = struct {
		result1 k8s.DeploymentBuilder
	}{result1}
}

func (fake *K8sDeploymentBuilder) SetContainers(arg1 []v1.Container) k8s.DeploymentBuilder {
	var arg1Copy []v1.Container
	if arg1 != nil {
		arg1Copy = make([]v1.Container, len(arg1))
		copy(arg1Copy, arg1)
	}
	fake.setContainersMutex.Lock()
	ret, specificReturn := fake.setContainersReturnsOnCall[len(fake.setContainersArgsForCall)]
	fake.setContainersArgsForCall = append(fake.setContainersArgsForCall, struct {
		arg1 []v1.Container
	}{arg1Copy})
	stub := fake.SetContainersStub
	fakeReturns := fake.setContainersReturns
	fake.recordInvocation("SetContainers", []interface{}{arg1Copy})
	fake.setContainersMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *K8sDeploymentBuilder) SetContainersCallCount() int {
	fake.setContainersMutex.RLock()
	defer fake.setContainersMutex.RUnlock()
	return len(fake.setContainersArgsForCall)
}

func (fake *K8sDeploymentBuilder) SetContainersCalls(stub func([]v1.Container) k8s.DeploymentBuilder) {
	fake.setContainersMutex.Lock()
	defer fake.setContainersMutex.Unlock()
	fake.SetContainersStub = stub
}

func (fake *K8sDeploymentBuilder) SetContainersArgsForCall(i int) []v1.Container {
	fake.setContainersMutex.RLock()
	defer fake.setContainersMutex.RUnlock()
	argsForCall := fake.setContainersArgsForCall[i]
	return argsForCall.arg1
}

func (fake *K8sDeploymentBuilder) SetContainersReturns(result1 k8s.DeploymentBuilder) {
	fake.setContainersMutex.Lock()
	defer fake.setContainersMutex.Unlock()
	fake.SetContainersStub = nil
	fake.setContainersReturns = struct {
		result1 k8s.DeploymentBuilder
	}{result1}
}

func (fake *K8sDeploymentBuilder) SetContainersReturnsOnCall(i int, result1 k8s.DeploymentBuilder) {
	fake.setContainersMutex.Lock()
	defer fake.setContainersMutex.Unlock()
	fake.SetContainersStub = nil
	if fake.setContainersReturnsOnCall == nil {
		fake.setContainersReturnsOnCall = make(map[int]struct {
			result1 k8s.DeploymentBuilder
		})
	}
	fake.setContainersReturnsOnCall[i] = struct {
		result1 k8s.DeploymentBuilder
	}{result1}
}

func (fake *K8sDeploymentBuilder) SetContainersBuilder(arg1 k8s.HasBuildContainers) k8s.DeploymentBuilder {
	fake.setContainersBuilderMutex.Lock()
	ret, specificReturn := fake.setContainersBuilderReturnsOnCall[len(fake.setContainersBuilderArgsForCall)]
	fake.setContainersBuilderArgsForCall = append(fake.setContainersBuilderArgsForCall, struct {
		arg1 k8s.HasBuildContainers
	}{arg1})
	stub := fake.SetContainersBuilderStub
	fakeReturns := fake.setContainersBuilderReturns
	fake.recordInvocation("SetContainersBuilder", []interface{}{arg1})
	fake.setContainersBuilderMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *K8sDeploymentBuilder) SetContainersBuilderCallCount() int {
	fake.setContainersBuilderMutex.RLock()
	defer fake.setContainersBuilderMutex.RUnlock()
	return len(fake.setContainersBuilderArgsForCall)
}

func (fake *K8sDeploymentBuilder) SetContainersBuilderCalls(stub func(k8s.HasBuildContainers) k8s.DeploymentBuilder) {
	fake.setContainersBuilderMutex.Lock()
	defer fake.setContainersBuilderMutex.Unlock()
	fake.SetContainersBuilderStub = stub
}

func (fake *K8sDeploymentBuilder) SetContainersBuilderArgsForCall(i int) k8s.HasBuildContainers {
	fake.setContainersBuilderMutex.RLock()
	defer fake.setContainersBuilderMutex.RUnlock()
	argsForCall := fake.setContainersBuilderArgsForCall[i]
	return argsForCall.arg1
}

func (fake *K8sDeploymentBuilder) SetContainersBuilderReturns(result1 k8s.DeploymentBuilder) {
	fake.setContainersBuilderMutex.Lock()
	defer fake.setContainersBuilderMutex.Unlock()
	fake.SetContainersBuilderStub = nil
	fake.setContainersBuilderReturns = struct {
		result1 k8s.DeploymentBuilder
	}{result1}
}

func (fake *K8sDeploymentBuilder) SetContainersBuilderReturnsOnCall(i int, result1 k8s.DeploymentBuilder) {
	fake.setContainersBuilderMutex.Lock()
	defer fake.setContainersBuilderMutex.Unlock()
	fake.SetContainersBuilderStub = nil
	if fake.setContainersBuilderReturnsOnCall == nil {
		fake.setContainersBuilderReturnsOnCall = make(map[int]struct {
			result1 k8s.DeploymentBuilder
		})
	}
	fake.setContainersBuilderReturnsOnCall[i] = struct {
		result1 k8s.DeploymentBuilder
	}{result1}
}

func (fake *K8sDeploymentBuilder) SetImagePullSecrets(arg1 []string) k8s.DeploymentBuilder {
	var arg1Copy []string
	if arg1 != nil {
		arg1Copy = make([]string, len(arg1))
		copy(arg1Copy, arg1)
	}
	fake.setImagePullSecretsMutex.Lock()
	ret, specificReturn := fake.setImagePullSecretsReturnsOnCall[len(fake.setImagePullSecretsArgsForCall)]
	fake.setImagePullSecretsArgsForCall = append(fake.setImagePullSecretsArgsForCall, struct {
		arg1 []string
	}{arg1Copy})
	stub := fake.SetImagePullSecretsStub
	fakeReturns := fake.setImagePullSecretsReturns
	fake.recordInvocation("SetImagePullSecrets", []interface{}{arg1Copy})
	fake.setImagePullSecretsMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *K8sDeploymentBuilder) SetImagePullSecretsCallCount() int {
	fake.setImagePullSecretsMutex.RLock()
	defer fake.setImagePullSecretsMutex.RUnlock()
	return len(fake.setImagePullSecretsArgsForCall)
}

func (fake *K8sDeploymentBuilder) SetImagePullSecretsCalls(stub func([]string) k8s.DeploymentBuilder) {
	fake.setImagePullSecretsMutex.Lock()
	defer fake.setImagePullSecretsMutex.Unlock()
	fake.SetImagePullSecretsStub = stub
}

func (fake *K8sDeploymentBuilder) SetImagePullSecretsArgsForCall(i int) []string {
	fake.setImagePullSecretsMutex.RLock()
	defer fake.setImagePullSecretsMutex.RUnlock()
	argsForCall := fake.setImagePullSecretsArgsForCall[i]
	return argsForCall.arg1
}

func (fake *K8sDeploymentBuilder) SetImagePullSecretsReturns(result1 k8s.DeploymentBuilder) {
	fake.setImagePullSecretsMutex.Lock()
	defer fake.setImagePullSecretsMutex.Unlock()
	fake.SetImagePullSecretsStub = nil
	fake.setImagePullSecretsReturns = struct {
		result1 k8s.DeploymentBuilder
	}{result1}
}

func (fake *K8sDeploymentBuilder) SetImagePullSecretsReturnsOnCall(i int, result1 k8s.DeploymentBuilder) {
	fake.setImagePullSecretsMutex.Lock()
	defer fake.setImagePullSecretsMutex.Unlock()
	fake.SetImagePullSecretsStub = nil
	if fake.setImagePullSecretsReturnsOnCall == nil {
		fake.setImagePullSecretsReturnsOnCall = make(map[int]struct {
			result1 k8s.DeploymentBuilder
		})
	}
	fake.setImagePullSecretsReturnsOnCall[i] = struct {
		result1 k8s.DeploymentBuilder
	}{result1}
}

func (fake *K8sDeploymentBuilder) SetName(arg1 k8s.Name) k8s.DeploymentBuilder {
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

func (fake *K8sDeploymentBuilder) SetNameCallCount() int {
	fake.setNameMutex.RLock()
	defer fake.setNameMutex.RUnlock()
	return len(fake.setNameArgsForCall)
}

func (fake *K8sDeploymentBuilder) SetNameCalls(stub func(k8s.Name) k8s.DeploymentBuilder) {
	fake.setNameMutex.Lock()
	defer fake.setNameMutex.Unlock()
	fake.SetNameStub = stub
}

func (fake *K8sDeploymentBuilder) SetNameArgsForCall(i int) k8s.Name {
	fake.setNameMutex.RLock()
	defer fake.setNameMutex.RUnlock()
	argsForCall := fake.setNameArgsForCall[i]
	return argsForCall.arg1
}

func (fake *K8sDeploymentBuilder) SetNameReturns(result1 k8s.DeploymentBuilder) {
	fake.setNameMutex.Lock()
	defer fake.setNameMutex.Unlock()
	fake.SetNameStub = nil
	fake.setNameReturns = struct {
		result1 k8s.DeploymentBuilder
	}{result1}
}

func (fake *K8sDeploymentBuilder) SetNameReturnsOnCall(i int, result1 k8s.DeploymentBuilder) {
	fake.setNameMutex.Lock()
	defer fake.setNameMutex.Unlock()
	fake.SetNameStub = nil
	if fake.setNameReturnsOnCall == nil {
		fake.setNameReturnsOnCall = make(map[int]struct {
			result1 k8s.DeploymentBuilder
		})
	}
	fake.setNameReturnsOnCall[i] = struct {
		result1 k8s.DeploymentBuilder
	}{result1}
}

func (fake *K8sDeploymentBuilder) SetObjectMeta(arg1 v1b.ObjectMeta) k8s.DeploymentBuilder {
	fake.setObjectMetaMutex.Lock()
	ret, specificReturn := fake.setObjectMetaReturnsOnCall[len(fake.setObjectMetaArgsForCall)]
	fake.setObjectMetaArgsForCall = append(fake.setObjectMetaArgsForCall, struct {
		arg1 v1b.ObjectMeta
	}{arg1})
	stub := fake.SetObjectMetaStub
	fakeReturns := fake.setObjectMetaReturns
	fake.recordInvocation("SetObjectMeta", []interface{}{arg1})
	fake.setObjectMetaMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *K8sDeploymentBuilder) SetObjectMetaCallCount() int {
	fake.setObjectMetaMutex.RLock()
	defer fake.setObjectMetaMutex.RUnlock()
	return len(fake.setObjectMetaArgsForCall)
}

func (fake *K8sDeploymentBuilder) SetObjectMetaCalls(stub func(v1b.ObjectMeta) k8s.DeploymentBuilder) {
	fake.setObjectMetaMutex.Lock()
	defer fake.setObjectMetaMutex.Unlock()
	fake.SetObjectMetaStub = stub
}

func (fake *K8sDeploymentBuilder) SetObjectMetaArgsForCall(i int) v1b.ObjectMeta {
	fake.setObjectMetaMutex.RLock()
	defer fake.setObjectMetaMutex.RUnlock()
	argsForCall := fake.setObjectMetaArgsForCall[i]
	return argsForCall.arg1
}

func (fake *K8sDeploymentBuilder) SetObjectMetaReturns(result1 k8s.DeploymentBuilder) {
	fake.setObjectMetaMutex.Lock()
	defer fake.setObjectMetaMutex.Unlock()
	fake.SetObjectMetaStub = nil
	fake.setObjectMetaReturns = struct {
		result1 k8s.DeploymentBuilder
	}{result1}
}

func (fake *K8sDeploymentBuilder) SetObjectMetaReturnsOnCall(i int, result1 k8s.DeploymentBuilder) {
	fake.setObjectMetaMutex.Lock()
	defer fake.setObjectMetaMutex.Unlock()
	fake.SetObjectMetaStub = nil
	if fake.setObjectMetaReturnsOnCall == nil {
		fake.setObjectMetaReturnsOnCall = make(map[int]struct {
			result1 k8s.DeploymentBuilder
		})
	}
	fake.setObjectMetaReturnsOnCall[i] = struct {
		result1 k8s.DeploymentBuilder
	}{result1}
}

func (fake *K8sDeploymentBuilder) SetObjectMetaBuilder(arg1 k8s.HasBuildObjectMeta) k8s.DeploymentBuilder {
	fake.setObjectMetaBuilderMutex.Lock()
	ret, specificReturn := fake.setObjectMetaBuilderReturnsOnCall[len(fake.setObjectMetaBuilderArgsForCall)]
	fake.setObjectMetaBuilderArgsForCall = append(fake.setObjectMetaBuilderArgsForCall, struct {
		arg1 k8s.HasBuildObjectMeta
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

func (fake *K8sDeploymentBuilder) SetObjectMetaBuilderCallCount() int {
	fake.setObjectMetaBuilderMutex.RLock()
	defer fake.setObjectMetaBuilderMutex.RUnlock()
	return len(fake.setObjectMetaBuilderArgsForCall)
}

func (fake *K8sDeploymentBuilder) SetObjectMetaBuilderCalls(stub func(k8s.HasBuildObjectMeta) k8s.DeploymentBuilder) {
	fake.setObjectMetaBuilderMutex.Lock()
	defer fake.setObjectMetaBuilderMutex.Unlock()
	fake.SetObjectMetaBuilderStub = stub
}

func (fake *K8sDeploymentBuilder) SetObjectMetaBuilderArgsForCall(i int) k8s.HasBuildObjectMeta {
	fake.setObjectMetaBuilderMutex.RLock()
	defer fake.setObjectMetaBuilderMutex.RUnlock()
	argsForCall := fake.setObjectMetaBuilderArgsForCall[i]
	return argsForCall.arg1
}

func (fake *K8sDeploymentBuilder) SetObjectMetaBuilderReturns(result1 k8s.DeploymentBuilder) {
	fake.setObjectMetaBuilderMutex.Lock()
	defer fake.setObjectMetaBuilderMutex.Unlock()
	fake.SetObjectMetaBuilderStub = nil
	fake.setObjectMetaBuilderReturns = struct {
		result1 k8s.DeploymentBuilder
	}{result1}
}

func (fake *K8sDeploymentBuilder) SetObjectMetaBuilderReturnsOnCall(i int, result1 k8s.DeploymentBuilder) {
	fake.setObjectMetaBuilderMutex.Lock()
	defer fake.setObjectMetaBuilderMutex.Unlock()
	fake.SetObjectMetaBuilderStub = nil
	if fake.setObjectMetaBuilderReturnsOnCall == nil {
		fake.setObjectMetaBuilderReturnsOnCall = make(map[int]struct {
			result1 k8s.DeploymentBuilder
		})
	}
	fake.setObjectMetaBuilderReturnsOnCall[i] = struct {
		result1 k8s.DeploymentBuilder
	}{result1}
}

func (fake *K8sDeploymentBuilder) SetReplicas(arg1 int32) k8s.DeploymentBuilder {
	fake.setReplicasMutex.Lock()
	ret, specificReturn := fake.setReplicasReturnsOnCall[len(fake.setReplicasArgsForCall)]
	fake.setReplicasArgsForCall = append(fake.setReplicasArgsForCall, struct {
		arg1 int32
	}{arg1})
	stub := fake.SetReplicasStub
	fakeReturns := fake.setReplicasReturns
	fake.recordInvocation("SetReplicas", []interface{}{arg1})
	fake.setReplicasMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *K8sDeploymentBuilder) SetReplicasCallCount() int {
	fake.setReplicasMutex.RLock()
	defer fake.setReplicasMutex.RUnlock()
	return len(fake.setReplicasArgsForCall)
}

func (fake *K8sDeploymentBuilder) SetReplicasCalls(stub func(int32) k8s.DeploymentBuilder) {
	fake.setReplicasMutex.Lock()
	defer fake.setReplicasMutex.Unlock()
	fake.SetReplicasStub = stub
}

func (fake *K8sDeploymentBuilder) SetReplicasArgsForCall(i int) int32 {
	fake.setReplicasMutex.RLock()
	defer fake.setReplicasMutex.RUnlock()
	argsForCall := fake.setReplicasArgsForCall[i]
	return argsForCall.arg1
}

func (fake *K8sDeploymentBuilder) SetReplicasReturns(result1 k8s.DeploymentBuilder) {
	fake.setReplicasMutex.Lock()
	defer fake.setReplicasMutex.Unlock()
	fake.SetReplicasStub = nil
	fake.setReplicasReturns = struct {
		result1 k8s.DeploymentBuilder
	}{result1}
}

func (fake *K8sDeploymentBuilder) SetReplicasReturnsOnCall(i int, result1 k8s.DeploymentBuilder) {
	fake.setReplicasMutex.Lock()
	defer fake.setReplicasMutex.Unlock()
	fake.SetReplicasStub = nil
	if fake.setReplicasReturnsOnCall == nil {
		fake.setReplicasReturnsOnCall = make(map[int]struct {
			result1 k8s.DeploymentBuilder
		})
	}
	fake.setReplicasReturnsOnCall[i] = struct {
		result1 k8s.DeploymentBuilder
	}{result1}
}

func (fake *K8sDeploymentBuilder) SetServiceAccountName(arg1 string) k8s.DeploymentBuilder {
	fake.setServiceAccountNameMutex.Lock()
	ret, specificReturn := fake.setServiceAccountNameReturnsOnCall[len(fake.setServiceAccountNameArgsForCall)]
	fake.setServiceAccountNameArgsForCall = append(fake.setServiceAccountNameArgsForCall, struct {
		arg1 string
	}{arg1})
	stub := fake.SetServiceAccountNameStub
	fakeReturns := fake.setServiceAccountNameReturns
	fake.recordInvocation("SetServiceAccountName", []interface{}{arg1})
	fake.setServiceAccountNameMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *K8sDeploymentBuilder) SetServiceAccountNameCallCount() int {
	fake.setServiceAccountNameMutex.RLock()
	defer fake.setServiceAccountNameMutex.RUnlock()
	return len(fake.setServiceAccountNameArgsForCall)
}

func (fake *K8sDeploymentBuilder) SetServiceAccountNameCalls(stub func(string) k8s.DeploymentBuilder) {
	fake.setServiceAccountNameMutex.Lock()
	defer fake.setServiceAccountNameMutex.Unlock()
	fake.SetServiceAccountNameStub = stub
}

func (fake *K8sDeploymentBuilder) SetServiceAccountNameArgsForCall(i int) string {
	fake.setServiceAccountNameMutex.RLock()
	defer fake.setServiceAccountNameMutex.RUnlock()
	argsForCall := fake.setServiceAccountNameArgsForCall[i]
	return argsForCall.arg1
}

func (fake *K8sDeploymentBuilder) SetServiceAccountNameReturns(result1 k8s.DeploymentBuilder) {
	fake.setServiceAccountNameMutex.Lock()
	defer fake.setServiceAccountNameMutex.Unlock()
	fake.SetServiceAccountNameStub = nil
	fake.setServiceAccountNameReturns = struct {
		result1 k8s.DeploymentBuilder
	}{result1}
}

func (fake *K8sDeploymentBuilder) SetServiceAccountNameReturnsOnCall(i int, result1 k8s.DeploymentBuilder) {
	fake.setServiceAccountNameMutex.Lock()
	defer fake.setServiceAccountNameMutex.Unlock()
	fake.SetServiceAccountNameStub = nil
	if fake.setServiceAccountNameReturnsOnCall == nil {
		fake.setServiceAccountNameReturnsOnCall = make(map[int]struct {
			result1 k8s.DeploymentBuilder
		})
	}
	fake.setServiceAccountNameReturnsOnCall[i] = struct {
		result1 k8s.DeploymentBuilder
	}{result1}
}

func (fake *K8sDeploymentBuilder) SetVolumes(arg1 []v1.Volume) k8s.DeploymentBuilder {
	var arg1Copy []v1.Volume
	if arg1 != nil {
		arg1Copy = make([]v1.Volume, len(arg1))
		copy(arg1Copy, arg1)
	}
	fake.setVolumesMutex.Lock()
	ret, specificReturn := fake.setVolumesReturnsOnCall[len(fake.setVolumesArgsForCall)]
	fake.setVolumesArgsForCall = append(fake.setVolumesArgsForCall, struct {
		arg1 []v1.Volume
	}{arg1Copy})
	stub := fake.SetVolumesStub
	fakeReturns := fake.setVolumesReturns
	fake.recordInvocation("SetVolumes", []interface{}{arg1Copy})
	fake.setVolumesMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *K8sDeploymentBuilder) SetVolumesCallCount() int {
	fake.setVolumesMutex.RLock()
	defer fake.setVolumesMutex.RUnlock()
	return len(fake.setVolumesArgsForCall)
}

func (fake *K8sDeploymentBuilder) SetVolumesCalls(stub func([]v1.Volume) k8s.DeploymentBuilder) {
	fake.setVolumesMutex.Lock()
	defer fake.setVolumesMutex.Unlock()
	fake.SetVolumesStub = stub
}

func (fake *K8sDeploymentBuilder) SetVolumesArgsForCall(i int) []v1.Volume {
	fake.setVolumesMutex.RLock()
	defer fake.setVolumesMutex.RUnlock()
	argsForCall := fake.setVolumesArgsForCall[i]
	return argsForCall.arg1
}

func (fake *K8sDeploymentBuilder) SetVolumesReturns(result1 k8s.DeploymentBuilder) {
	fake.setVolumesMutex.Lock()
	defer fake.setVolumesMutex.Unlock()
	fake.SetVolumesStub = nil
	fake.setVolumesReturns = struct {
		result1 k8s.DeploymentBuilder
	}{result1}
}

func (fake *K8sDeploymentBuilder) SetVolumesReturnsOnCall(i int, result1 k8s.DeploymentBuilder) {
	fake.setVolumesMutex.Lock()
	defer fake.setVolumesMutex.Unlock()
	fake.SetVolumesStub = nil
	if fake.setVolumesReturnsOnCall == nil {
		fake.setVolumesReturnsOnCall = make(map[int]struct {
			result1 k8s.DeploymentBuilder
		})
	}
	fake.setVolumesReturnsOnCall[i] = struct {
		result1 k8s.DeploymentBuilder
	}{result1}
}

func (fake *K8sDeploymentBuilder) Validate(arg1 context.Context) error {
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

func (fake *K8sDeploymentBuilder) ValidateCallCount() int {
	fake.validateMutex.RLock()
	defer fake.validateMutex.RUnlock()
	return len(fake.validateArgsForCall)
}

func (fake *K8sDeploymentBuilder) ValidateCalls(stub func(context.Context) error) {
	fake.validateMutex.Lock()
	defer fake.validateMutex.Unlock()
	fake.ValidateStub = stub
}

func (fake *K8sDeploymentBuilder) ValidateArgsForCall(i int) context.Context {
	fake.validateMutex.RLock()
	defer fake.validateMutex.RUnlock()
	argsForCall := fake.validateArgsForCall[i]
	return argsForCall.arg1
}

func (fake *K8sDeploymentBuilder) ValidateReturns(result1 error) {
	fake.validateMutex.Lock()
	defer fake.validateMutex.Unlock()
	fake.ValidateStub = nil
	fake.validateReturns = struct {
		result1 error
	}{result1}
}

func (fake *K8sDeploymentBuilder) ValidateReturnsOnCall(i int, result1 error) {
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

func (fake *K8sDeploymentBuilder) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *K8sDeploymentBuilder) recordInvocation(key string, args []interface{}) {
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

var _ k8s.DeploymentBuilder = new(K8sDeploymentBuilder)
