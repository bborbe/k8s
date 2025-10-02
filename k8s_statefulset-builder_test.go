// Copyright (c) 2024 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package k8s_test

import (
	"context"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	appsv1 "k8s.io/api/apps/v1"

	"github.com/bborbe/k8s"
)

var _ = Describe("StatefulSet Builder", func() {
	var statefulSetBuilder k8s.StatefulSetBuilder
	var statefulSet *appsv1.StatefulSet
	var err error
	var ctx context.Context
	BeforeEach(func() {
		ctx = context.Background()
		objectMetaBuilder := k8s.NewObjectMetaBuilder()
		objectMetaBuilder.SetName("my-name")
		objectMetaBuilder.SetNamespace("my-namespace")

		containerBuilder := k8s.NewContainerBuilder()
		containerBuilder.SetName("my-container")

		containersBuilder := k8s.NewContainersBuilder()
		containersBuilder.AddContainerBuilder(containerBuilder)

		statefulSetBuilder = k8s.NewStatefulSetBuilder()
		statefulSetBuilder.SetObjectMetaBuilder(objectMetaBuilder)
		statefulSetBuilder.SetContainersBuilder(containersBuilder)
		statefulSetBuilder.SetName("my-name")
	})
	Context("Build", func() {
		JustBeforeEach(func() {
			statefulSet, err = statefulSetBuilder.Build(ctx)
		})
		Context("valid", func() {
			It("returns no error", func() {
				Expect(err).To(BeNil())
			})
			It("returns statefulSet", func() {
				Expect(statefulSet).NotTo(BeNil())
			})
			It("returns correct VolumeClaimTemplates", func() {
				Expect(statefulSet).NotTo(BeNil())
				Expect(statefulSet.Spec.VolumeClaimTemplates).To(HaveLen(1))
				Expect(
					*statefulSet.Spec.VolumeClaimTemplates[0].Spec.StorageClassName,
				).To(Equal("standard"))
			})
			It("returns correct ImagePullSecrets", func() {
				Expect(statefulSet).NotTo(BeNil())
				Expect(statefulSet.Spec.Template.Spec.ImagePullSecrets).To(HaveLen(1))
				Expect(statefulSet.Spec.Template.Spec.ImagePullSecrets[0].Name).To(Equal("docker"))
			})
			Context("SetImagePullSecrets", func() {
				BeforeEach(func() {
					statefulSetBuilder.SetImagePullSecrets([]string{"docker-registry"})
				})
				It("returns correct ImagePullSecrets", func() {
					Expect(statefulSet).NotTo(BeNil())
					Expect(statefulSet.Spec.Template.Spec.ImagePullSecrets).To(HaveLen(1))
					Expect(
						statefulSet.Spec.Template.Spec.ImagePullSecrets[0].Name,
					).To(Equal("docker-registry"))
				})
			})
		})
		Context("without name", func() {
			BeforeEach(func() {
				statefulSetBuilder.SetName("")
			})
			It("returns error", func() {
				Expect(err).NotTo(BeNil())
			})
			It("returns no statefulSet", func() {
				Expect(statefulSet).To(BeNil())
			})
		})
	})
})
