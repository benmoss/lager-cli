package chug_test

import (
	"time"

	. "github.com/benmoss/lager-cli/chug"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("TimeParser", func() {
	Context("when passed an empty string", func() {
		It("should return zero", func() {
			Expect(ParseTimeFlag("")).Should(BeZero())
		})
	})

	Context("when passed a unix timestamp", func() {
		It("should return that time", func() {
			Expect(ParseTimeFlag("1123491.1238")).Should(Equal(time.Unix(1123491, 123800000)))
		})
	})

	Context("when passed a duration", func() {
		It("should return a time relative to the current time", func() {
			duration := -(time.Hour + 15*time.Minute + 3*time.Second)
			Expect(ParseTimeFlag("-1h15m3s")).Should(BeTemporally("~", time.Now().Add(duration), time.Second))
		})
	})

	Context("when passed a chug-formatted timestamp", func() {
		It("should return that time", func() {
			expectedTime, err := time.Parse("02/01/06 15:04:05.99 MST", "10/03/17 07:35:14.43 PST")
			Expect(err).ShouldNot(HaveOccurred())
			Expect(ParseTimeFlag("10/03/17 07:35:14.43 PST")).Should(Equal(expectedTime))
		})
	})

	Context("when passed anything else", func() {
		It("should error", func() {
			t, err := ParseTimeFlag("abc")
			Expect(err).Should(HaveOccurred())
			Expect(t).Should(BeZero())
		})
	})
})
