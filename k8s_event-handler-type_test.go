// Copyright (c) 2024 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package k8s_test

import (
	"context"

	"github.com/bborbe/k8s"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

type MyType struct {
	Name string
}

func (m MyType) Equal(other k8s.Type) bool {
	return m.Identifier() == other.Identifier()
}

func (m MyType) Validate(ctx context.Context) error {
	return nil
}

func (m MyType) Identifier() k8s.Identifier {
	return k8s.Identifier(m.Name)
}

func (m MyType) String() string {
	return m.Name
}

var _ = Describe("EventHandler", func() {
	var ctx context.Context
	var err error
	var eventHandlerAlert k8s.EventHandler[MyType]
	var alerts []MyType
	var alertA, alertB MyType
	BeforeEach(func() {
		ctx = context.Background()
		alertA = MyType{
			Name: "a",
		}
		alertB = MyType{
			Name: "b",
		}
		eventHandlerAlert = k8s.NewEventHandler[MyType]()
	})
	Context("Get", func() {
		JustBeforeEach(func() {
			alerts, err = eventHandlerAlert.Get(ctx)
		})
		Context("empty", func() {
			It("returns no alerts", func() {
				Expect(alerts).To(HaveLen(0))
			})
			It("returns no error", func() {
				Expect(err).To(BeNil())
			})
		})

		Context("add", func() {
			BeforeEach(func() {
				err = eventHandlerAlert.OnAdd(ctx, alertA)
				Expect(err).To(BeNil())
			})
			It("returns one alert", func() {
				Expect(alerts).To(HaveLen(1))
			})
			It("returns no error", func() {
				Expect(err).To(BeNil())
			})
		})
		Context("update", func() {
			BeforeEach(func() {
				err = eventHandlerAlert.OnUpdate(ctx, alertA, alertB)
				Expect(err).To(BeNil())
			})
			It("returns one alert", func() {
				Expect(alerts).To(HaveLen(1))
			})
			It("returns no error", func() {
				Expect(err).To(BeNil())
			})
		})
		Context("delete", func() {
			BeforeEach(func() {
				err = eventHandlerAlert.OnAdd(ctx, alertA)
				Expect(err).To(BeNil())
				err = eventHandlerAlert.OnDelete(ctx, alertA)
				Expect(err).To(BeNil())
			})
			It("returns no alerts", func() {
				Expect(alerts).To(HaveLen(0))
			})
			It("returns no error", func() {
				Expect(err).To(BeNil())
			})
		})
	})
})
