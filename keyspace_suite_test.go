package keyspace_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestKeyspace(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Keyspace Suite")
}
