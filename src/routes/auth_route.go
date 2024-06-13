package routes

import (
	"github.com/AisAif/recipe-management-rest-api/src/http/controllers"
	"github.com/AisAif/recipe-management-rest-api/src/models"
	"github.com/AisAif/recipe-management-rest-api/src/services"
	"github.com/AisAif/recipe-management-rest-api/src/utils"
	"github.com/gin-gonic/gin"
)

func Auth(r *gin.RouterGroup) {
	authService := services.NewAuthService(models.DB, utils.InitValidator())
	authController := controllers.NewAuthController(authService)

	r.POST("register", authController.Register)
}
