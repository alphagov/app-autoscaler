package custom_metrics_cred_helper_plugin_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestCustomMetricsCredHelper(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Credentials Helper Suite")
}
