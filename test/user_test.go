package test

import (
	"net/http"
	"net/http/httptest"

	"github.com/AisAif/recipe-management-rest-api/src/routes"
	"github.com/gin-gonic/gin"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var routerForUser *gin.Engine = routes.InitRouter()

var _ = Describe("Auth", func() {
	var w *httptest.ResponseRecorder
	var req *http.Request

	BeforeEach(func() {

		w = httptest.NewRecorder()
	})

	Context("Get Current User", func() {

		BeforeEach(func() {
			RemoveAllData()
			CreateUser()
		})

		It("should return 200", func() {
			userToken := GetUserToken(routerForUser, w)
			req, _ = http.NewRequest("GET", "/users/current", nil)
			req.Header.Set("Content-Type", "application/json")
			req.Header.Set("Authorization", userToken)
			routerForUser.ServeHTTP(w, req)

			Expect(w.Code).To(Equal(http.StatusOK))
			Expect(w.Body.String()).To(ContainSubstring(`Success`))
		})

		It("should return 401", func() {

			req, _ = http.NewRequest("GET", "/users/current", nil)
			req.Header.Set("Content-Type", "application/json")
			routerForUser.ServeHTTP(w, req)

			Expect(w.Code).To(Equal(http.StatusUnauthorized))
			Expect(w.Body.String()).To(ContainSubstring(`No Authorization header provided`))
		})

		It("should return 401", func() {

			req, _ = http.NewRequest("GET", "/users/current", nil)
			req.Header.Set("Content-Type", "application/json")
			req.Header.Set("Authorization", "tokensalah")
			routerForUser.ServeHTTP(w, req)

			Expect(w.Code).To(Equal(http.StatusUnauthorized))
			Expect(w.Body.String()).To(ContainSubstring(`Invalid token`))
		})
	})
})
