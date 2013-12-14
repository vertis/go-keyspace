package keyspace_test

import (
	"github.com/vertis/go-keyspace"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Keyspace", func() {
  It("Calling next on a new keyspace, should yield the first entry", func() {
     k := keyspace.New([]byte("012345679"), 2)
     val, _ := k.Next()
     Expect(val).To(Equal([]byte("00")), "Should be the first entry")
  })
})
