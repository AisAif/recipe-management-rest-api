package middleware

import (
	"errors"

	"github.com/AisAif/recipe-management-rest-api/src/utils"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/iancoleman/strcase"
	"github.com/spf13/viper"
)

func GlobalErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		err := c.Errors.Last()
		if err != nil {
			switch err.Err.(type) {
			case validator.ValidationErrors:
				var ve validator.ValidationErrors
				if errors.As(err, &ve) {
					out := make([]utils.ValidationError, len(ve))
					for i, fe := range ve {
						out[i] = utils.ValidationError{
							Field:   strcase.ToSnake(fe.Field()),
							Message: utils.MessageForTag(fe.Tag(), fe.Param()),
						}
					}
					c.JSON(400, gin.H{"errors": out})
				}

				return
			default:
				if viper.GetString("GIN_MODE") == "debug" {
					c.JSON(500, gin.H{
						"message": "Internal server error.",
						"error":   err,
					})
				} else {
					c.JSON(500, gin.H{
						"message": "Internal server error.",
					})
				}
			}

			c.Abort()
		}
	}
}
