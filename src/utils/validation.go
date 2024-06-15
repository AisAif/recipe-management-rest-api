package utils

import (
	"github.com/AisAif/recipe-management-rest-api/src/models"
	"github.com/go-playground/validator/v10"
)

type ValidationError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func InitValidator() *validator.Validate {
	validate := validator.New(validator.WithRequiredStructEnabled())

	validate.RegisterValidation("username_exists", func(fl validator.FieldLevel) bool {
		username := fl.Field().String()

		var user models.User
		models.DB.First(&user, "username = ?", username)

		return user.Username == ""
	})

	return validate
}

func MessageForTag(tag string, param string) string {
	switch tag {
	case "required":
		return "REQUIRED"
	case "min":
		return "MIN:" + param
	case "max":
		return "MAX:" + param
	case "username_exists":
		return "USERNAME_EXISTS"
	default:
		return "UNKNOWN"
	}
}
