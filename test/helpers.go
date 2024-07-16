package test

import (
	"bytes"
	"encoding/json"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"net/textproto"
	"os"
	"strings"

	"github.com/AisAif/recipe-management-rest-api/src/http/resources"
	"github.com/AisAif/recipe-management-rest-api/src/models"
	"github.com/AisAif/recipe-management-rest-api/src/utils"
	"github.com/gin-gonic/gin"
)

func RemoveAllData() {
	models.DB.Exec("DELETE FROM recipes")
	models.DB.Exec("DELETE FROM users")
}

func CreateUser() {
	hash, _ := utils.HashPassword("testtest")
	models.DB.Create(&models.User{
		Username: "test",
		Name:     "test",
		Password: hash,
	})
}

func GetUserToken(router *gin.Engine) string {
	w := httptest.NewRecorder()
	exampleUser := models.User{
		Username: "test",
		Password: "testtest",
	}
	userJson, _ := json.Marshal(exampleUser)

	req, _ := http.NewRequest("POST", "/auth/login", strings.NewReader(string(userJson)))
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	var body resources.Resource[any]
	json.Unmarshal(w.Body.Bytes(), &body)

	return body.Data.(map[string]interface{})["token"].(string)
}

func GetRecipe(routerForRecipe *gin.Engine) models.Recipe {

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	_ = writer.WriteField("title", "test")
	_ = writer.WriteField("content", "test")

	file, _ := os.Open("./assets/test_asset.jpg")
	defer file.Close()

	partHeaders := textproto.MIMEHeader{}
	partHeaders.Set("Content-Disposition", `form-data; name="image"; filename="test_asset.jpg"`)
	partHeaders.Set("Content-Type", "image/jpeg")
	part, _ := writer.CreatePart(partHeaders)

	_, _ = io.Copy(part, file)

	_ = writer.Close()

	userToken := GetUserToken(routerForRecipe)
	req, _ := http.NewRequest("POST", "/recipes", body)
	req.Header.Set("Content-Type", writer.FormDataContentType())
	req.Header.Set("Authorization", userToken)

	w := httptest.NewRecorder()
	routerForRecipe.ServeHTTP(w, req)

	var recipe = models.Recipe{}
	models.DB.Last(&recipe)

	return recipe
}
