// Copyright (c) 2025 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package k8s_test

import (
	"context"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	batchv1 "k8s.io/api/batch/v1"
	corev1 "k8s.io/api/core/v1"

	"github.com/bborbe/k8s"
)

var _ = Describe("CronJob Builder", func() {
	var cronJobBuilder k8s.CronJobBuilder
	var cronJob *batchv1.CronJob
	var err error
	var ctx context.Context
	var objectMetaBuilder k8s.ObjectMetaBuilder
	var podSpecBuilder k8s.PodSpecBuilder
	var containersBuilder k8s.ContainersBuilder
	BeforeEach(func() {
		ctx = context.Background()

		objectMetaBuilder = k8s.NewObjectMetaBuilder()
		objectMetaBuilder.SetName("my-object")
		objectMetaBuilder.SetNamespace("my-namespace")

		containersBuilder = k8s.NewContainersBuilder()
		containersBuilder.SetContainers([]corev1.Container{
			{
				Name: "service",
			},
		})

		podSpecBuilder = k8s.NewPodSpecBuilder()
		podSpecBuilder.SetContainersBuilder(containersBuilder)
		podSpecBuilder.SetRestartPolicy(corev1.RestartPolicyOnFailure)

		cronJobBuilder = k8s.NewCronJobBuilder()
		cronJobBuilder.SetObjectMetaBuild(objectMetaBuilder)
		cronJobBuilder.SetPodSpecBuilder(podSpecBuilder)
		cronJobBuilder.SetCronExpression("0 0 * * *")
	})
	Context("Build", func() {
		JustBeforeEach(func() {
			cronJob, err = cronJobBuilder.Build(ctx)
		})
		It("returns no error", func() {
			Expect(err).To(BeNil())
		})
		It("returns cronjob", func() {
			Expect(cronJob).NotTo(BeNil())
		})
		It("sets correct TypeMeta", func() {
			Expect(cronJob).NotTo(BeNil())
			Expect(cronJob.TypeMeta.Kind).To(Equal("CronJob"))
			Expect(cronJob.TypeMeta.APIVersion).To(Equal("batch/v1"))
		})
		It("sets correct ObjectMeta", func() {
			Expect(cronJob).NotTo(BeNil())
			Expect(cronJob.ObjectMeta.Name).To(Equal("my-object"))
			Expect(cronJob.ObjectMeta.Namespace).To(Equal("my-namespace"))
		})
		It("sets correct schedule", func() {
			Expect(cronJob).NotTo(BeNil())
			Expect(cronJob.Spec.Schedule).To(Equal("0 0 * * *"))
		})
		It("sets default values", func() {
			Expect(cronJob).NotTo(BeNil())
			Expect(*cronJob.Spec.SuccessfulJobsHistoryLimit).To(Equal(int32(1)))
			Expect(*cronJob.Spec.FailedJobsHistoryLimit).To(Equal(int32(2)))
		})

		Context("with custom configuration", func() {
			BeforeEach(func() {
				cronJobBuilder.SetCronExpression("0 */6 * * *")
				cronJobBuilder.SetParallelism(3)
				cronJobBuilder.SetBackoffLimit(10)
				cronJobBuilder.SetCompletions(5)
				cronJobBuilder.SetImage("nginx:latest")
				cronJobBuilder.SetLoglevel(2)
				cronJobBuilder.SetEnv([]corev1.EnvVar{
					{Name: "ENV_VAR", Value: "test"},
				})
			})

			It("sets custom cron expression", func() {
				Expect(cronJob.Spec.Schedule).To(Equal("0 */6 * * *"))
			})

			It("sets custom job spec values", func() {
				Expect(*cronJob.Spec.JobTemplate.Spec.Parallelism).To(Equal(int32(3)))
				Expect(*cronJob.Spec.JobTemplate.Spec.BackoffLimit).To(Equal(int32(10)))
				Expect(*cronJob.Spec.JobTemplate.Spec.Completions).To(Equal(int32(5)))
			})
		})
	})

	Context("validation", func() {
		JustBeforeEach(func() {
			err = cronJobBuilder.Validate(ctx)
		})
		It("returns no error", func() {
			Expect(err).To(BeNil())
		})
		Context("without ObjectMeta", func() {
			BeforeEach(func() {
				cronJobBuilder.SetObjectMetaBuild(nil)
			})
			It("returns validation error", func() {
				Expect(err).NotTo(BeNil())
			})
		})
		Context("without PodSpec", func() {
			BeforeEach(func() {
				cronJobBuilder.SetPodSpecBuilder(nil)
			})
			It("returns validation error", func() {
				Expect(err).NotTo(BeNil())
			})
		})
	})
})
