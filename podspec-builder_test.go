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
		podSpec, err = podSpecBuilder.Build(ctx)
	})
	It("returns no error", func() {
		Expect(err).To(BeNil())
	})
	It("returns podSpec", func() {
		Expect(podSpec).NotTo(BeNil())
	})
})
