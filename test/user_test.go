package test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"

	"github.com/AisAif/recipe-management-rest-api/src/models"
	"github.com/AisAif/recipe-management-rest-api/src/routes"
	"github.com/AisAif/recipe-management-rest-api/src/utils"
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
			userToken := GetUserToken(routerForUser)
			req, _ = http.NewRequest("GET", "/users/current", nil)
			req.Header.Set("Content-Type", "application/json")
			req.Header.Set("Authorization", userToken)
			routerForUser.ServeHTTP(w, req)

			Expect(w.Body.String()).To(ContainSubstring(`Success`))
			Expect(w.Code).To(Equal(http.StatusOK))
		})

		It("should return 401 when no token", func() {

			req, _ = http.NewRequest("GET", "/users/current", nil)
			req.Header.Set("Content-Type", "application/json")
			routerForUser.ServeHTTP(w, req)

			Expect(w.Code).To(Equal(http.StatusUnauthorized))
			Expect(w.Body.String()).To(ContainSubstring(`No Authorization header provided`))
		})

		It("should return 401 when invalid token", func() {

			req, _ = http.NewRequest("GET", "/users/current", nil)
			req.Header.Set("Content-Type", "application/json")
			req.Header.Set("Authorization", "tokensalah")
			routerForUser.ServeHTTP(w, req)

			Expect(w.Code).To(Equal(http.StatusUnauthorized))
			Expect(w.Body.String()).To(ContainSubstring(`Invalid token`))
		})
	})

	Context("Update Current User", func() {
		BeforeEach(func() {
			RemoveAllData()
			CreateUser()
		})

		It("should return 200", func() {
			updateData := map[string]interface{}{
				"name":     "test1",
				"password": "testtest1",
			}
			userJson, _ := json.Marshal(updateData)

			userToken := GetUserToken(routerForUser)
			req, _ = http.NewRequest("PATCH", "/users/current", strings.NewReader(string(userJson)))
			req.Header.Set("Content-Type", "application/json")
			req.Header.Set("Authorization", userToken)
			routerForUser.ServeHTTP(w, req)

			Expect(w.Code).To(Equal(http.StatusOK))
			Expect(w.Body.String()).To(ContainSubstring(`Updated successfully`))

			var user models.User
			models.DB.Find(&user, "username = ?", "test")

			Expect(user.Name).To(Equal("test1"))
			err := utils.CheckPasswordHash("testtest1", user.Password)
			Expect(err).To(BeNil())
		})

		It("should return 401 when validation error", func() {
			updateData := map[string]interface{}{
				"name":     "te",
				"password": "tes",
			}
			userJson, _ := json.Marshal(updateData)

			userToken := GetUserToken(routerForUser)
			req, _ = http.NewRequest("PATCH", "/users/current", strings.NewReader(string(userJson)))
			req.Header.Set("Content-Type", "application/json")
			req.Header.Set("Authorization", userToken)
			routerForUser.ServeHTTP(w, req)

			Expect(w.Code).To(Equal(http.StatusBadRequest))
			Expect(w.Body.String()).To(ContainSubstring(`MIN:3`))
			Expect(w.Body.String()).To(ContainSubstring(`MIN:8`))
		})
	})
})
