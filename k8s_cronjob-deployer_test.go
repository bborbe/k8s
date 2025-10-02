// Copyright (c) 2025 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package k8s_test

import (
	"context"
	stderrors "errors"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	batchv1 "k8s.io/api/batch/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/kubernetes/fake"
	ktesting "k8s.io/client-go/testing"

	"github.com/bborbe/k8s"
	"github.com/bborbe/k8s/mocks"
)

var _ = Describe("CronJob Deployer", func() {
	var deployer k8s.CronJobDeployer
	var fakeClient *fake.Clientset
	var ctx context.Context
	var cronJob batchv1.CronJob

	BeforeEach(func() {
		ctx = context.Background()
		fakeClient = fake.NewSimpleClientset()
		deployer = k8s.NewCronJobDeployer(fakeClient)

		cronJob = batchv1.CronJob{
			ObjectMeta: metav1.ObjectMeta{
				Name:      "test-cronjob",
				Namespace: "default",
			},
			Spec: batchv1.CronJobSpec{
				Schedule: "0 0 * * *",
				JobTemplate: batchv1.JobTemplateSpec{
					Spec: batchv1.JobSpec{
						Template: corev1.PodTemplateSpec{
							Spec: corev1.PodSpec{
								Containers: []corev1.Container{
									{
										Name:  "test-container",
										Image: "nginx:latest",
									},
								},
								RestartPolicy: corev1.RestartPolicyOnFailure,
							},
						},
					},
				},
			},
		}
	})

	Context("Deploy", func() {
		var err error

		JustBeforeEach(func() {
			err = deployer.Deploy(ctx, cronJob)
		})

		Context("when cronjob does not exist", func() {
			It("returns no error", func() {
				Expect(err).To(BeNil())
			})

			It("creates the cronjob", func() {
				Expect(err).To(BeNil())
				actions := fakeClient.Actions()
				Expect(actions).To(HaveLen(2))
				Expect(actions[0].GetVerb()).To(Equal("get"))
				Expect(actions[1].GetVerb()).To(Equal("create"))
			})
		})

		Context("when cronjob already exists", func() {
			BeforeEach(func() {
				_, err := fakeClient.BatchV1().
					CronJobs("default").
					Create(ctx, &cronJob, metav1.CreateOptions{})
				Expect(err).To(BeNil())
			})

			It("returns no error", func() {
				Expect(err).To(BeNil())
			})

			It("updates the cronjob", func() {
				Expect(err).To(BeNil())
				actions := fakeClient.Actions()
				Expect(actions).To(HaveLen(3))
				Expect(actions[0].GetVerb()).To(Equal("create"))
				Expect(actions[1].GetVerb()).To(Equal("get"))
				Expect(actions[2].GetVerb()).To(Equal("update"))
			})
		})

		Context("when create fails", func() {
			BeforeEach(func() {
				fakeClient.PrependReactor(
					"create",
					"cronjobs",
					func(action ktesting.Action) (handled bool, ret runtime.Object, err error) {
						return true, nil, stderrors.New("create failed")
					},
				)
			})

			It("returns error", func() {
				Expect(err).NotTo(BeNil())
				Expect(err.Error()).To(ContainSubstring("create cronjob failed"))
			})
		})

		Context("when update fails", func() {
			BeforeEach(func() {
				_, err := fakeClient.BatchV1().
					CronJobs("default").
					Create(ctx, &cronJob, metav1.CreateOptions{})
				Expect(err).To(BeNil())
				fakeClient.PrependReactor(
					"update",
					"cronjobs",
					func(action ktesting.Action) (handled bool, ret runtime.Object, err error) {
						return true, nil, stderrors.New("update failed")
					},
				)
			})

			It("returns error", func() {
				Expect(err).NotTo(BeNil())
				Expect(err.Error()).To(ContainSubstring("update deployment failed"))
			})
		})
	})

	Context("Undeploy", func() {
		var err error

		JustBeforeEach(func() {
			err = deployer.Undeploy(ctx, "default", "test-cronjob")
		})

		Context("when cronjob exists", func() {
			BeforeEach(func() {
				_, err := fakeClient.BatchV1().
					CronJobs("default").
					Create(ctx, &cronJob, metav1.CreateOptions{})
				Expect(err).To(BeNil())
			})

			It("returns no error", func() {
				Expect(err).To(BeNil())
			})

			It("deletes the cronjob", func() {
				Expect(err).To(BeNil())
				actions := fakeClient.Actions()
				Expect(actions).To(HaveLen(3))
				Expect(actions[0].GetVerb()).To(Equal("create"))
				Expect(actions[1].GetVerb()).To(Equal("get"))
				Expect(actions[2].GetVerb()).To(Equal("delete"))
			})
		})

		Context("when cronjob does not exist", func() {
			It("returns no error", func() {
				Expect(err).To(BeNil())
			})

			It("does not attempt to delete", func() {
				Expect(err).To(BeNil())
				actions := fakeClient.Actions()
				Expect(actions).To(HaveLen(1))
				Expect(actions[0].GetVerb()).To(Equal("get"))
			})
		})

		Context("when delete fails", func() {
			BeforeEach(func() {
				_, err := fakeClient.BatchV1().
					CronJobs("default").
					Create(ctx, &cronJob, metav1.CreateOptions{})
				Expect(err).To(BeNil())
				fakeClient.PrependReactor(
					"delete",
					"cronjobs",
					func(action ktesting.Action) (handled bool, ret runtime.Object, err error) {
						return true, nil, stderrors.New("delete failed")
					},
				)
			})

			It("returns error", func() {
				Expect(err).NotTo(BeNil())
				Expect(err.Error()).To(ContainSubstring("delete failed"))
			})
		})
	})

	Context("with mock clientset", func() {
		var mockClientset *mocks.K8sInterface
		var mockBatchV1 *mocks.K8sBatchV1Interface
		var mockCronJobs *mocks.K8sCronJobInterface

		BeforeEach(func() {
			mockClientset = &mocks.K8sInterface{}
			mockBatchV1 = &mocks.K8sBatchV1Interface{}
			mockCronJobs = &mocks.K8sCronJobInterface{}

			mockClientset.BatchV1Returns(mockBatchV1)
			mockBatchV1.CronJobsReturns(mockCronJobs)

			deployer = k8s.NewCronJobDeployer(mockClientset)
		})

		Context("Deploy", func() {
			It("calls the correct methods", func() {
				err := deployer.Deploy(ctx, cronJob)
				Expect(err).To(BeNil())
				Expect(mockClientset.BatchV1CallCount()).To(Equal(2))
				Expect(mockBatchV1.CronJobsCallCount()).To(Equal(2))
				Expect(mockCronJobs.GetCallCount()).To(Equal(1))
			})
		})

		Context("Undeploy", func() {
			It("calls the correct methods", func() {
				err := deployer.Undeploy(ctx, "default", "test-cronjob")
				Expect(err).To(BeNil())
				Expect(mockClientset.BatchV1CallCount()).To(Equal(2))
				Expect(mockBatchV1.CronJobsCallCount()).To(Equal(2))
				Expect(mockCronJobs.GetCallCount()).To(Equal(1))
			})
		})
	})
})
