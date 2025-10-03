// Copyright (c) 2025 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package k8s_test

import (
	"context"
	"errors"

	libtime "github.com/bborbe/time"
	timemocks "github.com/bborbe/time/mocks"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/watch"

	"github.com/bborbe/k8s"
	"github.com/bborbe/k8s/mocks"
)

var _ = Describe("ServiceWatcher", func() {
	var serviceWatcher k8s.ServiceWatcher
	var clientset *mocks.K8sInterface
	var coreV1 *mocks.K8sCoreV1Interface
	var serviceInterface *mocks.K8sServiceInterface
	var serviceEventProcessor *mocks.K8sServiceEventProcessor
	var waiterDuration *timemocks.WaiterDuration
	var ctx context.Context
	var cancel context.CancelFunc
	var namespace k8s.Namespace
	var fakeWatcher *watch.FakeWatcher
	var err error

	BeforeEach(func() {
		ctx, cancel = context.WithCancel(context.Background())
		clientset = &mocks.K8sInterface{}
		coreV1 = &mocks.K8sCoreV1Interface{}
		serviceInterface = &mocks.K8sServiceInterface{}
		serviceEventProcessor = &mocks.K8sServiceEventProcessor{}
		waiterDuration = &timemocks.WaiterDuration{}
		namespace = k8s.Namespace("test-namespace")
		fakeWatcher = watch.NewFake()

		clientset.CoreV1Returns(coreV1)
		coreV1.ServicesReturns(serviceInterface)
		serviceInterface.WatchReturns(fakeWatcher, nil)

		serviceWatcher = k8s.NewServiceWatcher(
			clientset,
			serviceEventProcessor,
			waiterDuration,
			namespace,
		)

		// Default: waiter returns context.Canceled to exit the loop
		waiterDuration.WaitReturns(context.Canceled)
	})

	AfterEach(func() {
		if fakeWatcher != nil {
			fakeWatcher.Stop()
		}
		cancel()
	})

	Describe("Watch", func() {
		JustBeforeEach(func() {
			err = serviceWatcher.Watch(ctx)
		})

		Context("when context is already canceled", func() {
			BeforeEach(func() {
				cancel()
			})

			It("returns context.Canceled error", func() {
				Expect(err).To(Equal(context.Canceled))
			})

			It("does not call Watch on service interface", func() {
				Expect(serviceInterface.WatchCallCount()).To(Equal(0))
			})
		})

		Context("when watch fails to start", func() {
			BeforeEach(func() {
				serviceInterface.WatchReturns(nil, errors.New("watch failed"))
			})

			It("returns an error", func() {
				Expect(err).NotTo(BeNil())
				Expect(err.Error()).To(ContainSubstring("watch failed"))
			})
		})

		Context("when receiving Added event", func() {
			var service corev1.Service

			BeforeEach(func() {
				service = corev1.Service{
					ObjectMeta: metav1.ObjectMeta{
						Name:      "test-service",
						Namespace: "test-namespace",
					},
				}

				serviceEventProcessor.OnUpdateReturns(nil)

				// Start watch in background and send event
				go func() {
					fakeWatcher.Add(&service)
					fakeWatcher.Stop()
				}()
			})

			It("calls OnUpdate on service event processor", func() {
				Expect(serviceEventProcessor.OnUpdateCallCount()).To(Equal(1))
				_, receivedService := serviceEventProcessor.OnUpdateArgsForCall(0)
				Expect(receivedService.Name).To(Equal("test-service"))
				Expect(receivedService.Namespace).To(Equal("test-namespace"))
			})
		})

		Context("when receiving Modified event", func() {
			var service corev1.Service

			BeforeEach(func() {
				service = corev1.Service{
					ObjectMeta: metav1.ObjectMeta{
						Name:      "test-service",
						Namespace: "test-namespace",
					},
				}

				serviceEventProcessor.OnUpdateReturns(nil)

				go func() {
					fakeWatcher.Modify(&service)
					fakeWatcher.Stop()
				}()
			})

			It("calls OnUpdate on service event processor", func() {
				Expect(serviceEventProcessor.OnUpdateCallCount()).To(Equal(1))
				_, receivedService := serviceEventProcessor.OnUpdateArgsForCall(0)
				Expect(receivedService.Name).To(Equal("test-service"))
			})
		})

		Context("when receiving Deleted event", func() {
			var service corev1.Service

			BeforeEach(func() {
				service = corev1.Service{
					ObjectMeta: metav1.ObjectMeta{
						Name:      "test-service",
						Namespace: "test-namespace",
					},
				}

				serviceEventProcessor.OnDeleteReturns(nil)

				go func() {
					fakeWatcher.Delete(&service)
					fakeWatcher.Stop()
				}()
			})

			It("calls OnDelete on service event processor", func() {
				Expect(serviceEventProcessor.OnDeleteCallCount()).To(Equal(1))
				_, receivedService := serviceEventProcessor.OnDeleteArgsForCall(0)
				Expect(receivedService.Name).To(Equal("test-service"))
			})
		})

		Context("when OnUpdate returns an error", func() {
			var service corev1.Service

			BeforeEach(func() {
				service = corev1.Service{
					ObjectMeta: metav1.ObjectMeta{
						Name:      "test-service",
						Namespace: "test-namespace",
					},
				}

				serviceEventProcessor.OnUpdateReturns(errors.New("update failed"))

				go func() {
					fakeWatcher.Add(&service)
				}()
			})

			It("returns an error", func() {
				Expect(err).NotTo(BeNil())
				Expect(err.Error()).To(ContainSubstring("on delete failed"))
			})
		})

		Context("when OnDelete returns an error", func() {
			var service corev1.Service

			BeforeEach(func() {
				service = corev1.Service{
					ObjectMeta: metav1.ObjectMeta{
						Name:      "test-service",
						Namespace: "test-namespace",
					},
				}

				serviceEventProcessor.OnDeleteReturns(errors.New("delete failed"))

				go func() {
					fakeWatcher.Delete(&service)
				}()
			})

			It("returns an error", func() {
				Expect(err).NotTo(BeNil())
				Expect(err.Error()).To(ContainSubstring("on delete failed"))
			})
		})

		Context("when receiving Bookmark event", func() {
			BeforeEach(func() {
				go func() {
					fakeWatcher.Action(watch.Bookmark, nil)
					fakeWatcher.Stop()
				}()
			})

			It("does not call event processor", func() {
				Expect(serviceEventProcessor.OnUpdateCallCount()).To(Equal(0))
				Expect(serviceEventProcessor.OnDeleteCallCount()).To(Equal(0))
			})
		})

		Context("when result channel is closed", func() {
			BeforeEach(func() {
				go func() {
					fakeWatcher.Stop()
				}()
			})

			It("returns an error from wait", func() {
				Expect(err).NotTo(BeNil())
				Expect(err.Error()).To(ContainSubstring("wait failed"))
			})

			It("calls Wait on waiter duration", func() {
				Expect(waiterDuration.WaitCallCount()).To(Equal(1))
				_, duration := waiterDuration.WaitArgsForCall(0)
				Expect(duration).To(Equal(5 * libtime.Second))
			})
		})

		Context("when context is canceled during watch", func() {
			BeforeEach(func() {
				// Cancel context after a brief moment
				go func() {
					cancel()
				}()
			})

			It("returns context.Canceled", func() {
				Expect(err).To(MatchError(context.Canceled))
			})
		})

		Context("when waiter duration fails", func() {
			BeforeEach(func() {
				waiterDuration.WaitReturns(errors.New("wait failed"))

				go func() {
					fakeWatcher.Stop()
				}()
			})

			It("returns an error from wait", func() {
				Expect(err).NotTo(BeNil())
				Expect(err.Error()).To(ContainSubstring("wait failed"))
			})
		})
	})

	Describe("Constructor", func() {
		It("creates a new service watcher", func() {
			watcher := k8s.NewServiceWatcher(
				clientset,
				serviceEventProcessor,
				waiterDuration,
				namespace,
			)
			Expect(watcher).NotTo(BeNil())
		})
	})
})
