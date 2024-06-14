package test

import (
	"github.com/AisAif/recipe-management-rest-api/src/models"
	"github.com/AisAif/recipe-management-rest-api/src/utils"
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
