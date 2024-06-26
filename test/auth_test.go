package test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"

	"github.com/AisAif/recipe-management-rest-api/src/models"
	"github.com/AisAif/recipe-management-rest-api/src/routes"
	"github.com/gin-gonic/gin"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/rs/zerolog/log"
)

var routerForAuth *gin.Engine = routes.InitRouter()

var _ = Describe("Auth", func() {
	var w *httptest.ResponseRecorder
	var req *http.Request

	BeforeEach(func() {
		w = httptest.NewRecorder()
	})

	Context("Register", func() {

		BeforeEach(func() {
			RemoveAllData()
		})

		It("should return 201", func() {
			exampleUser := models.User{
				Username: "test",
				Name:     "test",
				Password: "testtest",
			}
			userJson, _ := json.Marshal(exampleUser)

			req, _ = http.NewRequest("POST", "/auth/register", strings.NewReader(string(userJson)))
			req.Header.Set("Content-Type", "application/json")
			routerForAuth.ServeHTTP(w, req)

			log.Info().Msg(w.Body.String())
			Expect(w.Code).To(Equal(http.StatusCreated))
			Expect(w.Body.String()).To(ContainSubstring(`registered`))
		})

		It("should return 400: all fields are required", func() {
			exampleUser := models.User{
				Username: "",
				Name:     "",
				Password: "",
			}
			userJson, _ := json.Marshal(exampleUser)

			req, _ = http.NewRequest("POST", "/auth/register", strings.NewReader(string(userJson)))
			req.Header.Set("Content-Type", "application/json")
			routerForAuth.ServeHTTP(w, req)

			log.Info().Msg(w.Body.String())
			Expect(w.Code).To(Equal(http.StatusBadRequest))
			Expect(w.Body.String()).To(ContainSubstring(`username`))
			Expect(w.Body.String()).To(ContainSubstring(`password`))
			Expect(w.Body.String()).To(ContainSubstring(`name`))
		})

		It("should return 400: validation error for username and password ", func() {
			exampleUser := models.User{
				Username: "te",
				Name:     "test",
				Password: "test",
			}
			userJson, _ := json.Marshal(exampleUser)

			req, _ = http.NewRequest("POST", "/auth/register", strings.NewReader(string(userJson)))
			req.Header.Set("Content-Type", "application/json")
			routerForAuth.ServeHTTP(w, req)

			log.Info().Msg(w.Body.String())
			Expect(w.Code).To(Equal(http.StatusBadRequest))
			Expect(w.Body.String()).To(ContainSubstring(`username`))
			Expect(w.Body.String()).To(ContainSubstring(`password`))
		})

		It("should return 400: username is exist", func() {
			CreateUser()
			exampleUser := models.User{
				Username: "test",
				Name:     "test",
				Password: "testtest",
			}
			userJson, _ := json.Marshal(exampleUser)

			req, _ = http.NewRequest("POST", "/auth/register", strings.NewReader(string(userJson)))
			req.Header.Set("Content-Type", "application/json")
			routerForAuth.ServeHTTP(w, req)

			log.Info().Msg(w.Body.String())
			Expect(w.Code).To(Equal(http.StatusBadRequest))
			Expect(w.Body.String()).To(ContainSubstring(`EXISTS`))
		})
	})

	Context("Login", func() {

		BeforeEach(func() {
			RemoveAllData()
			CreateUser()
		})

		It("should return 200", func() {
			exampleUser := models.User{
				Username: "test",
				Password: "testtest",
			}
			userJson, _ := json.Marshal(exampleUser)

			req, _ = http.NewRequest("POST", "/auth/login", strings.NewReader(string(userJson)))
			req.Header.Set("Content-Type", "application/json")
			routerForAuth.ServeHTTP(w, req)

			log.Info().Msg(w.Body.String())
			Expect(w.Code).To(Equal(http.StatusOK))
			Expect(w.Body.String()).To(ContainSubstring(`token`))
		})

		It("should return 400: all fields are required", func() {
			exampleUser := models.User{
				Username: "",
				Password: "",
			}
			userJson, _ := json.Marshal(exampleUser)

			req, _ = http.NewRequest("POST", "/auth/login", strings.NewReader(string(userJson)))
			req.Header.Set("Content-Type", "application/json")
			routerForAuth.ServeHTTP(w, req)

			log.Info().Msg(w.Body.String())
			Expect(w.Code).To(Equal(http.StatusBadRequest))
			Expect(w.Body.String()).To(ContainSubstring(`username`))
			Expect(w.Body.String()).To(ContainSubstring(`password`))
		})

		It("should return 404: user not found ", func() {
			exampleUser := models.User{
				Username: "testt",
				Password: "testtest",
			}
			userJson, _ := json.Marshal(exampleUser)

			req, _ = http.NewRequest("POST", "/auth/login", strings.NewReader(string(userJson)))
			req.Header.Set("Content-Type", "application/json")
			routerForAuth.ServeHTTP(w, req)

			log.Info().Msg(w.Body.String())
			Expect(w.Code).To(Equal(http.StatusNotFound))
			Expect(w.Body.String()).To(ContainSubstring(`NOT_FOUND`))
		})

		It("should return 400: password is wrong", func() {
			exampleUser := models.User{
				Username: "test",
				Password: "testtests",
			}
			userJson, _ := json.Marshal(exampleUser)

			req, _ = http.NewRequest("POST", "/auth/login", strings.NewReader(string(userJson)))
			req.Header.Set("Content-Type", "application/json")
			routerForAuth.ServeHTTP(w, req)

			log.Info().Msg(w.Body.String())
			Expect(w.Code).To(Equal(http.StatusBadRequest))
			Expect(w.Body.String()).To(ContainSubstring(`INVALID_CREDENTIALS`))
		})
	})
})
