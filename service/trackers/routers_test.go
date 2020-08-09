package trackers_test

import (
	"errors"
	"net/http"
	"net/http/httptest"

	"github.com/golang/mock/gomock"
	"github.com/request-limit/clients/redis"
	redismock "github.com/request-limit/clients/redis/mocks"

	"github.com/gin-gonic/gin"
	"github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/request-limit/service/trackers"
)

var _ = Describe("routes", func() {
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
		// prepare tracker handler
		var e error
		trackerHandler, e = trackers.NewTrackers(
			mockRedis, 10, requestLimit,
		)
		if e != nil {
			panic(e)
		}
		trackerGroup := engine.Group("")
		trackerGroup.Use(trackerHandler.RateLimitMiddleware)
		trackers.TrackersRegister(trackerGroup, trackerHandler)
	})
	AfterEach(func() {
		ctrl.Finish()
	})
	Describe("TrackWithTrackMiddleware", func() {
		var (
			path     string
			body     string
			method   string
			mockIP   string
			recorder *httptest.ResponseRecorder
		)
		BeforeEach(func() {
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
				mockRedis.EXPECT().Get(gomock.Any(), "ip:127.0.0.1").Return("59", nil)
				mockRedis.EXPECT().INCR(gomock.Any(), "ip:127.0.0.1").Return(int64(2), nil)
				mockRedis.EXPECT().Get(gomock.Any(), "ip:127.0.0.1").Return("60", nil)
			})
			It("should returns 200 without error", func() {
				Expect(string(body)).Should(Equal("{\"tries\":\"60\"}"))
				Expect(recorder.Code).Should(Equal(http.StatusOK))
			})
		})
		When("given a request over request limit", func() {
			BeforeEach(func() {
				mockRedis.EXPECT().Get(gomock.Any(), "ip:127.0.0.1").Return("60", nil)
				mockRedis.EXPECT().INCR(gomock.Any(), "ip:127.0.0.1").Return(int64(61), nil)
			})
			It("should returns 403", func() {
				Expect(recorder.Code).Should(Equal(http.StatusForbidden))
			})
		})
		When("given a blank IP request", func() {
			BeforeEach(func() {
				mockIP = ""
			})
			It("should returns 400", func() {
				Expect(recorder.Code).Should(Equal(http.StatusBadRequest))
			})
		})
		When("given request GET cmd got redis server error", func() {
			BeforeEach(func() {
				mockRedis.EXPECT().Get(gomock.Any(), "ip:127.0.0.1").Return("", errors.New("errors occur"))
			})
			It("should returns 500", func() {
				Expect(recorder.Code).Should(Equal(http.StatusInternalServerError))
			})
		})
		When("given request INCR cmd got redis server error", func() {
			BeforeEach(func() {
				mockRedis.EXPECT().Get(gomock.Any(), "ip:127.0.0.1").Return("2", nil)
				mockRedis.EXPECT().INCR(gomock.Any(), "ip:127.0.0.1").Return(int64(0), errors.New("errors occur"))
			})
			It("should returns 403", func() {
				Expect(recorder.Code).Should(Equal(http.StatusInternalServerError))
			})
		})
		When("given request INCRAndExpire cmd got redis server error", func() {
			BeforeEach(func() {
				mockRedis.EXPECT().Get(gomock.Any(), "ip:127.0.0.1").Return("2", redis.NotFoundError)
				mockRedis.EXPECT().INCRAndExpire(gomock.Any(), "ip:127.0.0.1", gomock.Any()).Return(int64(0), errors.New("errors occur"))
			})
			It("should returns 403", func() {
				Expect(recorder.Code).Should(Equal(http.StatusInternalServerError))
			})
		})
	})
})
