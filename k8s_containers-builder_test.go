// Copyright (c) 2024 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package k8s_test

import (
	"context"
	"errors"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	"k8s.io/apimachinery/pkg/util/intstr"

	"github.com/bborbe/k8s"
)

var _ = Describe("Containers Builder", func() {
	var containersBuilder k8s.ContainersBuilder
	var containers []corev1.Container
	var err error
	var ctx context.Context

	BeforeEach(func() {
		ctx = context.Background()
		containersBuilder = k8s.NewContainersBuilder()
	})

	Describe("Build", func() {
		JustBeforeEach(func() {
			containers, err = containersBuilder.Build(ctx)
		})

		Context("with no containers", func() {
			It("returns validation error", func() {
				Expect(err).NotTo(BeNil())
				Expect(err.Error()).To(ContainSubstring("ContainerBuilders"))
			})

			It("returns nil containers", func() {
				Expect(containers).To(BeNil())
			})
		})

		Context("with single container", func() {
			BeforeEach(func() {
				container := corev1.Container{
					Name:  "test-container",
					Image: "test-image:latest",
					Ports: []corev1.ContainerPort{
						{
							Name:          "http",
							ContainerPort: 8080,
							Protocol:      corev1.ProtocolTCP,
						},
					},
				}
				containersBuilder.SetContainers([]corev1.Container{container})
			})

			It("returns no error", func() {
				Expect(err).To(BeNil())
			})

			It("returns single container", func() {
				Expect(containers).To(HaveLen(1))
				Expect(containers[0].Name).To(Equal("test-container"))
				Expect(containers[0].Image).To(Equal("test-image:latest"))
				Expect(containers[0].Ports).To(HaveLen(1))
				Expect(containers[0].Ports[0].Name).To(Equal("http"))
				Expect(containers[0].Ports[0].ContainerPort).To(Equal(int32(8080)))
			})
		})

		Context("with multiple containers", func() {
			BeforeEach(func() {
				containers := []corev1.Container{
					{
						Name:  "app-container",
						Image: "app:v1.0.0",
						Ports: []corev1.ContainerPort{
							{
								Name:          "http",
								ContainerPort: 8080,
								Protocol:      corev1.ProtocolTCP,
							},
						},
						Resources: corev1.ResourceRequirements{
							Requests: corev1.ResourceList{
								corev1.ResourceCPU:    resource.MustParse("100m"),
								corev1.ResourceMemory: resource.MustParse("128Mi"),
							},
							Limits: corev1.ResourceList{
								corev1.ResourceCPU:    resource.MustParse("500m"),
								corev1.ResourceMemory: resource.MustParse("512Mi"),
							},
						},
					},
					{
						Name:  "sidecar-container",
						Image: "sidecar:v1.0.0",
						Ports: []corev1.ContainerPort{
							{
								Name:          "metrics",
								ContainerPort: 9090,
								Protocol:      corev1.ProtocolTCP,
							},
						},
						Resources: corev1.ResourceRequirements{
							Requests: corev1.ResourceList{
								corev1.ResourceCPU:    resource.MustParse("50m"),
								corev1.ResourceMemory: resource.MustParse("64Mi"),
							},
						},
					},
				}
				containersBuilder.SetContainers(containers)
			})

			It("returns no error", func() {
				Expect(err).To(BeNil())
			})

			It("returns all containers", func() {
				Expect(containers).To(HaveLen(2))
				Expect(containers[0].Name).To(Equal("app-container"))
				Expect(containers[0].Image).To(Equal("app:v1.0.0"))
				Expect(containers[1].Name).To(Equal("sidecar-container"))
				Expect(containers[1].Image).To(Equal("sidecar:v1.0.0"))
			})

			It("preserves container resources", func() {
				Expect(containers[0].Resources.Requests).To(HaveKeyWithValue(corev1.ResourceCPU, resource.MustParse("100m")))
				Expect(containers[0].Resources.Limits).To(HaveKeyWithValue(corev1.ResourceMemory, resource.MustParse("512Mi")))
				Expect(containers[1].Resources.Requests).To(HaveKeyWithValue(corev1.ResourceCPU, resource.MustParse("50m")))
			})

			It("preserves container ports", func() {
				Expect(containers[0].Ports).To(HaveLen(1))
				Expect(containers[0].Ports[0].ContainerPort).To(Equal(int32(8080)))
				Expect(containers[1].Ports).To(HaveLen(1))
				Expect(containers[1].Ports[0].ContainerPort).To(Equal(int32(9090)))
			})
		})

		Context("with container builders", func() {
			BeforeEach(func() {
				containerBuilder1 := k8s.NewContainerBuilder()
				containerBuilder1.SetName("builder-container-1")
				containerBuilder1.SetImage("builder-image-1:latest")

				containerBuilder2 := k8s.NewContainerBuilder()
				containerBuilder2.SetName("builder-container-2")
				containerBuilder2.SetImage("builder-image-2:latest")

				containersBuilder.AddContainerBuilder(containerBuilder1)
				containersBuilder.AddContainerBuilder(containerBuilder2)
			})

			It("returns no error", func() {
				Expect(err).To(BeNil())
			})

			It("returns containers from builders", func() {
				Expect(containers).To(HaveLen(2))
				Expect(containers[0].Name).To(Equal("builder-container-1"))
				Expect(containers[0].Image).To(Equal("builder-image-1:latest"))
				Expect(containers[1].Name).To(Equal("builder-container-2"))
				Expect(containers[1].Image).To(Equal("builder-image-2:latest"))
			})
		})

		Context("with mixed container builders", func() {
			BeforeEach(func() {
				containerBuilder1 := k8s.NewContainerBuilder()
				containerBuilder1.SetName("builder-container")
				containerBuilder1.SetImage("builder-image:latest")

				containerBuilder2 := k8s.NewContainerBuilder()
				containerBuilder2.SetName("another-builder")
				containerBuilder2.SetImage("another-image:latest")

				containersBuilder.SetContainerBuilders([]k8s.HasBuildContainer{containerBuilder1, containerBuilder2})
			})

			It("returns no error", func() {
				Expect(err).To(BeNil())
			})

			It("returns containers from all builders", func() {
				Expect(containers).To(HaveLen(2))
				Expect(containers[0].Name).To(Equal("builder-container"))
				Expect(containers[1].Name).To(Equal("another-builder"))
			})
		})

		Context("with containers having environment variables", func() {
			BeforeEach(func() {
				container := corev1.Container{
					Name:  "env-container",
					Image: "env-image:latest",
					Env: []corev1.EnvVar{
						{
							Name:  "DATABASE_URL",
							Value: "postgresql://localhost:5432/mydb",
						},
						{
							Name: "API_KEY",
							ValueFrom: &corev1.EnvVarSource{
								SecretKeyRef: &corev1.SecretKeySelector{
									LocalObjectReference: corev1.LocalObjectReference{Name: "api-secret"},
									Key:                  "key",
								},
							},
						},
					},
				}
				containersBuilder.SetContainers([]corev1.Container{container})
			})

			It("returns no error", func() {
				Expect(err).To(BeNil())
			})

			It("preserves environment variables", func() {
				Expect(containers).To(HaveLen(1))
				Expect(containers[0].Env).To(HaveLen(2))
				Expect(containers[0].Env[0].Name).To(Equal("DATABASE_URL"))
				Expect(containers[0].Env[0].Value).To(Equal("postgresql://localhost:5432/mydb"))
				Expect(containers[0].Env[1].Name).To(Equal("API_KEY"))
				Expect(containers[0].Env[1].ValueFrom.SecretKeyRef.Name).To(Equal("api-secret"))
			})
		})

		Context("with containers having volume mounts", func() {
			BeforeEach(func() {
				container := corev1.Container{
					Name:  "volume-container",
					Image: "volume-image:latest",
					VolumeMounts: []corev1.VolumeMount{
						{
							Name:      "config-volume",
							MountPath: "/etc/config",
							ReadOnly:  true,
						},
						{
							Name:      "data-volume",
							MountPath: "/data",
							ReadOnly:  false,
						},
					},
				}
				containersBuilder.SetContainers([]corev1.Container{container})
			})

			It("returns no error", func() {
				Expect(err).To(BeNil())
			})

			It("preserves volume mounts", func() {
				Expect(containers).To(HaveLen(1))
				Expect(containers[0].VolumeMounts).To(HaveLen(2))
				Expect(containers[0].VolumeMounts[0].Name).To(Equal("config-volume"))
				Expect(containers[0].VolumeMounts[0].MountPath).To(Equal("/etc/config"))
				Expect(containers[0].VolumeMounts[0].ReadOnly).To(BeTrue())
				Expect(containers[0].VolumeMounts[1].Name).To(Equal("data-volume"))
				Expect(containers[0].VolumeMounts[1].ReadOnly).To(BeFalse())
			})
		})

		Context("with containers having probes", func() {
			BeforeEach(func() {
				container := corev1.Container{
					Name:  "probe-container",
					Image: "probe-image:latest",
					LivenessProbe: &corev1.Probe{
						ProbeHandler: corev1.ProbeHandler{
							HTTPGet: &corev1.HTTPGetAction{
								Path: "/health",
								Port: intstr.FromInt(8080),
							},
						},
						InitialDelaySeconds: 30,
						PeriodSeconds:       10,
					},
					ReadinessProbe: &corev1.Probe{
						ProbeHandler: corev1.ProbeHandler{
							HTTPGet: &corev1.HTTPGetAction{
								Path: "/ready",
								Port: intstr.FromInt(8080),
							},
						},
						InitialDelaySeconds: 5,
						PeriodSeconds:       5,
					},
				}
				containersBuilder.SetContainers([]corev1.Container{container})
			})

			It("returns no error", func() {
				Expect(err).To(BeNil())
			})

			It("preserves probes", func() {
				Expect(containers).To(HaveLen(1))
				Expect(containers[0].LivenessProbe).NotTo(BeNil())
				Expect(containers[0].LivenessProbe.HTTPGet.Path).To(Equal("/health"))
				Expect(containers[0].LivenessProbe.InitialDelaySeconds).To(Equal(int32(30)))
				Expect(containers[0].ReadinessProbe).NotTo(BeNil())
				Expect(containers[0].ReadinessProbe.HTTPGet.Path).To(Equal("/ready"))
				Expect(containers[0].ReadinessProbe.InitialDelaySeconds).To(Equal(int32(5)))
			})
		})

		Context("when container builder fails", func() {
			BeforeEach(func() {
				failingBuilder := k8s.HasBuildContainerFunc(func(ctx context.Context) (*corev1.Container, error) {
					return nil, errors.New("container build failed")
				})
				containersBuilder.AddContainerBuilder(failingBuilder)
			})

			It("returns an error", func() {
				Expect(err).NotTo(BeNil())
				Expect(err.Error()).To(ContainSubstring("build container failed"))
			})

			It("returns nil containers", func() {
				Expect(containers).To(BeNil())
			})
		})

		Context("with complex container configuration", func() {
			BeforeEach(func() {
				container := corev1.Container{
					Name:       "complex-container",
					Image:      "complex-image:latest",
					Command:    []string{"/bin/sh"},
					Args:       []string{"-c", "echo 'Hello World'"},
					WorkingDir: "/app",
					Ports: []corev1.ContainerPort{
						{
							Name:          "http",
							ContainerPort: 8080,
							Protocol:      corev1.ProtocolTCP,
						},
						{
							Name:          "grpc",
							ContainerPort: 9090,
							Protocol:      corev1.ProtocolTCP,
						},
					},
					Env: []corev1.EnvVar{
						{
							Name:  "ENV_VAR",
							Value: "value",
						},
					},
					Resources: corev1.ResourceRequirements{
						Requests: corev1.ResourceList{
							corev1.ResourceCPU:    resource.MustParse("100m"),
							corev1.ResourceMemory: resource.MustParse("128Mi"),
						},
						Limits: corev1.ResourceList{
							corev1.ResourceCPU:    resource.MustParse("500m"),
							corev1.ResourceMemory: resource.MustParse("512Mi"),
						},
					},
					VolumeMounts: []corev1.VolumeMount{
						{
							Name:      "config",
							MountPath: "/config",
							ReadOnly:  true,
						},
					},
					SecurityContext: &corev1.SecurityContext{
						RunAsUser:    &[]int64{1000}[0],
						RunAsGroup:   &[]int64{1000}[0],
						RunAsNonRoot: &[]bool{true}[0],
					},
				}
				containersBuilder.SetContainers([]corev1.Container{container})
			})

			It("returns no error", func() {
				Expect(err).To(BeNil())
			})

			It("preserves all container configuration", func() {
				Expect(containers).To(HaveLen(1))
				container := containers[0]
				Expect(container.Name).To(Equal("complex-container"))
				Expect(container.Image).To(Equal("complex-image:latest"))
				Expect(container.Command).To(Equal([]string{"/bin/sh"}))
				Expect(container.Args).To(Equal([]string{"-c", "echo 'Hello World'"}))
				Expect(container.WorkingDir).To(Equal("/app"))
				Expect(container.Ports).To(HaveLen(2))
				Expect(container.Env).To(HaveLen(1))
				Expect(container.Resources.Requests).To(HaveKeyWithValue(corev1.ResourceCPU, resource.MustParse("100m")))
				Expect(container.VolumeMounts).To(HaveLen(1))
				Expect(container.SecurityContext).NotTo(BeNil())
				Expect(*container.SecurityContext.RunAsUser).To(Equal(int64(1000)))
			})
		})
	})

	Describe("Validation", func() {
		JustBeforeEach(func() {
			err = containersBuilder.Validate(ctx)
		})

		Context("with no containers", func() {
			It("returns validation error", func() {
				Expect(err).NotTo(BeNil())
				Expect(err.Error()).To(ContainSubstring("ContainerBuilders"))
			})
		})

		Context("with containers", func() {
			BeforeEach(func() {
				container := corev1.Container{
					Name:  "valid-container",
					Image: "valid-image:latest",
				}
				containersBuilder.SetContainers([]corev1.Container{container})
			})

			It("returns no error", func() {
				Expect(err).To(BeNil())
			})
		})

		Context("with container builders", func() {
			BeforeEach(func() {
				containerBuilder := k8s.NewContainerBuilder()
				containerBuilder.SetName("valid-builder")
				containerBuilder.SetImage("valid-image:latest")
				containersBuilder.AddContainerBuilder(containerBuilder)
			})

			It("returns no error", func() {
				Expect(err).To(BeNil())
			})
		})
	})

	Describe("Method chaining", func() {
		It("allows fluent interface", func() {
			container1 := corev1.Container{
				Name:  "chain-container-1",
				Image: "chain-image-1:latest",
			}
			container2 := corev1.Container{
				Name:  "chain-container-2",
				Image: "chain-image-2:latest",
			}

			containerBuilder := k8s.NewContainerBuilder()
			containerBuilder.SetName("chain-builder")
			containerBuilder.SetImage("chain-builder-image:latest")

			containers, err := k8s.NewContainersBuilder().
				SetContainers([]corev1.Container{container1, container2}).
				AddContainerBuilder(containerBuilder).
				Build(ctx)

			Expect(err).To(BeNil())
			Expect(containers).To(HaveLen(3))
			Expect(containers[0].Name).To(Equal("chain-container-1"))
			Expect(containers[1].Name).To(Equal("chain-container-2"))
			Expect(containers[2].Name).To(Equal("chain-builder"))
		})
	})

	Describe("Constructor", func() {
		It("creates a new containers builder", func() {
			builder := k8s.NewContainersBuilder()
			Expect(builder).NotTo(BeNil())
		})
	})

	Describe("SetContainerBuilders", func() {
		Context("with multiple builders", func() {
			BeforeEach(func() {
				builder1 := k8s.NewContainerBuilder()
				builder1.SetName("builder-1")
				builder1.SetImage("image-1:latest")

				builder2 := k8s.NewContainerBuilder()
				builder2.SetName("builder-2")
				builder2.SetImage("image-2:latest")

				containersBuilder.SetContainerBuilders([]k8s.HasBuildContainer{builder1, builder2})
			})

			JustBeforeEach(func() {
				containers, err = containersBuilder.Build(ctx)
			})

			It("returns no error", func() {
				Expect(err).To(BeNil())
			})

			It("replaces existing builders", func() {
				Expect(containers).To(HaveLen(2))
				Expect(containers[0].Name).To(Equal("builder-1"))
				Expect(containers[1].Name).To(Equal("builder-2"))
			})
		})

		Context("with empty builders", func() {
			BeforeEach(func() {
				// First add a builder
				builder := k8s.NewContainerBuilder()
				builder.SetName("initial-builder")
				builder.SetImage("initial-image:latest")
				containersBuilder.AddContainerBuilder(builder)

				// Then set empty builders
				containersBuilder.SetContainerBuilders([]k8s.HasBuildContainer{})
			})

			JustBeforeEach(func() {
				containers, err = containersBuilder.Build(ctx)
			})

			It("returns validation error", func() {
				Expect(err).NotTo(BeNil())
				Expect(err.Error()).To(ContainSubstring("ContainerBuilders"))
			})
		})
	})
})
