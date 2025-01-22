// Copyright (c) 2024 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package k8s_test

import (
	"context"

	"github.com/bborbe/k8s"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	batchv1 "k8s.io/api/batch/v1"
)

var _ = Describe("Job Builder", func() {
	var jobBuilder k8s.JobBuilder
	var job *batchv1.Job
	var err error
	var ctx context.Context
	BeforeEach(func() {
		ctx = context.Background()
		jobBuilder = k8s.NewJobBuilder()
		job, err = jobBuilder.Build(ctx)
	})
	It("returns no error", func() {
		Expect(err).To(BeNil())
	})
	It("returns job", func() {
		Expect(job).NotTo(BeNil())
	})
})
