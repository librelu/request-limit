package redis_test

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	redisclient "github.com/request-limit/clients/redis"
)

var _ = Describe("redis", func() {
	Describe("NewClient", func() {
		var (
			address, password                      string
			db, maxRetries                         int
			readTimeout, writeTimeout, dialTimeout time.Duration
			redisClient                            *redis.Client
			err                                    error
			ctx                                    context.Context
		)
		BeforeEach(func() {
			ctx = context.Background()
			address = "127.0.0.1:6379"
			password = ""
			db = 0
			maxRetries = 0
			readTimeout = 3000
			writeTimeout = 3000
			dialTimeout = 3000
		})
		JustBeforeEach(func() {
			client := redis.NewClient(&redis.Options{
				Addr: address,
			})
			if e := client.Ping(ctx).Err(); e != nil {
				panic(e)
			}

			redisClient, err = redisclient.NewClient(address, password, db, maxRetries, readTimeout, writeTimeout, dialTimeout)
		})
		When("given correct input", func() {
			It("should pass the flow without error", func() {
				Expect(err).Should(BeNil())
				Expect(redisClient).ShouldNot(BeNil())
			})
		})
		When("given blank address", func() {
			BeforeEach(func() {
				address = ""
			})
			It("should returns error", func() {
				Expect(err).ShouldNot(BeNil())
				Expect(redisClient).Should(BeNil())
			})
		})
		When("given negative db", func() {
			BeforeEach(func() {
				db = -100
			})
			It("should returns error", func() {
				Expect(err).ShouldNot(BeNil())
				Expect(redisClient).Should(BeNil())
			})
		})
		When("given negative maxRetries", func() {
			BeforeEach(func() {
				maxRetries = -100
			})
			It("should returns error", func() {
				Expect(err).ShouldNot(BeNil())
				Expect(redisClient).Should(BeNil())
			})
		})
		When("given negative readTimeout", func() {
			BeforeEach(func() {
				readTimeout = -100
			})
			It("should returns error", func() {
				Expect(err).ShouldNot(BeNil())
				Expect(redisClient).Should(BeNil())
			})
		})
		When("given negative writeTimeout", func() {
			BeforeEach(func() {
				writeTimeout = -100
			})
			It("should returns error", func() {
				Expect(err).ShouldNot(BeNil())
				Expect(redisClient).Should(BeNil())
			})
		})
		When("given negative dialTimeout", func() {
			BeforeEach(func() {
				dialTimeout = -100
			})
			It("should returns error", func() {
				Expect(err).ShouldNot(BeNil())
				Expect(redisClient).Should(BeNil())
			})
		})
	})
})
