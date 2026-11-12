// Copyright (c) 2025 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package k8s

import (
	"context"
	"time"

	"github.com/bborbe/errors"
	libtime "github.com/bborbe/time"
	"github.com/golang/glog"
	corev1 "k8s.io/api/core/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/kubernetes"
)

//counterfeiter:generate -o mocks/k8s-pod-watcher.go --fake-name K8sPodWatcher . PodWatcher
type PodWatcher interface {
	Watch(ctx context.Context) error
}

// NewPodWatcher creates a Kubernetes Pod watcher that monitors Pod
// resources in the specified namespace. It invokes callbacks on the provided
// PodEventProcessor for Add/Modify/Delete events.
//
// The watcher delivers the objects that already exist when it starts, then
// re-delivers the full set every defaultResyncInterval, so a missed watch
// event repairs itself without a restart. Delivery is at-least-once: each
// reconnect re-delivers the full set.
//
// Watch still returns when the watch connection closes, so NewPodWatcherRetry
// reconnects as before. Use NewPodWatcherWithResync to set the interval.
func NewPodWatcher(
	clientset kubernetes.Interface,
	podManager PodEventProcessor,
	namespace Namespace,
) PodWatcher {
	return NewPodWatcherWithResync(
		clientset,
		podManager,
		namespace,
		defaultResyncInterval,
	)
}

// NewPodWatcherWithResync is NewPodWatcher with an explicit resync interval.
// A non-positive interval disables the periodic resync while the startup list
// still runs; it exists so time.NewTicker is never constructed with a
// non-positive duration.
func NewPodWatcherWithResync(
	clientset kubernetes.Interface,
	podManager PodEventProcessor,
	namespace Namespace,
	resyncInterval libtime.Duration,
) PodWatcher {
	return &podStatusMonitoring{
		clientset:      clientset,
		namespace:      namespace,
		podManager:     podManager,
		resyncInterval: resyncInterval,
	}
}

type podStatusMonitoring struct {
	clientset      kubernetes.Interface
	namespace      Namespace
	podManager     PodEventProcessor
	resyncInterval libtime.Duration
}

// reconcile lists every Pod in the namespace and delivers each one to the event
// processor. An empty ListOptions is a true full list, which is what a resync
// needs.
func (s *podStatusMonitoring) reconcile(ctx context.Context) error {
	list, err := s.clientset.CoreV1().
		Pods(s.namespace.String()).
		List(ctx, metav1.ListOptions{})
	if err != nil {
		return errors.Wrap(ctx, err, "list pods failed")
	}
	for _, pod := range list.Items {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}
		if err := s.podManager.OnUpdate(ctx, pod); err != nil {
			return errors.Wrap(ctx, err, "on update failed")
		}
	}
	return nil
}

//nolint:dupl // the secret and pod watchers are intentionally parallel; see docs/dod.md
func (s *podStatusMonitoring) Watch(ctx context.Context) error {
	// Check if context is already canceled before starting watch
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
	}

	if err := s.reconcile(ctx); err != nil {
		return err
	}

	glog.V(2).Infof("watch pods started (resync interval %s)", s.resyncInterval)
	watchInf, err := s.clientset.CoreV1().
		Pods(s.namespace.String()).
		Watch(ctx, metav1.ListOptions{})
	if err != nil {
		return errors.Wrap(ctx, err, "watch failed")
	}
	defer watchInf.Stop()

	// A nil channel blocks forever in select, so a non-positive interval simply
	// never takes the resync branch.
	var resyncC <-chan time.Time
	if s.resyncInterval > 0 {
		ticker := time.NewTicker(s.resyncInterval.Duration())
		defer ticker.Stop()
		resyncC = ticker.C
	}

	resultChan := watchInf.ResultChan()

	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-resyncC:
			if err := s.reconcile(ctx); err != nil {
				return err
			}
		case event, ok := <-resultChan:
			if !ok {
				glog.V(2).Infof("result channel closed")
				return ErrResultChannelClosed
			}
			switch event.Type {
			case watch.Error:
				return apierrors.FromObject(event.Object)
			case watch.Added, watch.Modified:
				switch o := event.Object.(type) {
				case *corev1.Pod:
					if err := s.podManager.OnUpdate(ctx, *o); err != nil {
						return errors.Wrap(ctx, err, "on update failed")
					}
				default:
					return errors.Errorf(ctx, "unknown object type %T", event.Object)
				}
			case watch.Deleted:
				switch o := event.Object.(type) {
				case *corev1.Pod:
					if err := s.podManager.OnDelete(ctx, *o); err != nil {
						return errors.Wrap(ctx, err, "on delete failed")
					}
				default:
					return errors.Errorf(ctx, "unknown object type %T", event.Object)
				}
			case watch.Bookmark:
			default:
				return errors.Wrapf(ctx, ErrUnknownEventType, "event type: %s", event.Type)
			}
		}
	}
}
