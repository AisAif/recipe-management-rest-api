package middleware

import (
	"errors"
	"net/http"
	"reflect"

	"github.com/AisAif/recipe-management-rest-api/src/http/resources"
	"github.com/AisAif/recipe-management-rest-api/src/utils"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/iancoleman/strcase"
	"github.com/spf13/viper"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func GlobalErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		err := c.Errors.Last()
		if err != nil {
			if reflect.TypeOf(err.Err) == reflect.TypeOf(validator.ValidationErrors{}) {
				var ve validator.ValidationErrors
				if errors.As(err, &ve) {
					out := make([]utils.ValidationError, len(ve))
					for i, fe := range ve {
						out[i] = utils.ValidationError{
							Field:   strcase.ToSnake(fe.Field()),
							Message: utils.MessageForTag(fe.Tag(), fe.Param()),
						}
					}
					c.JSON(http.StatusBadRequest, resources.Resource[any]{
						Message: "Bad Request",
						Errors:  out,
					})
				}
			} else if errors.Is(err, gorm.ErrRecordNotFound) {
				c.JSON(http.StatusNotFound, resources.Resource[any]{
					Message: "Bad Request",
					Errors:  "NOT_FOUND",
				})
			} else if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
				c.JSON(http.StatusBadRequest, resources.Resource[any]{
					Message: "Bad Request",
					Errors: map[string]string{
						"username": "INVALID_CREDENTIALS",
					},
				})
			} else {
				if viper.GetString("GIN_MODE") == "debug" {
					c.JSON(http.StatusInternalServerError, resources.Resource[any]{
						Message: "Internal server error",
						Errors:  err.Error(),
					})
				} else {
					c.JSON(http.StatusInternalServerError, resources.Resource[any]{
						Message: "Internal server error",
					})
				}
			}

			c.Abort()
		}
	}
}
