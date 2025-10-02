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

	"github.com/bborbe/k8s"
	"github.com/bborbe/k8s/mocks"
)

var _ = Describe("ConfigMap Deployer", func() {
	var configMapDeployer k8s.ConfigMapDeployer
	var clientset *mocks.K8sInterface
	var coreV1 *mocks.K8sCoreV1Interface
	var configMapInterface *mocks.K8sConfigMapInterface
	var ctx context.Context
	var configMap corev1.ConfigMap
	var err error

	BeforeEach(func() {
		ctx = context.Background()
		clientset = &mocks.K8sInterface{}
		coreV1 = &mocks.K8sCoreV1Interface{}
		configMapInterface = &mocks.K8sConfigMapInterface{}

		clientset.CoreV1Returns(coreV1)
		coreV1.ConfigMapsReturns(configMapInterface)

		configMapDeployer = k8s.NewConfigMapDeployer(clientset)

		configMap = corev1.ConfigMap{
			ObjectMeta: metav1.ObjectMeta{
				Name:      "test-configmap",
				Namespace: "test-namespace",
				Labels: map[string]string{
					"app": "test-app",
				},
			},
			Data: map[string]string{
				"config.yaml": "database:\n  host: localhost\n  port: 5432",
				"app.json":    `{"name": "test-app", "version": "1.0.0"}`,
			},
			BinaryData: map[string][]byte{
				"binary.dat": []byte{0x89, 0x50, 0x4E, 0x47},
			},
		}
	})

	Describe("Get", func() {
		var namespace k8s.Namespace
		var name k8s.Name
		var retrievedConfigMap *corev1.ConfigMap

		BeforeEach(func() {
			namespace = k8s.Namespace("test-namespace")
			name = k8s.Name("test-configmap")
		})

		JustBeforeEach(func() {
			retrievedConfigMap, err = configMapDeployer.Get(ctx, namespace, name)
		})

		Context("when configmap exists", func() {
			BeforeEach(func() {
				configMapInterface.GetReturns(&configMap, nil)
			})

			It("returns no error", func() {
				Expect(err).To(BeNil())
			})

			It("returns the configmap", func() {
				Expect(retrievedConfigMap).NotTo(BeNil())
				Expect(retrievedConfigMap.Name).To(Equal("test-configmap"))
				Expect(retrievedConfigMap.Namespace).To(Equal("test-namespace"))
				Expect(
					retrievedConfigMap.Data,
				).To(HaveKeyWithValue("config.yaml", "database:\n  host: localhost\n  port: 5432"))
			})

			It("calls Get with correct parameters", func() {
				Expect(configMapInterface.GetCallCount()).To(Equal(1))
				_, configMapName, _ := configMapInterface.GetArgsForCall(0)
				Expect(configMapName).To(Equal("test-configmap"))
			})
		})

		Context("when configmap does not exist", func() {
			BeforeEach(func() {
				notFoundError := k8s_errors.NewNotFound(schema.GroupResource{
					Group:    "",
					Resource: "configmaps",
				}, "test-configmap")
				configMapInterface.GetReturns(nil, notFoundError)
			})

			It("returns an error", func() {
				Expect(err).NotTo(BeNil())
				Expect(
					err.Error(),
				).To(ContainSubstring("get configmap(test-configmap) in namespace(test-namespace) failed"))
			})

			It("returns nil configmap", func() {
				Expect(retrievedConfigMap).To(BeNil())
			})
		})

		Context("when Get returns an unexpected error", func() {
			BeforeEach(func() {
				configMapInterface.GetReturns(nil, errors.New("unexpected error"))
			})

			It("returns an error", func() {
				Expect(err).NotTo(BeNil())
				Expect(
					err.Error(),
				).To(ContainSubstring("get configmap(test-configmap) in namespace(test-namespace) failed"))
			})
		})
	})

	Describe("Deploy", func() {
		JustBeforeEach(func() {
			err = configMapDeployer.Deploy(ctx, configMap)
		})

		Context("when configmap does not exist", func() {
			BeforeEach(func() {
				notFoundError := k8s_errors.NewNotFound(schema.GroupResource{
					Group:    "",
					Resource: "configmaps",
				}, "test-configmap")
				configMapInterface.GetReturns(nil, notFoundError)
				configMapInterface.CreateReturns(&configMap, nil)
			})

			It("returns no error", func() {
				Expect(err).To(BeNil())
			})

			It("calls Get to check if configmap exists", func() {
				Expect(configMapInterface.GetCallCount()).To(Equal(1))
				_, name, _ := configMapInterface.GetArgsForCall(0)
				Expect(name).To(Equal("test-configmap"))
			})

			It("calls Create to create the configmap", func() {
				Expect(configMapInterface.CreateCallCount()).To(Equal(1))
				_, createdConfigMap, _ := configMapInterface.CreateArgsForCall(0)
				Expect(createdConfigMap.Name).To(Equal("test-configmap"))
				Expect(createdConfigMap.Namespace).To(Equal("test-namespace"))
				Expect(
					createdConfigMap.Data,
				).To(HaveKeyWithValue("config.yaml", "database:\n  host: localhost\n  port: 5432"))
				Expect(
					createdConfigMap.BinaryData,
				).To(HaveKeyWithValue("binary.dat", []byte{0x89, 0x50, 0x4E, 0x47}))
			})

			It("does not call Update", func() {
				Expect(configMapInterface.UpdateCallCount()).To(Equal(0))
			})
		})

		Context("when configmap already exists", func() {
			BeforeEach(func() {
				existingConfigMap := configMap
				existingConfigMap.ResourceVersion = "123"
				configMapInterface.GetReturns(&existingConfigMap, nil)
				configMapInterface.UpdateReturns(&configMap, nil)
			})

			It("returns no error", func() {
				Expect(err).To(BeNil())
			})

			It("calls Get to check if configmap exists", func() {
				Expect(configMapInterface.GetCallCount()).To(Equal(1))
				_, name, _ := configMapInterface.GetArgsForCall(0)
				Expect(name).To(Equal("test-configmap"))
			})

			It("calls Update to update the configmap", func() {
				Expect(configMapInterface.UpdateCallCount()).To(Equal(1))
				_, updatedConfigMap, _ := configMapInterface.UpdateArgsForCall(0)
				Expect(updatedConfigMap.Name).To(Equal("test-configmap"))
				Expect(updatedConfigMap.Namespace).To(Equal("test-namespace"))
			})

			It("preserves existing ResourceVersion in update", func() {
				Expect(configMapInterface.UpdateCallCount()).To(Equal(1))
				_, updatedConfigMap, _ := configMapInterface.UpdateArgsForCall(0)
				Expect(updatedConfigMap.ResourceVersion).To(Equal("123"))
			})

			It("does not call Create", func() {
				Expect(configMapInterface.CreateCallCount()).To(Equal(0))
			})
		})

		Context("when configmap exists with different data", func() {
			BeforeEach(func() {
				existingConfigMap := configMap
				existingConfigMap.ResourceVersion = "123"
				existingConfigMap.Data = map[string]string{
					"old-config.yaml": "old: value",
				}
				configMapInterface.GetReturns(&existingConfigMap, nil)
				configMapInterface.UpdateReturns(&configMap, nil)
			})

			It("updates the configmap data", func() {
				Expect(configMapInterface.UpdateCallCount()).To(Equal(1))
				_, updatedConfigMap, _ := configMapInterface.UpdateArgsForCall(0)
				Expect(
					updatedConfigMap.Data,
				).To(HaveKeyWithValue("config.yaml", "database:\n  host: localhost\n  port: 5432"))
				Expect(
					updatedConfigMap.Data,
				).To(HaveKeyWithValue("app.json", `{"name": "test-app", "version": "1.0.0"}`))
				Expect(updatedConfigMap.Data).NotTo(HaveKey("old-config.yaml"))
			})
		})

		Context("when configmap exists with different binary data", func() {
			BeforeEach(func() {
				existingConfigMap := configMap
				existingConfigMap.ResourceVersion = "123"
				existingConfigMap.BinaryData = map[string][]byte{
					"old-binary.dat": []byte{0x01, 0x02, 0x03},
				}
				configMapInterface.GetReturns(&existingConfigMap, nil)
				configMapInterface.UpdateReturns(&configMap, nil)
			})

			It("updates the configmap binary data", func() {
				Expect(configMapInterface.UpdateCallCount()).To(Equal(1))
				_, updatedConfigMap, _ := configMapInterface.UpdateArgsForCall(0)
				Expect(
					updatedConfigMap.BinaryData,
				).To(HaveKeyWithValue("binary.dat", []byte{0x89, 0x50, 0x4E, 0x47}))
				Expect(updatedConfigMap.BinaryData).NotTo(HaveKey("old-binary.dat"))
			})
		})

		Context("when configmap exists with different labels", func() {
			BeforeEach(func() {
				existingConfigMap := configMap
				existingConfigMap.ResourceVersion = "123"
				existingConfigMap.Labels = map[string]string{
					"old-label": "old-value",
				}
				configMapInterface.GetReturns(&existingConfigMap, nil)
				configMapInterface.UpdateReturns(&configMap, nil)
			})

			It("updates the configmap labels", func() {
				Expect(configMapInterface.UpdateCallCount()).To(Equal(1))
				_, updatedConfigMap, _ := configMapInterface.UpdateArgsForCall(0)
				Expect(updatedConfigMap.Labels).To(HaveKeyWithValue("app", "test-app"))
				Expect(updatedConfigMap.Labels).NotTo(HaveKey("old-label"))
			})
		})

		Context("when Get returns an unexpected error", func() {
			BeforeEach(func() {
				configMapInterface.GetReturns(nil, errors.New("unexpected error"))
			})

			It("calls Create assuming configmap doesn't exist", func() {
				Expect(configMapInterface.CreateCallCount()).To(Equal(1))
			})
		})

		Context("when Create fails", func() {
			BeforeEach(func() {
				notFoundError := k8s_errors.NewNotFound(schema.GroupResource{
					Group:    "",
					Resource: "configmaps",
				}, "test-configmap")
				configMapInterface.GetReturns(nil, notFoundError)
				configMapInterface.CreateReturns(nil, errors.New("create failed"))
			})

			It("returns an error", func() {
				Expect(err).NotTo(BeNil())
				Expect(err.Error()).To(ContainSubstring("create configmap failed"))
			})

			It("calls Create", func() {
				Expect(configMapInterface.CreateCallCount()).To(Equal(1))
			})
		})

		Context("when Update fails", func() {
			BeforeEach(func() {
				existingConfigMap := configMap
				existingConfigMap.ResourceVersion = "123"
				configMapInterface.GetReturns(&existingConfigMap, nil)
				configMapInterface.UpdateReturns(nil, errors.New("update failed"))
			})

			It("returns an error", func() {
				Expect(err).NotTo(BeNil())
				Expect(err.Error()).To(ContainSubstring("update configmap failed"))
			})

			It("calls Update", func() {
				Expect(configMapInterface.UpdateCallCount()).To(Equal(1))
			})
		})

		Context("with context cancellation", func() {
			var cancelCtx context.Context
			var cancel context.CancelFunc

			BeforeEach(func() {
				cancelCtx, cancel = context.WithCancel(ctx)
				cancel() // Cancel immediately
				ctx = cancelCtx

				configMapInterface.GetReturns(nil, context.Canceled)
			})

			It("handles context cancellation gracefully", func() {
				Expect(configMapInterface.CreateCallCount()).To(Equal(1))
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
			name = k8s.Name("test-configmap")
		})

		JustBeforeEach(func() {
			err = configMapDeployer.Undeploy(ctx, namespace, name)
		})

		Context("when configmap exists", func() {
			BeforeEach(func() {
				configMapInterface.GetReturns(&configMap, nil)
				configMapInterface.DeleteReturns(nil)
			})

			It("returns no error", func() {
				Expect(err).To(BeNil())
			})

			It("calls Get to check if configmap exists", func() {
				Expect(configMapInterface.GetCallCount()).To(Equal(1))
				_, configMapName, _ := configMapInterface.GetArgsForCall(0)
				Expect(configMapName).To(Equal("test-configmap"))
			})

			It("calls Delete to remove the configmap", func() {
				Expect(configMapInterface.DeleteCallCount()).To(Equal(1))
				_, deletedName, _ := configMapInterface.DeleteArgsForCall(0)
				Expect(deletedName).To(Equal("test-configmap"))
			})
		})

		Context("when configmap does not exist", func() {
			BeforeEach(func() {
				notFoundError := k8s_errors.NewNotFound(schema.GroupResource{
					Group:    "",
					Resource: "configmaps",
				}, "test-configmap")
				configMapInterface.GetReturns(nil, notFoundError)
			})

			It("returns no error", func() {
				Expect(err).To(BeNil())
			})

			It("calls Get to check if configmap exists", func() {
				Expect(configMapInterface.GetCallCount()).To(Equal(1))
			})

			It("does not call Delete", func() {
				Expect(configMapInterface.DeleteCallCount()).To(Equal(0))
			})
		})

		Context("when Get returns an unexpected error", func() {
			BeforeEach(func() {
				configMapInterface.GetReturns(nil, errors.New("unexpected error"))
			})

			It("does not call Delete", func() {
				Expect(configMapInterface.DeleteCallCount()).To(Equal(0))
			})

			It("returns no error (treats as not found)", func() {
				Expect(err).To(BeNil())
			})
		})

		Context("when Delete fails", func() {
			BeforeEach(func() {
				configMapInterface.GetReturns(&configMap, nil)
				configMapInterface.DeleteReturns(errors.New("delete failed"))
			})

			It("returns an error", func() {
				Expect(err).NotTo(BeNil())
				Expect(err.Error()).To(ContainSubstring("delete failed"))
			})

			It("calls Delete", func() {
				Expect(configMapInterface.DeleteCallCount()).To(Equal(1))
			})
		})

		Context("with context cancellation", func() {
			var cancelCtx context.Context
			var cancel context.CancelFunc

			BeforeEach(func() {
				cancelCtx, cancel = context.WithCancel(ctx)
				cancel() // Cancel immediately
				ctx = cancelCtx

				configMapInterface.GetReturns(nil, context.Canceled)
			})

			It("handles context cancellation gracefully", func() {
				Expect(err).To(BeNil())
				Expect(configMapInterface.DeleteCallCount()).To(Equal(0))
			})
		})
	})

	Describe("Constructor", func() {
		It("creates a new configmap deployer", func() {
			deployer := k8s.NewConfigMapDeployer(clientset)
			Expect(deployer).NotTo(BeNil())
		})
	})

	Describe("mergeConfigMap functionality", func() {
		Context("when updating existing configmap", func() {
			BeforeEach(func() {
				existingConfigMap := configMap
				existingConfigMap.ResourceVersion = "456"

				// Create a new configmap with different data
				newConfigMap := configMap
				newConfigMap.Data = map[string]string{
					"new-config.yaml": "new: configuration",
				}
				newConfigMap.BinaryData = map[string][]byte{
					"new-binary.dat": []byte{0xAB, 0xCD, 0xEF},
				}

				configMapInterface.GetReturns(&existingConfigMap, nil)
				configMapInterface.UpdateReturns(&newConfigMap, nil)

				configMap = newConfigMap
			})

			JustBeforeEach(func() {
				err = configMapDeployer.Deploy(ctx, configMap)
			})

			It("merges the configmaps correctly", func() {
				Expect(err).To(BeNil())
				Expect(configMapInterface.UpdateCallCount()).To(Equal(1))
				_, updatedConfigMap, _ := configMapInterface.UpdateArgsForCall(0)

				// Should preserve existing ResourceVersion
				Expect(updatedConfigMap.ResourceVersion).To(Equal("456"))

				// Should update the new data
				Expect(
					updatedConfigMap.Data,
				).To(HaveKeyWithValue("new-config.yaml", "new: configuration"))
				Expect(
					updatedConfigMap.BinaryData,
				).To(HaveKeyWithValue("new-binary.dat", []byte{0xAB, 0xCD, 0xEF}))
			})
		})
	})

	Describe("ConfigMap data validation scenarios", func() {
		JustBeforeEach(func() {
			err = configMapDeployer.Deploy(ctx, configMap)
		})

		Context("when deploying with large data", func() {
			BeforeEach(func() {
				// Create a configmap with large data (close to 1MB limit)
				largeData := make([]byte, 1024*1024-1000) // Just under 1MB
				for i := range largeData {
					largeData[i] = byte(i % 256)
				}

				configMap.BinaryData = map[string][]byte{
					"large-file.bin": largeData,
				}

				notFoundError := k8s_errors.NewNotFound(schema.GroupResource{
					Group:    "",
					Resource: "configmaps",
				}, "test-configmap")
				configMapInterface.GetReturns(nil, notFoundError)
				configMapInterface.CreateReturns(&configMap, nil)
			})

			It("handles large data correctly", func() {
				Expect(err).To(BeNil())
				Expect(configMapInterface.CreateCallCount()).To(Equal(1))
				_, createdConfigMap, _ := configMapInterface.CreateArgsForCall(0)
				Expect(createdConfigMap.BinaryData).To(HaveKey("large-file.bin"))
				Expect(
					len(createdConfigMap.BinaryData["large-file.bin"]),
				).To(Equal(1024*1024 - 1000))
			})
		})

		Context("when deploying with empty data", func() {
			BeforeEach(func() {
				configMap.Data = map[string]string{}
				configMap.BinaryData = map[string][]byte{}

				notFoundError := k8s_errors.NewNotFound(schema.GroupResource{
					Group:    "",
					Resource: "configmaps",
				}, "test-configmap")
				configMapInterface.GetReturns(nil, notFoundError)
				configMapInterface.CreateReturns(&configMap, nil)
			})

			It("handles empty data correctly", func() {
				Expect(err).To(BeNil())
				Expect(configMapInterface.CreateCallCount()).To(Equal(1))
				_, createdConfigMap, _ := configMapInterface.CreateArgsForCall(0)
				Expect(createdConfigMap.Data).To(BeEmpty())
				Expect(createdConfigMap.BinaryData).To(BeEmpty())
			})
		})
	})
})
