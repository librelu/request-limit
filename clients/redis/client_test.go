package redis_test

import (
	"time"

	"github.com/alicebob/miniredis"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	redisclient "github.com/request-limit/clients/redis"
)

var _ = Describe("redis", func() {
	var (
		mr *miniredis.Miniredis
	)
	BeforeEach(func() {
		var e error
		mr, e = miniredis.Run()
		if e != nil {
			panic(e)
		}
	})
	AfterEach(func() {
		mr.Close()
	})
	Describe("NewClient", func() {
		var (
			address, password                      string
			db, maxRetries                         int
			readTimeout, writeTimeout, dialTimeout time.Duration
			redisHandler                           redisclient.Handler
			err                                    error
		)
		BeforeEach(func() {
			address = mr.Addr()
			password = ""
			db = 0
			maxRetries = 0
			readTimeout = 3000
			writeTimeout = 3000
			dialTimeout = 3000
		})
		JustBeforeEach(func() {
			redisHandler, err = redisclient.NewClient(address, password, db, maxRetries, readTimeout, writeTimeout, dialTimeout)
		})
		When("given correct input", func() {
			It("should pass the flow without error", func() {
				Expect(err).Should(BeNil())
				Expect(redisHandler).ShouldNot(BeNil())
			})
		})
		When("given blank address", func() {
			BeforeEach(func() {
				address = ""
			})
			It("should returns error", func() {
				Expect(err).ShouldNot(BeNil())
				Expect(redisHandler).Should(BeNil())
			})
		})
		When("given negative db", func() {
			BeforeEach(func() {
				db = -100
			})
			It("should returns error", func() {
				Expect(err).ShouldNot(BeNil())
				Expect(redisHandler).Should(BeNil())
			})
		})
		When("given negative maxRetries", func() {
			BeforeEach(func() {
				maxRetries = -100
			})
			It("should returns error", func() {
				Expect(err).ShouldNot(BeNil())
				Expect(redisHandler).Should(BeNil())
			})
		})
		When("given negative readTimeout", func() {
			BeforeEach(func() {
				readTimeout = -100
			})
			It("should returns error", func() {
				Expect(err).ShouldNot(BeNil())
				Expect(redisHandler).Should(BeNil())
			})
		})
		When("given negative writeTimeout", func() {
			BeforeEach(func() {
				writeTimeout = -100
			})
			It("should returns error", func() {
				Expect(err).ShouldNot(BeNil())
				Expect(redisHandler).Should(BeNil())
			})
		})
		When("given negative dialTimeout", func() {
			BeforeEach(func() {
				dialTimeout = -100
			})
			It("should returns error", func() {
				Expect(err).ShouldNot(BeNil())
				Expect(redisHandler).Should(BeNil())
			})
		})
	})
})
