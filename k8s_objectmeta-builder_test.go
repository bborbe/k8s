// Copyright (c) 2024 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package k8s_test

import (
	"context"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/bborbe/k8s"
)

var _ = Describe("ObjectMeta Builder", func() {
	var objectMetaBuilder k8s.ObjectMetaBuilder
	var objectMeta *metav1.ObjectMeta
	var err error
	var ctx context.Context
	BeforeEach(func() {
		ctx = context.Background()
		objectMetaBuilder = k8s.NewObjectMetaBuilder()
		objectMetaBuilder.SetName("my-name")
		objectMetaBuilder.SetNamespace("my-namespace")
	})
	Context("Build", func() {
		JustBeforeEach(func() {
			objectMeta, err = objectMetaBuilder.Build(ctx)
		})
		Context("default", func() {
			It("returns no error", func() {
				Expect(err).To(BeNil())
			})
			It("returns objectMeta", func() {
				Expect(objectMeta).NotTo(BeNil())
			})
			It("has no finalizer", func() {
				Expect(objectMeta).NotTo(BeNil())
				Expect(objectMeta.Finalizers).To(HaveLen(0))
			})
		})
		Context("with finalizer", func() {
			BeforeEach(func() {
				objectMetaBuilder = objectMetaBuilder.SetFinalizers([]string{"my-finalizer"})
			})
			It("returns no error", func() {
				Expect(err).To(BeNil())
			})
			It("returns objectMeta", func() {
				Expect(objectMeta).NotTo(BeNil())
			})
			It("has finalizer", func() {
				Expect(objectMeta).NotTo(BeNil())
				Expect(objectMeta.Finalizers).To(Equal([]string{"my-finalizer"}))
			})
		})
	})

})
