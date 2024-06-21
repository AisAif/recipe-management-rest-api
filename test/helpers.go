package test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"

	"github.com/AisAif/recipe-management-rest-api/src/http/resources"
	"github.com/AisAif/recipe-management-rest-api/src/models"
	"github.com/AisAif/recipe-management-rest-api/src/utils"
	"github.com/gin-gonic/gin"
)

func RemoveAllData() {
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
