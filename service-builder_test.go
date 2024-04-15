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

var _ = Describe("Service Builder", func() {
	var serviceBuilder k8s.ServiceBuilder
	var service *corev1.Service
	var err error
	var ctx context.Context
	BeforeEach(func() {
		ctx = context.Background()

		objectMetaBuilder := k8s.NewObjectMetaBuilder()
		objectMetaBuilder.SetName("my-ingress-name")
		objectMetaBuilder.SetNamespace("my-namespace")

		serviceBuilder = k8s.NewServiceBuilder()
		serviceBuilder.SetObjectMetaBuilder(objectMetaBuilder)
		serviceBuilder.SetName("my-service")
	})
	Context("Build", func() {
		JustBeforeEach(func() {
			service, err = serviceBuilder.Build(ctx)
		})
		Context("Valid", func() {
			It("returns no error", func() {
				Expect(err).To(BeNil())
			})
			It("returns service", func() {
				Expect(service).NotTo(BeNil())
			})
		})
	})
})
