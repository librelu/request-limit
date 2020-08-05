package trackers_test

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/request-limit/clients/redis"
	redismock "github.com/request-limit/clients/redis/mocks"
	"github.com/request-limit/service/trackers"
)

var _ = Describe("handlers", func() {
	var (
		engine         *gin.Engine
		mockRedis      *redismock.MockHandler
		ctrl           *gomock.Controller
		requestLimit   int64
		trackerHandler trackers.Handler
	)
	BeforeEach(func() {
		requestLimit = 10
		// prepare gin server
		gin.SetMode(gin.TestMode)
		engine = gin.New()
		engine.Use(gin.Logger())
		engine.Use(gin.Recovery())
		// prepare redis
		ctrl = gomock.NewController(ginkgo.GinkgoT())
		mockRedis = redismock.NewMockHandler(ctrl)
	})
	AfterEach(func() {
		ctrl.Finish()
	})
	Describe("NewTrackers", func() {
		var (
			redisHandler    redis.Handler
			expiredDuration time.Duration
			requestLimit    int64
			trackerHandler  trackers.Handler
			err             error
		)
		BeforeEach(func() {
			redisHandler = mockRedis
			expiredDuration = 2000
			requestLimit = 10
		})
		JustBeforeEach(func() {
			trackerHandler, err = trackers.NewTrackers(redisHandler, expiredDuration, requestLimit)
		})
		When("given validate input", func() {
			It("should returns Handler without error", func() {
				Expect(err).Should(BeNil())
				Expect(trackerHandler).ShouldNot(BeNil())
			})
		})
		When("given redis is nil", func() {
			BeforeEach(func() {
				redisHandler = nil
			})
			It("should returns error", func() {
				Expect(err).ShouldNot(BeNil())
				Expect(trackerHandler).Should(BeNil())
			})
		})
		When("given negative expired duration", func() {
			BeforeEach(func() {
				expiredDuration = -2000
			})
			It("should returns error", func() {
				Expect(err).ShouldNot(BeNil())
				Expect(trackerHandler).Should(BeNil())
			})
		})
		When("given expired duration zero", func() {
			BeforeEach(func() {
				expiredDuration = 0
			})
			It("should returns error", func() {
				Expect(err).ShouldNot(BeNil())
				Expect(trackerHandler).Should(BeNil())
			})
		})
		When("given negative requestLimit zero", func() {
			BeforeEach(func() {
				requestLimit = -100
			})
			It("should returns error", func() {
				Expect(err).ShouldNot(BeNil())
				Expect(trackerHandler).Should(BeNil())
			})
		})
		When("given requestLimit zero", func() {
			BeforeEach(func() {
				requestLimit = 0
			})
			It("should returns error", func() {
				Expect(err).ShouldNot(BeNil())
				Expect(trackerHandler).Should(BeNil())
			})
		})
	})
	Describe("Track", func() {
		var (
			path     string
			body     string
			method   string
			mockIP   string
			recorder *httptest.ResponseRecorder
		)
		BeforeEach(func() {
			// prepare tracker handler
			var e error
			trackerHandler, e = trackers.NewTrackers(
				mockRedis, 10, requestLimit,
			)
			if e != nil {
				panic(e)
			}
			trackerGroup := engine.Group("")
			trackers.TrackersRegister(trackerGroup, trackerHandler)
			// prepare input
			method = http.MethodGet
			path = "/track"
			mockIP = "127.0.0.1"
		})
		JustBeforeEach(func() {
			recorder = httptest.NewRecorder()
			req, err := http.NewRequest(method, path, nil)
			req.Header.Set("X-Forwarded-For", mockIP)
			if err != nil {
				panic(err)
			}
			engine.ServeHTTP(recorder, req)
			body = recorder.Body.String()
		})
		When("given a correct path and method", func() {
			BeforeEach(func() {
				mockRedis.EXPECT().Get(gomock.Any(), "ip:127.0.0.1").Return("60", nil)
			})
			It("should returns 200 without error", func() {
				Expect(string(body)).Should(Equal("{\"tries\":\"60\"}"))
				Expect(recorder.Code).Should(Equal(http.StatusOK))
			})
		})
		When("given a blank IP request", func() {
			BeforeEach(func() {
				mockIP = ""
			})
			It("should returns 403", func() {
				Expect(recorder.Code).Should(Equal(http.StatusForbidden))
			})
		})
		When("given facing GET cmd redis server error", func() {
			BeforeEach(func() {
				mockRedis.EXPECT().Get(gomock.Any(), "ip:127.0.0.1").Return("", errors.New("errors occur"))
			})
			It("should returns 500", func() {
				Expect(recorder.Code).Should(Equal(http.StatusInternalServerError))
			})
		})
	})
})
