package healthchecks_test

import (
	"net/http"
	"net/http/httptest"

	"github.com/gin-gonic/gin"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/request-limit/service/healthchecks"
)

var _ = Describe("routes", func() {
	var (
		engine *gin.Engine
	)
	BeforeEach(func() {
		gin.SetMode(gin.TestMode)
		engine = gin.New()
		engine.Use(gin.Logger())
		engine.Use(gin.Recovery())
	})
	Describe("HealthCheck", func() {
		var (
			path     string
			body     string
			method   string
			recorder *httptest.ResponseRecorder
		)
		BeforeEach(func() {
			healthchecks.HealthChecksRegister(engine.Group(""))
		})
		JustBeforeEach(func() {
			recorder = httptest.NewRecorder()
			req, err := http.NewRequest(method, path, nil)
			if err != nil {
				panic(err)
			}
			engine.ServeHTTP(recorder, req)
			body = recorder.Body.String()
		})
		When("given a correct path and method", func() {
			BeforeEach(func() {
				method = http.MethodGet
				path = "/healthcheck"
			})
			It("should returns 200 without error", func() {
				Expect(recorder.Code).Should(Equal(http.StatusOK))
				Expect(body).Should(Equal("null"))
			})
		})
	})
})
