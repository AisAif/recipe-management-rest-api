package middleware

import (
	"net/http"

	"github.com/AisAif/recipe-management-rest-api/src/http/resources"
	"github.com/AisAif/recipe-management-rest-api/src/models"
	"github.com/AisAif/recipe-management-rest-api/src/utils"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")

		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, resources.Resource[any]{
				Message: "No Authorization header provided",
			})
			c.Abort()
			return
		}

		user, err := utils.VerifyJWT(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, resources.Resource[any]{
				Message: "Invalid token",
			})
			c.Abort()
			return
		}

		var userData models.User
		err = models.DB.Where("username = ?", user.Username).First(&userData).Error
		if err != nil {
			c.JSON(http.StatusUnauthorized, resources.Resource[any]{
				Message: "Invalid token",
			})
			c.Abort()
			return
		}

		c.Set("user", resources.UserResource{
			Username: userData.Username,
			Name:     userData.Name,
		})

		c.Next()
	}
}
