// Copyright (c) 2024 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package k8s_test

import (
	"context"
	stderrors "errors"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	v1 "k8s.io/api/networking/v1"
	k8s_errors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	"github.com/bborbe/k8s"
	"github.com/bborbe/k8s/mocks"
)

var _ = Describe("IngressDeployer", func() {
	var ctx context.Context
	var err error
	var ingressDeployer k8s.IngressDeployer
	var newIngress v1.Ingress
	var k8sInterface *mocks.K8sInterface
	var k8sNetworkingV1Interface *mocks.K8sNetworkingV1Interface
	var ingressInterface *mocks.K8sIngressInterface
	BeforeEach(func() {
		ctx = context.Background()
		newIngress = v1.Ingress{}

		ingressInterface = &mocks.K8sIngressInterface{}

		k8sNetworkingV1Interface = &mocks.K8sNetworkingV1Interface{}
		k8sNetworkingV1Interface.IngressesReturns(ingressInterface)

		k8sInterface = &mocks.K8sInterface{}
		k8sInterface.NetworkingV1Returns(k8sNetworkingV1Interface)

		ingressDeployer = k8s.NewIngressDeployer(k8sInterface)
	})
	Context("Deploy", func() {
		BeforeEach(func() {
			newIngress = v1.Ingress{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "test-ingress",
					Namespace: "test-namespace",
					Labels:    map[string]string{"new": "label"},
				},
			}
		})
		JustBeforeEach(func() {
			err = ingressDeployer.Deploy(ctx, newIngress)
		})
		Context("when ingress does not exist", func() {
			BeforeEach(func() {
				ingressInterface.GetReturns(nil, stderrors.New("banana"))
			})
			It("returns no error", func() {
				Expect(err).To(BeNil())
			})
			It("calls create", func() {
				Expect(ingressInterface.CreateCallCount()).To(Equal(1))
				_, ingress, _ := ingressInterface.CreateArgsForCall(0)
				Expect(ingress.Name).To(Equal("test-ingress"))
			})
			It("calls not update", func() {
				Expect(ingressInterface.UpdateCallCount()).To(Equal(0))
			})
		})
		Context("when ingress already exists and is different", func() {
			BeforeEach(func() {
				ingressInterface.GetReturns(&v1.Ingress{
					TypeMeta: metav1.TypeMeta{},
					ObjectMeta: metav1.ObjectMeta{
						Name:      "test-ingress",
						Namespace: "test-namespace",
						Labels:    map[string]string{"old": "label"},
					},
					Spec:   v1.IngressSpec{},
					Status: v1.IngressStatus{},
				}, nil)
			})
			It("returns no error", func() {
				Expect(err).To(BeNil())
			})
			It("calls not create", func() {
				Expect(ingressInterface.CreateCallCount()).To(Equal(0))
			})
			It("calls update", func() {
				Expect(ingressInterface.UpdateCallCount()).To(Equal(1))
				_, ingress, _ := ingressInterface.UpdateArgsForCall(0)
				Expect(ingress.Name).To(Equal("test-ingress"))
				Expect(ingress.Labels).To(HaveKeyWithValue("new", "label"))
			})
		})
		Context("when ingress already exists and is the same", func() {
			BeforeEach(func() {
				ingressInterface.GetReturns(&v1.Ingress{
					TypeMeta: metav1.TypeMeta{},
					ObjectMeta: metav1.ObjectMeta{
						Name:      "test-ingress",
						Namespace: "test-namespace",
						Labels:    map[string]string{"new": "label"},
					},
					Spec:   v1.IngressSpec{},
					Status: v1.IngressStatus{},
				}, nil)
			})
			It("returns no error", func() {
				Expect(err).To(BeNil())
			})
			It("calls not create", func() {
				Expect(ingressInterface.CreateCallCount()).To(Equal(0))
			})
			It("calls not update", func() {
				Expect(ingressInterface.UpdateCallCount()).To(Equal(0))
			})
		})
	})
	Describe("Undeploy", func() {
		var namespace k8s.Namespace
		var name k8s.Name

		BeforeEach(func() {
			namespace = k8s.Namespace("test-namespace")
			name = k8s.Name("test-ingress")
		})

		JustBeforeEach(func() {
			err = ingressDeployer.Undeploy(ctx, namespace, name)
		})

		Context("when ingress exists", func() {
			BeforeEach(func() {
				ingressInterface.GetReturns(&v1.Ingress{}, nil)
				ingressInterface.DeleteReturns(nil)
			})

			It("returns no error", func() {
				Expect(err).To(BeNil())
			})

			It("calls Get to check if ingress exists", func() {
				Expect(ingressInterface.GetCallCount()).To(Equal(1))
				_, ingressName, _ := ingressInterface.GetArgsForCall(0)
				Expect(ingressName).To(Equal("test-ingress"))
			})

			It("calls Delete to remove the ingress", func() {
				Expect(ingressInterface.DeleteCallCount()).To(Equal(1))
				_, deletedName, _ := ingressInterface.DeleteArgsForCall(0)
				Expect(deletedName).To(Equal("test-ingress"))
			})
		})

		Context("when ingress does not exist", func() {
			BeforeEach(func() {
				notFoundError := k8s_errors.NewNotFound(schema.GroupResource{
					Group:    "networking.k8s.io",
					Resource: "ingresses",
				}, "test-ingress")
				ingressInterface.GetReturns(nil, notFoundError)
			})

			It("returns no error", func() {
				Expect(err).To(BeNil())
			})

			It("calls Get to check if ingress exists", func() {
				Expect(ingressInterface.GetCallCount()).To(Equal(1))
			})

			It("does not call Delete", func() {
				Expect(ingressInterface.DeleteCallCount()).To(Equal(0))
			})
		})
	})
})
