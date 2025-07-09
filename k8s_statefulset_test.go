// Copyright (c) 2025 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package k8s_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/bborbe/k8s"
)

var _ = Describe("StatefulSets", func() {
	Context("ParseStatefulSetsFromString", func() {
		It("parses a single statefulset", func() {
			statefulsets := k8s.ParseStatefulSetsFromString("my-statefulset")
			Expect(statefulsets).To(HaveLen(1))
			Expect(statefulsets[0]).To(Equal(k8s.StatefulSet("my-statefulset")))
		})
		It("parses multiple statefulsets", func() {
			statefulsets := k8s.ParseStatefulSetsFromString("statefulset-a,statefulset-b")
			Expect(statefulsets).To(HaveLen(2))
			Expect(statefulsets[0]).To(Equal(k8s.StatefulSet("statefulset-a")))
			Expect(statefulsets[1]).To(Equal(k8s.StatefulSet("statefulset-b")))
		})
		It("handles empty string", func() {
			statefulsets := k8s.ParseStatefulSetsFromString("")
			Expect(statefulsets).To(HaveLen(0))
		})
	})

	Context("ParseStatefulSets", func() {
		It("parses a slice of strings", func() {
			statefulsets := k8s.ParseStatefulSets([]string{"statefulset-a", "statefulset-b"})
			Expect(statefulsets).To(HaveLen(2))
			Expect(statefulsets[0]).To(Equal(k8s.StatefulSet("statefulset-a")))
			Expect(statefulsets[1]).To(Equal(k8s.StatefulSet("statefulset-b")))
		})
		It("handles empty slice", func() {
			statefulsets := k8s.ParseStatefulSets([]string{})
			Expect(statefulsets).To(HaveLen(0))
		})
	})

	Context("Contains", func() {
		var statefulsets k8s.StatefulSets
		BeforeEach(func() {
			statefulsets = k8s.StatefulSets{"statefulset-a", "statefulset-b"}
		})
		It("returns true if statefulset exists", func() {
			Expect(statefulsets.Contains("statefulset-a")).To(BeTrue())
		})
		It("returns false if statefulset does not exist", func() {
			Expect(statefulsets.Contains("statefulset-c")).To(BeFalse())
		})
	})
})
