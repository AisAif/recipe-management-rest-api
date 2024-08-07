package routes

import (
	"github.com/AisAif/recipe-management-rest-api/src/http/controllers"
	"github.com/AisAif/recipe-management-rest-api/src/http/middleware"
	"github.com/AisAif/recipe-management-rest-api/src/models"
	"github.com/AisAif/recipe-management-rest-api/src/services"
	"github.com/AisAif/recipe-management-rest-api/src/utils"
	"github.com/gin-gonic/gin"
)

func Recipe(r *gin.RouterGroup) {
	recipeService := services.NewRecipeService(models.DB, utils.InitValidator())
	recipeController := controllers.NewRecipeController(recipeService)

	r.GET("", recipeController.PublishedList)

	authRouter := r.Use(middleware.AuthMiddleware())
	authRouter.POST("", recipeController.Create)
	authRouter.PATCH("/:recipe_id", recipeController.Update)
	authRouter.PATCH("/:recipe_id/toggle-publish", recipeController.TogglePublish)
	authRouter.DELETE("/:recipe_id", recipeController.Delete)
	authRouter.GET("/current", recipeController.Current)
}
