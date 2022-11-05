package expconv_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestExpconv(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Expconv Suite")
}
