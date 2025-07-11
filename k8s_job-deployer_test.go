// Copyright (c) 2024 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package k8s_test

import (
	"context"
	"errors"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	batchv1 "k8s.io/api/batch/v1"
	corev1 "k8s.io/api/core/v1"
	k8s_errors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	"github.com/bborbe/k8s"
	"github.com/bborbe/k8s/mocks"
)

var _ = Describe("Job Deployer", func() {
	var jobDeployer k8s.JobDeployer
	var clientset *mocks.K8sInterface
	var batchV1 *mocks.K8sBatchV1Interface
	var jobInterface *mocks.K8sJobInterface
	var ctx context.Context
	var job batchv1.Job
	var err error

	BeforeEach(func() {
		ctx = context.Background()
		clientset = &mocks.K8sInterface{}
		batchV1 = &mocks.K8sBatchV1Interface{}
		jobInterface = &mocks.K8sJobInterface{}

		clientset.BatchV1Returns(batchV1)
		batchV1.JobsReturns(jobInterface)

		jobDeployer = k8s.NewJobDeployer(clientset)

		job = batchv1.Job{
			ObjectMeta: metav1.ObjectMeta{
				Name:      "test-job",
				Namespace: "test-namespace",
			},
			Spec: batchv1.JobSpec{
				Template: corev1.PodTemplateSpec{
					Spec: corev1.PodSpec{
						Containers: []corev1.Container{
							{
								Name:  "test-container",
								Image: "test-image:latest",
							},
						},
						RestartPolicy: corev1.RestartPolicyOnFailure,
					},
				},
			},
		}
	})

	Describe("Deploy", func() {
		JustBeforeEach(func() {
			err = jobDeployer.Deploy(ctx, job)
		})

		Context("when job does not exist", func() {
			BeforeEach(func() {
				// Get call returns NotFound error (job doesn't exist)
				notFoundError := k8s_errors.NewNotFound(schema.GroupResource{
					Group:    "batch",
					Resource: "jobs",
				}, "test-job")
				jobInterface.GetReturns(nil, notFoundError)

				// Create call succeeds
				jobInterface.CreateReturns(&job, nil)
			})

			It("returns no error", func() {
				Expect(err).To(BeNil())
			})

			It("calls Get to check if job exists", func() {
				Expect(jobInterface.GetCallCount()).To(Equal(1))
				_, name, _ := jobInterface.GetArgsForCall(0)
				Expect(name).To(Equal("test-job"))
			})

			It("calls Create to create the job", func() {
				Expect(jobInterface.CreateCallCount()).To(Equal(1))
				_, createdJob, _ := jobInterface.CreateArgsForCall(0)
				Expect(createdJob.Name).To(Equal("test-job"))
				Expect(createdJob.Namespace).To(Equal("test-namespace"))
			})

			It("does not call Delete", func() {
				Expect(jobInterface.DeleteCallCount()).To(Equal(0))
			})
		})

		Context("when job already exists", func() {
			BeforeEach(func() {
				// Get call returns existing job (job exists)
				existingJob := job
				existingJob.ResourceVersion = "123"
				jobInterface.GetReturns(&existingJob, nil)
			})

			It("returns JobAlreadyExistsError", func() {
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(ContainSubstring("job already exists"))
			})

			It("calls Get to check if job exists", func() {
				Expect(jobInterface.GetCallCount()).To(Equal(1))
				_, name, _ := jobInterface.GetArgsForCall(0)
				Expect(name).To(Equal("test-job"))
			})

			It("does not call Create", func() {
				Expect(jobInterface.CreateCallCount()).To(Equal(0))
			})

			It("does not call Delete", func() {
				Expect(jobInterface.DeleteCallCount()).To(Equal(0))
			})
		})

		Context("when create fails", func() {
			BeforeEach(func() {
				// Get call returns NotFound error (job doesn't exist)
				notFoundError := k8s_errors.NewNotFound(schema.GroupResource{
					Group:    "batch",
					Resource: "jobs",
				}, "test-job")
				jobInterface.GetReturns(nil, notFoundError)

				// Create fails
				jobInterface.CreateReturns(nil, errors.New("create failed"))
			})

			It("returns an error", func() {
				Expect(err).NotTo(BeNil())
				Expect(err.Error()).To(ContainSubstring("create job failed"))
			})

			It("calls Get to check if job exists", func() {
				Expect(jobInterface.GetCallCount()).To(Equal(1))
			})

			It("calls Create", func() {
				Expect(jobInterface.CreateCallCount()).To(Equal(1))
			})

			It("does not call Delete", func() {
				Expect(jobInterface.DeleteCallCount()).To(Equal(0))
			})
		})

		Context("with context cancellation", func() {
			var cancelCtx context.Context
			var cancel context.CancelFunc

			BeforeEach(func() {
				cancelCtx, cancel = context.WithCancel(ctx)
				cancel() // Cancel immediately
				ctx = cancelCtx

				jobInterface.GetReturns(nil, context.Canceled)
			})

			It("handles context cancellation gracefully", func() {
				Expect(err).To(BeNil())
				Expect(jobInterface.CreateCallCount()).To(Equal(1))
				// Get returns context.Canceled, so undeploy treats it as not found
				// Then create is called, which would also fail with context.Canceled
			})
		})
	})

	Describe("Undeploy", func() {
		var namespace k8s.Namespace
		var name k8s.Name

		BeforeEach(func() {
			namespace = k8s.Namespace("test-namespace")
			name = k8s.Name("test-job")
		})

		JustBeforeEach(func() {
			err = jobDeployer.Undeploy(ctx, namespace, name)
		})

		Context("when job exists", func() {
			BeforeEach(func() {
				jobInterface.GetReturns(&job, nil)
				jobInterface.DeleteReturns(nil)
			})

			It("returns no error", func() {
				Expect(err).To(BeNil())
			})

			It("calls Get to check if job exists", func() {
				Expect(jobInterface.GetCallCount()).To(Equal(1))
				_, jobName, _ := jobInterface.GetArgsForCall(0)
				Expect(jobName).To(Equal("test-job"))
			})

			It("calls Delete to remove the job", func() {
				Expect(jobInterface.DeleteCallCount()).To(Equal(1))
				_, deletedName, opts := jobInterface.DeleteArgsForCall(0)
				Expect(deletedName).To(Equal("test-job"))
				Expect(opts.PropagationPolicy).NotTo(BeNil())
				Expect(*opts.PropagationPolicy).To(Equal(metav1.DeletePropagationForeground))
			})
		})

		Context("when job does not exist", func() {
			BeforeEach(func() {
				notFoundError := k8s_errors.NewNotFound(schema.GroupResource{
					Group:    "batch",
					Resource: "jobs",
				}, "test-job")
				jobInterface.GetReturns(nil, notFoundError)
			})

			It("returns no error", func() {
				Expect(err).To(BeNil())
			})

			It("calls Get to check if job exists", func() {
				Expect(jobInterface.GetCallCount()).To(Equal(1))
			})

			It("does not call Delete", func() {
				Expect(jobInterface.DeleteCallCount()).To(Equal(0))
			})
		})

		Context("when Get returns an unexpected error", func() {
			BeforeEach(func() {
				jobInterface.GetReturns(nil, errors.New("unexpected error"))
			})

			It("does not call Delete", func() {
				Expect(jobInterface.DeleteCallCount()).To(Equal(0))
			})

			It("returns no error (treats as not found)", func() {
				Expect(err).To(BeNil())
			})
		})

		Context("when Delete fails", func() {
			BeforeEach(func() {
				jobInterface.GetReturns(&job, nil)
				jobInterface.DeleteReturns(errors.New("delete failed"))
			})

			It("returns an error", func() {
				Expect(err).NotTo(BeNil())
				Expect(err.Error()).To(ContainSubstring("delete failed"))
			})

			It("calls Delete", func() {
				Expect(jobInterface.DeleteCallCount()).To(Equal(1))
			})
		})

		Context("with context cancellation", func() {
			var cancelCtx context.Context
			var cancel context.CancelFunc

			BeforeEach(func() {
				cancelCtx, cancel = context.WithCancel(ctx)
				cancel() // Cancel immediately
				ctx = cancelCtx

				jobInterface.GetReturns(nil, context.Canceled)
			})

			It("handles context cancellation gracefully", func() {
				Expect(err).To(BeNil())
				Expect(jobInterface.DeleteCallCount()).To(Equal(0))
			})
		})
	})

	Describe("Constructor", func() {
		It("creates a new job deployer", func() {
			deployer := k8s.NewJobDeployer(clientset)
			Expect(deployer).NotTo(BeNil())
		})
	})

	Describe("Deploy behavior", func() {
		Context("when deploying multiple times", func() {
			BeforeEach(func() {
				// First deployment - job doesn't exist
				notFoundError := k8s_errors.NewNotFound(schema.GroupResource{
					Group:    "batch",
					Resource: "jobs",
				}, "test-job")
				jobInterface.GetReturnsOnCall(0, nil, notFoundError)
				jobInterface.CreateReturnsOnCall(0, &job, nil)

				// Second deployment - job now exists
				existingJob := job
				existingJob.ResourceVersion = "123"
				jobInterface.GetReturnsOnCall(1, &existingJob, nil)
			})

			It("handles multiple deployments correctly", func() {
				// First deployment succeeds
				err1 := jobDeployer.Deploy(ctx, job)
				Expect(err1).To(BeNil())

				// Second deployment fails with JobAlreadyExistsError
				err2 := jobDeployer.Deploy(ctx, job)
				Expect(err2).To(HaveOccurred())
				Expect(err2.Error()).To(ContainSubstring("job already exists"))

				// Should have called Get twice, Create once, no Delete calls
				Expect(jobInterface.GetCallCount()).To(Equal(2))
				Expect(jobInterface.DeleteCallCount()).To(Equal(0))
				Expect(jobInterface.CreateCallCount()).To(Equal(1))
			})
		})

		Context("with different job specifications", func() {
			BeforeEach(func() {
				// Job doesn't exist
				notFoundError := k8s_errors.NewNotFound(schema.GroupResource{
					Group:    "batch",
					Resource: "jobs",
				}, "test-job")
				jobInterface.GetReturns(nil, notFoundError)

				// Modify job spec
				job.Spec.Parallelism = &[]int32{2}[0]
				job.Spec.Completions = &[]int32{3}[0]
				jobInterface.CreateReturns(&job, nil)
			})

			JustBeforeEach(func() {
				err = jobDeployer.Deploy(ctx, job)
			})

			It("deploys job with custom specifications", func() {
				Expect(err).To(BeNil())
				Expect(jobInterface.CreateCallCount()).To(Equal(1))
				_, createdJob, _ := jobInterface.CreateArgsForCall(0)
				Expect(*createdJob.Spec.Parallelism).To(Equal(int32(2)))
				Expect(*createdJob.Spec.Completions).To(Equal(int32(3)))
			})
		})
	})

	Describe("Undeploy with PropagationPolicy", func() {
		JustBeforeEach(func() {
			err = jobDeployer.Undeploy(ctx, "test-namespace", "test-job")
		})

		Context("when deleting job", func() {
			BeforeEach(func() {
				jobInterface.GetReturns(&job, nil)
				jobInterface.DeleteReturns(nil)
			})

			It("uses foreground propagation policy", func() {
				Expect(err).To(BeNil())
				Expect(jobInterface.DeleteCallCount()).To(Equal(1))
				_, _, opts := jobInterface.DeleteArgsForCall(0)
				Expect(opts.PropagationPolicy).NotTo(BeNil())
				Expect(*opts.PropagationPolicy).To(Equal(metav1.DeletePropagationForeground))
			})
		})
	})
})
