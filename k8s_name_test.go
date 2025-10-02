// Copyright (c) 2024 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package k8s_test

import (
	"context"

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

var _ = DescribeTable(
	"NameFromPod",
	func(pod corev1.Pod, expectedName k8s.Name) {
		name := k8s.NameFromPod(pod)
		Expect(name).To(Equal(expectedName))
	},
	Entry(
		"deployment",
		createTestPod(
			"raw-fetcher-646d746df5-tdzls",
			map[string]string{"pod-template-hash": "646d746df5"},
		),
		k8s.Name("raw-fetcher"),
	),
	Entry(
		"statefulset",
		createTestPod("raw-fetcher-0", map[string]string{}),
		k8s.Name("raw-fetcher"),
	),
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

var _ = Describe("Name methods", func() {
	It("Validate returns error for empty name", func() {
		var n k8s.Name = ""
		err := n.Validate(context.Background())
		Expect(err).To(HaveOccurred())
	})

	It("Validate returns error for name longer than 253", func() {
		var n k8s.Name = k8s.Name(make([]byte, 254))
		err := n.Validate(context.Background())
		Expect(err).To(HaveOccurred())
	})

	It("Validate returns nil for valid name", func() {
		var n k8s.Name = "valid-name"
		err := n.Validate(context.Background())
		Expect(err).ToNot(HaveOccurred())
	})

	It("String returns the string value", func() {
		var n k8s.Name = "foo-bar"
		Expect(n.String()).To(Equal("foo-bar"))
	})

	It("Bytes returns the byte slice", func() {
		var n k8s.Name = "foo-bar"
		Expect(n.Bytes()).To(Equal([]byte("foo-bar")))
	})

	It("Ptr returns a pointer to the Name", func() {
		var n k8s.Name = "foo-bar"
		ptr := n.Ptr()
		Expect(ptr).ToNot(BeNil())
		Expect(*ptr).To(Equal(n))
	})

	It("Add returns a new Name with concatenated value", func() {
		var n k8s.Name = "foo"
		result := n.Add("bar")
		Expect(result).To(Equal(k8s.Name("foo-bar")))
	})
})
