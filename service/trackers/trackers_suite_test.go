package trackers_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestTrackers(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Trackers Suite")
}
