// Copyright (c) 2024 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package k8s_test

import (
	"context"

	"github.com/bborbe/k8s"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"k8s.io/api/networking/v1"
)

var _ = Describe("Ingress Builder", func() {
	var ingressBuilder k8s.IngressBuilder
	var ingress *v1.Ingress
	var err error
	var ctx context.Context
	BeforeEach(func() {
		ctx = context.Background()

		objectMetaBuilder := k8s.NewObjectMetaBuilder()
		objectMetaBuilder.SetName("my-ingress-name")
		objectMetaBuilder.SetNamespace("my-namespace")

		ingressBuilder = k8s.NewIngressBuilder()
		ingressBuilder.SetServiceName("my-service-name")
		ingressBuilder.SetHost("myname.example.com")
		ingressBuilder.SetObjectMetaBuilder(objectMetaBuilder)
	})
	Context("Build", func() {
		JustBeforeEach(func() {
			ingress, err = ingressBuilder.Build(ctx)
		})
		Context("valid", func() {
			It("returns no error", func() {
				Expect(err).To(BeNil())
			})
			It("returns correct ingress", func() {
				Expect(ingress).NotTo(BeNil())
				Expect(ingress.Name).To(Equal("my-ingress-name"))
				Expect(ingress.ObjectMeta.Name).To(Equal("my-ingress-name"))
				Expect(ingress.Spec.Rules[0].Host).To(Equal("myname.example.com"))
				Expect(*ingress.Spec.IngressClassName).To(Equal("traefik"))
				Expect(ingress.Spec.Rules[0].IngressRuleValue.HTTP.Paths[0].Backend.Service.Name).To(Equal("my-service-name"))
			})
		})

		Context("without host", func() {
			BeforeEach(func() {
				ingressBuilder.SetHost("")
			})
			It("returns error", func() {
				Expect(err).NotTo(BeNil())
			})
			It("returns no ingress", func() {
				Expect(ingress).To(BeNil())
			})
		})
	})
})
