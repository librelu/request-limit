package redis_test

import (
	"context"
	"time"

	"github.com/alicebob/miniredis"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/request-limit/clients/redis"
	redisclient "github.com/request-limit/clients/redis"
)

var _ = Describe("redis", func() {
	var (
		mr           *miniredis.Miniredis
		redisHandler redis.Handler
	)
	BeforeEach(func() {
		var e error
		mr, e = miniredis.Run()
		if e != nil {
			panic(e)
		}
		redisHandler, e = redisclient.NewClient(
			mr.Addr(), "", 0, 0, 3000, 3000, 3000)
		if e != nil {
			panic(e)
		}
	})
	AfterEach(func() {
		mr.FlushAll()
		mr.Close()
	})
	Describe("Get", func() {
		var (
			getResult string
			err       error
			key       string
		)
		BeforeEach(func() {
			mr.Set("hello", "world")
			key = "hello"
		})
		JustBeforeEach(func() {
			getResult, err = redisHandler.Get(context.Background(), key)
		})
		When("given existed key", func() {
			It("should get correct value", func() {
				Expect(err).Should(BeNil())
				Expect(getResult).Should(Equal("world"))
			})
		})
		When("given not existed key", func() {
			BeforeEach(func() {
				key = "not found key"
			})
			It("should returns NotFoundError", func() {
				Expect(err).Should(Equal(redisclient.NotFoundError))
				Expect(getResult).Should(Equal(""))
			})
		})
	})
	Describe("INCRAndExpire", func() {
		var (
			durationTime time.Duration
			key          string
			err          error
		)
		BeforeEach(func() {
			durationTime = 1
			key = "hello"
		})
		JustBeforeEach(func() {
			err = redisHandler.INCRAndExpire(context.Background(), key, time.Second*durationTime)
		})
		When("given key with expired time", func() {
			It("should returns correct value", func() {
				Expect(err).Should(BeNil())
				Expect(mr.Get(key)).Should(Equal("1"))
			})
			It("should expired after sleep", func() {
				mr.FastForward(time.Second * 1)
				result, err := mr.Get(key)
				Expect(err.Error()).Should(Equal("ERR no such key"))
				Expect(result).Should(BeZero())
			})
		})
	})
	Describe("INCR", func() {
		var (
			key string
			err error
		)
		BeforeEach(func() {
			key = "hello"
		})
		JustBeforeEach(func() {
			err = redisHandler.INCR(context.Background(), key)
		})
		When("given key existed", func() {
			It("should returns increpted value", func() {
				Expect(err).Should(BeNil())
				Expect(mr.Get(key)).Should(Equal("1"))
			})
		})
	})
})
