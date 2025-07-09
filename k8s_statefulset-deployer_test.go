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
	k8s_errors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	"github.com/bborbe/k8s"
	"github.com/bborbe/k8s/mocks"
)

var _ = Describe("StatefulSet Deployer", func() {
	var statefulSetDeployer k8s.StatefulSetDeployer
	var clientset *mocks.K8sInterface
	var appsV1 *mocks.K8sAppsV1Interface
	var statefulSetInterface *mocks.K8sStatefulSetInterface
	var ctx context.Context
	var statefulSet appsv1.StatefulSet
	var err error

	BeforeEach(func() {
		ctx = context.Background()
		clientset = &mocks.K8sInterface{}
		appsV1 = &mocks.K8sAppsV1Interface{}
		statefulSetInterface = &mocks.K8sStatefulSetInterface{}

		clientset.AppsV1Returns(appsV1)
		appsV1.StatefulSetsReturns(statefulSetInterface)

		statefulSetDeployer = k8s.NewStatefulSetDeployer(clientset)

		statefulSet = appsv1.StatefulSet{
			ObjectMeta: metav1.ObjectMeta{
				Name:      "test-statefulset",
				Namespace: "test-namespace",
			},
			Spec: appsv1.StatefulSetSpec{
				Replicas:    &[]int32{3}[0],
				ServiceName: "test-service",
				Selector: &metav1.LabelSelector{
					MatchLabels: map[string]string{
						"app": "test-app",
					},
				},
				Template: corev1.PodTemplateSpec{
					ObjectMeta: metav1.ObjectMeta{
						Labels: map[string]string{
							"app": "test-app",
						},
					},
					Spec: corev1.PodSpec{
						Containers: []corev1.Container{
							{
								Name:  "test-container",
								Image: "test-image:latest",
							},
						},
					},
				},
				VolumeClaimTemplates: []corev1.PersistentVolumeClaim{
					{
						ObjectMeta: metav1.ObjectMeta{
							Name: "data",
						},
						Spec: corev1.PersistentVolumeClaimSpec{
							AccessModes: []corev1.PersistentVolumeAccessMode{
								corev1.ReadWriteOnce,
							},
							Resources: corev1.VolumeResourceRequirements{
								Requests: corev1.ResourceList{
									corev1.ResourceStorage: resource.MustParse("10Gi"),
								},
							},
						},
					},
				},
			},
		}
	})

	Describe("Deploy", func() {
		JustBeforeEach(func() {
			err = statefulSetDeployer.Deploy(ctx, statefulSet)
		})

		Context("when statefulSet does not exist", func() {
			BeforeEach(func() {
				notFoundError := k8s_errors.NewNotFound(schema.GroupResource{
					Group:    "apps",
					Resource: "statefulsets",
				}, "test-statefulset")
				statefulSetInterface.GetReturns(nil, notFoundError)
				statefulSetInterface.CreateReturns(&statefulSet, nil)
			})

			It("returns no error", func() {
				Expect(err).To(BeNil())
			})

			It("calls Get to check if statefulSet exists", func() {
				Expect(statefulSetInterface.GetCallCount()).To(Equal(1))
				_, name, _ := statefulSetInterface.GetArgsForCall(0)
				Expect(name).To(Equal("test-statefulset"))
			})

			It("calls Create to create the statefulSet", func() {
				Expect(statefulSetInterface.CreateCallCount()).To(Equal(1))
				_, createdStatefulSet, _ := statefulSetInterface.CreateArgsForCall(0)
				Expect(createdStatefulSet.Name).To(Equal("test-statefulset"))
				Expect(createdStatefulSet.Namespace).To(Equal("test-namespace"))
				Expect(createdStatefulSet.Spec.ServiceName).To(Equal("test-service"))
				Expect(*createdStatefulSet.Spec.Replicas).To(Equal(int32(3)))
			})

			It("does not call Update", func() {
				Expect(statefulSetInterface.UpdateCallCount()).To(Equal(0))
			})
		})

		Context("when statefulSet already exists", func() {
			BeforeEach(func() {
				existingStatefulSet := statefulSet
				existingStatefulSet.ResourceVersion = "123"
				statefulSetInterface.GetReturns(&existingStatefulSet, nil)
				statefulSetInterface.UpdateReturns(&statefulSet, nil)
			})

			It("returns no error", func() {
				Expect(err).To(BeNil())
			})

			It("calls Get to check if statefulSet exists", func() {
				Expect(statefulSetInterface.GetCallCount()).To(Equal(1))
				_, name, _ := statefulSetInterface.GetArgsForCall(0)
				Expect(name).To(Equal("test-statefulset"))
			})

			It("calls Update to update the statefulSet", func() {
				Expect(statefulSetInterface.UpdateCallCount()).To(Equal(1))
				_, updatedStatefulSet, _ := statefulSetInterface.UpdateArgsForCall(0)
				Expect(updatedStatefulSet.Name).To(Equal("test-statefulset"))
				Expect(updatedStatefulSet.Namespace).To(Equal("test-namespace"))
			})

			It("does not call Create", func() {
				Expect(statefulSetInterface.CreateCallCount()).To(Equal(0))
			})
		})

		Context("when statefulSet exists with different replica count", func() {
			BeforeEach(func() {
				existingStatefulSet := statefulSet
				existingStatefulSet.ResourceVersion = "123"
				existingStatefulSet.Spec.Replicas = &[]int32{5}[0]
				statefulSetInterface.GetReturns(&existingStatefulSet, nil)
				statefulSetInterface.UpdateReturns(&statefulSet, nil)
			})

			It("updates the replica count", func() {
				Expect(statefulSetInterface.UpdateCallCount()).To(Equal(1))
				_, updatedStatefulSet, _ := statefulSetInterface.UpdateArgsForCall(0)
				Expect(*updatedStatefulSet.Spec.Replicas).To(Equal(int32(3)))
			})
		})

		Context("when statefulSet exists with different volume claim templates", func() {
			BeforeEach(func() {
				existingStatefulSet := statefulSet
				existingStatefulSet.ResourceVersion = "123"
				existingStatefulSet.Spec.VolumeClaimTemplates = []corev1.PersistentVolumeClaim{
					{
						ObjectMeta: metav1.ObjectMeta{
							Name: "old-data",
						},
						Spec: corev1.PersistentVolumeClaimSpec{
							AccessModes: []corev1.PersistentVolumeAccessMode{
								corev1.ReadWriteOnce,
							},
							Resources: corev1.VolumeResourceRequirements{
								Requests: corev1.ResourceList{
									corev1.ResourceStorage: resource.MustParse("5Gi"),
								},
							},
						},
					},
				}
				statefulSetInterface.GetReturns(&existingStatefulSet, nil)
				statefulSetInterface.UpdateReturns(&statefulSet, nil)
			})

			It("updates the volume claim templates", func() {
				Expect(statefulSetInterface.UpdateCallCount()).To(Equal(1))
				_, updatedStatefulSet, _ := statefulSetInterface.UpdateArgsForCall(0)
				Expect(updatedStatefulSet.Spec.VolumeClaimTemplates).To(HaveLen(1))
				Expect(updatedStatefulSet.Spec.VolumeClaimTemplates[0].Name).To(Equal("data"))
				storageRequest := updatedStatefulSet.Spec.VolumeClaimTemplates[0].Spec.Resources.Requests[corev1.ResourceStorage]
				Expect(storageRequest.String()).To(Equal("10Gi"))
			})
		})

		Context("when statefulSet exists with different update strategy", func() {
			BeforeEach(func() {
				existingStatefulSet := statefulSet
				existingStatefulSet.ResourceVersion = "123"
				existingStatefulSet.Spec.UpdateStrategy = appsv1.StatefulSetUpdateStrategy{
					Type: appsv1.OnDeleteStatefulSetStrategyType,
				}
				statefulSetInterface.GetReturns(&existingStatefulSet, nil)

				// Update the new statefulSet with RollingUpdate strategy
				statefulSet.Spec.UpdateStrategy = appsv1.StatefulSetUpdateStrategy{
					Type: appsv1.RollingUpdateStatefulSetStrategyType,
					RollingUpdate: &appsv1.RollingUpdateStatefulSetStrategy{
						Partition: &[]int32{0}[0],
					},
				}
				statefulSetInterface.UpdateReturns(&statefulSet, nil)
			})

			It("updates the update strategy", func() {
				Expect(statefulSetInterface.UpdateCallCount()).To(Equal(1))
				_, updatedStatefulSet, _ := statefulSetInterface.UpdateArgsForCall(0)
				Expect(updatedStatefulSet.Spec.UpdateStrategy.Type).To(Equal(appsv1.RollingUpdateStatefulSetStrategyType))
				Expect(updatedStatefulSet.Spec.UpdateStrategy.RollingUpdate).NotTo(BeNil())
				Expect(*updatedStatefulSet.Spec.UpdateStrategy.RollingUpdate.Partition).To(Equal(int32(0)))
			})
		})

		Context("when Get returns an unexpected error", func() {
			BeforeEach(func() {
				statefulSetInterface.GetReturns(nil, errors.New("unexpected error"))
			})

			It("calls Create assuming statefulSet doesn't exist", func() {
				Expect(statefulSetInterface.CreateCallCount()).To(Equal(1))
			})
		})

		Context("when Create fails", func() {
			BeforeEach(func() {
				notFoundError := k8s_errors.NewNotFound(schema.GroupResource{
					Group:    "apps",
					Resource: "statefulsets",
				}, "test-statefulset")
				statefulSetInterface.GetReturns(nil, notFoundError)
				statefulSetInterface.CreateReturns(nil, errors.New("create failed"))
			})

			It("returns an error", func() {
				Expect(err).NotTo(BeNil())
				Expect(err.Error()).To(ContainSubstring("create statefulSet failed"))
			})

			It("calls Create", func() {
				Expect(statefulSetInterface.CreateCallCount()).To(Equal(1))
			})
		})

		Context("when Update fails", func() {
			BeforeEach(func() {
				existingStatefulSet := statefulSet
				statefulSetInterface.GetReturns(&existingStatefulSet, nil)
				statefulSetInterface.UpdateReturns(nil, errors.New("update failed"))
			})

			It("returns an error", func() {
				Expect(err).NotTo(BeNil())
				Expect(err.Error()).To(ContainSubstring("update statefulSet failed"))
			})

			It("calls Update", func() {
				Expect(statefulSetInterface.UpdateCallCount()).To(Equal(1))
			})
		})

		Context("with context cancellation", func() {
			var cancelCtx context.Context
			var cancel context.CancelFunc

			BeforeEach(func() {
				cancelCtx, cancel = context.WithCancel(ctx)
				cancel() // Cancel immediately
				ctx = cancelCtx

				statefulSetInterface.GetReturns(nil, context.Canceled)
			})

			It("handles context cancellation gracefully", func() {
				Expect(statefulSetInterface.CreateCallCount()).To(Equal(1))
				// Even if Get returns context.Canceled, the code treats it as not found
				// and tries to create, which would also fail with context.Canceled
			})
		})
	})

	Describe("Undeploy", func() {
		var namespace k8s.Namespace
		var name k8s.Name

		BeforeEach(func() {
			namespace = k8s.Namespace("test-namespace")
			name = k8s.Name("test-statefulset")
		})

		JustBeforeEach(func() {
			err = statefulSetDeployer.Undeploy(ctx, namespace, name)
		})

		Context("when statefulSet exists", func() {
			BeforeEach(func() {
				statefulSetInterface.GetReturns(&statefulSet, nil)
				statefulSetInterface.DeleteReturns(nil)
			})

			It("returns no error", func() {
				Expect(err).To(BeNil())
			})

			It("calls Get to check if statefulSet exists", func() {
				Expect(statefulSetInterface.GetCallCount()).To(Equal(1))
				_, statefulSetName, _ := statefulSetInterface.GetArgsForCall(0)
				Expect(statefulSetName).To(Equal("test-statefulset"))
			})

			It("calls Delete to remove the statefulSet", func() {
				Expect(statefulSetInterface.DeleteCallCount()).To(Equal(1))
				_, deletedName, _ := statefulSetInterface.DeleteArgsForCall(0)
				Expect(deletedName).To(Equal("test-statefulset"))
			})
		})

		Context("when statefulSet does not exist", func() {
			BeforeEach(func() {
				notFoundError := k8s_errors.NewNotFound(schema.GroupResource{
					Group:    "apps",
					Resource: "statefulsets",
				}, "test-statefulset")
				statefulSetInterface.GetReturns(nil, notFoundError)
			})

			It("returns no error", func() {
				Expect(err).To(BeNil())
			})

			It("calls Get to check if statefulSet exists", func() {
				Expect(statefulSetInterface.GetCallCount()).To(Equal(1))
			})

			It("does not call Delete", func() {
				Expect(statefulSetInterface.DeleteCallCount()).To(Equal(0))
			})
		})

		Context("when Get returns an unexpected error", func() {
			BeforeEach(func() {
				statefulSetInterface.GetReturns(nil, errors.New("unexpected error"))
			})

			It("does not call Delete", func() {
				Expect(statefulSetInterface.DeleteCallCount()).To(Equal(0))
			})

			It("returns no error (treats as not found)", func() {
				Expect(err).To(BeNil())
			})
		})

		Context("when Delete fails", func() {
			BeforeEach(func() {
				statefulSetInterface.GetReturns(&statefulSet, nil)
				statefulSetInterface.DeleteReturns(errors.New("delete failed"))
			})

			It("returns an error", func() {
				Expect(err).NotTo(BeNil())
				Expect(err.Error()).To(ContainSubstring("delete failed"))
			})

			It("calls Delete", func() {
				Expect(statefulSetInterface.DeleteCallCount()).To(Equal(1))
			})
		})

		Context("with context cancellation", func() {
			var cancelCtx context.Context
			var cancel context.CancelFunc

			BeforeEach(func() {
				cancelCtx, cancel = context.WithCancel(ctx)
				cancel() // Cancel immediately
				ctx = cancelCtx

				statefulSetInterface.GetReturns(nil, context.Canceled)
			})

			It("handles context cancellation gracefully", func() {
				Expect(err).To(BeNil())
				Expect(statefulSetInterface.DeleteCallCount()).To(Equal(0))
			})
		})
	})

	Describe("Constructor", func() {
		It("creates a new statefulSet deployer", func() {
			deployer := k8s.NewStatefulSetDeployer(clientset)
			Expect(deployer).NotTo(BeNil())
		})
	})

	Describe("StatefulSet specific scenarios", func() {
		JustBeforeEach(func() {
			err = statefulSetDeployer.Deploy(ctx, statefulSet)
		})

		Context("when deploying with persistent volume claims", func() {
			BeforeEach(func() {
				notFoundError := k8s_errors.NewNotFound(schema.GroupResource{
					Group:    "apps",
					Resource: "statefulsets",
				}, "test-statefulset")
				statefulSetInterface.GetReturns(nil, notFoundError)
				statefulSetInterface.CreateReturns(&statefulSet, nil)
			})

			It("creates statefulSet with volume claim templates", func() {
				Expect(err).To(BeNil())
				Expect(statefulSetInterface.CreateCallCount()).To(Equal(1))
				_, createdStatefulSet, _ := statefulSetInterface.CreateArgsForCall(0)
				Expect(createdStatefulSet.Spec.VolumeClaimTemplates).To(HaveLen(1))
				Expect(createdStatefulSet.Spec.VolumeClaimTemplates[0].Name).To(Equal("data"))
				Expect(createdStatefulSet.Spec.VolumeClaimTemplates[0].Spec.AccessModes).To(ContainElement(corev1.ReadWriteOnce))
			})
		})

		Context("when deploying with service name", func() {
			BeforeEach(func() {
				notFoundError := k8s_errors.NewNotFound(schema.GroupResource{
					Group:    "apps",
					Resource: "statefulsets",
				}, "test-statefulset")
				statefulSetInterface.GetReturns(nil, notFoundError)
				statefulSetInterface.CreateReturns(&statefulSet, nil)
			})

			It("creates statefulSet with correct service name", func() {
				Expect(err).To(BeNil())
				Expect(statefulSetInterface.CreateCallCount()).To(Equal(1))
				_, createdStatefulSet, _ := statefulSetInterface.CreateArgsForCall(0)
				Expect(createdStatefulSet.Spec.ServiceName).To(Equal("test-service"))
			})
		})
	})
})
