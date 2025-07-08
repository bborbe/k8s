// Copyright (c) 2024 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package k8s_test

import (
	"context"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	corev1 "k8s.io/api/core/v1"

	"github.com/bborbe/k8s"
)

var _ = Describe("PodSpec Builder", func() {
	var podSpecBuilder k8s.PodSpecBuilder
	var podSpec *corev1.PodSpec
	var err error
	var ctx context.Context
	BeforeEach(func() {
		ctx = context.Background()
		podSpecBuilder = k8s.NewPodSpecBuilder()
	})
	JustBeforeEach(func() {
		podSpec, err = podSpecBuilder.Build(ctx)
	})
	It("returns no error", func() {
		Expect(err).To(BeNil())
	})
	It("returns podSpec", func() {
		Expect(podSpec).NotTo(BeNil())
		Expect(podSpec.ImagePullSecrets).To(HaveLen(1))
		Expect(podSpec.ImagePullSecrets[0].Name).To(Equal("docker"))
	})
	Context("custom pull secret", func() {
		BeforeEach(func() {
			podSpecBuilder.SetImagePullSecrets([]string{"docker-test"})
		})
		It("returns podSpec", func() {
			Expect(podSpec).NotTo(BeNil())
			Expect(podSpec.ImagePullSecrets).To(HaveLen(1))
			Expect(podSpec.ImagePullSecrets[0].Name).To(Equal("docker-test"))
		})
	})
	Context("with priority class", func() {
		BeforeEach(func() {
			podSpecBuilder.SetPriorityClassName("my-class")
		})
		It("returns podSpec with priority class", func() {
			Expect(podSpec).NotTo(BeNil())
			Expect(podSpec.PriorityClassName).To(Equal("my-class"))
		})
	})
})
