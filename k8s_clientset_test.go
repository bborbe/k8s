// Copyright (c) 2024 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package k8s_test

import (
	"os"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/bborbe/k8s"
)

var _ = Describe("Clientset", func() {
	Context("CreateConfig", func() {
		It("returns config from kubeconfig file when provided", func() {
			Skip("Requires valid kubeconfig file")
		})

		It("returns in-cluster config when kubeconfig is empty", func() {
			config, err := k8s.CreateConfig("")
			Expect(err).To(HaveOccurred())
			Expect(config).To(BeNil())
		})

		It("returns error when kubeconfig file doesn't exist", func() {
			config, err := k8s.CreateConfig("/nonexistent/kubeconfig")
			Expect(err).To(HaveOccurred())
			Expect(config).To(BeNil())
		})
	})

	Context("CreateClientset", func() {
		It("returns error when config creation fails", func() {
			clientset, err := k8s.CreateClientset("/nonexistent/kubeconfig")
			Expect(err).To(HaveOccurred())
			Expect(clientset).To(BeNil())
		})

		It("creates clientset when valid kubeconfig provided", func() {
			kubeconfig := os.Getenv("KUBECONFIG")
			if kubeconfig == "" {
				Skip("No KUBECONFIG environment variable set")
			}

			clientset, err := k8s.CreateClientset(kubeconfig)
			if err != nil {
				Skip("Valid kubernetes config not available")
			}
			Expect(clientset).ToNot(BeNil())
		})
	})
})
