// Copyright (c) 2024 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package k8s_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/bborbe/k8s"
)

func createTestPod(name string, labels map[string]string) corev1.Pod {
	return corev1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name:   name,
			Labels: labels,
		},
	}
}

var _ = DescribeTable("NameFromPod",
	func(pod corev1.Pod, expectedName k8s.Name) {
		name := k8s.NameFromPod(pod)
		Expect(name).To(Equal(expectedName))
	},
	Entry("deployment", createTestPod("raw-fetcher-646d746df5-tdzls", map[string]string{"pod-template-hash": "646d746df5"}), k8s.Name("raw-fetcher")),
	Entry("statefulset", createTestPod("raw-fetcher-0", map[string]string{}), k8s.Name("raw-fetcher")),
)

var _ = DescribeTable("BuildName",
	func(parts []string, expectedName k8s.Name) {
		name := k8s.BuildName(parts...)
		Expect(name).To(Equal(expectedName))
	},
	Entry("simple", []string{"my-valid-name"}, k8s.Name("my-valid-name")),
	Entry("toLower", []string{"my-VALID-name"}, k8s.Name("my-valid-name")),
	Entry("replace invalid chart with dash", []string{"my!valid_name"}, k8s.Name("my-valid-name")),
	Entry("multidash", []string{"my-----valid-----name"}, k8s.Name("my-valid-name")),
	Entry("leading dash", []string{"---my-valid-name"}, k8s.Name("my-valid-name")),
	Entry("following dash", []string{"my-valid-name---"}, k8s.Name("my-valid-name")),
)

var _ = DescribeTable("BuildName",
	func(parts []string, expectedName k8s.Name) {
		name := k8s.BuildName(parts...)
		Expect(name).To(Equal(expectedName))
	},
	Entry("simple", []string{"my-valid-name"}, k8s.Name("my-valid-name")),
	Entry("toLower", []string{"my-VALID-name"}, k8s.Name("my-valid-name")),
	Entry("replace invalid chart with dash", []string{"my!valid_name"}, k8s.Name("my-valid-name")),
	Entry("multidash", []string{"my-----valid-----name"}, k8s.Name("my-valid-name")),
	Entry("leading dash", []string{"---my-valid-name"}, k8s.Name("my-valid-name")),
	Entry("following dash", []string{"my-valid-name---"}, k8s.Name("my-valid-name")),
	Entry("with number", []string{"my-valid-name-f16"}, k8s.Name("my-valid-name-f16")),
)
