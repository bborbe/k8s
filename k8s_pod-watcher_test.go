// Copyright (c) 2025 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package k8s_test

import (
	"context"
	"errors"
	"time"

	libtime "github.com/bborbe/time"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/watch"

	"github.com/bborbe/k8s"
	"github.com/bborbe/k8s/mocks"
)

var _ = Describe("PodWatcher", func() {
	var podWatcher k8s.PodWatcher
	var clientset *mocks.K8sInterface
	var coreV1 *mocks.K8sCoreV1Interface
	var podInterface *mocks.K8sPodInterface
	var podEventProcessor *mocks.K8sPodEventProcessor
	var ctx context.Context
	var cancel context.CancelFunc
	var namespace k8s.Namespace
	var fakeWatcher *watch.FakeWatcher
	var err error

	BeforeEach(func() {
		ctx, cancel = context.WithCancel(context.Background())
		clientset = &mocks.K8sInterface{}
		coreV1 = &mocks.K8sCoreV1Interface{}
		podInterface = &mocks.K8sPodInterface{}
		podEventProcessor = &mocks.K8sPodEventProcessor{}
		namespace = k8s.Namespace("test-namespace")
		fakeWatcher = watch.NewFake()

		clientset.CoreV1Returns(coreV1)
		coreV1.PodsReturns(podInterface)
		podInterface.ListReturns(&corev1.PodList{}, nil)
		podInterface.WatchReturns(fakeWatcher, nil)

		podWatcher = k8s.NewPodWatcher(
			clientset,
			podEventProcessor,
			namespace,
		)
	})

	AfterEach(func() {
		if fakeWatcher != nil {
			fakeWatcher.Stop()
		}
		cancel()
	})

	Describe("Watch", func() {
		JustBeforeEach(func() {
			err = podWatcher.Watch(ctx)
		})

		Context("when context is already canceled", func() {
			BeforeEach(func() {
				cancel()
			})

			It("returns context.Canceled error", func() {
				Expect(err).To(Equal(context.Canceled))
			})

			It("does not call Watch on pod interface", func() {
				Expect(podInterface.WatchCallCount()).To(Equal(0))
			})
		})

		Context("when watch fails to start", func() {
			BeforeEach(func() {
				podInterface.WatchReturns(nil, errors.New("watch failed"))
			})

			It("returns an error", func() {
				Expect(err).NotTo(BeNil())
				Expect(err.Error()).To(ContainSubstring("watch failed"))
			})
		})

		Context("when receiving Added event", func() {
			var pod corev1.Pod

			BeforeEach(func() {
				pod = corev1.Pod{
					ObjectMeta: metav1.ObjectMeta{
						Name:      "test-pod",
						Namespace: "test-namespace",
					},
				}

				podEventProcessor.OnUpdateReturns(nil)

				// Start watch in background and send event
				go func() {
					fakeWatcher.Add(&pod)
					fakeWatcher.Stop()
				}()
			})

			It("calls OnUpdate on pod event processor", func() {
				Expect(podEventProcessor.OnUpdateCallCount()).To(Equal(1))
				_, receivedPod := podEventProcessor.OnUpdateArgsForCall(0)
				Expect(receivedPod.Name).To(Equal("test-pod"))
				Expect(receivedPod.Namespace).To(Equal("test-namespace"))
			})
		})

		Context("when receiving Modified event", func() {
			var pod corev1.Pod

			BeforeEach(func() {
				pod = corev1.Pod{
					ObjectMeta: metav1.ObjectMeta{
						Name:      "test-pod",
						Namespace: "test-namespace",
					},
				}

				podEventProcessor.OnUpdateReturns(nil)

				go func() {
					fakeWatcher.Modify(&pod)
					fakeWatcher.Stop()
				}()
			})

			It("calls OnUpdate on pod event processor", func() {
				Expect(podEventProcessor.OnUpdateCallCount()).To(Equal(1))
				_, receivedPod := podEventProcessor.OnUpdateArgsForCall(0)
				Expect(receivedPod.Name).To(Equal("test-pod"))
			})
		})

		Context("when receiving Deleted event", func() {
			var pod corev1.Pod

			BeforeEach(func() {
				pod = corev1.Pod{
					ObjectMeta: metav1.ObjectMeta{
						Name:      "test-pod",
						Namespace: "test-namespace",
					},
				}

				podEventProcessor.OnDeleteReturns(nil)

				go func() {
					fakeWatcher.Delete(&pod)
					fakeWatcher.Stop()
				}()
			})

			It("calls OnDelete on pod event processor", func() {
				Expect(podEventProcessor.OnDeleteCallCount()).To(Equal(1))
				_, receivedPod := podEventProcessor.OnDeleteArgsForCall(0)
				Expect(receivedPod.Name).To(Equal("test-pod"))
			})
		})

		Context("when OnUpdate returns an error", func() {
			var pod corev1.Pod

			BeforeEach(func() {
				pod = corev1.Pod{
					ObjectMeta: metav1.ObjectMeta{
						Name:      "test-pod",
						Namespace: "test-namespace",
					},
				}

				podEventProcessor.OnUpdateReturns(errors.New("update failed"))

				go func() {
					fakeWatcher.Add(&pod)
				}()
			})

			It("returns an error", func() {
				Expect(err).NotTo(BeNil())
				Expect(err.Error()).To(ContainSubstring("on update failed"))
			})
		})

		Context("when OnDelete returns an error", func() {
			var pod corev1.Pod

			BeforeEach(func() {
				pod = corev1.Pod{
					ObjectMeta: metav1.ObjectMeta{
						Name:      "test-pod",
						Namespace: "test-namespace",
					},
				}

				podEventProcessor.OnDeleteReturns(errors.New("delete failed"))

				go func() {
					fakeWatcher.Delete(&pod)
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
				Expect(podEventProcessor.OnUpdateCallCount()).To(Equal(0))
				Expect(podEventProcessor.OnDeleteCallCount()).To(Equal(0))
			})
		})

		Context("when result channel is closed", func() {
			BeforeEach(func() {
				go func() {
					fakeWatcher.Stop()
				}()
			})

			It("returns ErrResultChannelClosed", func() {
				Expect(err).To(MatchError(k8s.ErrResultChannelClosed))
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

		Context("when objects already exist before the watch starts", func() {
			BeforeEach(func() {
				podInterface.ListReturns(&corev1.PodList{
					Items: []corev1.Pod{
						{
							ObjectMeta: metav1.ObjectMeta{
								Name:      "existing-pod",
								Namespace: "test-namespace",
							},
						},
					},
				}, nil)
				podEventProcessor.OnUpdateReturns(nil)
				go func() {
					fakeWatcher.Stop()
				}()
			})

			It("delivers the existing pod", func() {
				Expect(podEventProcessor.OnUpdateCallCount()).To(Equal(1))
				_, receivedPod := podEventProcessor.OnUpdateArgsForCall(0)
				Expect(receivedPod.Name).To(Equal("existing-pod"))
			})
		})

		Context("when the initial list fails", func() {
			BeforeEach(func() {
				podInterface.ListReturns(nil, errors.New("list failed"))
			})

			It("returns an error", func() {
				Expect(err).NotTo(BeNil())
				Expect(err.Error()).To(ContainSubstring("list pods failed"))
			})

			It("does not start the watch", func() {
				Expect(podInterface.WatchCallCount()).To(Equal(0))
			})
		})

		Context("when the event processor fails during the initial list", func() {
			BeforeEach(func() {
				podInterface.ListReturns(&corev1.PodList{
					Items: []corev1.Pod{
						{ObjectMeta: metav1.ObjectMeta{Name: "existing-pod"}},
					},
				}, nil)
				podEventProcessor.OnUpdateReturns(errors.New("update failed"))
			})

			It("returns an error", func() {
				Expect(err).NotTo(BeNil())
				Expect(err.Error()).To(ContainSubstring("on update failed"))
			})
		})

		Context("when the context is canceled during the initial list", func() {
			BeforeEach(func() {
				podInterface.ListReturns(&corev1.PodList{
					Items: []corev1.Pod{
						{ObjectMeta: metav1.ObjectMeta{Name: "first-pod"}},
						{ObjectMeta: metav1.ObjectMeta{Name: "second-pod"}},
					},
				}, nil)
				podEventProcessor.OnUpdateStub = func(ctx context.Context, pod corev1.Pod) error {
					cancel()
					return nil
				}
			})

			It("stops at the cancellation", func() {
				Expect(err).To(MatchError(context.Canceled))
				Expect(podEventProcessor.OnUpdateCallCount()).To(Equal(1))
			})
		})

		Context("when the resync interval elapses", func() {
			BeforeEach(func() {
				podWatcher = k8s.NewPodWatcherWithResync(
					clientset,
					podEventProcessor,
					namespace,
					libtime.Duration(10*time.Millisecond),
				)
				podInterface.ListReturns(&corev1.PodList{
					Items: []corev1.Pod{
						{ObjectMeta: metav1.ObjectMeta{Name: "existing-pod"}},
					},
				}, nil)
				podEventProcessor.OnUpdateReturns(nil)
				go func() {
					time.Sleep(200 * time.Millisecond)
					fakeWatcher.Stop()
				}()
			})

			It("re-delivers the full set on resync", func() {
				Expect(podEventProcessor.OnUpdateCallCount()).To(BeNumerically(">=", 2))
			})
		})

		Context("when the resync interval is non-positive", func() {
			BeforeEach(func() {
				podWatcher = k8s.NewPodWatcherWithResync(
					clientset,
					podEventProcessor,
					namespace,
					libtime.Duration(0),
				)
				podInterface.ListReturns(&corev1.PodList{
					Items: []corev1.Pod{
						{ObjectMeta: metav1.ObjectMeta{Name: "existing-pod"}},
					},
				}, nil)
				podEventProcessor.OnUpdateReturns(nil)
				go func() {
					fakeWatcher.Stop()
				}()
			})

			It("delivers the startup list only", func() {
				Expect(podEventProcessor.OnUpdateCallCount()).To(Equal(1))
			})
		})
	})

	Describe("Constructor", func() {
		It("creates a new pod watcher", func() {
			watcher := k8s.NewPodWatcher(
				clientset,
				podEventProcessor,
				namespace,
			)
			Expect(watcher).NotTo(BeNil())
		})
	})
})
