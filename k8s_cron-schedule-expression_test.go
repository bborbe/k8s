// Copyright (c) 2025 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package k8s_test

import (
	"context"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/bborbe/k8s"
)

var _ = Describe("CronScheduleExpression", func() {
	var ctx context.Context
	var cronExpression k8s.CronScheduleExpression
	var err error

	BeforeEach(func() {
		ctx = context.Background()
	})

	Describe("String", func() {
		Context("with valid expression", func() {
			BeforeEach(func() {
				cronExpression = k8s.CronScheduleExpression("0 0 * * *")
			})

			It("returns the string representation", func() {
				Expect(cronExpression.String()).To(Equal("0 0 * * *"))
			})
		})

		Context("with empty expression", func() {
			BeforeEach(func() {
				cronExpression = k8s.CronScheduleExpression("")
			})

			It("returns empty string", func() {
				Expect(cronExpression.String()).To(Equal(""))
			})
		})
	})

	Describe("Validate", func() {
		JustBeforeEach(func() {
			err = cronExpression.Validate(ctx)
		})

		Context("with valid daily expression", func() {
			BeforeEach(func() {
				cronExpression = k8s.CronScheduleExpression("0 0 * * *")
			})

			It("returns no error", func() {
				Expect(err).To(BeNil())
			})
		})

		Context("with valid hourly expression", func() {
			BeforeEach(func() {
				cronExpression = k8s.CronScheduleExpression("0 * * * *")
			})

			It("returns no error", func() {
				Expect(err).To(BeNil())
			})
		})

		Context("with valid every 15 minutes expression", func() {
			BeforeEach(func() {
				cronExpression = k8s.CronScheduleExpression("*/15 * * * *")
			})

			It("returns no error", func() {
				Expect(err).To(BeNil())
			})
		})

		Context("with valid weekly expression", func() {
			BeforeEach(func() {
				cronExpression = k8s.CronScheduleExpression("0 0 * * 0")
			})

			It("returns no error", func() {
				Expect(err).To(BeNil())
			})
		})

		Context("with valid monthly expression", func() {
			BeforeEach(func() {
				cronExpression = k8s.CronScheduleExpression("0 0 1 * *")
			})

			It("returns no error", func() {
				Expect(err).To(BeNil())
			})
		})

		Context("with valid yearly expression", func() {
			BeforeEach(func() {
				cronExpression = k8s.CronScheduleExpression("0 0 1 1 *")
			})

			It("returns no error", func() {
				Expect(err).To(BeNil())
			})
		})

		Context("with empty expression", func() {
			BeforeEach(func() {
				cronExpression = k8s.CronScheduleExpression("")
			})

			It("returns validation error", func() {
				Expect(err).NotTo(BeNil())
				Expect(err.Error()).To(ContainSubstring("CronScheduleExpression empty"))
			})
		})

		Context("with invalid expression - too few fields", func() {
			BeforeEach(func() {
				cronExpression = k8s.CronScheduleExpression("0 0 *")
			})

			It("returns validation error", func() {
				Expect(err).NotTo(BeNil())
				Expect(err.Error()).To(ContainSubstring("parse CronScheduleExpression failed"))
			})
		})

		Context("with invalid expression - too many fields", func() {
			BeforeEach(func() {
				cronExpression = k8s.CronScheduleExpression("0 0 * * * * *")
			})

			It("returns validation error", func() {
				Expect(err).NotTo(BeNil())
				Expect(err.Error()).To(ContainSubstring("parse CronScheduleExpression failed"))
			})
		})

		Context("with invalid expression - invalid minute", func() {
			BeforeEach(func() {
				cronExpression = k8s.CronScheduleExpression("60 0 * * *")
			})

			It("returns validation error", func() {
				Expect(err).NotTo(BeNil())
				Expect(err.Error()).To(ContainSubstring("parse CronScheduleExpression failed"))
			})
		})

		Context("with invalid expression - invalid hour", func() {
			BeforeEach(func() {
				cronExpression = k8s.CronScheduleExpression("0 24 * * *")
			})

			It("returns validation error", func() {
				Expect(err).NotTo(BeNil())
				Expect(err.Error()).To(ContainSubstring("parse CronScheduleExpression failed"))
			})
		})

		Context("with invalid expression - invalid day of month", func() {
			BeforeEach(func() {
				cronExpression = k8s.CronScheduleExpression("0 0 32 * *")
			})

			It("returns validation error", func() {
				Expect(err).NotTo(BeNil())
				Expect(err.Error()).To(ContainSubstring("parse CronScheduleExpression failed"))
			})
		})

		Context("with invalid expression - invalid month", func() {
			BeforeEach(func() {
				cronExpression = k8s.CronScheduleExpression("0 0 * 13 *")
			})

			It("returns validation error", func() {
				Expect(err).NotTo(BeNil())
				Expect(err.Error()).To(ContainSubstring("parse CronScheduleExpression failed"))
			})
		})

		Context("with invalid expression - invalid day of week", func() {
			BeforeEach(func() {
				cronExpression = k8s.CronScheduleExpression("0 0 * * 8")
			})

			It("returns validation error", func() {
				Expect(err).NotTo(BeNil())
				Expect(err.Error()).To(ContainSubstring("parse CronScheduleExpression failed"))
			})
		})

		Context("with invalid expression - non-numeric characters", func() {
			BeforeEach(func() {
				cronExpression = k8s.CronScheduleExpression("a b c d e")
			})

			It("returns validation error", func() {
				Expect(err).NotTo(BeNil())
				Expect(err.Error()).To(ContainSubstring("parse CronScheduleExpression failed"))
			})
		})

		Context("with invalid expression - missing spaces", func() {
			BeforeEach(func() {
				cronExpression = k8s.CronScheduleExpression("00***")
			})

			It("returns validation error", func() {
				Expect(err).NotTo(BeNil())
				Expect(err.Error()).To(ContainSubstring("parse CronScheduleExpression failed"))
			})
		})
	})
})
