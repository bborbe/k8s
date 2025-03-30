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
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

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
		JustBeforeEach(func() {
			err = ingressDeployer.Deploy(ctx, newIngress)
		})
		Context("Create", func() {
			BeforeEach(func() {
				ingressInterface.GetReturns(nil, stderrors.New("banana"))
			})
			It("returns no error", func() {
				Expect(err).To(BeNil())
			})
			It("calls create", func() {
				Expect(ingressInterface.CreateCallCount()).To(Equal(1))
			})
			It("calls not update", func() {
				Expect(ingressInterface.UpdateCallCount()).To(Equal(0))
			})
		})
		Context("Update", func() {
			BeforeEach(func() {
				ingressInterface.GetReturns(&v1.Ingress{
					TypeMeta: metav1.TypeMeta{},
					ObjectMeta: metav1.ObjectMeta{
						Labels: map[string]string{"new": "label"},
					},
					Spec:   v1.IngressSpec{},
					Status: v1.IngressStatus{},
				}, nil)
			})
			It("returns no error", func() {
				Expect(err).To(BeNil())
			})
			It("calls create", func() {
				Expect(ingressInterface.CreateCallCount()).To(Equal(0))
			})
			It("calls not update", func() {
				Expect(ingressInterface.UpdateCallCount()).To(Equal(1))
			})
		})
		Context("Unchanged", func() {
			BeforeEach(func() {
				ingressInterface.GetReturns(&v1.Ingress{
					TypeMeta:   metav1.TypeMeta{},
					ObjectMeta: metav1.ObjectMeta{},
					Spec:       v1.IngressSpec{},
					Status:     v1.IngressStatus{},
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
})
