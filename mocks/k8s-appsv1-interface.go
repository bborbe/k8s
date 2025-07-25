// Code generated by counterfeiter. DO NOT EDIT.
package mocks

import (
	"sync"

	"github.com/bborbe/k8s"
	v1 "k8s.io/client-go/kubernetes/typed/apps/v1"
	"k8s.io/client-go/rest"
)

type K8sAppsV1Interface struct {
	ControllerRevisionsStub        func(string) v1.ControllerRevisionInterface
	controllerRevisionsMutex       sync.RWMutex
	controllerRevisionsArgsForCall []struct {
		arg1 string
	}
	controllerRevisionsReturns struct {
		result1 v1.ControllerRevisionInterface
	}
	controllerRevisionsReturnsOnCall map[int]struct {
		result1 v1.ControllerRevisionInterface
	}
	DaemonSetsStub        func(string) v1.DaemonSetInterface
	daemonSetsMutex       sync.RWMutex
	daemonSetsArgsForCall []struct {
		arg1 string
	}
	daemonSetsReturns struct {
		result1 v1.DaemonSetInterface
	}
	daemonSetsReturnsOnCall map[int]struct {
		result1 v1.DaemonSetInterface
	}
	DeploymentsStub        func(string) v1.DeploymentInterface
	deploymentsMutex       sync.RWMutex
	deploymentsArgsForCall []struct {
		arg1 string
	}
	deploymentsReturns struct {
		result1 v1.DeploymentInterface
	}
	deploymentsReturnsOnCall map[int]struct {
		result1 v1.DeploymentInterface
	}
	RESTClientStub        func() rest.Interface
	rESTClientMutex       sync.RWMutex
	rESTClientArgsForCall []struct {
	}
	rESTClientReturns struct {
		result1 rest.Interface
	}
	rESTClientReturnsOnCall map[int]struct {
		result1 rest.Interface
	}
	ReplicaSetsStub        func(string) v1.ReplicaSetInterface
	replicaSetsMutex       sync.RWMutex
	replicaSetsArgsForCall []struct {
		arg1 string
	}
	replicaSetsReturns struct {
		result1 v1.ReplicaSetInterface
	}
	replicaSetsReturnsOnCall map[int]struct {
		result1 v1.ReplicaSetInterface
	}
	StatefulSetsStub        func(string) v1.StatefulSetInterface
	statefulSetsMutex       sync.RWMutex
	statefulSetsArgsForCall []struct {
		arg1 string
	}
	statefulSetsReturns struct {
		result1 v1.StatefulSetInterface
	}
	statefulSetsReturnsOnCall map[int]struct {
		result1 v1.StatefulSetInterface
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *K8sAppsV1Interface) ControllerRevisions(arg1 string) v1.ControllerRevisionInterface {
	fake.controllerRevisionsMutex.Lock()
	ret, specificReturn := fake.controllerRevisionsReturnsOnCall[len(fake.controllerRevisionsArgsForCall)]
	fake.controllerRevisionsArgsForCall = append(fake.controllerRevisionsArgsForCall, struct {
		arg1 string
	}{arg1})
	stub := fake.ControllerRevisionsStub
	fakeReturns := fake.controllerRevisionsReturns
	fake.recordInvocation("ControllerRevisions", []interface{}{arg1})
	fake.controllerRevisionsMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *K8sAppsV1Interface) ControllerRevisionsCallCount() int {
	fake.controllerRevisionsMutex.RLock()
	defer fake.controllerRevisionsMutex.RUnlock()
	return len(fake.controllerRevisionsArgsForCall)
}

func (fake *K8sAppsV1Interface) ControllerRevisionsCalls(stub func(string) v1.ControllerRevisionInterface) {
	fake.controllerRevisionsMutex.Lock()
	defer fake.controllerRevisionsMutex.Unlock()
	fake.ControllerRevisionsStub = stub
}

func (fake *K8sAppsV1Interface) ControllerRevisionsArgsForCall(i int) string {
	fake.controllerRevisionsMutex.RLock()
	defer fake.controllerRevisionsMutex.RUnlock()
	argsForCall := fake.controllerRevisionsArgsForCall[i]
	return argsForCall.arg1
}

func (fake *K8sAppsV1Interface) ControllerRevisionsReturns(result1 v1.ControllerRevisionInterface) {
	fake.controllerRevisionsMutex.Lock()
	defer fake.controllerRevisionsMutex.Unlock()
	fake.ControllerRevisionsStub = nil
	fake.controllerRevisionsReturns = struct {
		result1 v1.ControllerRevisionInterface
	}{result1}
}

func (fake *K8sAppsV1Interface) ControllerRevisionsReturnsOnCall(i int, result1 v1.ControllerRevisionInterface) {
	fake.controllerRevisionsMutex.Lock()
	defer fake.controllerRevisionsMutex.Unlock()
	fake.ControllerRevisionsStub = nil
	if fake.controllerRevisionsReturnsOnCall == nil {
		fake.controllerRevisionsReturnsOnCall = make(map[int]struct {
			result1 v1.ControllerRevisionInterface
		})
	}
	fake.controllerRevisionsReturnsOnCall[i] = struct {
		result1 v1.ControllerRevisionInterface
	}{result1}
}

func (fake *K8sAppsV1Interface) DaemonSets(arg1 string) v1.DaemonSetInterface {
	fake.daemonSetsMutex.Lock()
	ret, specificReturn := fake.daemonSetsReturnsOnCall[len(fake.daemonSetsArgsForCall)]
	fake.daemonSetsArgsForCall = append(fake.daemonSetsArgsForCall, struct {
		arg1 string
	}{arg1})
	stub := fake.DaemonSetsStub
	fakeReturns := fake.daemonSetsReturns
	fake.recordInvocation("DaemonSets", []interface{}{arg1})
	fake.daemonSetsMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *K8sAppsV1Interface) DaemonSetsCallCount() int {
	fake.daemonSetsMutex.RLock()
	defer fake.daemonSetsMutex.RUnlock()
	return len(fake.daemonSetsArgsForCall)
}

func (fake *K8sAppsV1Interface) DaemonSetsCalls(stub func(string) v1.DaemonSetInterface) {
	fake.daemonSetsMutex.Lock()
	defer fake.daemonSetsMutex.Unlock()
	fake.DaemonSetsStub = stub
}

func (fake *K8sAppsV1Interface) DaemonSetsArgsForCall(i int) string {
	fake.daemonSetsMutex.RLock()
	defer fake.daemonSetsMutex.RUnlock()
	argsForCall := fake.daemonSetsArgsForCall[i]
	return argsForCall.arg1
}

func (fake *K8sAppsV1Interface) DaemonSetsReturns(result1 v1.DaemonSetInterface) {
	fake.daemonSetsMutex.Lock()
	defer fake.daemonSetsMutex.Unlock()
	fake.DaemonSetsStub = nil
	fake.daemonSetsReturns = struct {
		result1 v1.DaemonSetInterface
	}{result1}
}

func (fake *K8sAppsV1Interface) DaemonSetsReturnsOnCall(i int, result1 v1.DaemonSetInterface) {
	fake.daemonSetsMutex.Lock()
	defer fake.daemonSetsMutex.Unlock()
	fake.DaemonSetsStub = nil
	if fake.daemonSetsReturnsOnCall == nil {
		fake.daemonSetsReturnsOnCall = make(map[int]struct {
			result1 v1.DaemonSetInterface
		})
	}
	fake.daemonSetsReturnsOnCall[i] = struct {
		result1 v1.DaemonSetInterface
	}{result1}
}

func (fake *K8sAppsV1Interface) Deployments(arg1 string) v1.DeploymentInterface {
	fake.deploymentsMutex.Lock()
	ret, specificReturn := fake.deploymentsReturnsOnCall[len(fake.deploymentsArgsForCall)]
	fake.deploymentsArgsForCall = append(fake.deploymentsArgsForCall, struct {
		arg1 string
	}{arg1})
	stub := fake.DeploymentsStub
	fakeReturns := fake.deploymentsReturns
	fake.recordInvocation("Deployments", []interface{}{arg1})
	fake.deploymentsMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *K8sAppsV1Interface) DeploymentsCallCount() int {
	fake.deploymentsMutex.RLock()
	defer fake.deploymentsMutex.RUnlock()
	return len(fake.deploymentsArgsForCall)
}

func (fake *K8sAppsV1Interface) DeploymentsCalls(stub func(string) v1.DeploymentInterface) {
	fake.deploymentsMutex.Lock()
	defer fake.deploymentsMutex.Unlock()
	fake.DeploymentsStub = stub
}

func (fake *K8sAppsV1Interface) DeploymentsArgsForCall(i int) string {
	fake.deploymentsMutex.RLock()
	defer fake.deploymentsMutex.RUnlock()
	argsForCall := fake.deploymentsArgsForCall[i]
	return argsForCall.arg1
}

func (fake *K8sAppsV1Interface) DeploymentsReturns(result1 v1.DeploymentInterface) {
	fake.deploymentsMutex.Lock()
	defer fake.deploymentsMutex.Unlock()
	fake.DeploymentsStub = nil
	fake.deploymentsReturns = struct {
		result1 v1.DeploymentInterface
	}{result1}
}

func (fake *K8sAppsV1Interface) DeploymentsReturnsOnCall(i int, result1 v1.DeploymentInterface) {
	fake.deploymentsMutex.Lock()
	defer fake.deploymentsMutex.Unlock()
	fake.DeploymentsStub = nil
	if fake.deploymentsReturnsOnCall == nil {
		fake.deploymentsReturnsOnCall = make(map[int]struct {
			result1 v1.DeploymentInterface
		})
	}
	fake.deploymentsReturnsOnCall[i] = struct {
		result1 v1.DeploymentInterface
	}{result1}
}

func (fake *K8sAppsV1Interface) RESTClient() rest.Interface {
	fake.rESTClientMutex.Lock()
	ret, specificReturn := fake.rESTClientReturnsOnCall[len(fake.rESTClientArgsForCall)]
	fake.rESTClientArgsForCall = append(fake.rESTClientArgsForCall, struct {
	}{})
	stub := fake.RESTClientStub
	fakeReturns := fake.rESTClientReturns
	fake.recordInvocation("RESTClient", []interface{}{})
	fake.rESTClientMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *K8sAppsV1Interface) RESTClientCallCount() int {
	fake.rESTClientMutex.RLock()
	defer fake.rESTClientMutex.RUnlock()
	return len(fake.rESTClientArgsForCall)
}

func (fake *K8sAppsV1Interface) RESTClientCalls(stub func() rest.Interface) {
	fake.rESTClientMutex.Lock()
	defer fake.rESTClientMutex.Unlock()
	fake.RESTClientStub = stub
}

func (fake *K8sAppsV1Interface) RESTClientReturns(result1 rest.Interface) {
	fake.rESTClientMutex.Lock()
	defer fake.rESTClientMutex.Unlock()
	fake.RESTClientStub = nil
	fake.rESTClientReturns = struct {
		result1 rest.Interface
	}{result1}
}

func (fake *K8sAppsV1Interface) RESTClientReturnsOnCall(i int, result1 rest.Interface) {
	fake.rESTClientMutex.Lock()
	defer fake.rESTClientMutex.Unlock()
	fake.RESTClientStub = nil
	if fake.rESTClientReturnsOnCall == nil {
		fake.rESTClientReturnsOnCall = make(map[int]struct {
			result1 rest.Interface
		})
	}
	fake.rESTClientReturnsOnCall[i] = struct {
		result1 rest.Interface
	}{result1}
}

func (fake *K8sAppsV1Interface) ReplicaSets(arg1 string) v1.ReplicaSetInterface {
	fake.replicaSetsMutex.Lock()
	ret, specificReturn := fake.replicaSetsReturnsOnCall[len(fake.replicaSetsArgsForCall)]
	fake.replicaSetsArgsForCall = append(fake.replicaSetsArgsForCall, struct {
		arg1 string
	}{arg1})
	stub := fake.ReplicaSetsStub
	fakeReturns := fake.replicaSetsReturns
	fake.recordInvocation("ReplicaSets", []interface{}{arg1})
	fake.replicaSetsMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *K8sAppsV1Interface) ReplicaSetsCallCount() int {
	fake.replicaSetsMutex.RLock()
	defer fake.replicaSetsMutex.RUnlock()
	return len(fake.replicaSetsArgsForCall)
}

func (fake *K8sAppsV1Interface) ReplicaSetsCalls(stub func(string) v1.ReplicaSetInterface) {
	fake.replicaSetsMutex.Lock()
	defer fake.replicaSetsMutex.Unlock()
	fake.ReplicaSetsStub = stub
}

func (fake *K8sAppsV1Interface) ReplicaSetsArgsForCall(i int) string {
	fake.replicaSetsMutex.RLock()
	defer fake.replicaSetsMutex.RUnlock()
	argsForCall := fake.replicaSetsArgsForCall[i]
	return argsForCall.arg1
}

func (fake *K8sAppsV1Interface) ReplicaSetsReturns(result1 v1.ReplicaSetInterface) {
	fake.replicaSetsMutex.Lock()
	defer fake.replicaSetsMutex.Unlock()
	fake.ReplicaSetsStub = nil
	fake.replicaSetsReturns = struct {
		result1 v1.ReplicaSetInterface
	}{result1}
}

func (fake *K8sAppsV1Interface) ReplicaSetsReturnsOnCall(i int, result1 v1.ReplicaSetInterface) {
	fake.replicaSetsMutex.Lock()
	defer fake.replicaSetsMutex.Unlock()
	fake.ReplicaSetsStub = nil
	if fake.replicaSetsReturnsOnCall == nil {
		fake.replicaSetsReturnsOnCall = make(map[int]struct {
			result1 v1.ReplicaSetInterface
		})
	}
	fake.replicaSetsReturnsOnCall[i] = struct {
		result1 v1.ReplicaSetInterface
	}{result1}
}

func (fake *K8sAppsV1Interface) StatefulSets(arg1 string) v1.StatefulSetInterface {
	fake.statefulSetsMutex.Lock()
	ret, specificReturn := fake.statefulSetsReturnsOnCall[len(fake.statefulSetsArgsForCall)]
	fake.statefulSetsArgsForCall = append(fake.statefulSetsArgsForCall, struct {
		arg1 string
	}{arg1})
	stub := fake.StatefulSetsStub
	fakeReturns := fake.statefulSetsReturns
	fake.recordInvocation("StatefulSets", []interface{}{arg1})
	fake.statefulSetsMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *K8sAppsV1Interface) StatefulSetsCallCount() int {
	fake.statefulSetsMutex.RLock()
	defer fake.statefulSetsMutex.RUnlock()
	return len(fake.statefulSetsArgsForCall)
}

func (fake *K8sAppsV1Interface) StatefulSetsCalls(stub func(string) v1.StatefulSetInterface) {
	fake.statefulSetsMutex.Lock()
	defer fake.statefulSetsMutex.Unlock()
	fake.StatefulSetsStub = stub
}

func (fake *K8sAppsV1Interface) StatefulSetsArgsForCall(i int) string {
	fake.statefulSetsMutex.RLock()
	defer fake.statefulSetsMutex.RUnlock()
	argsForCall := fake.statefulSetsArgsForCall[i]
	return argsForCall.arg1
}

func (fake *K8sAppsV1Interface) StatefulSetsReturns(result1 v1.StatefulSetInterface) {
	fake.statefulSetsMutex.Lock()
	defer fake.statefulSetsMutex.Unlock()
	fake.StatefulSetsStub = nil
	fake.statefulSetsReturns = struct {
		result1 v1.StatefulSetInterface
	}{result1}
}

func (fake *K8sAppsV1Interface) StatefulSetsReturnsOnCall(i int, result1 v1.StatefulSetInterface) {
	fake.statefulSetsMutex.Lock()
	defer fake.statefulSetsMutex.Unlock()
	fake.StatefulSetsStub = nil
	if fake.statefulSetsReturnsOnCall == nil {
		fake.statefulSetsReturnsOnCall = make(map[int]struct {
			result1 v1.StatefulSetInterface
		})
	}
	fake.statefulSetsReturnsOnCall[i] = struct {
		result1 v1.StatefulSetInterface
	}{result1}
}

func (fake *K8sAppsV1Interface) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *K8sAppsV1Interface) recordInvocation(key string, args []interface{}) {
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

var _ k8s.AppsV1Interface = new(K8sAppsV1Interface)
