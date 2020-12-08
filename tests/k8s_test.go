package tests

import (
	"context"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	v1 "k8s.io/api/core/v1"
	_ "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/envtest"
)

var testEnv *envtest.Environment
var _ = Describe("k8s", func() {
	Context("Cluster being in the desired state", func() {
		It("should have exact node count", func() {
			Eventually(func() bool {
				existingCluster := true
				testEnv = &envtest.Environment{UseExistingCluster: &existingCluster}
				cfg, err := testEnv.Start()
				Expect(err).ToNot(HaveOccurred())
				Expect(cfg).ToNot(BeNil())

				k8sClient, err := client.New(cfg, client.Options{})
				Expect(err).ToNot(HaveOccurred())
				Expect(k8sClient).ToNot(BeNil())

				nodes := &v1.NodeList{}
				err = k8sClient.List(context.TODO(), nodes)
				Expect(err).Should(BeNil())
				return len(nodes.Items) >= config.K8s.Nodes
			}, timeout, interval).Should(BeTrue())
		})
	})
})
