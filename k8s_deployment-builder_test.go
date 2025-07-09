// Copyright (c) 2024 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package k8s_test

import (
	"context"
	"errors"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"

	"github.com/bborbe/k8s"
)

var _ = Describe("Deployment Builder", func() {
	var deploymentBuilder k8s.DeploymentBuilder
	var deployment *appsv1.Deployment
	var err error
	var ctx context.Context
	BeforeEach(func() {
		ctx = context.Background()

		objectMetaBuilder := k8s.NewObjectMetaBuilder()
		objectMetaBuilder.SetName("my-name")
		objectMetaBuilder.SetNamespace("my-namespace")

		containerBuilder := k8s.NewContainerBuilder()
		containerBuilder.SetName("my-container")

		containersBuilder := k8s.NewContainersBuilder()
		containersBuilder.AddContainerBuilder(containerBuilder)

		deploymentBuilder = k8s.NewDeploymentBuilder()
		deploymentBuilder.SetObjectMetaBuilder(objectMetaBuilder)
		deploymentBuilder.SetContainersBuilder(containersBuilder)
	})
	Context("Build", func() {
		JustBeforeEach(func() {
			deployment, err = deploymentBuilder.Build(ctx)
		})
		It("returns no error", func() {
			Expect(err).To(BeNil())
		})
		It("returns deployment", func() {
			Expect(deployment).NotTo(BeNil())
		})
		It("returns correct ImagePullSecrets", func() {
			Expect(deployment).NotTo(BeNil())
			Expect(deployment.Spec.Template.Spec.ImagePullSecrets).To(HaveLen(1))
			Expect(deployment.Spec.Template.Spec.ImagePullSecrets[0].Name).To(Equal("docker"))
		})
		Context("SetImagePullSecrets", func() {
			BeforeEach(func() {
				deploymentBuilder.SetImagePullSecrets([]string{"docker-registry"})
			})
			It("returns correct ImagePullSecrets", func() {
				Expect(deployment).NotTo(BeNil())
				Expect(deployment.Spec.Template.Spec.ImagePullSecrets).To(HaveLen(1))
				Expect(deployment.Spec.Template.Spec.ImagePullSecrets[0].Name).To(Equal("docker-registry"))
			})
		})
	})

	Context("validation", func() {
		JustBeforeEach(func() {
			err = deploymentBuilder.Validate(ctx)
		})

		It("returns no error", func() {
			Expect(err).To(BeNil())
		})

		Context("without ObjectMeta", func() {
			BeforeEach(func() {
				deploymentBuilder.SetObjectMetaBuilder(nil)
			})

			It("returns validation error", func() {
				Expect(err).NotTo(BeNil())
				Expect(err.Error()).To(ContainSubstring("ObjectMeta"))
			})
		})

		Context("without ContainersBuilder", func() {
			BeforeEach(func() {
				deploymentBuilder.SetContainersBuilder(nil)
			})

			It("returns validation error", func() {
				Expect(err).NotTo(BeNil())
				Expect(err.Error()).To(ContainSubstring("ContainersBuilder"))
			})
		})
	})

	Context("replica configuration", func() {
		Context("with zero replicas", func() {
			BeforeEach(func() {
				deploymentBuilder.SetReplicas(0)
			})

			JustBeforeEach(func() {
				deployment, err = deploymentBuilder.Build(ctx)
			})

			It("returns no error", func() {
				Expect(err).To(BeNil())
			})

			It("sets replicas to 0", func() {
				Expect(deployment).NotTo(BeNil())
				Expect(*deployment.Spec.Replicas).To(Equal(int32(0)))
			})
		})

		Context("with high replica count", func() {
			BeforeEach(func() {
				deploymentBuilder.SetReplicas(100)
			})

			JustBeforeEach(func() {
				deployment, err = deploymentBuilder.Build(ctx)
			})

			It("returns no error", func() {
				Expect(err).To(BeNil())
			})

			It("sets replicas to 100", func() {
				Expect(deployment).NotTo(BeNil())
				Expect(*deployment.Spec.Replicas).To(Equal(int32(100)))
			})
		})

		Context("with negative replica count", func() {
			BeforeEach(func() {
				deploymentBuilder.SetReplicas(-1)
			})

			JustBeforeEach(func() {
				deployment, err = deploymentBuilder.Build(ctx)
			})

			It("returns no error (Kubernetes will handle validation)", func() {
				Expect(err).To(BeNil())
			})

			It("sets replicas to -1", func() {
				Expect(deployment).NotTo(BeNil())
				Expect(*deployment.Spec.Replicas).To(Equal(int32(-1)))
			})
		})
	})

	Context("volume configuration", func() {
		Context("with single volume", func() {
			BeforeEach(func() {
				volume := corev1.Volume{
					Name: "test-volume",
					VolumeSource: corev1.VolumeSource{
						EmptyDir: &corev1.EmptyDirVolumeSource{},
					},
				}
				deploymentBuilder.AddVolumes(volume)
			})

			JustBeforeEach(func() {
				deployment, err = deploymentBuilder.Build(ctx)
			})

			It("returns no error", func() {
				Expect(err).To(BeNil())
			})

			It("includes the volume", func() {
				Expect(deployment).NotTo(BeNil())
				Expect(deployment.Spec.Template.Spec.Volumes).To(HaveLen(1))
				Expect(deployment.Spec.Template.Spec.Volumes[0].Name).To(Equal("test-volume"))
			})
		})

		Context("with multiple volumes", func() {
			BeforeEach(func() {
				volumes := []corev1.Volume{
					{
						Name: "volume1",
						VolumeSource: corev1.VolumeSource{
							EmptyDir: &corev1.EmptyDirVolumeSource{},
						},
					},
					{
						Name: "volume2",
						VolumeSource: corev1.VolumeSource{
							ConfigMap: &corev1.ConfigMapVolumeSource{
								LocalObjectReference: corev1.LocalObjectReference{Name: "config"},
							},
						},
					},
				}
				deploymentBuilder.SetVolumes(volumes)
			})

			JustBeforeEach(func() {
				deployment, err = deploymentBuilder.Build(ctx)
			})

			It("returns no error", func() {
				Expect(err).To(BeNil())
			})

			It("includes all volumes", func() {
				Expect(deployment).NotTo(BeNil())
				Expect(deployment.Spec.Template.Spec.Volumes).To(HaveLen(2))
				Expect(deployment.Spec.Template.Spec.Volumes[0].Name).To(Equal("volume1"))
				Expect(deployment.Spec.Template.Spec.Volumes[1].Name).To(Equal("volume2"))
			})
		})

		Context("adding volumes incrementally", func() {
			BeforeEach(func() {
				volume1 := corev1.Volume{
					Name: "volume1",
					VolumeSource: corev1.VolumeSource{
						EmptyDir: &corev1.EmptyDirVolumeSource{},
					},
				}
				volume2 := corev1.Volume{
					Name: "volume2",
					VolumeSource: corev1.VolumeSource{
						EmptyDir: &corev1.EmptyDirVolumeSource{},
					},
				}
				deploymentBuilder.AddVolumes(volume1)
				deploymentBuilder.AddVolumes(volume2)
			})

			JustBeforeEach(func() {
				deployment, err = deploymentBuilder.Build(ctx)
			})

			It("returns no error", func() {
				Expect(err).To(BeNil())
			})

			It("includes all added volumes", func() {
				Expect(deployment).NotTo(BeNil())
				Expect(deployment.Spec.Template.Spec.Volumes).To(HaveLen(2))
				Expect(deployment.Spec.Template.Spec.Volumes[0].Name).To(Equal("volume1"))
				Expect(deployment.Spec.Template.Spec.Volumes[1].Name).To(Equal("volume2"))
			})
		})
	})

	Context("affinity configuration", func() {
		Context("with node affinity", func() {
			BeforeEach(func() {
				affinity := corev1.Affinity{
					NodeAffinity: &corev1.NodeAffinity{
						RequiredDuringSchedulingIgnoredDuringExecution: &corev1.NodeSelector{
							NodeSelectorTerms: []corev1.NodeSelectorTerm{
								{
									MatchExpressions: []corev1.NodeSelectorRequirement{
										{
											Key:      "kubernetes.io/arch",
											Operator: corev1.NodeSelectorOpIn,
											Values:   []string{"amd64"},
										},
									},
								},
							},
						},
					},
				}
				deploymentBuilder.SetAffinity(affinity)
			})

			JustBeforeEach(func() {
				deployment, err = deploymentBuilder.Build(ctx)
			})

			It("returns no error", func() {
				Expect(err).To(BeNil())
			})

			It("includes the affinity", func() {
				Expect(deployment).NotTo(BeNil())
				Expect(deployment.Spec.Template.Spec.Affinity).NotTo(BeNil())
				Expect(deployment.Spec.Template.Spec.Affinity.NodeAffinity).NotTo(BeNil())
				Expect(deployment.Spec.Template.Spec.Affinity.NodeAffinity.RequiredDuringSchedulingIgnoredDuringExecution).NotTo(BeNil())
			})
		})

		Context("with pod affinity", func() {
			BeforeEach(func() {
				affinity := corev1.Affinity{
					PodAffinity: &corev1.PodAffinity{
						RequiredDuringSchedulingIgnoredDuringExecution: []corev1.PodAffinityTerm{
							{
								LabelSelector: &metav1.LabelSelector{
									MatchLabels: map[string]string{
										"app": "database",
									},
								},
								TopologyKey: "kubernetes.io/hostname",
							},
						},
					},
				}
				deploymentBuilder.SetAffinity(affinity)
			})

			JustBeforeEach(func() {
				deployment, err = deploymentBuilder.Build(ctx)
			})

			It("returns no error", func() {
				Expect(err).To(BeNil())
			})

			It("includes the pod affinity", func() {
				Expect(deployment).NotTo(BeNil())
				Expect(deployment.Spec.Template.Spec.Affinity).NotTo(BeNil())
				Expect(deployment.Spec.Template.Spec.Affinity.PodAffinity).NotTo(BeNil())
				Expect(deployment.Spec.Template.Spec.Affinity.PodAffinity.RequiredDuringSchedulingIgnoredDuringExecution).To(HaveLen(1))
			})
		})
	})

	Context("image pull secrets configuration", func() {
		Context("with multiple image pull secrets", func() {
			BeforeEach(func() {
				deploymentBuilder.SetImagePullSecrets([]string{"docker-registry", "private-registry"})
			})

			JustBeforeEach(func() {
				deployment, err = deploymentBuilder.Build(ctx)
			})

			It("returns no error", func() {
				Expect(err).To(BeNil())
			})

			It("includes all image pull secrets", func() {
				Expect(deployment).NotTo(BeNil())
				Expect(deployment.Spec.Template.Spec.ImagePullSecrets).To(HaveLen(2))
				Expect(deployment.Spec.Template.Spec.ImagePullSecrets[0].Name).To(Equal("docker-registry"))
				Expect(deployment.Spec.Template.Spec.ImagePullSecrets[1].Name).To(Equal("private-registry"))
			})
		})

		Context("adding image pull secrets incrementally", func() {
			BeforeEach(func() {
				deploymentBuilder.AddImagePullSecrets("registry1", "registry2")
			})

			JustBeforeEach(func() {
				deployment, err = deploymentBuilder.Build(ctx)
			})

			It("returns no error", func() {
				Expect(err).To(BeNil())
			})

			It("includes all image pull secrets including defaults", func() {
				Expect(deployment).NotTo(BeNil())
				Expect(deployment.Spec.Template.Spec.ImagePullSecrets).To(HaveLen(3))
				Expect(deployment.Spec.Template.Spec.ImagePullSecrets[0].Name).To(Equal("docker"))
				Expect(deployment.Spec.Template.Spec.ImagePullSecrets[1].Name).To(Equal("registry1"))
				Expect(deployment.Spec.Template.Spec.ImagePullSecrets[2].Name).To(Equal("registry2"))
			})
		})

		Context("with empty image pull secrets", func() {
			BeforeEach(func() {
				deploymentBuilder.SetImagePullSecrets([]string{})
			})

			JustBeforeEach(func() {
				deployment, err = deploymentBuilder.Build(ctx)
			})

			It("returns no error", func() {
				Expect(err).To(BeNil())
			})

			It("has no image pull secrets", func() {
				Expect(deployment).NotTo(BeNil())
				Expect(deployment.Spec.Template.Spec.ImagePullSecrets).To(BeEmpty())
			})
		})
	})

	Context("service account configuration", func() {
		Context("with service account name", func() {
			BeforeEach(func() {
				deploymentBuilder.SetServiceAccountName("custom-service-account")
			})

			JustBeforeEach(func() {
				deployment, err = deploymentBuilder.Build(ctx)
			})

			It("returns no error", func() {
				Expect(err).To(BeNil())
			})

			It("sets the service account name", func() {
				Expect(deployment).NotTo(BeNil())
				Expect(deployment.Spec.Template.Spec.ServiceAccountName).To(Equal("custom-service-account"))
			})
		})

		Context("with empty service account name", func() {
			BeforeEach(func() {
				deploymentBuilder.SetServiceAccountName("")
			})

			JustBeforeEach(func() {
				deployment, err = deploymentBuilder.Build(ctx)
			})

			It("returns no error", func() {
				Expect(err).To(BeNil())
			})

			It("has empty service account name", func() {
				Expect(deployment).NotTo(BeNil())
				Expect(deployment.Spec.Template.Spec.ServiceAccountName).To(BeEmpty())
			})
		})
	})

	Context("deployment strategy configuration", func() {
		JustBeforeEach(func() {
			deployment, err = deploymentBuilder.Build(ctx)
		})

		It("returns no error", func() {
			Expect(err).To(BeNil())
		})

		It("uses rolling update strategy", func() {
			Expect(deployment).NotTo(BeNil())
			Expect(deployment.Spec.Strategy.Type).To(Equal(appsv1.RollingUpdateDeploymentStrategyType))
			Expect(deployment.Spec.Strategy.RollingUpdate).NotTo(BeNil())
		})

		It("sets correct rolling update parameters", func() {
			Expect(deployment).NotTo(BeNil())
			Expect(deployment.Spec.Strategy.RollingUpdate.MaxUnavailable).To(Equal(&intstr.IntOrString{
				Type:   intstr.Int,
				IntVal: 1,
			}))
			Expect(deployment.Spec.Strategy.RollingUpdate.MaxSurge).To(Equal(&intstr.IntOrString{
				Type:   intstr.Int,
				IntVal: 1,
			}))
		})
	})

	Context("selector configuration", func() {
		Context("with custom name", func() {
			BeforeEach(func() {
				deploymentBuilder.SetName(k8s.Name("custom-app"))
			})

			JustBeforeEach(func() {
				deployment, err = deploymentBuilder.Build(ctx)
			})

			It("returns no error", func() {
				Expect(err).To(BeNil())
			})

			It("sets selector with custom name", func() {
				Expect(deployment).NotTo(BeNil())
				Expect(deployment.Spec.Selector.MatchLabels).To(HaveKeyWithValue("app", "custom-app"))
			})

			It("sets pod template labels with custom name", func() {
				Expect(deployment).NotTo(BeNil())
				Expect(deployment.Spec.Template.Labels).To(HaveKeyWithValue("app", "custom-app"))
			})
		})

		Context("with component", func() {
			BeforeEach(func() {
				deploymentBuilder.SetComponent("backend")
			})

			JustBeforeEach(func() {
				deployment, err = deploymentBuilder.Build(ctx)
			})

			It("returns no error", func() {
				Expect(err).To(BeNil())
			})

			It("sets component label", func() {
				Expect(deployment).NotTo(BeNil())
				Expect(deployment.Spec.Template.Labels).To(HaveKeyWithValue("component", "backend"))
			})
		})
	})

	Context("prometheus annotations", func() {
		JustBeforeEach(func() {
			deployment, err = deploymentBuilder.Build(ctx)
		})

		It("returns no error", func() {
			Expect(err).To(BeNil())
		})

		It("includes prometheus annotations", func() {
			Expect(deployment).NotTo(BeNil())
			annotations := deployment.Spec.Template.Annotations
			Expect(annotations).To(HaveKeyWithValue("prometheus.io/path", "/metrics"))
			Expect(annotations).To(HaveKeyWithValue("prometheus.io/port", "9090"))
			Expect(annotations).To(HaveKeyWithValue("prometheus.io/scheme", "http"))
			Expect(annotations).To(HaveKeyWithValue("prometheus.io/scrape", "true"))
		})
	})

	Context("error handling", func() {
		Context("when ObjectMeta build fails", func() {
			BeforeEach(func() {
				deploymentBuilder.SetObjectMetaBuilder(k8s.HasBuildObjectMetaFunc(func(ctx context.Context) (*metav1.ObjectMeta, error) {
					return nil, errors.New("object meta build failed")
				}))
			})

			JustBeforeEach(func() {
				deployment, err = deploymentBuilder.Build(ctx)
			})

			It("returns an error", func() {
				Expect(err).NotTo(BeNil())
				Expect(err.Error()).To(ContainSubstring("build objectMeta failed"))
			})

			It("returns nil deployment", func() {
				Expect(deployment).To(BeNil())
			})
		})

		Context("when containers build fails", func() {
			BeforeEach(func() {
				deploymentBuilder.SetContainersBuilder(k8s.HasBuildContainersFunc(func(ctx context.Context) ([]corev1.Container, error) {
					return nil, errors.New("containers build failed")
				}))
			})

			JustBeforeEach(func() {
				deployment, err = deploymentBuilder.Build(ctx)
			})

			It("returns an error", func() {
				Expect(err).NotTo(BeNil())
				Expect(err.Error()).To(ContainSubstring("build containers failed"))
			})

			It("returns nil deployment", func() {
				Expect(deployment).To(BeNil())
			})
		})
	})

	Context("method chaining", func() {
		It("allows fluent interface", func() {
			objectMeta := metav1.ObjectMeta{
				Name:      "chain-deployment",
				Namespace: "default",
			}
			containers := []corev1.Container{
				{
					Name:  "chain-container",
					Image: "chain-image:latest",
				},
			}

			deployment, err := k8s.NewDeploymentBuilder().
				SetObjectMeta(objectMeta).
				SetContainers(containers).
				SetName(k8s.Name("chain-app")).
				SetReplicas(3).
				SetComponent("frontend").
				SetServiceAccountName("chain-sa").
				SetImagePullSecrets([]string{"chain-registry"}).
				Build(ctx)

			Expect(err).To(BeNil())
			Expect(deployment).NotTo(BeNil())
			Expect(*deployment.Spec.Replicas).To(Equal(int32(3)))
			Expect(deployment.Spec.Template.Labels).To(HaveKeyWithValue("component", "frontend"))
			Expect(deployment.Spec.Template.Labels).To(HaveKeyWithValue("app", "chain-app"))
			Expect(deployment.Spec.Template.Spec.ServiceAccountName).To(Equal("chain-sa"))
			Expect(deployment.Spec.Template.Spec.ImagePullSecrets).To(HaveLen(1))
			Expect(deployment.Spec.Template.Spec.ImagePullSecrets[0].Name).To(Equal("chain-registry"))
		})
	})
})
