// Copyright (c) 2024 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package k8s_test

import (
	"context"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"k8s.io/api/core/v1"

	"github.com/bborbe/k8s"
)

var _ = Describe("Container Builder", func() {
	var containerBuilder k8s.ContainerBuilder
	var err error
	var ctx context.Context
	var container *v1.Container
	BeforeEach(func() {
		ctx = context.Background()
		containerBuilder = k8s.NewContainerBuilder()
		containerBuilder.SetName("my-container")
	})
	Context("Build", func() {
		JustBeforeEach(func() {
			container, err = containerBuilder.Build(ctx)
		})
		It("returns no error", func() {
			Expect(err).To(BeNil())
		})
		It("returns container", func() {
			Expect(container).NotTo(BeNil())
		})
	})
})
