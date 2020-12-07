package tests

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"net"
)

var _ = Describe("Checking for Network connectivity ", func() {
	Context("when running within the DMZ", func() {
		It("Should be able to connect to the DNS server 8.8.8.8", func() {
			Eventually(func() bool {
				_, err := net.Dial("udp", "8.8.8.8:53")
				return err == nil
			}, timeout, interval).Should(BeTrue())
		})
		It("NETWORK Should be able to resolve DNS lookup ", func() {
			Eventually(func() bool {
				_, err := net.LookupIP("google.com")
				return err == nil
			}, timeout, interval).Should(BeTrue())
		})
	})
})
