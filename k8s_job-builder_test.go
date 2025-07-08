// Copyright (c) 2024 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package k8s_test

import (
	"context"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	batchv1 "k8s.io/api/batch/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/bborbe/k8s"
)

var _ = Describe("Job Builder", func() {
	var jobBuilder k8s.JobBuilder
	var job *batchv1.Job
	var err error
	var ctx context.Context

	BeforeEach(func() {
		ctx = context.Background()
		jobBuilder = k8s.NewJobBuilder()
	})

	Context("Build", func() {
		JustBeforeEach(func() {
			job, err = jobBuilder.Build(ctx)
		})

		Context("with minimal configuration", func() {
			BeforeEach(func() {
				objectMeta := metav1.ObjectMeta{
					Name:      "test-job",
					Namespace: "default",
				}
				jobBuilder.SetObjectMeta(objectMeta)
				jobBuilder.SetName("test-job")
			})

			It("returns no error", func() {
				Expect(err).To(BeNil())
			})

			It("returns job", func() {
				Expect(job).NotTo(BeNil())
			})

			It("sets correct TypeMeta", func() {
				Expect(job.TypeMeta.Kind).To(Equal("Job"))
				Expect(job.TypeMeta.APIVersion).To(Equal("batch/v1"))
			})

			It("sets correct ObjectMeta", func() {
				Expect(job.ObjectMeta.Name).To(Equal("test-job"))
				Expect(job.ObjectMeta.Namespace).To(Equal("default"))
			})

			It("sets default values", func() {
				Expect(*job.Spec.BackoffLimit).To(Equal(int32(4)))
				Expect(*job.Spec.Completions).To(Equal(int32(1)))
				Expect(*job.Spec.Parallelism).To(Equal(int32(1)))
				Expect(*job.Spec.TTLSecondsAfterFinished).To(Equal(int32(600)))
				Expect(*job.Spec.CompletionMode).To(Equal(batchv1.NonIndexedCompletion))
				Expect(*job.Spec.PodReplacementPolicy).To(Equal(batchv1.TerminatingOrFailed))
			})

			It("sets app label", func() {
				Expect(job.Spec.Template.ObjectMeta.Labels).To(HaveKeyWithValue("app", "test-job"))
			})
		})

		Context("with custom configuration", func() {
			BeforeEach(func() {
				objectMeta := metav1.ObjectMeta{
					Name:      "custom-job",
					Namespace: "custom-namespace",
				}
				podSpec := corev1.PodSpec{
					Containers: []corev1.Container{
						{
							Name:  "test-container",
							Image: "test-image:latest",
						},
					},
					RestartPolicy: corev1.RestartPolicyNever,
				}

				jobBuilder.SetObjectMeta(objectMeta)
				jobBuilder.SetName("custom-job")
				jobBuilder.SetPodSpec(podSpec)
				jobBuilder.SetBackoffLimit(6)
				jobBuilder.SetCompletions(3)
				jobBuilder.SetParallelism(2)
				jobBuilder.SetComponent("worker")
				jobBuilder.AddLabel("environment", "test")
				jobBuilder.AddLabel("version", "v1.0.0")
			})

			It("returns no error", func() {
				Expect(err).To(BeNil())
			})

			It("sets custom backoff limit", func() {
				Expect(*job.Spec.BackoffLimit).To(Equal(int32(6)))
			})

			It("sets custom completions", func() {
				Expect(*job.Spec.Completions).To(Equal(int32(3)))
			})

			It("sets custom parallelism", func() {
				Expect(*job.Spec.Parallelism).To(Equal(int32(2)))
			})

			It("sets custom pod spec", func() {
				Expect(job.Spec.Template.Spec.Containers).To(HaveLen(1))
				Expect(job.Spec.Template.Spec.Containers[0].Name).To(Equal("test-container"))
				Expect(job.Spec.Template.Spec.Containers[0].Image).To(Equal("test-image:latest"))
				Expect(job.Spec.Template.Spec.RestartPolicy).To(Equal(corev1.RestartPolicyNever))
			})

			It("sets custom labels", func() {
				Expect(job.Spec.Template.ObjectMeta.Labels).To(HaveKeyWithValue("app", "custom-job"))
				Expect(job.Spec.Template.ObjectMeta.Labels).To(HaveKeyWithValue("component", "worker"))
				Expect(job.Spec.Template.ObjectMeta.Labels).To(HaveKeyWithValue("environment", "test"))
				Expect(job.Spec.Template.ObjectMeta.Labels).To(HaveKeyWithValue("version", "v1.0.0"))
			})
		})

		Context("with SetLabels", func() {
			BeforeEach(func() {
				objectMeta := metav1.ObjectMeta{
					Name:      "labels-job",
					Namespace: "default",
				}
				labels := map[string]string{
					"team":        "platform",
					"environment": "production",
					"tier":        "backend",
				}

				jobBuilder.SetObjectMeta(objectMeta)
				jobBuilder.SetName("labels-job")
				jobBuilder.SetLabels(labels)
			})

			It("returns no error", func() {
				Expect(err).To(BeNil())
			})

			It("sets all labels including app label", func() {
				Expect(job.Spec.Template.ObjectMeta.Labels).To(HaveKeyWithValue("app", "labels-job"))
				Expect(job.Spec.Template.ObjectMeta.Labels).To(HaveKeyWithValue("team", "platform"))
				Expect(job.Spec.Template.ObjectMeta.Labels).To(HaveKeyWithValue("environment", "production"))
				Expect(job.Spec.Template.ObjectMeta.Labels).To(HaveKeyWithValue("tier", "backend"))
			})
		})

		Context("validation", func() {
			Context("without ObjectMeta", func() {
				BeforeEach(func() {
					jobBuilder.SetName("no-meta-job")
				})

				It("returns validation error", func() {
					Expect(err).To(HaveOccurred())
					Expect(err.Error()).To(ContainSubstring("validate jobBuilder failed"))
				})

				It("returns nil job", func() {
					Expect(job).To(BeNil())
				})
			})
		})
	})

	Describe("Method chaining", func() {
		It("allows fluent interface", func() {
			objectMeta := metav1.ObjectMeta{
				Name:      "chain-job",
				Namespace: "default",
			}
			podSpec := corev1.PodSpec{
				Containers: []corev1.Container{
					{
						Name:  "chain-container",
						Image: "chain-image:latest",
					},
				},
			}

			job, err := k8s.NewJobBuilder().
				SetObjectMeta(objectMeta).
				SetName("chain-job").
				SetPodSpec(podSpec).
				SetBackoffLimit(10).
				SetCompletions(5).
				SetParallelism(3).
				SetComponent("batch").
				AddLabel("priority", "high").
				Build(ctx)

			Expect(err).To(BeNil())
			Expect(job).NotTo(BeNil())
			Expect(*job.Spec.BackoffLimit).To(Equal(int32(10)))
			Expect(*job.Spec.Completions).To(Equal(int32(5)))
			Expect(*job.Spec.Parallelism).To(Equal(int32(3)))
			Expect(job.Spec.Template.ObjectMeta.Labels).To(HaveKeyWithValue("component", "batch"))
			Expect(job.Spec.Template.ObjectMeta.Labels).To(HaveKeyWithValue("priority", "high"))
		})
	})
})
