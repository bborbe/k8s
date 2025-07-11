// Copyright (c) 2025 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package test

import (
	. "github.com/onsi/gomega"
	"sigs.k8s.io/yaml"
)

/*
MarshalAsYaml is a test utility function that marshals Kubernetes objects to YAML format.
It automatically fails the test if marshaling fails.

Usage example:

import (

	"strings"
	libk8stest "github.com/bborbe/k8s/test"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gbytes"

)

	It("has correct YAML content", func() {
		// Marshal any Kubernetes object to YAML
		bytes := libk8stest.MarshalAsYaml(service)

		// Verify the YAML contains expected content
		Expect(gbytes.BufferWithBytes(bytes)).To(gbytes.Say(strings.TrimSpace(`

apiVersion: v1
kind: Service
metadata:

	name: my-service
	namespace: default

spec:

	ports:
	- port: 80
	  targetPort: 8080
	selector:
	  app: my-app

`)))
})
*/
func MarshalAsYaml(input interface{}) []byte {
	bytes, err := yaml.Marshal(input)
	Expect(err).To(BeNil())
	return bytes
}
