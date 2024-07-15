package utils

import (
	"mime/multipart"
	"strconv"
	"strings"

	"github.com/AisAif/recipe-management-rest-api/src/models"
	"github.com/go-playground/validator/v10"
)

type ValidationError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func InitValidator() *validator.Validate {
	validate := validator.New(validator.WithRequiredStructEnabled())

	// custom validations

	validate.RegisterValidation("username_exists", func(fl validator.FieldLevel) bool {
		username := fl.Field().String()

		var user models.User
		models.DB.First(&user, "username = ?", username)

		return user.Username == ""
	})

	validate.RegisterValidation("file_type", func(fl validator.FieldLevel) bool {
		fileHeader, ok := fl.Field().Interface().(multipart.FileHeader)
		if !ok {
			return false
		}

		fileType := fileHeader.Header.Get("Content-Type")
		allowedTypes := strings.Split(fl.Param(), ":")

		for _, t := range allowedTypes {
			t = strings.TrimSpace(t)
			if strings.HasSuffix(t, "/*") {
				if strings.HasPrefix(fileType, strings.TrimSuffix(t, "/*")) {
					return true
				}
			} else {
				if t == fileType {
					return true
				}
			}
		}

		return false
	})

	validate.RegisterValidation("max_size", func(fl validator.FieldLevel) bool {
		maxFileSize, err := strconv.Atoi(fl.Param())
		if err != nil {
			return false
		}

		fileHeader, ok := fl.Field().Interface().(multipart.FileHeader)
		if !ok {
			return false
		}

		return int(fileHeader.Size) <= maxFileSize*1024
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
	case "file_type":
		return "FILE_TYPE:" + param
	case "max_size":
		return "MAX_SIZE:" + param
	default:
		return "UNKNOWN"
	}
}
