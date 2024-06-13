package test

import "github.com/AisAif/recipe-management-rest-api/src/models"

func RemoveAllData() {
	models.DB.Exec("DELETE FROM users")
}
