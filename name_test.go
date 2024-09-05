// Copyright (c) 2024 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package k8s_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/bborbe/k8s"
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
