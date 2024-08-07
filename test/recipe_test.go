package test

import (
	"bytes"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"net/textproto"
	"os"
	"strconv"

	"github.com/AisAif/recipe-management-rest-api/src/models"
	"github.com/AisAif/recipe-management-rest-api/src/routes"
	"github.com/gin-gonic/gin"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var routerForRecipe *gin.Engine = routes.InitRouter()

var _ = Describe("Recipe", func() {
	var w *httptest.ResponseRecorder
	var req *http.Request

	BeforeEach(func() {

		w = httptest.NewRecorder()
	})

	Context("Create Recipe", func() {
		BeforeEach(func() {
			RemoveAllData()
			CreateUser()
		})

		It("should return 400 when validation error", func() {
			// multipart form
			body := &bytes.Buffer{}
			writer := multipart.NewWriter(body)
			_ = writer.WriteField("title", "")
			_ = writer.WriteField("content", "")

			file, _ := os.Open("./assets/test_asset.pptx")
			defer file.Close()

			part, _ := writer.CreateFormFile("image", "test_asset.pptx")

			_, _ = io.Copy(part, file)
			_ = writer.Close()

			userToken := GetUserToken(routerForRecipe)
			req, _ = http.NewRequest("POST", "/recipes", body)
			req.Header.Set("Content-Type", writer.FormDataContentType())
			req.Header.Set("Authorization", userToken)
			routerForRecipe.ServeHTTP(w, req)

			Expect(w.Code).To(Equal(http.StatusBadRequest))
			Expect(w.Body.String()).To(ContainSubstring(`REQUIRED`))
			Expect(w.Body.String()).To(ContainSubstring(`REQUIRED`))
			Expect(w.Body.String()).To(ContainSubstring(`FILE_TYPE:image/*`))
		})

		It("should return 200", func() {
			// multipart form
			body := &bytes.Buffer{}
			writer := multipart.NewWriter(body)
			_ = writer.WriteField("title", "test")
			_ = writer.WriteField("content", "test")

			file, err := os.Open("./assets/test_asset.jpg")
			Expect(err).To(BeNil())
			defer file.Close()

			partHeaders := textproto.MIMEHeader{}
			partHeaders.Set("Content-Disposition", `form-data; name="image"; filename="test_asset.jpg"`)
			partHeaders.Set("Content-Type", "image/jpeg")
			part, err := writer.CreatePart(partHeaders)
			Expect(err).To(BeNil())

			_, err = io.Copy(part, file)
			Expect(err).To(BeNil())

			_ = writer.Close()

			userToken := GetUserToken(routerForRecipe)
			req, _ = http.NewRequest("POST", "/recipes", body)
			req.Header.Set("Content-Type", writer.FormDataContentType())
			req.Header.Set("Authorization", userToken)
			routerForRecipe.ServeHTTP(w, req)

			Expect(w.Code).To(Equal(http.StatusCreated))
		})
	})

	Context("Update Recipe", func() {
		BeforeEach(func() {
			RemoveAllData()
			CreateUser()
		})

		It("should return 400 when validation error", func() {
			// get recipe
			recipe := GetRecipe(routerForRecipe)

			// multipart form
			body := &bytes.Buffer{}
			writer := multipart.NewWriter(body)
			_ = writer.WriteField("title", "aa")

			file, _ := os.Open("./assets/test_asset.pptx")
			defer file.Close()

			part, _ := writer.CreateFormFile("image", "test_asset.pptx")

			_, _ = io.Copy(part, file)
			_ = writer.Close()

			userToken := GetUserToken(routerForRecipe)
			req, _ = http.NewRequest("PATCH", "/recipes/"+strconv.FormatUint(recipe.ID, 10), body)
			req.Header.Set("Content-Type", writer.FormDataContentType())
			req.Header.Set("Authorization", userToken)
			routerForRecipe.ServeHTTP(w, req)

			Expect(w.Code).To(Equal(http.StatusBadRequest))
			Expect(w.Body.String()).To(ContainSubstring(`MIN:3`))
			Expect(w.Body.String()).To(ContainSubstring(`FILE_TYPE:image/*`))
		})

		It("should return 404 when not found", func() {
			// get recipe
			recipe := GetRecipe(routerForRecipe)

			// multipart form
			body := &bytes.Buffer{}
			writer := multipart.NewWriter(body)
			_ = writer.WriteField("title", "test")
			_ = writer.WriteField("content", "test")

			file, err := os.Open("./assets/test_asset.jpg")
			Expect(err).To(BeNil())
			defer file.Close()

			partHeaders := textproto.MIMEHeader{}
			partHeaders.Set("Content-Disposition", `form-data; name="image"; filename="test_asset.jpg"`)
			partHeaders.Set("Content-Type", "image/jpeg")
			part, err := writer.CreatePart(partHeaders)
			Expect(err).To(BeNil())

			_, err = io.Copy(part, file)
			Expect(err).To(BeNil())

			_ = writer.Close()

			userToken := GetUserToken(routerForRecipe)
			req, _ = http.NewRequest("PATCH", "/recipes/"+strconv.FormatUint(recipe.ID+1, 10), body)
			req.Header.Set("Content-Type", writer.FormDataContentType())
			req.Header.Set("Authorization", userToken)
			routerForRecipe.ServeHTTP(w, req)

			Expect(w.Code).To(Equal(http.StatusNotFound))
		})

		It("should return 200", func() {
			// get recipe
			recipe := GetRecipe(routerForRecipe)

			// multipart form
			body := &bytes.Buffer{}
			writer := multipart.NewWriter(body)
			_ = writer.WriteField("title", "test")
			_ = writer.WriteField("content", "test")

			file, err := os.Open("./assets/test_asset.jpg")
			Expect(err).To(BeNil())
			defer file.Close()

			partHeaders := textproto.MIMEHeader{}
			partHeaders.Set("Content-Disposition", `form-data; name="image"; filename="test_asset.jpg"`)
			partHeaders.Set("Content-Type", "image/jpeg")
			part, err := writer.CreatePart(partHeaders)
			Expect(err).To(BeNil())

			_, err = io.Copy(part, file)
			Expect(err).To(BeNil())

			_ = writer.Close()

			userToken := GetUserToken(routerForRecipe)
			req, _ = http.NewRequest("PATCH", "/recipes/"+strconv.FormatUint(recipe.ID, 10), body)
			req.Header.Set("Content-Type", writer.FormDataContentType())
			req.Header.Set("Authorization", userToken)
			routerForRecipe.ServeHTTP(w, req)

			Expect(w.Code).To(Equal(http.StatusOK))
		})
	})

	Context("Delete Recipe", func() {
		BeforeEach(func() {
			RemoveAllData()
			CreateUser()
		})

		It("should return 404", func() {
			// get recipe
			recipe := GetRecipe(routerForRecipe)

			userToken := GetUserToken(routerForRecipe)
			req, _ = http.NewRequest("DELETE", "/recipes/"+strconv.FormatUint(recipe.ID+1, 10), nil)
			req.Header.Set("Authorization", userToken)
			routerForRecipe.ServeHTTP(w, req)

			Expect(w.Code).To(Equal(http.StatusNotFound))
			Expect(w.Body.String()).To(ContainSubstring(`NOT_FOUND`))
		})

		It("should return 200", func() {
			// get recipe
			recipe := GetRecipe(routerForRecipe)

			userToken := GetUserToken(routerForRecipe)
			req, _ = http.NewRequest("DELETE", "/recipes/"+strconv.FormatUint(recipe.ID, 10), nil)
			req.Header.Set("Authorization", userToken)
			routerForRecipe.ServeHTTP(w, req)

			Expect(w.Code).To(Equal(http.StatusOK))
			Expect(w.Body.String()).To(ContainSubstring(`Successfully deleted`))
		})
	})

	Context("Get Current User's Recipes", func() {
		BeforeEach(func() {
			RemoveAllData()
			CreateUser()
		})

		It("should return 200", func() {
			// make recipes
			GetRecipe(routerForRecipe)
			GetRecipe(routerForRecipe)
			GetRecipe(routerForRecipe)

			userToken := GetUserToken(routerForRecipe)
			req, _ = http.NewRequest("GET", "/recipes/current", nil)
			req.Header.Set("Authorization", userToken)
			routerForRecipe.ServeHTTP(w, req)

			Expect(w.Code).To(Equal(http.StatusOK))
			Expect(w.Body.String()).To(ContainSubstring(`Successfully fetched`))
			Expect(w.Body.String()).To(ContainSubstring(`test`))
		})
	})

	Context("Get All Published Recipes", func() {
		BeforeEach(func() {
			RemoveAllData()
			CreateUser()
		})

		It("should return 200", func() {
			// make recipes
			recipe := GetRecipe(routerForRecipe)
			recipe.IsPublic = true

			models.DB.Save(&recipe)

			userToken := GetUserToken(routerForRecipe)
			req, _ = http.NewRequest("GET", "/recipes", nil)
			req.Header.Set("Authorization", userToken)
			routerForRecipe.ServeHTTP(w, req)

			Expect(w.Code).To(Equal(http.StatusOK))
			Expect(w.Body.String()).To(ContainSubstring(`Successfully fetched`))
			Expect(w.Body.String()).To(ContainSubstring(`test`))
			Expect(w.Body.String()).To(ContainSubstring(`"is_public":true`))
		})
	})
})
