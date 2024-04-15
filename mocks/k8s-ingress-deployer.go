// Code generated by counterfeiter. DO NOT EDIT.
package mocks

import (
	"context"
	"sync"

	"github.com/bborbe/k8s"
	v1 "k8s.io/api/networking/v1"
)

type K8sIngressDeployer struct {
	DeployStub        func(context.Context, v1.Ingress) error
	deployMutex       sync.RWMutex
	deployArgsForCall []struct {
		arg1 context.Context
		arg2 v1.Ingress
	}
	deployReturns struct {
		result1 error
	}
	deployReturnsOnCall map[int]struct {
		result1 error
	}
	UndeployStub        func(context.Context, k8s.Namespace, string) error
	undeployMutex       sync.RWMutex
	undeployArgsForCall []struct {
		arg1 context.Context
		arg2 k8s.Namespace
		arg3 string
	}
	undeployReturns struct {
		result1 error
	}
	undeployReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *K8sIngressDeployer) Deploy(arg1 context.Context, arg2 v1.Ingress) error {
	fake.deployMutex.Lock()
	ret, specificReturn := fake.deployReturnsOnCall[len(fake.deployArgsForCall)]
	fake.deployArgsForCall = append(fake.deployArgsForCall, struct {
		arg1 context.Context
		arg2 v1.Ingress
	}{arg1, arg2})
	stub := fake.DeployStub
	fakeReturns := fake.deployReturns
	fake.recordInvocation("Deploy", []interface{}{arg1, arg2})
	fake.deployMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *K8sIngressDeployer) DeployCallCount() int {
	fake.deployMutex.RLock()
	defer fake.deployMutex.RUnlock()
	return len(fake.deployArgsForCall)
}

func (fake *K8sIngressDeployer) DeployCalls(stub func(context.Context, v1.Ingress) error) {
	fake.deployMutex.Lock()
	defer fake.deployMutex.Unlock()
	fake.DeployStub = stub
}

func (fake *K8sIngressDeployer) DeployArgsForCall(i int) (context.Context, v1.Ingress) {
	fake.deployMutex.RLock()
	defer fake.deployMutex.RUnlock()
	argsForCall := fake.deployArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *K8sIngressDeployer) DeployReturns(result1 error) {
	fake.deployMutex.Lock()
	defer fake.deployMutex.Unlock()
	fake.DeployStub = nil
	fake.deployReturns = struct {
		result1 error
	}{result1}
}

func (fake *K8sIngressDeployer) DeployReturnsOnCall(i int, result1 error) {
	fake.deployMutex.Lock()
	defer fake.deployMutex.Unlock()
	fake.DeployStub = nil
	if fake.deployReturnsOnCall == nil {
		fake.deployReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.deployReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *K8sIngressDeployer) Undeploy(arg1 context.Context, arg2 k8s.Namespace, arg3 string) error {
	fake.undeployMutex.Lock()
	ret, specificReturn := fake.undeployReturnsOnCall[len(fake.undeployArgsForCall)]
	fake.undeployArgsForCall = append(fake.undeployArgsForCall, struct {
		arg1 context.Context
		arg2 k8s.Namespace
		arg3 string
	}{arg1, arg2, arg3})
	stub := fake.UndeployStub
	fakeReturns := fake.undeployReturns
	fake.recordInvocation("Undeploy", []interface{}{arg1, arg2, arg3})
	fake.undeployMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *K8sIngressDeployer) UndeployCallCount() int {
	fake.undeployMutex.RLock()
	defer fake.undeployMutex.RUnlock()
	return len(fake.undeployArgsForCall)
}

func (fake *K8sIngressDeployer) UndeployCalls(stub func(context.Context, k8s.Namespace, string) error) {
	fake.undeployMutex.Lock()
	defer fake.undeployMutex.Unlock()
	fake.UndeployStub = stub
}

func (fake *K8sIngressDeployer) UndeployArgsForCall(i int) (context.Context, k8s.Namespace, string) {
	fake.undeployMutex.RLock()
	defer fake.undeployMutex.RUnlock()
	argsForCall := fake.undeployArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *K8sIngressDeployer) UndeployReturns(result1 error) {
	fake.undeployMutex.Lock()
	defer fake.undeployMutex.Unlock()
	fake.UndeployStub = nil
	fake.undeployReturns = struct {
		result1 error
	}{result1}
}

func (fake *K8sIngressDeployer) UndeployReturnsOnCall(i int, result1 error) {
	fake.undeployMutex.Lock()
	defer fake.undeployMutex.Unlock()
	fake.UndeployStub = nil
	if fake.undeployReturnsOnCall == nil {
		fake.undeployReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.undeployReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *K8sIngressDeployer) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.deployMutex.RLock()
	defer fake.deployMutex.RUnlock()
	fake.undeployMutex.RLock()
	defer fake.undeployMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *K8sIngressDeployer) recordInvocation(key string, args []interface{}) {
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

var _ k8s.IngressDeployer = new(K8sIngressDeployer)
