// Copyright (c) 2025 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package k8s

import (
	"context"

	"github.com/bborbe/errors"
	"github.com/bborbe/run"
	libtime "github.com/bborbe/time"
	"github.com/golang/glog"
	corev1 "k8s.io/api/core/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/kubernetes"
)

//counterfeiter:generate -o mocks/k8s-service-watcher.go --fake-name K8sServiceWatcher . ServiceWatcher
type ServiceWatcher interface {
	Watch(ctx context.Context) error
}

func NewServiceWatcher(
	clientset kubernetes.Interface,
	serviceManager ServiceEventProcessor,
	waiterDuration libtime.WaiterDuration,
	namespace Namespace,
) ServiceWatcher {
	return &serviceStatusMonitoring{
		waiterDuration: waiterDuration,
		clientset:      clientset,
		namespace:      namespace,
		serviceManager: serviceManager,
	}
}

type serviceStatusMonitoring struct {
	clientset      kubernetes.Interface
	namespace      Namespace
	serviceManager ServiceEventProcessor
	waiterDuration libtime.WaiterDuration
}

func (s *serviceStatusMonitoring) Watch(ctx context.Context) error {
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			glog.V(2).Infof("watch services started")
			if err := s.watch(ctx); err != nil {
				return errors.Wrapf(ctx, err, "watch failed")
			}
			if err := s.waiterDuration.Wait(ctx, 5*libtime.Second); err != nil {
				return errors.Wrapf(ctx, err, "wait failed")
			}
			glog.V(2).Infof("watch services completed")
		}
	}
}

func (s *serviceStatusMonitoring) watch(ctx context.Context) error {
	watchInf, err := s.clientset.CoreV1().
		Services(s.namespace.String()).
		Watch(ctx, metav1.ListOptions{})
	if err != nil {
		return errors.Wrap(ctx, err, "watch failed")
	}
	resultChan := watchInf.ResultChan()
	return run.CancelOnFirstFinish(
		ctx,
		func(ctx context.Context) error {
			select {
			case <-ctx.Done():
				watchInf.Stop()
				return ctx.Err()
			}
		},
		func(ctx context.Context) error {
			for {
				select {
				case <-ctx.Done():
					return ctx.Err()
				case event, ok := <-resultChan:
					if !ok {
						glog.V(2).Infof("result channel closed => return")
						return nil
					}
					if event.Type == watch.Error {
						return apierrors.FromObject(event.Object)
					}
					switch event.Type {
					case watch.Added, watch.Modified:
						switch o := event.Object.(type) {
						case *corev1.Service:
							if err := s.serviceManager.OnUpdate(ctx, *o); err != nil {
								return errors.Wrap(ctx, err, "on delete failed")
							}
						default:
							return errors.Errorf(ctx, "unknown object type %T", event.Object)
						}
					case watch.Deleted:
						switch o := event.Object.(type) {
						case *corev1.Service:
							if err := s.serviceManager.OnDelete(ctx, *o); err != nil {
								return errors.Wrap(ctx, err, "on delete failed")
							}
						default:
							return errors.Errorf(ctx, "unknown object type %T", event.Object)
						}
					case watch.Bookmark:
					case watch.Error:
					default:
						return errors.Errorf(ctx, "unknown event type")
					}
				}
			}
		},
	)
}
