// Copyright (c) 2024 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package k8s_test

import (
	"context"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	corev1 "k8s.io/api/core/v1"

	"github.com/bborbe/k8s"
)

var _ = Describe("Env Builder", func() {
	var envBuilder k8s.EnvBuilder
	var envVars []corev1.EnvVar
	var err error
	var ctx context.Context

	BeforeEach(func() {
		ctx = context.Background()
		envBuilder = k8s.NewEnvBuilder()
	})

	Describe("Build", func() {
		JustBeforeEach(func() {
			envVars, err = envBuilder.Build(ctx)
		})

		Context("with no environment variables", func() {
			It("returns no error", func() {
				Expect(err).To(BeNil())
			})

			It("returns empty slice", func() {
				Expect(envVars).To(BeEmpty())
			})
		})

		Context("with single environment variable", func() {
			BeforeEach(func() {
				envBuilder.Add("TEST_VAR", "test_value")
			})

			It("returns no error", func() {
				Expect(err).To(BeNil())
			})

			It("returns single environment variable", func() {
				Expect(envVars).To(HaveLen(1))
				Expect(envVars[0].Name).To(Equal("TEST_VAR"))
				Expect(envVars[0].Value).To(Equal("test_value"))
				Expect(envVars[0].ValueFrom).To(BeNil())
			})
		})

		Context("with multiple environment variables", func() {
			BeforeEach(func() {
				envBuilder.Add("VAR1", "value1")
				envBuilder.Add("VAR2", "value2")
				envBuilder.Add("VAR3", "value3")
			})

			It("returns no error", func() {
				Expect(err).To(BeNil())
			})

			It("returns all environment variables", func() {
				Expect(envVars).To(HaveLen(3))
				Expect(envVars[0].Name).To(Equal("VAR1"))
				Expect(envVars[0].Value).To(Equal("value1"))
				Expect(envVars[1].Name).To(Equal("VAR2"))
				Expect(envVars[1].Value).To(Equal("value2"))
				Expect(envVars[2].Name).To(Equal("VAR3"))
				Expect(envVars[2].Value).To(Equal("value3"))
			})
		})

		Context("with environment variables from secrets", func() {
			BeforeEach(func() {
				envBuilder.AddSecret("DATABASE_PASSWORD", "db-secret", "password")
				envBuilder.AddSecret("API_KEY", "api-secret", "key")
			})

			It("returns no error", func() {
				Expect(err).To(BeNil())
			})

			It("returns environment variables with secret references", func() {
				Expect(envVars).To(HaveLen(2))

				// First secret env var
				Expect(envVars[0].Name).To(Equal("DATABASE_PASSWORD"))
				Expect(envVars[0].Value).To(BeEmpty())
				Expect(envVars[0].ValueFrom).NotTo(BeNil())
				Expect(envVars[0].ValueFrom.SecretKeyRef).NotTo(BeNil())
				Expect(envVars[0].ValueFrom.SecretKeyRef.Name).To(Equal("db-secret"))
				Expect(envVars[0].ValueFrom.SecretKeyRef.Key).To(Equal("password"))

				// Second secret env var
				Expect(envVars[1].Name).To(Equal("API_KEY"))
				Expect(envVars[1].Value).To(BeEmpty())
				Expect(envVars[1].ValueFrom).NotTo(BeNil())
				Expect(envVars[1].ValueFrom.SecretKeyRef).NotTo(BeNil())
				Expect(envVars[1].ValueFrom.SecretKeyRef.Name).To(Equal("api-secret"))
				Expect(envVars[1].ValueFrom.SecretKeyRef.Key).To(Equal("key"))
			})
		})

		Context("with environment variables from field references", func() {
			BeforeEach(func() {
				envBuilder.AddFieldRef("POD_NAME", "v1", "metadata.name")
				envBuilder.AddFieldRef("POD_NAMESPACE", "v1", "metadata.namespace")
				envBuilder.AddFieldRef("NODE_NAME", "v1", "spec.nodeName")
			})

			It("returns no error", func() {
				Expect(err).To(BeNil())
			})

			It("returns environment variables with field references", func() {
				Expect(envVars).To(HaveLen(3))

				// Pod name field ref
				Expect(envVars[0].Name).To(Equal("POD_NAME"))
				Expect(envVars[0].Value).To(BeEmpty())
				Expect(envVars[0].ValueFrom).NotTo(BeNil())
				Expect(envVars[0].ValueFrom.FieldRef).NotTo(BeNil())
				Expect(envVars[0].ValueFrom.FieldRef.APIVersion).To(Equal("v1"))
				Expect(envVars[0].ValueFrom.FieldRef.FieldPath).To(Equal("metadata.name"))

				// Pod namespace field ref
				Expect(envVars[1].Name).To(Equal("POD_NAMESPACE"))
				Expect(envVars[1].ValueFrom.FieldRef.FieldPath).To(Equal("metadata.namespace"))

				// Node name field ref
				Expect(envVars[2].Name).To(Equal("NODE_NAME"))
				Expect(envVars[2].ValueFrom.FieldRef.FieldPath).To(Equal("spec.nodeName"))
			})
		})

		Context("with mixed environment variable types", func() {
			BeforeEach(func() {
				envBuilder.Add("SIMPLE_VAR", "simple_value")
				envBuilder.AddSecret("SECRET_VAR", "my-secret", "secret-key")
				envBuilder.AddFieldRef("FIELD_VAR", "v1", "metadata.name")
				envBuilder.Add("ANOTHER_VAR", "another_value")
			})

			It("returns no error", func() {
				Expect(err).To(BeNil())
			})

			It("returns all environment variables in order", func() {
				Expect(envVars).To(HaveLen(4))

				// Simple value
				Expect(envVars[0].Name).To(Equal("SIMPLE_VAR"))
				Expect(envVars[0].Value).To(Equal("simple_value"))
				Expect(envVars[0].ValueFrom).To(BeNil())

				// Secret reference
				Expect(envVars[1].Name).To(Equal("SECRET_VAR"))
				Expect(envVars[1].Value).To(BeEmpty())
				Expect(envVars[1].ValueFrom.SecretKeyRef).NotTo(BeNil())
				Expect(envVars[1].ValueFrom.SecretKeyRef.Name).To(Equal("my-secret"))

				// Field reference
				Expect(envVars[2].Name).To(Equal("FIELD_VAR"))
				Expect(envVars[2].Value).To(BeEmpty())
				Expect(envVars[2].ValueFrom.FieldRef).NotTo(BeNil())
				Expect(envVars[2].ValueFrom.FieldRef.FieldPath).To(Equal("metadata.name"))

				// Another simple value
				Expect(envVars[3].Name).To(Equal("ANOTHER_VAR"))
				Expect(envVars[3].Value).To(Equal("another_value"))
				Expect(envVars[3].ValueFrom).To(BeNil())
			})
		})

		Context("with empty values", func() {
			BeforeEach(func() {
				envBuilder.Add("EMPTY_VAR", "")
				envBuilder.Add("SPACE_VAR", " ")
				envBuilder.Add("NORMAL_VAR", "normal_value")
			})

			It("returns no error", func() {
				Expect(err).To(BeNil())
			})

			It("preserves empty and space values", func() {
				Expect(envVars).To(HaveLen(3))
				Expect(envVars[0].Name).To(Equal("EMPTY_VAR"))
				Expect(envVars[0].Value).To(Equal(""))
				Expect(envVars[1].Name).To(Equal("SPACE_VAR"))
				Expect(envVars[1].Value).To(Equal(" "))
				Expect(envVars[2].Name).To(Equal("NORMAL_VAR"))
				Expect(envVars[2].Value).To(Equal("normal_value"))
			})
		})

		Context("with special characters in values", func() {
			BeforeEach(func() {
				envBuilder.Add("SPECIAL_CHARS", "value with spaces and !@#$%^&*()")
				envBuilder.Add("MULTILINE", "line1\nline2\nline3")
				envBuilder.Add("UNICODE", "héllo wörld 🌍")
				envBuilder.Add("JSON", `{"key": "value", "number": 42}`)
			})

			It("returns no error", func() {
				Expect(err).To(BeNil())
			})

			It("preserves special characters", func() {
				Expect(envVars).To(HaveLen(4))
				Expect(envVars[0].Value).To(Equal("value with spaces and !@#$%^&*()"))
				Expect(envVars[1].Value).To(Equal("line1\nline2\nline3"))
				Expect(envVars[2].Value).To(Equal("héllo wörld 🌍"))
				Expect(envVars[3].Value).To(Equal(`{"key": "value", "number": 42}`))
			})
		})

		Context("with duplicate environment variable names", func() {
			BeforeEach(func() {
				envBuilder.Add("DUPLICATE", "value1")
				envBuilder.Add("DUPLICATE", "value2")
				envBuilder.AddSecret("DUPLICATE", "secret", "key")
			})

			It("returns no error", func() {
				Expect(err).To(BeNil())
			})

			It("includes all variables (Kubernetes handles duplicates)", func() {
				Expect(envVars).To(HaveLen(3))
				Expect(envVars[0].Name).To(Equal("DUPLICATE"))
				Expect(envVars[0].Value).To(Equal("value1"))
				Expect(envVars[1].Name).To(Equal("DUPLICATE"))
				Expect(envVars[1].Value).To(Equal("value2"))
				Expect(envVars[2].Name).To(Equal("DUPLICATE"))
				Expect(envVars[2].ValueFrom.SecretKeyRef).NotTo(BeNil())
			})
		})

		Context("with different API versions for field refs", func() {
			BeforeEach(func() {
				envBuilder.AddFieldRef("POD_NAME_V1", "v1", "metadata.name")
				envBuilder.AddFieldRef("POD_NAME_EMPTY", "", "metadata.name")
				envBuilder.AddFieldRef("POD_NAME_CUSTOM", "custom/v1", "metadata.name")
			})

			It("returns no error", func() {
				Expect(err).To(BeNil())
			})

			It("preserves API versions", func() {
				Expect(envVars).To(HaveLen(3))
				Expect(envVars[0].ValueFrom.FieldRef.APIVersion).To(Equal("v1"))
				Expect(envVars[1].ValueFrom.FieldRef.APIVersion).To(Equal(""))
				Expect(envVars[2].ValueFrom.FieldRef.APIVersion).To(Equal("custom/v1"))
			})
		})

		Context("with complex field paths", func() {
			BeforeEach(func() {
				envBuilder.AddFieldRef("POD_IP", "v1", "status.podIP")
				envBuilder.AddFieldRef("HOST_IP", "v1", "status.hostIP")
				envBuilder.AddFieldRef("SERVICE_ACCOUNT", "v1", "spec.serviceAccountName")
				envBuilder.AddFieldRef("CPU_LIMIT", "v1", "spec.containers[0].resources.limits.cpu")
			})

			It("returns no error", func() {
				Expect(err).To(BeNil())
			})

			It("preserves complex field paths", func() {
				Expect(envVars).To(HaveLen(4))
				Expect(envVars[0].ValueFrom.FieldRef.FieldPath).To(Equal("status.podIP"))
				Expect(envVars[1].ValueFrom.FieldRef.FieldPath).To(Equal("status.hostIP"))
				Expect(envVars[2].ValueFrom.FieldRef.FieldPath).To(Equal("spec.serviceAccountName"))
				Expect(
					envVars[3].ValueFrom.FieldRef.FieldPath,
				).To(Equal("spec.containers[0].resources.limits.cpu"))
			})
		})

		Context("with secret references to different secrets", func() {
			BeforeEach(func() {
				envBuilder.AddSecret("DB_PASSWORD", "database-secret", "password")
				envBuilder.AddSecret("DB_USERNAME", "database-secret", "username")
				envBuilder.AddSecret("API_KEY", "api-secret", "key")
				envBuilder.AddSecret("JWT_SECRET", "jwt-secret", "secret")
			})

			It("returns no error", func() {
				Expect(err).To(BeNil())
			})

			It("preserves secret references", func() {
				Expect(envVars).To(HaveLen(4))

				// Database password
				Expect(envVars[0].ValueFrom.SecretKeyRef.Name).To(Equal("database-secret"))
				Expect(envVars[0].ValueFrom.SecretKeyRef.Key).To(Equal("password"))

				// Database username (same secret, different key)
				Expect(envVars[1].ValueFrom.SecretKeyRef.Name).To(Equal("database-secret"))
				Expect(envVars[1].ValueFrom.SecretKeyRef.Key).To(Equal("username"))

				// API key (different secret)
				Expect(envVars[2].ValueFrom.SecretKeyRef.Name).To(Equal("api-secret"))
				Expect(envVars[2].ValueFrom.SecretKeyRef.Key).To(Equal("key"))

				// JWT secret (different secret)
				Expect(envVars[3].ValueFrom.SecretKeyRef.Name).To(Equal("jwt-secret"))
				Expect(envVars[3].ValueFrom.SecretKeyRef.Key).To(Equal("secret"))
			})
		})
	})

	Describe("Validation", func() {
		JustBeforeEach(func() {
			// EnvBuilder validation is internal, so we test it through Build
			_, err = envBuilder.Build(ctx)
		})

		It("returns no error", func() {
			Expect(err).To(BeNil())
		})

		Context("with environment variables", func() {
			BeforeEach(func() {
				envBuilder.Add("TEST", "value")
			})

			It("returns no error", func() {
				Expect(err).To(BeNil())
			})
		})

		Context("with no environment variables", func() {
			It("returns no error", func() {
				Expect(err).To(BeNil())
			})
		})
	})

	Describe("Method chaining", func() {
		It("allows fluent interface", func() {
			envVars, err := k8s.NewEnvBuilder().
				Add("VAR1", "value1").
				AddSecret("SECRET_VAR", "secret-name", "secret-key").
				AddFieldRef("FIELD_VAR", "v1", "metadata.name").
				Add("VAR2", "value2").
				Build(ctx)

			Expect(err).To(BeNil())
			Expect(envVars).To(HaveLen(4))
			Expect(envVars[0].Name).To(Equal("VAR1"))
			Expect(envVars[0].Value).To(Equal("value1"))
			Expect(envVars[1].Name).To(Equal("SECRET_VAR"))
			Expect(envVars[1].ValueFrom.SecretKeyRef.Name).To(Equal("secret-name"))
			Expect(envVars[2].Name).To(Equal("FIELD_VAR"))
			Expect(envVars[2].ValueFrom.FieldRef.FieldPath).To(Equal("metadata.name"))
			Expect(envVars[3].Name).To(Equal("VAR2"))
			Expect(envVars[3].Value).To(Equal("value2"))
		})
	})

	Describe("Constructor", func() {
		It("creates a new env builder", func() {
			builder := k8s.NewEnvBuilder()
			Expect(builder).NotTo(BeNil())
		})

		It("starts with empty environment variables", func() {
			builder := k8s.NewEnvBuilder()
			envVars, err := builder.Build(ctx)
			Expect(err).To(BeNil())
			Expect(envVars).To(BeEmpty())
		})
	})

	Describe("Edge cases", func() {
		Context("with very long variable names and values", func() {
			BeforeEach(func() {
				longName := "VERY_LONG_VARIABLE_NAME_THAT_EXCEEDS_NORMAL_LENGTH_BOUNDARIES_AND_CONTINUES_FOR_A_WHILE"
				longValue := "This is a very long value that contains a lot of text and might be used for configuration purposes or other scenarios where long strings are needed in environment variables"
				envBuilder.Add(longName, longValue)
			})

			JustBeforeEach(func() {
				envVars, err = envBuilder.Build(ctx)
			})

			It("returns no error", func() {
				Expect(err).To(BeNil())
			})

			It("preserves long names and values", func() {
				Expect(envVars).To(HaveLen(1))
				Expect(len(envVars[0].Name)).To(BeNumerically(">", 50))
				Expect(len(envVars[0].Value)).To(BeNumerically(">", 100))
			})
		})

		Context("with empty secret name or key", func() {
			BeforeEach(func() {
				envBuilder.AddSecret("EMPTY_SECRET", "", "key")
				envBuilder.AddSecret("EMPTY_KEY", "secret", "")
				envBuilder.AddSecret("BOTH_EMPTY", "", "")
			})

			JustBeforeEach(func() {
				envVars, err = envBuilder.Build(ctx)
			})

			It("returns no error", func() {
				Expect(err).To(BeNil())
			})

			It("preserves empty secret names and keys", func() {
				Expect(envVars).To(HaveLen(3))
				Expect(envVars[0].ValueFrom.SecretKeyRef.Name).To(Equal(""))
				Expect(envVars[0].ValueFrom.SecretKeyRef.Key).To(Equal("key"))
				Expect(envVars[1].ValueFrom.SecretKeyRef.Name).To(Equal("secret"))
				Expect(envVars[1].ValueFrom.SecretKeyRef.Key).To(Equal(""))
				Expect(envVars[2].ValueFrom.SecretKeyRef.Name).To(Equal(""))
				Expect(envVars[2].ValueFrom.SecretKeyRef.Key).To(Equal(""))
			})
		})

		Context("with empty field ref parameters", func() {
			BeforeEach(func() {
				envBuilder.AddFieldRef("EMPTY_API", "", "metadata.name")
				envBuilder.AddFieldRef("EMPTY_PATH", "v1", "")
				envBuilder.AddFieldRef("BOTH_EMPTY", "", "")
			})

			JustBeforeEach(func() {
				envVars, err = envBuilder.Build(ctx)
			})

			It("returns no error", func() {
				Expect(err).To(BeNil())
			})

			It("preserves empty field ref parameters", func() {
				Expect(envVars).To(HaveLen(3))
				Expect(envVars[0].ValueFrom.FieldRef.APIVersion).To(Equal(""))
				Expect(envVars[0].ValueFrom.FieldRef.FieldPath).To(Equal("metadata.name"))
				Expect(envVars[1].ValueFrom.FieldRef.APIVersion).To(Equal("v1"))
				Expect(envVars[1].ValueFrom.FieldRef.FieldPath).To(Equal(""))
				Expect(envVars[2].ValueFrom.FieldRef.APIVersion).To(Equal(""))
				Expect(envVars[2].ValueFrom.FieldRef.FieldPath).To(Equal(""))
			})
		})
	})
})
