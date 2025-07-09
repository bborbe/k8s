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
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	"github.com/bborbe/k8s"
	"github.com/bborbe/k8s/mocks"
)

var _ = Describe("Deployment Deployer", func() {
	var deploymentDeployer k8s.DeploymentDeployer
	var clientset *mocks.K8sInterface
	var appsV1 *mocks.K8sAppsV1Interface
	var deploymentInterface *mocks.K8sDeploymentInterface
	var ctx context.Context
	var deployment appsv1.Deployment
	var err error

	BeforeEach(func() {
		ctx = context.Background()
		clientset = &mocks.K8sInterface{}
		appsV1 = &mocks.K8sAppsV1Interface{}
		deploymentInterface = &mocks.K8sDeploymentInterface{}

		clientset.AppsV1Returns(appsV1)
		appsV1.DeploymentsReturns(deploymentInterface)

		deploymentDeployer = k8s.NewDeploymentDeployer(clientset)

		deployment = appsv1.Deployment{
			ObjectMeta: metav1.ObjectMeta{
				Name:      "test-deployment",
				Namespace: "test-namespace",
			},
			Spec: appsv1.DeploymentSpec{
				Replicas: &[]int32{1}[0],
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
			},
		}
	})

	Describe("Deploy", func() {
		JustBeforeEach(func() {
			err = deploymentDeployer.Deploy(ctx, deployment)
		})

		Context("when deployment does not exist", func() {
			BeforeEach(func() {
				notFoundError := k8s_errors.NewNotFound(schema.GroupResource{
					Group:    "apps",
					Resource: "deployments",
				}, "test-deployment")
				deploymentInterface.GetReturns(nil, notFoundError)
				deploymentInterface.CreateReturns(&deployment, nil)
			})

			It("returns no error", func() {
				Expect(err).To(BeNil())
			})

			It("calls Get to check if deployment exists", func() {
				Expect(deploymentInterface.GetCallCount()).To(Equal(1))
				_, name, _ := deploymentInterface.GetArgsForCall(0)
				Expect(name).To(Equal("test-deployment"))
			})

			It("calls Create to create the deployment", func() {
				Expect(deploymentInterface.CreateCallCount()).To(Equal(1))
				_, createdDeployment, _ := deploymentInterface.CreateArgsForCall(0)
				Expect(createdDeployment.Name).To(Equal("test-deployment"))
				Expect(createdDeployment.Namespace).To(Equal("test-namespace"))
			})

			It("does not call Update", func() {
				Expect(deploymentInterface.UpdateCallCount()).To(Equal(0))
			})
		})

		Context("when deployment already exists", func() {
			BeforeEach(func() {
				existingDeployment := deployment
				existingDeployment.ResourceVersion = "123"
				deploymentInterface.GetReturns(&existingDeployment, nil)
				deploymentInterface.UpdateReturns(&deployment, nil)
			})

			It("returns no error", func() {
				Expect(err).To(BeNil())
			})

			It("calls Get to check if deployment exists", func() {
				Expect(deploymentInterface.GetCallCount()).To(Equal(1))
				_, name, _ := deploymentInterface.GetArgsForCall(0)
				Expect(name).To(Equal("test-deployment"))
			})

			It("calls Update to update the deployment", func() {
				Expect(deploymentInterface.UpdateCallCount()).To(Equal(1))
				_, updatedDeployment, _ := deploymentInterface.UpdateArgsForCall(0)
				Expect(updatedDeployment.Name).To(Equal("test-deployment"))
				Expect(updatedDeployment.Namespace).To(Equal("test-namespace"))
			})

			It("does not call Create", func() {
				Expect(deploymentInterface.CreateCallCount()).To(Equal(0))
			})
		})

		Context("when Get returns an unexpected error", func() {
			BeforeEach(func() {
				deploymentInterface.GetReturns(nil, errors.New("unexpected error"))
			})

			It("calls Create assuming deployment doesn't exist", func() {
				Expect(deploymentInterface.CreateCallCount()).To(Equal(1))
			})
		})

		Context("when Create fails", func() {
			BeforeEach(func() {
				notFoundError := k8s_errors.NewNotFound(schema.GroupResource{
					Group:    "apps",
					Resource: "deployments",
				}, "test-deployment")
				deploymentInterface.GetReturns(nil, notFoundError)
				deploymentInterface.CreateReturns(nil, errors.New("create failed"))
			})

			It("returns an error", func() {
				Expect(err).NotTo(BeNil())
				Expect(err.Error()).To(ContainSubstring("create deployment failed"))
			})

			It("calls Create", func() {
				Expect(deploymentInterface.CreateCallCount()).To(Equal(1))
			})
		})

		Context("when Update fails", func() {
			BeforeEach(func() {
				existingDeployment := deployment
				deploymentInterface.GetReturns(&existingDeployment, nil)
				deploymentInterface.UpdateReturns(nil, errors.New("update failed"))
			})

			It("returns an error", func() {
				Expect(err).NotTo(BeNil())
				Expect(err.Error()).To(ContainSubstring("update deployment failed"))
			})

			It("calls Update", func() {
				Expect(deploymentInterface.UpdateCallCount()).To(Equal(1))
			})
		})

		Context("with context cancellation", func() {
			var cancelCtx context.Context
			var cancel context.CancelFunc

			BeforeEach(func() {
				cancelCtx, cancel = context.WithCancel(ctx)
				cancel() // Cancel immediately
				ctx = cancelCtx

				deploymentInterface.GetReturns(nil, context.Canceled)
			})

			It("handles context cancellation gracefully", func() {
				Expect(deploymentInterface.CreateCallCount()).To(Equal(1))
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
			name = k8s.Name("test-deployment")
		})

		JustBeforeEach(func() {
			err = deploymentDeployer.Undeploy(ctx, namespace, name)
		})

		Context("when deployment exists", func() {
			BeforeEach(func() {
				deploymentInterface.GetReturns(&deployment, nil)
				deploymentInterface.DeleteReturns(nil)
			})

			It("returns no error", func() {
				Expect(err).To(BeNil())
			})

			It("calls Get to check if deployment exists", func() {
				Expect(deploymentInterface.GetCallCount()).To(Equal(1))
				_, deploymentName, _ := deploymentInterface.GetArgsForCall(0)
				Expect(deploymentName).To(Equal("test-deployment"))
			})

			It("calls Delete to remove the deployment", func() {
				Expect(deploymentInterface.DeleteCallCount()).To(Equal(1))
				_, deletedName, _ := deploymentInterface.DeleteArgsForCall(0)
				Expect(deletedName).To(Equal("test-deployment"))
			})
		})

		Context("when deployment does not exist", func() {
			BeforeEach(func() {
				notFoundError := k8s_errors.NewNotFound(schema.GroupResource{
					Group:    "apps",
					Resource: "deployments",
				}, "test-deployment")
				deploymentInterface.GetReturns(nil, notFoundError)
			})

			It("returns no error", func() {
				Expect(err).To(BeNil())
			})

			It("calls Get to check if deployment exists", func() {
				Expect(deploymentInterface.GetCallCount()).To(Equal(1))
			})

			It("does not call Delete", func() {
				Expect(deploymentInterface.DeleteCallCount()).To(Equal(0))
			})
		})

		Context("when Get returns an unexpected error", func() {
			BeforeEach(func() {
				deploymentInterface.GetReturns(nil, errors.New("unexpected error"))
			})

			It("does not call Delete", func() {
				Expect(deploymentInterface.DeleteCallCount()).To(Equal(0))
			})

			It("returns no error (treats as not found)", func() {
				Expect(err).To(BeNil())
			})
		})

		Context("when Delete fails", func() {
			BeforeEach(func() {
				deploymentInterface.GetReturns(&deployment, nil)
				deploymentInterface.DeleteReturns(errors.New("delete failed"))
			})

			It("returns an error", func() {
				Expect(err).NotTo(BeNil())
				Expect(err.Error()).To(ContainSubstring("delete failed"))
			})

			It("calls Delete", func() {
				Expect(deploymentInterface.DeleteCallCount()).To(Equal(1))
			})
		})

		Context("with context cancellation", func() {
			var cancelCtx context.Context
			var cancel context.CancelFunc

			BeforeEach(func() {
				cancelCtx, cancel = context.WithCancel(ctx)
				cancel() // Cancel immediately
				ctx = cancelCtx

				deploymentInterface.GetReturns(nil, context.Canceled)
			})

			It("handles context cancellation gracefully", func() {
				Expect(err).To(BeNil())
				Expect(deploymentInterface.DeleteCallCount()).To(Equal(0))
			})
		})
	})

	Describe("Constructor", func() {
		It("creates a new deployment deployer", func() {
			deployer := k8s.NewDeploymentDeployer(clientset)
			Expect(deployer).NotTo(BeNil())
		})
	})
})
