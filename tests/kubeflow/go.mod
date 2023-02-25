module github.com/naveensrinivasan/improved-invention/tests/kubeflow

go 1.15

require (
	github.com/kubeflow/kubeflow/components/notebook-controller v0.0.0-20201201202452-3c3552c7d7bc
	github.com/onsi/ginkgo v1.14.2
	github.com/onsi/gomega v1.10.3
	k8s.io/api v0.0.0-20190409021203-6e4e0e4f393b
	k8s.io/apiextensions-apiserver v0.0.0-20190409022649-727a075fdec8
	k8s.io/apimachinery v0.15.7
	k8s.io/client-go v11.0.1-0.20190409021438-1a26190bd76a+incompatible
	sigs.k8s.io/controller-runtime v0.2.0
)
