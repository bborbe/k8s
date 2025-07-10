package k8s_test

import (
	"context"
	"github.com/bborbe/k8s"
	corev1 "k8s.io/api/core/v1"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	batchv1 "k8s.io/api/batch/v1"
)

var _ = Describe("CronJob Builder", func() {
	var cronJobBuilder k8s.CronJobBuilder
	var cronJob *batchv1.CronJob
	var err error
	var ctx context.Context
	var objectMetaBuilder k8s.ObjectMetaBuilder
	var podSpecBuilder k8s.PodSpecBuilder
	var containersBuilder k8s.ContainersBuilder
	BeforeEach(func() {
		ctx = context.Background()

		objectMetaBuilder = k8s.NewObjectMetaBuilder()
		objectMetaBuilder.SetName("my-object")
		objectMetaBuilder.SetNamespace("my-namespace")

		containersBuilder = k8s.NewContainersBuilder()
		containersBuilder.SetContainers([]corev1.Container{
			{
				Name: "service",
			},
		})

		podSpecBuilder = k8s.NewPodSpecBuilder()
		podSpecBuilder.SetContainersBuilder(containersBuilder)
		podSpecBuilder.SetRestartPolicy(corev1.RestartPolicyOnFailure)

		cronJobBuilder = k8s.NewCronJobBuilder()
		cronJobBuilder.SetObjectMetaBuild(objectMetaBuilder)
		cronJobBuilder.SetPodSpecBuilder(podSpecBuilder)
	})
	Context("Build", func() {
		JustBeforeEach(func() {
			cronJob, err = cronJobBuilder.Build(ctx)
		})
		It("returns no error", func() {
			Expect(err).To(BeNil())
		})
		Context("default", func() {
			It("returns cronjob", func() {
				Expect(cronJob).NotTo(BeNil())
			})
		})
	})
})
