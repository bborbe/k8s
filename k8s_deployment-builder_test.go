// Copyright (c) 2024 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package k8s_test

import (
	"context"

	"github.com/bborbe/k8s"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	appsv1 "k8s.io/api/apps/v1"
)

var _ = Describe("Deployment Builder", func() {
	var deploymentBuilder k8s.DeploymentBuilder
	var deployment *appsv1.Deployment
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

		deploymentBuilder = k8s.NewDeploymentBuilder()
		deploymentBuilder.SetObjectMetaBuilder(objectMetaBuilder)
		deploymentBuilder.SetContainersBuilder(containersBuilder)
	})
	Context("Build", func() {
		JustBeforeEach(func() {
			deployment, err = deploymentBuilder.Build(ctx)
		})
		It("returns no error", func() {
			Expect(err).To(BeNil())
		})
		It("returns deployment", func() {
			Expect(deployment).NotTo(BeNil())
		})
		It("returns correct ImagePullSecrets", func() {
			Expect(deployment).NotTo(BeNil())
			Expect(deployment.Spec.Template.Spec.ImagePullSecrets).To(HaveLen(1))
			Expect(deployment.Spec.Template.Spec.ImagePullSecrets[0].Name).To(Equal("docker"))
		})
	})
})
