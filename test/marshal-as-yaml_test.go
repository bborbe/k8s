// Copyright (c) 2025 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package test_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	corev1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	libk8stest "github.com/bborbe/k8s/test"
)

var _ = Describe("MarshalAsYaml", func() {
	It("marshals Kubernetes objects to YAML", func() {
		deployment := &corev1.Deployment{
			TypeMeta: metav1.TypeMeta{
				APIVersion: "apps/v1",
				Kind:       "Deployment",
			},
			ObjectMeta: metav1.ObjectMeta{
				Name:      "test-deployment",
				Namespace: "default",
			},
		}

		result := libk8stest.MarshalAsYaml(deployment)

		Expect(result).ToNot(BeEmpty())
		Expect(string(result)).To(ContainSubstring("apiVersion: apps/v1"))
		Expect(string(result)).To(ContainSubstring("kind: Deployment"))
		Expect(string(result)).To(ContainSubstring("name: test-deployment"))
		Expect(string(result)).To(ContainSubstring("namespace: default"))
	})

	It("marshals simple objects to YAML", func() {
		simpleObj := map[string]interface{}{
			"name":  "test",
			"value": 42,
		}

		result := libk8stest.MarshalAsYaml(simpleObj)

		Expect(result).ToNot(BeEmpty())
		Expect(string(result)).To(ContainSubstring("name: test"))
		Expect(string(result)).To(ContainSubstring("value: 42"))
	})

	It("marshals empty objects", func() {
		emptyObj := map[string]interface{}{}

		result := libk8stest.MarshalAsYaml(emptyObj)

		Expect(result).To(Equal([]byte("{}\n")))
	})
})
