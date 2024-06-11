// Code generated by counterfeiter. DO NOT EDIT.
package mocks

import (
	"context"
	"sync"

	"github.com/bborbe/k8s"
	v1 "k8s.io/api/networking/v1"
	v1b "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/watch"
	v1a "k8s.io/client-go/applyconfigurations/networking/v1"
)

type K8sIngressInterface struct {
	ApplyStub        func(context.Context, *v1a.IngressApplyConfiguration, v1b.ApplyOptions) (*v1.Ingress, error)
	applyMutex       sync.RWMutex
	applyArgsForCall []struct {
		arg1 context.Context
		arg2 *v1a.IngressApplyConfiguration
		arg3 v1b.ApplyOptions
	}
	applyReturns struct {
		result1 *v1.Ingress
		result2 error
	}
	applyReturnsOnCall map[int]struct {
		result1 *v1.Ingress
		result2 error
	}
	ApplyStatusStub        func(context.Context, *v1a.IngressApplyConfiguration, v1b.ApplyOptions) (*v1.Ingress, error)
	applyStatusMutex       sync.RWMutex
	applyStatusArgsForCall []struct {
		arg1 context.Context
		arg2 *v1a.IngressApplyConfiguration
		arg3 v1b.ApplyOptions
	}
	applyStatusReturns struct {
		result1 *v1.Ingress
		result2 error
	}
	applyStatusReturnsOnCall map[int]struct {
		result1 *v1.Ingress
		result2 error
	}
	CreateStub        func(context.Context, *v1.Ingress, v1b.CreateOptions) (*v1.Ingress, error)
	createMutex       sync.RWMutex
	createArgsForCall []struct {
		arg1 context.Context
		arg2 *v1.Ingress
		arg3 v1b.CreateOptions
	}
	createReturns struct {
		result1 *v1.Ingress
		result2 error
	}
	createReturnsOnCall map[int]struct {
		result1 *v1.Ingress
		result2 error
	}
	DeleteStub        func(context.Context, string, v1b.DeleteOptions) error
	deleteMutex       sync.RWMutex
	deleteArgsForCall []struct {
		arg1 context.Context
		arg2 string
		arg3 v1b.DeleteOptions
	}
	deleteReturns struct {
		result1 error
	}
	deleteReturnsOnCall map[int]struct {
		result1 error
	}
	DeleteCollectionStub        func(context.Context, v1b.DeleteOptions, v1b.ListOptions) error
	deleteCollectionMutex       sync.RWMutex
	deleteCollectionArgsForCall []struct {
		arg1 context.Context
		arg2 v1b.DeleteOptions
		arg3 v1b.ListOptions
	}
	deleteCollectionReturns struct {
		result1 error
	}
	deleteCollectionReturnsOnCall map[int]struct {
		result1 error
	}
	GetStub        func(context.Context, string, v1b.GetOptions) (*v1.Ingress, error)
	getMutex       sync.RWMutex
	getArgsForCall []struct {
		arg1 context.Context
		arg2 string
		arg3 v1b.GetOptions
	}
	getReturns struct {
		result1 *v1.Ingress
		result2 error
	}
	getReturnsOnCall map[int]struct {
		result1 *v1.Ingress
		result2 error
	}
	ListStub        func(context.Context, v1b.ListOptions) (*v1.IngressList, error)
	listMutex       sync.RWMutex
	listArgsForCall []struct {
		arg1 context.Context
		arg2 v1b.ListOptions
	}
	listReturns struct {
		result1 *v1.IngressList
		result2 error
	}
	listReturnsOnCall map[int]struct {
		result1 *v1.IngressList
		result2 error
	}
	PatchStub        func(context.Context, string, types.PatchType, []byte, v1b.PatchOptions, ...string) (*v1.Ingress, error)
	patchMutex       sync.RWMutex
	patchArgsForCall []struct {
		arg1 context.Context
		arg2 string
		arg3 types.PatchType
		arg4 []byte
		arg5 v1b.PatchOptions
		arg6 []string
	}
	patchReturns struct {
		result1 *v1.Ingress
		result2 error
	}
	patchReturnsOnCall map[int]struct {
		result1 *v1.Ingress
		result2 error
	}
	UpdateStub        func(context.Context, *v1.Ingress, v1b.UpdateOptions) (*v1.Ingress, error)
	updateMutex       sync.RWMutex
	updateArgsForCall []struct {
		arg1 context.Context
		arg2 *v1.Ingress
		arg3 v1b.UpdateOptions
	}
	updateReturns struct {
		result1 *v1.Ingress
		result2 error
	}
	updateReturnsOnCall map[int]struct {
		result1 *v1.Ingress
		result2 error
	}
	UpdateStatusStub        func(context.Context, *v1.Ingress, v1b.UpdateOptions) (*v1.Ingress, error)
	updateStatusMutex       sync.RWMutex
	updateStatusArgsForCall []struct {
		arg1 context.Context
		arg2 *v1.Ingress
		arg3 v1b.UpdateOptions
	}
	updateStatusReturns struct {
		result1 *v1.Ingress
		result2 error
	}
	updateStatusReturnsOnCall map[int]struct {
		result1 *v1.Ingress
		result2 error
	}
	WatchStub        func(context.Context, v1b.ListOptions) (watch.Interface, error)
	watchMutex       sync.RWMutex
	watchArgsForCall []struct {
		arg1 context.Context
		arg2 v1b.ListOptions
	}
	watchReturns struct {
		result1 watch.Interface
		result2 error
	}
	watchReturnsOnCall map[int]struct {
		result1 watch.Interface
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *K8sIngressInterface) Apply(arg1 context.Context, arg2 *v1a.IngressApplyConfiguration, arg3 v1b.ApplyOptions) (*v1.Ingress, error) {
	fake.applyMutex.Lock()
	ret, specificReturn := fake.applyReturnsOnCall[len(fake.applyArgsForCall)]
	fake.applyArgsForCall = append(fake.applyArgsForCall, struct {
		arg1 context.Context
		arg2 *v1a.IngressApplyConfiguration
		arg3 v1b.ApplyOptions
	}{arg1, arg2, arg3})
	stub := fake.ApplyStub
	fakeReturns := fake.applyReturns
	fake.recordInvocation("Apply", []interface{}{arg1, arg2, arg3})
	fake.applyMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *K8sIngressInterface) ApplyCallCount() int {
	fake.applyMutex.RLock()
	defer fake.applyMutex.RUnlock()
	return len(fake.applyArgsForCall)
}

func (fake *K8sIngressInterface) ApplyCalls(stub func(context.Context, *v1a.IngressApplyConfiguration, v1b.ApplyOptions) (*v1.Ingress, error)) {
	fake.applyMutex.Lock()
	defer fake.applyMutex.Unlock()
	fake.ApplyStub = stub
}

func (fake *K8sIngressInterface) ApplyArgsForCall(i int) (context.Context, *v1a.IngressApplyConfiguration, v1b.ApplyOptions) {
	fake.applyMutex.RLock()
	defer fake.applyMutex.RUnlock()
	argsForCall := fake.applyArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *K8sIngressInterface) ApplyReturns(result1 *v1.Ingress, result2 error) {
	fake.applyMutex.Lock()
	defer fake.applyMutex.Unlock()
	fake.ApplyStub = nil
	fake.applyReturns = struct {
		result1 *v1.Ingress
		result2 error
	}{result1, result2}
}

func (fake *K8sIngressInterface) ApplyReturnsOnCall(i int, result1 *v1.Ingress, result2 error) {
	fake.applyMutex.Lock()
	defer fake.applyMutex.Unlock()
	fake.ApplyStub = nil
	if fake.applyReturnsOnCall == nil {
		fake.applyReturnsOnCall = make(map[int]struct {
			result1 *v1.Ingress
			result2 error
		})
	}
	fake.applyReturnsOnCall[i] = struct {
		result1 *v1.Ingress
		result2 error
	}{result1, result2}
}

func (fake *K8sIngressInterface) ApplyStatus(arg1 context.Context, arg2 *v1a.IngressApplyConfiguration, arg3 v1b.ApplyOptions) (*v1.Ingress, error) {
	fake.applyStatusMutex.Lock()
	ret, specificReturn := fake.applyStatusReturnsOnCall[len(fake.applyStatusArgsForCall)]
	fake.applyStatusArgsForCall = append(fake.applyStatusArgsForCall, struct {
		arg1 context.Context
		arg2 *v1a.IngressApplyConfiguration
		arg3 v1b.ApplyOptions
	}{arg1, arg2, arg3})
	stub := fake.ApplyStatusStub
	fakeReturns := fake.applyStatusReturns
	fake.recordInvocation("ApplyStatus", []interface{}{arg1, arg2, arg3})
	fake.applyStatusMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *K8sIngressInterface) ApplyStatusCallCount() int {
	fake.applyStatusMutex.RLock()
	defer fake.applyStatusMutex.RUnlock()
	return len(fake.applyStatusArgsForCall)
}

func (fake *K8sIngressInterface) ApplyStatusCalls(stub func(context.Context, *v1a.IngressApplyConfiguration, v1b.ApplyOptions) (*v1.Ingress, error)) {
	fake.applyStatusMutex.Lock()
	defer fake.applyStatusMutex.Unlock()
	fake.ApplyStatusStub = stub
}

func (fake *K8sIngressInterface) ApplyStatusArgsForCall(i int) (context.Context, *v1a.IngressApplyConfiguration, v1b.ApplyOptions) {
	fake.applyStatusMutex.RLock()
	defer fake.applyStatusMutex.RUnlock()
	argsForCall := fake.applyStatusArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *K8sIngressInterface) ApplyStatusReturns(result1 *v1.Ingress, result2 error) {
	fake.applyStatusMutex.Lock()
	defer fake.applyStatusMutex.Unlock()
	fake.ApplyStatusStub = nil
	fake.applyStatusReturns = struct {
		result1 *v1.Ingress
		result2 error
	}{result1, result2}
}

func (fake *K8sIngressInterface) ApplyStatusReturnsOnCall(i int, result1 *v1.Ingress, result2 error) {
	fake.applyStatusMutex.Lock()
	defer fake.applyStatusMutex.Unlock()
	fake.ApplyStatusStub = nil
	if fake.applyStatusReturnsOnCall == nil {
		fake.applyStatusReturnsOnCall = make(map[int]struct {
			result1 *v1.Ingress
			result2 error
		})
	}
	fake.applyStatusReturnsOnCall[i] = struct {
		result1 *v1.Ingress
		result2 error
	}{result1, result2}
}

func (fake *K8sIngressInterface) Create(arg1 context.Context, arg2 *v1.Ingress, arg3 v1b.CreateOptions) (*v1.Ingress, error) {
	fake.createMutex.Lock()
	ret, specificReturn := fake.createReturnsOnCall[len(fake.createArgsForCall)]
	fake.createArgsForCall = append(fake.createArgsForCall, struct {
		arg1 context.Context
		arg2 *v1.Ingress
		arg3 v1b.CreateOptions
	}{arg1, arg2, arg3})
	stub := fake.CreateStub
	fakeReturns := fake.createReturns
	fake.recordInvocation("Create", []interface{}{arg1, arg2, arg3})
	fake.createMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *K8sIngressInterface) CreateCallCount() int {
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	return len(fake.createArgsForCall)
}

func (fake *K8sIngressInterface) CreateCalls(stub func(context.Context, *v1.Ingress, v1b.CreateOptions) (*v1.Ingress, error)) {
	fake.createMutex.Lock()
	defer fake.createMutex.Unlock()
	fake.CreateStub = stub
}

func (fake *K8sIngressInterface) CreateArgsForCall(i int) (context.Context, *v1.Ingress, v1b.CreateOptions) {
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	argsForCall := fake.createArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *K8sIngressInterface) CreateReturns(result1 *v1.Ingress, result2 error) {
	fake.createMutex.Lock()
	defer fake.createMutex.Unlock()
	fake.CreateStub = nil
	fake.createReturns = struct {
		result1 *v1.Ingress
		result2 error
	}{result1, result2}
}

func (fake *K8sIngressInterface) CreateReturnsOnCall(i int, result1 *v1.Ingress, result2 error) {
	fake.createMutex.Lock()
	defer fake.createMutex.Unlock()
	fake.CreateStub = nil
	if fake.createReturnsOnCall == nil {
		fake.createReturnsOnCall = make(map[int]struct {
			result1 *v1.Ingress
			result2 error
		})
	}
	fake.createReturnsOnCall[i] = struct {
		result1 *v1.Ingress
		result2 error
	}{result1, result2}
}

func (fake *K8sIngressInterface) Delete(arg1 context.Context, arg2 string, arg3 v1b.DeleteOptions) error {
	fake.deleteMutex.Lock()
	ret, specificReturn := fake.deleteReturnsOnCall[len(fake.deleteArgsForCall)]
	fake.deleteArgsForCall = append(fake.deleteArgsForCall, struct {
		arg1 context.Context
		arg2 string
		arg3 v1b.DeleteOptions
	}{arg1, arg2, arg3})
	stub := fake.DeleteStub
	fakeReturns := fake.deleteReturns
	fake.recordInvocation("Delete", []interface{}{arg1, arg2, arg3})
	fake.deleteMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *K8sIngressInterface) DeleteCallCount() int {
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	return len(fake.deleteArgsForCall)
}

func (fake *K8sIngressInterface) DeleteCalls(stub func(context.Context, string, v1b.DeleteOptions) error) {
	fake.deleteMutex.Lock()
	defer fake.deleteMutex.Unlock()
	fake.DeleteStub = stub
}

func (fake *K8sIngressInterface) DeleteArgsForCall(i int) (context.Context, string, v1b.DeleteOptions) {
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	argsForCall := fake.deleteArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *K8sIngressInterface) DeleteReturns(result1 error) {
	fake.deleteMutex.Lock()
	defer fake.deleteMutex.Unlock()
	fake.DeleteStub = nil
	fake.deleteReturns = struct {
		result1 error
	}{result1}
}

func (fake *K8sIngressInterface) DeleteReturnsOnCall(i int, result1 error) {
	fake.deleteMutex.Lock()
	defer fake.deleteMutex.Unlock()
	fake.DeleteStub = nil
	if fake.deleteReturnsOnCall == nil {
		fake.deleteReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.deleteReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *K8sIngressInterface) DeleteCollection(arg1 context.Context, arg2 v1b.DeleteOptions, arg3 v1b.ListOptions) error {
	fake.deleteCollectionMutex.Lock()
	ret, specificReturn := fake.deleteCollectionReturnsOnCall[len(fake.deleteCollectionArgsForCall)]
	fake.deleteCollectionArgsForCall = append(fake.deleteCollectionArgsForCall, struct {
		arg1 context.Context
		arg2 v1b.DeleteOptions
		arg3 v1b.ListOptions
	}{arg1, arg2, arg3})
	stub := fake.DeleteCollectionStub
	fakeReturns := fake.deleteCollectionReturns
	fake.recordInvocation("DeleteCollection", []interface{}{arg1, arg2, arg3})
	fake.deleteCollectionMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *K8sIngressInterface) DeleteCollectionCallCount() int {
	fake.deleteCollectionMutex.RLock()
	defer fake.deleteCollectionMutex.RUnlock()
	return len(fake.deleteCollectionArgsForCall)
}

func (fake *K8sIngressInterface) DeleteCollectionCalls(stub func(context.Context, v1b.DeleteOptions, v1b.ListOptions) error) {
	fake.deleteCollectionMutex.Lock()
	defer fake.deleteCollectionMutex.Unlock()
	fake.DeleteCollectionStub = stub
}

func (fake *K8sIngressInterface) DeleteCollectionArgsForCall(i int) (context.Context, v1b.DeleteOptions, v1b.ListOptions) {
	fake.deleteCollectionMutex.RLock()
	defer fake.deleteCollectionMutex.RUnlock()
	argsForCall := fake.deleteCollectionArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *K8sIngressInterface) DeleteCollectionReturns(result1 error) {
	fake.deleteCollectionMutex.Lock()
	defer fake.deleteCollectionMutex.Unlock()
	fake.DeleteCollectionStub = nil
	fake.deleteCollectionReturns = struct {
		result1 error
	}{result1}
}

func (fake *K8sIngressInterface) DeleteCollectionReturnsOnCall(i int, result1 error) {
	fake.deleteCollectionMutex.Lock()
	defer fake.deleteCollectionMutex.Unlock()
	fake.DeleteCollectionStub = nil
	if fake.deleteCollectionReturnsOnCall == nil {
		fake.deleteCollectionReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.deleteCollectionReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *K8sIngressInterface) Get(arg1 context.Context, arg2 string, arg3 v1b.GetOptions) (*v1.Ingress, error) {
	fake.getMutex.Lock()
	ret, specificReturn := fake.getReturnsOnCall[len(fake.getArgsForCall)]
	fake.getArgsForCall = append(fake.getArgsForCall, struct {
		arg1 context.Context
		arg2 string
		arg3 v1b.GetOptions
	}{arg1, arg2, arg3})
	stub := fake.GetStub
	fakeReturns := fake.getReturns
	fake.recordInvocation("Get", []interface{}{arg1, arg2, arg3})
	fake.getMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *K8sIngressInterface) GetCallCount() int {
	fake.getMutex.RLock()
	defer fake.getMutex.RUnlock()
	return len(fake.getArgsForCall)
}

func (fake *K8sIngressInterface) GetCalls(stub func(context.Context, string, v1b.GetOptions) (*v1.Ingress, error)) {
	fake.getMutex.Lock()
	defer fake.getMutex.Unlock()
	fake.GetStub = stub
}

func (fake *K8sIngressInterface) GetArgsForCall(i int) (context.Context, string, v1b.GetOptions) {
	fake.getMutex.RLock()
	defer fake.getMutex.RUnlock()
	argsForCall := fake.getArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *K8sIngressInterface) GetReturns(result1 *v1.Ingress, result2 error) {
	fake.getMutex.Lock()
	defer fake.getMutex.Unlock()
	fake.GetStub = nil
	fake.getReturns = struct {
		result1 *v1.Ingress
		result2 error
	}{result1, result2}
}

func (fake *K8sIngressInterface) GetReturnsOnCall(i int, result1 *v1.Ingress, result2 error) {
	fake.getMutex.Lock()
	defer fake.getMutex.Unlock()
	fake.GetStub = nil
	if fake.getReturnsOnCall == nil {
		fake.getReturnsOnCall = make(map[int]struct {
			result1 *v1.Ingress
			result2 error
		})
	}
	fake.getReturnsOnCall[i] = struct {
		result1 *v1.Ingress
		result2 error
	}{result1, result2}
}

func (fake *K8sIngressInterface) List(arg1 context.Context, arg2 v1b.ListOptions) (*v1.IngressList, error) {
	fake.listMutex.Lock()
	ret, specificReturn := fake.listReturnsOnCall[len(fake.listArgsForCall)]
	fake.listArgsForCall = append(fake.listArgsForCall, struct {
		arg1 context.Context
		arg2 v1b.ListOptions
	}{arg1, arg2})
	stub := fake.ListStub
	fakeReturns := fake.listReturns
	fake.recordInvocation("List", []interface{}{arg1, arg2})
	fake.listMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *K8sIngressInterface) ListCallCount() int {
	fake.listMutex.RLock()
	defer fake.listMutex.RUnlock()
	return len(fake.listArgsForCall)
}

func (fake *K8sIngressInterface) ListCalls(stub func(context.Context, v1b.ListOptions) (*v1.IngressList, error)) {
	fake.listMutex.Lock()
	defer fake.listMutex.Unlock()
	fake.ListStub = stub
}

func (fake *K8sIngressInterface) ListArgsForCall(i int) (context.Context, v1b.ListOptions) {
	fake.listMutex.RLock()
	defer fake.listMutex.RUnlock()
	argsForCall := fake.listArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *K8sIngressInterface) ListReturns(result1 *v1.IngressList, result2 error) {
	fake.listMutex.Lock()
	defer fake.listMutex.Unlock()
	fake.ListStub = nil
	fake.listReturns = struct {
		result1 *v1.IngressList
		result2 error
	}{result1, result2}
}

func (fake *K8sIngressInterface) ListReturnsOnCall(i int, result1 *v1.IngressList, result2 error) {
	fake.listMutex.Lock()
	defer fake.listMutex.Unlock()
	fake.ListStub = nil
	if fake.listReturnsOnCall == nil {
		fake.listReturnsOnCall = make(map[int]struct {
			result1 *v1.IngressList
			result2 error
		})
	}
	fake.listReturnsOnCall[i] = struct {
		result1 *v1.IngressList
		result2 error
	}{result1, result2}
}

func (fake *K8sIngressInterface) Patch(arg1 context.Context, arg2 string, arg3 types.PatchType, arg4 []byte, arg5 v1b.PatchOptions, arg6 ...string) (*v1.Ingress, error) {
	var arg4Copy []byte
	if arg4 != nil {
		arg4Copy = make([]byte, len(arg4))
		copy(arg4Copy, arg4)
	}
	fake.patchMutex.Lock()
	ret, specificReturn := fake.patchReturnsOnCall[len(fake.patchArgsForCall)]
	fake.patchArgsForCall = append(fake.patchArgsForCall, struct {
		arg1 context.Context
		arg2 string
		arg3 types.PatchType
		arg4 []byte
		arg5 v1b.PatchOptions
		arg6 []string
	}{arg1, arg2, arg3, arg4Copy, arg5, arg6})
	stub := fake.PatchStub
	fakeReturns := fake.patchReturns
	fake.recordInvocation("Patch", []interface{}{arg1, arg2, arg3, arg4Copy, arg5, arg6})
	fake.patchMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3, arg4, arg5, arg6...)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *K8sIngressInterface) PatchCallCount() int {
	fake.patchMutex.RLock()
	defer fake.patchMutex.RUnlock()
	return len(fake.patchArgsForCall)
}

func (fake *K8sIngressInterface) PatchCalls(stub func(context.Context, string, types.PatchType, []byte, v1b.PatchOptions, ...string) (*v1.Ingress, error)) {
	fake.patchMutex.Lock()
	defer fake.patchMutex.Unlock()
	fake.PatchStub = stub
}

func (fake *K8sIngressInterface) PatchArgsForCall(i int) (context.Context, string, types.PatchType, []byte, v1b.PatchOptions, []string) {
	fake.patchMutex.RLock()
	defer fake.patchMutex.RUnlock()
	argsForCall := fake.patchArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4, argsForCall.arg5, argsForCall.arg6
}

func (fake *K8sIngressInterface) PatchReturns(result1 *v1.Ingress, result2 error) {
	fake.patchMutex.Lock()
	defer fake.patchMutex.Unlock()
	fake.PatchStub = nil
	fake.patchReturns = struct {
		result1 *v1.Ingress
		result2 error
	}{result1, result2}
}

func (fake *K8sIngressInterface) PatchReturnsOnCall(i int, result1 *v1.Ingress, result2 error) {
	fake.patchMutex.Lock()
	defer fake.patchMutex.Unlock()
	fake.PatchStub = nil
	if fake.patchReturnsOnCall == nil {
		fake.patchReturnsOnCall = make(map[int]struct {
			result1 *v1.Ingress
			result2 error
		})
	}
	fake.patchReturnsOnCall[i] = struct {
		result1 *v1.Ingress
		result2 error
	}{result1, result2}
}

func (fake *K8sIngressInterface) Update(arg1 context.Context, arg2 *v1.Ingress, arg3 v1b.UpdateOptions) (*v1.Ingress, error) {
	fake.updateMutex.Lock()
	ret, specificReturn := fake.updateReturnsOnCall[len(fake.updateArgsForCall)]
	fake.updateArgsForCall = append(fake.updateArgsForCall, struct {
		arg1 context.Context
		arg2 *v1.Ingress
		arg3 v1b.UpdateOptions
	}{arg1, arg2, arg3})
	stub := fake.UpdateStub
	fakeReturns := fake.updateReturns
	fake.recordInvocation("Update", []interface{}{arg1, arg2, arg3})
	fake.updateMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *K8sIngressInterface) UpdateCallCount() int {
	fake.updateMutex.RLock()
	defer fake.updateMutex.RUnlock()
	return len(fake.updateArgsForCall)
}

func (fake *K8sIngressInterface) UpdateCalls(stub func(context.Context, *v1.Ingress, v1b.UpdateOptions) (*v1.Ingress, error)) {
	fake.updateMutex.Lock()
	defer fake.updateMutex.Unlock()
	fake.UpdateStub = stub
}

func (fake *K8sIngressInterface) UpdateArgsForCall(i int) (context.Context, *v1.Ingress, v1b.UpdateOptions) {
	fake.updateMutex.RLock()
	defer fake.updateMutex.RUnlock()
	argsForCall := fake.updateArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *K8sIngressInterface) UpdateReturns(result1 *v1.Ingress, result2 error) {
	fake.updateMutex.Lock()
	defer fake.updateMutex.Unlock()
	fake.UpdateStub = nil
	fake.updateReturns = struct {
		result1 *v1.Ingress
		result2 error
	}{result1, result2}
}

func (fake *K8sIngressInterface) UpdateReturnsOnCall(i int, result1 *v1.Ingress, result2 error) {
	fake.updateMutex.Lock()
	defer fake.updateMutex.Unlock()
	fake.UpdateStub = nil
	if fake.updateReturnsOnCall == nil {
		fake.updateReturnsOnCall = make(map[int]struct {
			result1 *v1.Ingress
			result2 error
		})
	}
	fake.updateReturnsOnCall[i] = struct {
		result1 *v1.Ingress
		result2 error
	}{result1, result2}
}

func (fake *K8sIngressInterface) UpdateStatus(arg1 context.Context, arg2 *v1.Ingress, arg3 v1b.UpdateOptions) (*v1.Ingress, error) {
	fake.updateStatusMutex.Lock()
	ret, specificReturn := fake.updateStatusReturnsOnCall[len(fake.updateStatusArgsForCall)]
	fake.updateStatusArgsForCall = append(fake.updateStatusArgsForCall, struct {
		arg1 context.Context
		arg2 *v1.Ingress
		arg3 v1b.UpdateOptions
	}{arg1, arg2, arg3})
	stub := fake.UpdateStatusStub
	fakeReturns := fake.updateStatusReturns
	fake.recordInvocation("UpdateStatus", []interface{}{arg1, arg2, arg3})
	fake.updateStatusMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *K8sIngressInterface) UpdateStatusCallCount() int {
	fake.updateStatusMutex.RLock()
	defer fake.updateStatusMutex.RUnlock()
	return len(fake.updateStatusArgsForCall)
}

func (fake *K8sIngressInterface) UpdateStatusCalls(stub func(context.Context, *v1.Ingress, v1b.UpdateOptions) (*v1.Ingress, error)) {
	fake.updateStatusMutex.Lock()
	defer fake.updateStatusMutex.Unlock()
	fake.UpdateStatusStub = stub
}

func (fake *K8sIngressInterface) UpdateStatusArgsForCall(i int) (context.Context, *v1.Ingress, v1b.UpdateOptions) {
	fake.updateStatusMutex.RLock()
	defer fake.updateStatusMutex.RUnlock()
	argsForCall := fake.updateStatusArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *K8sIngressInterface) UpdateStatusReturns(result1 *v1.Ingress, result2 error) {
	fake.updateStatusMutex.Lock()
	defer fake.updateStatusMutex.Unlock()
	fake.UpdateStatusStub = nil
	fake.updateStatusReturns = struct {
		result1 *v1.Ingress
		result2 error
	}{result1, result2}
}

func (fake *K8sIngressInterface) UpdateStatusReturnsOnCall(i int, result1 *v1.Ingress, result2 error) {
	fake.updateStatusMutex.Lock()
	defer fake.updateStatusMutex.Unlock()
	fake.UpdateStatusStub = nil
	if fake.updateStatusReturnsOnCall == nil {
		fake.updateStatusReturnsOnCall = make(map[int]struct {
			result1 *v1.Ingress
			result2 error
		})
	}
	fake.updateStatusReturnsOnCall[i] = struct {
		result1 *v1.Ingress
		result2 error
	}{result1, result2}
}

func (fake *K8sIngressInterface) Watch(arg1 context.Context, arg2 v1b.ListOptions) (watch.Interface, error) {
	fake.watchMutex.Lock()
	ret, specificReturn := fake.watchReturnsOnCall[len(fake.watchArgsForCall)]
	fake.watchArgsForCall = append(fake.watchArgsForCall, struct {
		arg1 context.Context
		arg2 v1b.ListOptions
	}{arg1, arg2})
	stub := fake.WatchStub
	fakeReturns := fake.watchReturns
	fake.recordInvocation("Watch", []interface{}{arg1, arg2})
	fake.watchMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *K8sIngressInterface) WatchCallCount() int {
	fake.watchMutex.RLock()
	defer fake.watchMutex.RUnlock()
	return len(fake.watchArgsForCall)
}

func (fake *K8sIngressInterface) WatchCalls(stub func(context.Context, v1b.ListOptions) (watch.Interface, error)) {
	fake.watchMutex.Lock()
	defer fake.watchMutex.Unlock()
	fake.WatchStub = stub
}

func (fake *K8sIngressInterface) WatchArgsForCall(i int) (context.Context, v1b.ListOptions) {
	fake.watchMutex.RLock()
	defer fake.watchMutex.RUnlock()
	argsForCall := fake.watchArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *K8sIngressInterface) WatchReturns(result1 watch.Interface, result2 error) {
	fake.watchMutex.Lock()
	defer fake.watchMutex.Unlock()
	fake.WatchStub = nil
	fake.watchReturns = struct {
		result1 watch.Interface
		result2 error
	}{result1, result2}
}

func (fake *K8sIngressInterface) WatchReturnsOnCall(i int, result1 watch.Interface, result2 error) {
	fake.watchMutex.Lock()
	defer fake.watchMutex.Unlock()
	fake.WatchStub = nil
	if fake.watchReturnsOnCall == nil {
		fake.watchReturnsOnCall = make(map[int]struct {
			result1 watch.Interface
			result2 error
		})
	}
	fake.watchReturnsOnCall[i] = struct {
		result1 watch.Interface
		result2 error
	}{result1, result2}
}

func (fake *K8sIngressInterface) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.applyMutex.RLock()
	defer fake.applyMutex.RUnlock()
	fake.applyStatusMutex.RLock()
	defer fake.applyStatusMutex.RUnlock()
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	fake.deleteCollectionMutex.RLock()
	defer fake.deleteCollectionMutex.RUnlock()
	fake.getMutex.RLock()
	defer fake.getMutex.RUnlock()
	fake.listMutex.RLock()
	defer fake.listMutex.RUnlock()
	fake.patchMutex.RLock()
	defer fake.patchMutex.RUnlock()
	fake.updateMutex.RLock()
	defer fake.updateMutex.RUnlock()
	fake.updateStatusMutex.RLock()
	defer fake.updateStatusMutex.RUnlock()
	fake.watchMutex.RLock()
	defer fake.watchMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *K8sIngressInterface) recordInvocation(key string, args []interface{}) {
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

var _ k8s.IngressInterface = new(K8sIngressInterface)
