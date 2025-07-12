// Copyright (c) 2024 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package k8s_test

import (
	"context"
	"errors"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"k8s.io/client-go/tools/cache"

	"github.com/bborbe/k8s"
	"github.com/bborbe/k8s/mocks"
)

// TestResource is a test type that implements k8s.Type
type TestResource struct {
	Name      string
	Namespace string
	Replicas  int32
}

func (t TestResource) Equal(other k8s.Type) bool {
	return t.Identifier() == other.Identifier()
}

func (t TestResource) Validate(ctx context.Context) error {
	return nil
}

func (t TestResource) Identifier() k8s.Identifier {
	return k8s.Identifier(t.Namespace + "/" + t.Name)
}

func (t TestResource) String() string {
	return t.Namespace + "/" + t.Name
}

var _ = Describe("ResourceEventHandler", func() {
	var ctx context.Context
	var eventHandler *mocks.EventHandler[TestResource]
	var resourceEventHandler cache.ResourceEventHandler

	BeforeEach(func() {
		ctx = context.Background()
		eventHandler = &mocks.EventHandler[TestResource]{}
		resourceEventHandler = k8s.NewResourceEventHandler(ctx, eventHandler)
	})

	Context("AddFunc", func() {
		It("calls OnAdd when object is valid type", func() {
			resource := &TestResource{
				Name:      "test-resource",
				Namespace: "default",
				Replicas:  1,
			}

			eventHandler.OnAddReturns(nil)

			resourceEventHandler.OnAdd(resource, false)

			Expect(eventHandler.OnAddCallCount()).To(Equal(1))
			actualCtx, actualResource := eventHandler.OnAddArgsForCall(0)
			Expect(actualCtx).To(Equal(ctx))
			Expect(actualResource).To(Equal(*resource))
		})

		It("handles OnAdd error gracefully", func() {
			resource := &TestResource{
				Name:      "test-resource",
				Namespace: "default",
			}

			eventHandler.OnAddReturns(errors.New("test error"))

			Expect(func() {
				resourceEventHandler.OnAdd(resource, false)
			}).ToNot(Panic())

			Expect(eventHandler.OnAddCallCount()).To(Equal(1))
		})

		It("handles invalid object type gracefully", func() {
			invalidObj := "not a resource"

			Expect(func() {
				resourceEventHandler.OnAdd(invalidObj, false)
			}).ToNot(Panic())

			Expect(eventHandler.OnAddCallCount()).To(Equal(0))
		})
	})

	Context("UpdateFunc", func() {
		It("calls OnUpdate when objects are valid type", func() {
			oldResource := &TestResource{
				Name:      "test-resource",
				Namespace: "default",
				Replicas:  1,
			}
			newResource := &TestResource{
				Name:      "test-resource",
				Namespace: "default",
				Replicas:  3,
			}

			eventHandler.OnUpdateReturns(nil)

			resourceEventHandler.OnUpdate(oldResource, newResource)

			Expect(eventHandler.OnUpdateCallCount()).To(Equal(1))
			actualCtx, actualOldResource, actualNewResource := eventHandler.OnUpdateArgsForCall(0)
			Expect(actualCtx).To(Equal(ctx))
			Expect(actualOldResource).To(Equal(*oldResource))
			Expect(actualNewResource).To(Equal(*newResource))
		})

		It("handles OnUpdate error gracefully", func() {
			oldResource := &TestResource{}
			newResource := &TestResource{}

			eventHandler.OnUpdateReturns(errors.New("test error"))

			Expect(func() {
				resourceEventHandler.OnUpdate(oldResource, newResource)
			}).ToNot(Panic())

			Expect(eventHandler.OnUpdateCallCount()).To(Equal(1))
		})

		It("handles invalid old object type gracefully", func() {
			invalidOldObj := "not a resource"
			newResource := &TestResource{}

			Expect(func() {
				resourceEventHandler.OnUpdate(invalidOldObj, newResource)
			}).ToNot(Panic())

			Expect(eventHandler.OnUpdateCallCount()).To(Equal(0))
		})

		It("handles invalid new object type gracefully", func() {
			oldResource := &TestResource{}
			invalidNewObj := "not a resource"

			Expect(func() {
				resourceEventHandler.OnUpdate(oldResource, invalidNewObj)
			}).ToNot(Panic())

			Expect(eventHandler.OnUpdateCallCount()).To(Equal(0))
		})
	})

	Context("DeleteFunc", func() {
		It("calls OnDelete when object is valid type", func() {
			resource := &TestResource{
				Name:      "test-resource",
				Namespace: "default",
			}

			eventHandler.OnDeleteReturns(nil)

			resourceEventHandler.OnDelete(resource)

			Expect(eventHandler.OnDeleteCallCount()).To(Equal(1))
			actualCtx, actualResource := eventHandler.OnDeleteArgsForCall(0)
			Expect(actualCtx).To(Equal(ctx))
			Expect(actualResource).To(Equal(*resource))
		})

		It("handles OnDelete error gracefully", func() {
			resource := &TestResource{
				Name:      "test-resource",
				Namespace: "default",
			}

			eventHandler.OnDeleteReturns(errors.New("test error"))

			Expect(func() {
				resourceEventHandler.OnDelete(resource)
			}).ToNot(Panic())

			Expect(eventHandler.OnDeleteCallCount()).To(Equal(1))
		})

		It("handles invalid object type gracefully", func() {
			invalidObj := "not a resource"

			Expect(func() {
				resourceEventHandler.OnDelete(invalidObj)
			}).ToNot(Panic())

			Expect(eventHandler.OnDeleteCallCount()).To(Equal(0))
		})
	})
})
