// Copyright (c) 2025 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package k8s_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/bborbe/k8s"
)

var _ = Describe("Containers", func() {
	Context("ParseContainersFromString", func() {
		It("parses a single container", func() {
			containers := k8s.ParseContainersFromString("my-container")
			Expect(containers).To(HaveLen(1))
			Expect(containers[0]).To(Equal(k8s.Container("my-container")))
		})
		It("parses multiple containers", func() {
			containers := k8s.ParseContainersFromString("container-a,container-b")
			Expect(containers).To(HaveLen(2))
			Expect(containers[0]).To(Equal(k8s.Container("container-a")))
			Expect(containers[1]).To(Equal(k8s.Container("container-b")))
		})
		It("handles empty string", func() {
			containers := k8s.ParseContainersFromString("")
			Expect(containers).To(HaveLen(0))
		})
	})

	Context("ParseContainers", func() {
		It("parses a slice of strings", func() {
			containers := k8s.ParseContainers([]string{"container-a", "container-b"})
			Expect(containers).To(HaveLen(2))
			Expect(containers[0]).To(Equal(k8s.Container("container-a")))
			Expect(containers[1]).To(Equal(k8s.Container("container-b")))
		})
		It("handles empty slice", func() {
			containers := k8s.ParseContainers([]string{})
			Expect(containers).To(HaveLen(0))
		})
	})

	Context("Contains", func() {
		var containers k8s.Containers
		BeforeEach(func() {
			containers = k8s.Containers{"container-a", "container-b"}
		})
		It("returns true if container exists", func() {
			Expect(containers.Contains("container-a")).To(BeTrue())
		})
		It("returns false if container does not exist", func() {
			Expect(containers.Contains("container-c")).To(BeFalse())
		})
	})
})
