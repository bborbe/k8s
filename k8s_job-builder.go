// Copyright (c) 2024 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package k8s

import (
	"context"

	"github.com/bborbe/collection"
	"github.com/bborbe/errors"
	"github.com/bborbe/validation"
	batchv1 "k8s.io/api/batch/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

//counterfeiter:generate -o mocks/k8s-job-builder.go --fake-name K8sJobBuilder . JobBuilder
type JobBuilder interface {
	Build(ctx context.Context) (*batchv1.Job, error)
	SetBackoffLimit(backoffLimit int32) JobBuilder
	SetComponent(component string) JobBuilder
	SetCompletions(completions int32) JobBuilder
	AddLabel(key, value string) JobBuilder
	SetLabels(labels map[string]string) JobBuilder
	SetApp(app string) JobBuilder
	SetObjectMeta(objectMeta metav1.ObjectMeta) JobBuilder
	SetParallelism(parallelism int32) JobBuilder
	SetPodSpec(podSpec corev1.PodSpec) JobBuilder
}

func NewJobBuilder() JobBuilder {
	return &jobBuilder{
		labels:       map[string]string{},
		backoffLimit: collection.Ptr(int32(4)),
		completions:  collection.Ptr(int32(1)),
		parallelism:  collection.Ptr(int32(1)),
	}
}

type jobBuilder struct {
	objectMeta   metav1.ObjectMeta
	component    string
	labels       map[string]string
	podSpec      corev1.PodSpec
	backoffLimit *int32
	completions  *int32
	parallelism  *int32
}

func (j *jobBuilder) SetPodSpec(podSpec corev1.PodSpec) JobBuilder {
	j.podSpec = podSpec
	return j
}

func (j *jobBuilder) SetObjectMeta(objectMeta metav1.ObjectMeta) JobBuilder {
	j.objectMeta = objectMeta
	return j
}

func (j *jobBuilder) SetApp(app string) JobBuilder {
	return j.AddLabel("app", app)
}

func (j *jobBuilder) SetComponent(component string) JobBuilder {
	return j.AddLabel("component", component)
}

func (j *jobBuilder) SetLabels(labels map[string]string) JobBuilder {
	j.labels = labels
	return j
}

func (j *jobBuilder) AddLabel(key, value string) JobBuilder {
	j.labels[key] = value
	return j
}

func (j *jobBuilder) SetBackoffLimit(backoffLimit int32) JobBuilder {
	j.backoffLimit = &backoffLimit
	return j
}

func (j *jobBuilder) SetCompletions(completions int32) JobBuilder {
	j.completions = &completions
	return j
}

func (j *jobBuilder) SetParallelism(parallelism int32) JobBuilder {
	j.parallelism = &parallelism
	return j
}

func (j *jobBuilder) Validate(ctx context.Context) error {
	return validation.All{
		validation.Name("ObjectMeta", validation.NotEmptyString(j.objectMeta.Name)),
	}.Validate(ctx)
}

func (j *jobBuilder) Build(ctx context.Context) (*batchv1.Job, error) {
	if err := j.Validate(ctx); err != nil {
		return nil, errors.Wrapf(ctx, err, "validate jobBuilder failed")
	}

	return &batchv1.Job{
		TypeMeta: metav1.TypeMeta{
			Kind:       "Job",
			APIVersion: "batch/v1",
		},
		ObjectMeta: j.objectMeta,
		Spec: batchv1.JobSpec{
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Annotations: map[string]string{},
					Labels:      j.labels,
				},
				Spec: j.podSpec,
			},
			TTLSecondsAfterFinished: collection.Ptr(int32(600)),
			CompletionMode:          collection.Ptr(batchv1.NonIndexedCompletion),
			PodReplacementPolicy:    collection.Ptr(batchv1.TerminatingOrFailed),
			BackoffLimit:            j.backoffLimit,
			Completions:             j.completions,
			Parallelism:             j.parallelism,
		},
	}, nil
}
