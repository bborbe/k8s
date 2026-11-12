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

var _ = Describe("SecretWatcher", func() {
	var secretWatcher k8s.SecretWatcher
	var clientset *mocks.K8sInterface
	var coreV1 *mocks.K8sCoreV1Interface
	var secretInterface *mocks.K8sSecretInterface
	var secretEventProcessor *mocks.K8sSecretEventProcessor
	var ctx context.Context
	var cancel context.CancelFunc
	var namespace k8s.Namespace
	var fakeWatcher *watch.FakeWatcher
	var err error

	BeforeEach(func() {
		ctx, cancel = context.WithCancel(context.Background())
		clientset = &mocks.K8sInterface{}
		coreV1 = &mocks.K8sCoreV1Interface{}
		secretInterface = &mocks.K8sSecretInterface{}
		secretEventProcessor = &mocks.K8sSecretEventProcessor{}
		namespace = k8s.Namespace("test-namespace")
		fakeWatcher = watch.NewFake()

		clientset.CoreV1Returns(coreV1)
		coreV1.SecretsReturns(secretInterface)
		secretInterface.ListReturns(&corev1.SecretList{}, nil)
		secretInterface.WatchReturns(fakeWatcher, nil)

		secretWatcher = k8s.NewSecretWatcher(
			clientset,
			secretEventProcessor,
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
			err = secretWatcher.Watch(ctx)
		})

		Context("when context is already canceled", func() {
			BeforeEach(func() {
				cancel()
			})

			It("returns context.Canceled error", func() {
				Expect(err).To(Equal(context.Canceled))
			})

			It("does not call Watch on secret interface", func() {
				Expect(secretInterface.WatchCallCount()).To(Equal(0))
			})
		})

		Context("when watch fails to start", func() {
			BeforeEach(func() {
				secretInterface.WatchReturns(nil, errors.New("watch failed"))
			})

			It("returns an error", func() {
				Expect(err).NotTo(BeNil())
				Expect(err.Error()).To(ContainSubstring("watch failed"))
			})
		})

		Context("when receiving Added event", func() {
			var secret corev1.Secret

			BeforeEach(func() {
				secret = corev1.Secret{
					ObjectMeta: metav1.ObjectMeta{
						Name:      "test-secret",
						Namespace: "test-namespace",
					},
				}

				secretEventProcessor.OnUpdateReturns(nil)

				// Start watch in background and send event
				go func() {
					fakeWatcher.Add(&secret)
					fakeWatcher.Stop()
				}()
			})

			It("calls OnUpdate on secret event processor", func() {
				Expect(secretEventProcessor.OnUpdateCallCount()).To(Equal(1))
				_, receivedSecret := secretEventProcessor.OnUpdateArgsForCall(0)
				Expect(receivedSecret.Name).To(Equal("test-secret"))
				Expect(receivedSecret.Namespace).To(Equal("test-namespace"))
			})
		})

		Context("when receiving Modified event", func() {
			var secret corev1.Secret

			BeforeEach(func() {
				secret = corev1.Secret{
					ObjectMeta: metav1.ObjectMeta{
						Name:      "test-secret",
						Namespace: "test-namespace",
					},
				}

				secretEventProcessor.OnUpdateReturns(nil)

				go func() {
					fakeWatcher.Modify(&secret)
					fakeWatcher.Stop()
				}()
			})

			It("calls OnUpdate on secret event processor", func() {
				Expect(secretEventProcessor.OnUpdateCallCount()).To(Equal(1))
				_, receivedSecret := secretEventProcessor.OnUpdateArgsForCall(0)
				Expect(receivedSecret.Name).To(Equal("test-secret"))
			})
		})

		Context("when receiving Deleted event", func() {
			var secret corev1.Secret

			BeforeEach(func() {
				secret = corev1.Secret{
					ObjectMeta: metav1.ObjectMeta{
						Name:      "test-secret",
						Namespace: "test-namespace",
					},
				}

				secretEventProcessor.OnDeleteReturns(nil)

				go func() {
					fakeWatcher.Delete(&secret)
					fakeWatcher.Stop()
				}()
			})

			It("calls OnDelete on secret event processor", func() {
				Expect(secretEventProcessor.OnDeleteCallCount()).To(Equal(1))
				_, receivedSecret := secretEventProcessor.OnDeleteArgsForCall(0)
				Expect(receivedSecret.Name).To(Equal("test-secret"))
			})
		})

		Context("when OnUpdate returns an error", func() {
			var secret corev1.Secret

			BeforeEach(func() {
				secret = corev1.Secret{
					ObjectMeta: metav1.ObjectMeta{
						Name:      "test-secret",
						Namespace: "test-namespace",
					},
				}

				secretEventProcessor.OnUpdateReturns(errors.New("update failed"))

				go func() {
					fakeWatcher.Add(&secret)
				}()
			})

			It("returns an error", func() {
				Expect(err).NotTo(BeNil())
				Expect(err.Error()).To(ContainSubstring("on update failed"))
			})
		})

		Context("when OnDelete returns an error", func() {
			var secret corev1.Secret

			BeforeEach(func() {
				secret = corev1.Secret{
					ObjectMeta: metav1.ObjectMeta{
						Name:      "test-secret",
						Namespace: "test-namespace",
					},
				}

				secretEventProcessor.OnDeleteReturns(errors.New("delete failed"))

				go func() {
					fakeWatcher.Delete(&secret)
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
				Expect(secretEventProcessor.OnUpdateCallCount()).To(Equal(0))
				Expect(secretEventProcessor.OnDeleteCallCount()).To(Equal(0))
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
				secretInterface.ListReturns(&corev1.SecretList{
					Items: []corev1.Secret{
						{
							ObjectMeta: metav1.ObjectMeta{
								Name:      "existing-secret",
								Namespace: "test-namespace",
							},
						},
					},
				}, nil)
				secretEventProcessor.OnUpdateReturns(nil)
				go func() {
					fakeWatcher.Stop()
				}()
			})

			It("delivers the existing secret", func() {
				Expect(secretEventProcessor.OnUpdateCallCount()).To(Equal(1))
				_, receivedSecret := secretEventProcessor.OnUpdateArgsForCall(0)
				Expect(receivedSecret.Name).To(Equal("existing-secret"))
			})
		})

		Context("when the initial list fails", func() {
			BeforeEach(func() {
				secretInterface.ListReturns(nil, errors.New("list failed"))
			})

			It("returns an error", func() {
				Expect(err).NotTo(BeNil())
				Expect(err.Error()).To(ContainSubstring("list secrets failed"))
			})

			It("does not start the watch", func() {
				Expect(secretInterface.WatchCallCount()).To(Equal(0))
			})
		})

		Context("when the event processor fails during the initial list", func() {
			BeforeEach(func() {
				secretInterface.ListReturns(&corev1.SecretList{
					Items: []corev1.Secret{
						{ObjectMeta: metav1.ObjectMeta{Name: "existing-secret"}},
					},
				}, nil)
				secretEventProcessor.OnUpdateReturns(errors.New("update failed"))
			})

			It("returns an error", func() {
				Expect(err).NotTo(BeNil())
				Expect(err.Error()).To(ContainSubstring("on update failed"))
			})
		})

		Context("when the context is canceled during the initial list", func() {
			BeforeEach(func() {
				secretInterface.ListReturns(&corev1.SecretList{
					Items: []corev1.Secret{
						{ObjectMeta: metav1.ObjectMeta{Name: "first-secret"}},
						{ObjectMeta: metav1.ObjectMeta{Name: "second-secret"}},
					},
				}, nil)
				secretEventProcessor.OnUpdateStub = func(ctx context.Context, secret corev1.Secret) error {
					cancel()
					return nil
				}
			})

			It("stops at the cancellation", func() {
				Expect(err).To(MatchError(context.Canceled))
				Expect(secretEventProcessor.OnUpdateCallCount()).To(Equal(1))
			})
		})

		Context("when the resync interval elapses", func() {
			BeforeEach(func() {
				secretWatcher = k8s.NewSecretWatcherWithResync(
					clientset,
					secretEventProcessor,
					namespace,
					libtime.Duration(10*time.Millisecond),
				)
				secretInterface.ListReturns(&corev1.SecretList{
					Items: []corev1.Secret{
						{ObjectMeta: metav1.ObjectMeta{Name: "existing-secret"}},
					},
				}, nil)
				secretEventProcessor.OnUpdateReturns(nil)
				go func() {
					time.Sleep(200 * time.Millisecond)
					fakeWatcher.Stop()
				}()
			})

			It("re-delivers the full set on resync", func() {
				Expect(secretEventProcessor.OnUpdateCallCount()).To(BeNumerically(">=", 2))
			})
		})

		Context("when the resync interval is non-positive", func() {
			BeforeEach(func() {
				secretWatcher = k8s.NewSecretWatcherWithResync(
					clientset,
					secretEventProcessor,
					namespace,
					libtime.Duration(0),
				)
				secretInterface.ListReturns(&corev1.SecretList{
					Items: []corev1.Secret{
						{ObjectMeta: metav1.ObjectMeta{Name: "existing-secret"}},
					},
				}, nil)
				secretEventProcessor.OnUpdateReturns(nil)
				go func() {
					fakeWatcher.Stop()
				}()
			})

			It("delivers the startup list only", func() {
				Expect(secretEventProcessor.OnUpdateCallCount()).To(Equal(1))
			})
		})
	})

	Describe("Constructor", func() {
		It("creates a new secret watcher", func() {
			watcher := k8s.NewSecretWatcher(
				clientset,
				secretEventProcessor,
				namespace,
			)
			Expect(watcher).NotTo(BeNil())
		})
	})
})
