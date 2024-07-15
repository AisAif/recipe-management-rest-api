package routes

import (
	"github.com/AisAif/recipe-management-rest-api/src/http/controllers"
	"github.com/AisAif/recipe-management-rest-api/src/http/middleware"
	"github.com/AisAif/recipe-management-rest-api/src/models"
	"github.com/AisAif/recipe-management-rest-api/src/services"
	"github.com/AisAif/recipe-management-rest-api/src/utils"
	"github.com/gin-gonic/gin"
)

func User(r *gin.RouterGroup) {
	userService := services.NewUserService(models.DB, utils.InitValidator())
	userController := controllers.NewUserController(userService)

	authRouter := r.Use(middleware.AuthMiddleware())
	authRouter.GET("/current", userController.GetCurrent)
	authRouter.PATCH("/current", userController.UpdateCurrent)
}
