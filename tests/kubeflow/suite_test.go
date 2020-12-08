package kubeflow_test

import (
	"math/rand"
	"os"
	"testing"
	"time"

	nbv1beta1 "github.com/kubeflow/kubeflow/components/notebook-controller/api/v1beta1"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/log/zap"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	apiextensionsv1beta1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/envtest"
	"sigs.k8s.io/controller-runtime/pkg/envtest/printer"
)

// These tests use Ginkgo (BDD-style Go testing framework). Refer to
// http://onsi.github.io/ginkgo/ to learn more about Ginkgo.


var (
	useExistingCluster bool // flag for using an existing cluster for testing.
	cfg                *rest.Config
	k8sClient          client.Client // You'll be using this client in your tests.
	testEnv            *envtest.Environment
	clusterName        string // cluster that was dynamically created.
)

func init() {
	rand.Seed(time.Now().UnixNano())
	_ = apiextensionsv1beta1.AddToScheme(scheme.Scheme)
}

// randStringRunes is used for generating random string.
func randStringRunes(n int) string {
	letterRunes := []rune("abcdefghijklmnopqrstuvwxyz") // for generating random names for tests.
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

// getEnvDefault returns the value of the given environment variable or a
// default value if the given environment variable is not set.
func getEnvDefault(variable string, defaultVal string) string {
	envVar, exists := os.LookupEnv(variable)
	if !exists {
		return defaultVal
	}
	return envVar
}

// Setting up the ENV. This will run once before all the tests.
var _ = SynchronizedBeforeSuite(func() []byte {
	logf.SetLogger(zap.LoggerTo(GinkgoWriter, true))
	By("Creating a k8s cluster.")
	existingCluster := true
	testEnv = &envtest.Environment{UseExistingCluster: &existingCluster}
	var err error
	cfg, err = testEnv.Start()
	Expect(err).ToNot(HaveOccurred())
	Expect(cfg).ToNot(BeNil())

	err = nbv1beta1.AddToScheme(scheme.Scheme)
	Expect(err).NotTo(HaveOccurred())


	// +kubebuilder:scaffold:scheme

	k8sClient, err = client.New(cfg, client.Options{Scheme: scheme.Scheme})
	Expect(err).ToNot(HaveOccurred())
	Expect(k8sClient).ToNot(BeNil())

	return nil

}, func(data []byte) {
})

// Cleaning the ENV

func TestAPIs(t *testing.T) {
	RegisterFailHandler(Fail)

	RunSpecsWithDefaultAndCustomReporters(t,
		"Controller Suite",
		[]Reporter{printer.NewlineReporter{}})
}
