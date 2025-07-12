// Copyright (c) 2024 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package k8s_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/bborbe/k8s"
)

var _ = Describe("Namespace", func() {
	Context("String", func() {
		It("returns the namespace as string", func() {
			namespace := k8s.Namespace("default")
			Expect(namespace.String()).To(Equal("default"))
		})

		It("handles empty namespace", func() {
			namespace := k8s.Namespace("")
			Expect(namespace.String()).To(Equal(""))
		})

		It("handles namespaces with special characters", func() {
			namespace := k8s.Namespace("my-namespace-123")
			Expect(namespace.String()).To(Equal("my-namespace-123"))
		})
	})
})
