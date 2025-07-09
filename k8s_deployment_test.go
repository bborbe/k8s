// Copyright (c) 2025 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package k8s_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/bborbe/k8s"
)

var _ = Describe("Deployments", func() {
	Context("ParseDeploymentsFromString", func() {
		It("parses a single deployment", func() {
			deployments := k8s.ParseDeploymentsFromString("my-deployment")
			Expect(deployments).To(HaveLen(1))
			Expect(deployments[0]).To(Equal(k8s.Deployment("my-deployment")))
		})
		It("parses multiple deployments", func() {
			deployments := k8s.ParseDeploymentsFromString("deployment-a,deployment-b")
			Expect(deployments).To(HaveLen(2))
			Expect(deployments[0]).To(Equal(k8s.Deployment("deployment-a")))
			Expect(deployments[1]).To(Equal(k8s.Deployment("deployment-b")))
		})
		It("handles empty string", func() {
			deployments := k8s.ParseDeploymentsFromString("")
			Expect(deployments).To(HaveLen(0))
		})
	})

	Context("ParseDeployments", func() {
		It("parses a slice of strings", func() {
			deployments := k8s.ParseDeployments([]string{"deployment-a", "deployment-b"})
			Expect(deployments).To(HaveLen(2))
			Expect(deployments[0]).To(Equal(k8s.Deployment("deployment-a")))
			Expect(deployments[1]).To(Equal(k8s.Deployment("deployment-b")))
		})
		It("handles empty slice", func() {
			deployments := k8s.ParseDeployments([]string{})
			Expect(deployments).To(HaveLen(0))
		})
	})

	Context("Contains", func() {
		var deployments k8s.Deployments
		BeforeEach(func() {
			deployments = k8s.Deployments{"deployment-a", "deployment-b"}
		})
		It("returns true if deployment exists", func() {
			Expect(deployments.Contains("deployment-a")).To(BeTrue())
		})
		It("returns false if deployment does not exist", func() {
			Expect(deployments.Contains("deployment-c")).To(BeFalse())
		})
	})
})
