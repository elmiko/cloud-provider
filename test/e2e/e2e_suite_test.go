package e2e_suite_test

import (
	"testing"

	"k8s.io/kubernetes/test/e2e/cloud/external"
)

type EmptyCloudProvider struct{}

func TestCloudProvider(t *testing.T) {
	p := EmptyCloudProvider{}

	external.InitializeAndRunExternalCloudProviderTests(t, p)
}
