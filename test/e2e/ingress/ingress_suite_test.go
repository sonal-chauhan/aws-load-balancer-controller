package ingress

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	framework "github.com/sonal-chauhan/aws-load-balancer-controller/test/framework"
)

var tf *framework.Framework

func TestIngress(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Ingress Suite")
}

var _ = BeforeSuite(func() {
	var err error
	tf, err = framework.InitFramework()
	Expect(err).NotTo(HaveOccurred())
})
