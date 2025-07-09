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
	k8s_errors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/util/intstr"

	"github.com/bborbe/k8s"
	"github.com/bborbe/k8s/mocks"
)

var _ = Describe("Service Deployer", func() {
	var serviceDeployer k8s.ServiceDeployer
	var clientset *mocks.K8sInterface
	var coreV1 *mocks.K8sCoreV1Interface
	var serviceInterface *mocks.K8sServiceInterface
	var ctx context.Context
	var service corev1.Service
	var err error

	BeforeEach(func() {
		ctx = context.Background()
		clientset = &mocks.K8sInterface{}
		coreV1 = &mocks.K8sCoreV1Interface{}
		serviceInterface = &mocks.K8sServiceInterface{}

		clientset.CoreV1Returns(coreV1)
		coreV1.ServicesReturns(serviceInterface)

		serviceDeployer = k8s.NewServiceDeployer(clientset)

		service = corev1.Service{
			ObjectMeta: metav1.ObjectMeta{
				Name:      "test-service",
				Namespace: "test-namespace",
			},
			Spec: corev1.ServiceSpec{
				Type: corev1.ServiceTypeClusterIP,
				Selector: map[string]string{
					"app": "test-app",
				},
				Ports: []corev1.ServicePort{
					{
						Name:       "http",
						Port:       80,
						TargetPort: intstr.FromInt(8080),
						Protocol:   corev1.ProtocolTCP,
					},
				},
			},
		}
	})

	Describe("Deploy", func() {
		JustBeforeEach(func() {
			err = serviceDeployer.Deploy(ctx, service)
		})

		Context("when service does not exist", func() {
			BeforeEach(func() {
				notFoundError := k8s_errors.NewNotFound(schema.GroupResource{
					Group:    "",
					Resource: "services",
				}, "test-service")
				serviceInterface.GetReturns(nil, notFoundError)
				serviceInterface.CreateReturns(&service, nil)
			})

			It("returns no error", func() {
				Expect(err).To(BeNil())
			})

			It("calls Get to check if service exists", func() {
				Expect(serviceInterface.GetCallCount()).To(Equal(1))
				_, name, _ := serviceInterface.GetArgsForCall(0)
				Expect(name).To(Equal("test-service"))
			})

			It("calls Create to create the service", func() {
				Expect(serviceInterface.CreateCallCount()).To(Equal(1))
				_, createdService, _ := serviceInterface.CreateArgsForCall(0)
				Expect(createdService.Name).To(Equal("test-service"))
				Expect(createdService.Namespace).To(Equal("test-namespace"))
			})

			It("does not call Update", func() {
				Expect(serviceInterface.UpdateCallCount()).To(Equal(0))
			})
		})

		Context("when service already exists", func() {
			BeforeEach(func() {
				existingService := service
				existingService.ResourceVersion = "123"
				existingService.Spec.ClusterIP = "10.0.0.1"
				serviceInterface.GetReturns(&existingService, nil)
				serviceInterface.UpdateReturns(&service, nil)
			})

			It("returns no error", func() {
				Expect(err).To(BeNil())
			})

			It("calls Get to check if service exists", func() {
				Expect(serviceInterface.GetCallCount()).To(Equal(1))
				_, name, _ := serviceInterface.GetArgsForCall(0)
				Expect(name).To(Equal("test-service"))
			})

			It("calls Update to update the service", func() {
				Expect(serviceInterface.UpdateCallCount()).To(Equal(1))
				_, updatedService, _ := serviceInterface.UpdateArgsForCall(0)
				Expect(updatedService.Name).To(Equal("test-service"))
				Expect(updatedService.Namespace).To(Equal("test-namespace"))
			})

			It("preserves existing ClusterIP in update", func() {
				Expect(serviceInterface.UpdateCallCount()).To(Equal(1))
				_, updatedService, _ := serviceInterface.UpdateArgsForCall(0)
				Expect(updatedService.Spec.ClusterIP).To(Equal("10.0.0.1"))
			})

			It("preserves existing ResourceVersion in update", func() {
				Expect(serviceInterface.UpdateCallCount()).To(Equal(1))
				_, updatedService, _ := serviceInterface.UpdateArgsForCall(0)
				Expect(updatedService.ResourceVersion).To(Equal("123"))
			})

			It("does not call Create", func() {
				Expect(serviceInterface.CreateCallCount()).To(Equal(0))
			})
		})

		Context("when service exists with different service type", func() {
			BeforeEach(func() {
				existingService := service
				existingService.ResourceVersion = "123"
				existingService.Spec.ClusterIP = "10.0.0.1"
				existingService.Spec.Type = corev1.ServiceTypeNodePort
				serviceInterface.GetReturns(&existingService, nil)

				// New service has different type
				service.Spec.Type = corev1.ServiceTypeLoadBalancer
				serviceInterface.UpdateReturns(&service, nil)
			})

			It("updates the service type", func() {
				Expect(serviceInterface.UpdateCallCount()).To(Equal(1))
				_, updatedService, _ := serviceInterface.UpdateArgsForCall(0)
				Expect(updatedService.Spec.Type).To(Equal(corev1.ServiceTypeLoadBalancer))
			})
		})

		Context("when service exists with different ports", func() {
			BeforeEach(func() {
				existingService := service
				existingService.ResourceVersion = "123"
				existingService.Spec.ClusterIP = "10.0.0.1"
				existingService.Spec.Ports = []corev1.ServicePort{
					{
						Name:       "old-port",
						Port:       8080,
						TargetPort: intstr.FromInt(8080),
						Protocol:   corev1.ProtocolTCP,
					},
				}
				serviceInterface.GetReturns(&existingService, nil)
				serviceInterface.UpdateReturns(&service, nil)
			})

			It("updates the service ports", func() {
				Expect(serviceInterface.UpdateCallCount()).To(Equal(1))
				_, updatedService, _ := serviceInterface.UpdateArgsForCall(0)
				Expect(updatedService.Spec.Ports).To(HaveLen(1))
				Expect(updatedService.Spec.Ports[0].Name).To(Equal("http"))
				Expect(updatedService.Spec.Ports[0].Port).To(Equal(int32(80)))
			})
		})

		Context("when Get returns an unexpected error", func() {
			BeforeEach(func() {
				serviceInterface.GetReturns(nil, errors.New("unexpected error"))
			})

			It("calls Create assuming service doesn't exist", func() {
				Expect(serviceInterface.CreateCallCount()).To(Equal(1))
			})
		})

		Context("when Create fails", func() {
			BeforeEach(func() {
				notFoundError := k8s_errors.NewNotFound(schema.GroupResource{
					Group:    "",
					Resource: "services",
				}, "test-service")
				serviceInterface.GetReturns(nil, notFoundError)
				serviceInterface.CreateReturns(nil, errors.New("create failed"))
			})

			It("returns an error", func() {
				Expect(err).NotTo(BeNil())
				Expect(err.Error()).To(ContainSubstring("create service failed"))
			})

			It("calls Create", func() {
				Expect(serviceInterface.CreateCallCount()).To(Equal(1))
			})
		})

		Context("when Update fails", func() {
			BeforeEach(func() {
				existingService := service
				existingService.ResourceVersion = "123"
				serviceInterface.GetReturns(&existingService, nil)
				serviceInterface.UpdateReturns(nil, errors.New("update failed"))
			})

			It("returns an error", func() {
				Expect(err).NotTo(BeNil())
				Expect(err.Error()).To(ContainSubstring("update service failed"))
			})

			It("calls Update", func() {
				Expect(serviceInterface.UpdateCallCount()).To(Equal(1))
			})
		})

		Context("with context cancellation", func() {
			var cancelCtx context.Context
			var cancel context.CancelFunc

			BeforeEach(func() {
				cancelCtx, cancel = context.WithCancel(ctx)
				cancel() // Cancel immediately
				ctx = cancelCtx

				serviceInterface.GetReturns(nil, context.Canceled)
			})

			It("handles context cancellation gracefully", func() {
				Expect(serviceInterface.CreateCallCount()).To(Equal(1))
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
			name = k8s.Name("test-service")
		})

		JustBeforeEach(func() {
			err = serviceDeployer.Undeploy(ctx, namespace, name)
		})

		Context("when service exists", func() {
			BeforeEach(func() {
				serviceInterface.GetReturns(&service, nil)
				serviceInterface.DeleteReturns(nil)
			})

			It("returns no error", func() {
				Expect(err).To(BeNil())
			})

			It("calls Get to check if service exists", func() {
				Expect(serviceInterface.GetCallCount()).To(Equal(1))
				_, serviceName, _ := serviceInterface.GetArgsForCall(0)
				Expect(serviceName).To(Equal("test-service"))
			})

			It("calls Delete to remove the service", func() {
				Expect(serviceInterface.DeleteCallCount()).To(Equal(1))
				_, deletedName, _ := serviceInterface.DeleteArgsForCall(0)
				Expect(deletedName).To(Equal("test-service"))
			})
		})

		Context("when service does not exist", func() {
			BeforeEach(func() {
				notFoundError := k8s_errors.NewNotFound(schema.GroupResource{
					Group:    "",
					Resource: "services",
				}, "test-service")
				serviceInterface.GetReturns(nil, notFoundError)
			})

			It("returns no error", func() {
				Expect(err).To(BeNil())
			})

			It("calls Get to check if service exists", func() {
				Expect(serviceInterface.GetCallCount()).To(Equal(1))
			})

			It("does not call Delete", func() {
				Expect(serviceInterface.DeleteCallCount()).To(Equal(0))
			})
		})

		Context("when Get returns an unexpected error", func() {
			BeforeEach(func() {
				serviceInterface.GetReturns(nil, errors.New("unexpected error"))
			})

			It("does not call Delete", func() {
				Expect(serviceInterface.DeleteCallCount()).To(Equal(0))
			})

			It("returns no error (treats as not found)", func() {
				Expect(err).To(BeNil())
			})
		})

		Context("when Delete fails", func() {
			BeforeEach(func() {
				serviceInterface.GetReturns(&service, nil)
				serviceInterface.DeleteReturns(errors.New("delete failed"))
			})

			It("returns an error", func() {
				Expect(err).NotTo(BeNil())
				Expect(err.Error()).To(ContainSubstring("delete failed"))
			})

			It("calls Delete", func() {
				Expect(serviceInterface.DeleteCallCount()).To(Equal(1))
			})
		})

		Context("with context cancellation", func() {
			var cancelCtx context.Context
			var cancel context.CancelFunc

			BeforeEach(func() {
				cancelCtx, cancel = context.WithCancel(ctx)
				cancel() // Cancel immediately
				ctx = cancelCtx

				serviceInterface.GetReturns(nil, context.Canceled)
			})

			It("handles context cancellation gracefully", func() {
				Expect(err).To(BeNil())
				Expect(serviceInterface.DeleteCallCount()).To(Equal(0))
			})
		})
	})

	Describe("Constructor", func() {
		It("creates a new service deployer", func() {
			deployer := k8s.NewServiceDeployer(clientset)
			Expect(deployer).NotTo(BeNil())
		})
	})

	Describe("mergeService functionality", func() {
		Context("when updating existing service", func() {
			BeforeEach(func() {
				existingService := service
				existingService.ResourceVersion = "456"
				existingService.Spec.ClusterIP = "192.168.1.1"

				// Create a new service with different specs
				newService := service
				newService.Spec.Type = corev1.ServiceTypeNodePort
				newService.Spec.Ports = []corev1.ServicePort{
					{
						Name:       "https",
						Port:       443,
						TargetPort: intstr.FromInt(8443),
						Protocol:   corev1.ProtocolTCP,
					},
				}

				serviceInterface.GetReturns(&existingService, nil)
				serviceInterface.UpdateReturns(&newService, nil)

				service = newService
			})

			JustBeforeEach(func() {
				err = serviceDeployer.Deploy(ctx, service)
			})

			It("merges the services correctly", func() {
				Expect(err).To(BeNil())
				Expect(serviceInterface.UpdateCallCount()).To(Equal(1))
				_, updatedService, _ := serviceInterface.UpdateArgsForCall(0)

				// Should preserve existing ClusterIP and ResourceVersion
				Expect(updatedService.Spec.ClusterIP).To(Equal("192.168.1.1"))
				Expect(updatedService.ResourceVersion).To(Equal("456"))

				// Should update the new specifications
				Expect(updatedService.Spec.Type).To(Equal(corev1.ServiceTypeNodePort))
				Expect(updatedService.Spec.Ports).To(HaveLen(1))
				Expect(updatedService.Spec.Ports[0].Name).To(Equal("https"))
				Expect(updatedService.Spec.Ports[0].Port).To(Equal(int32(443)))
			})
		})
	})
})
