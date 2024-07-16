package controllers

import (
	"net/http"

	"github.com/AisAif/recipe-management-rest-api/src/http/requests"
	"github.com/AisAif/recipe-management-rest-api/src/http/resources"
	"github.com/AisAif/recipe-management-rest-api/src/services"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

type RecipeController interface {
	Create(c *gin.Context)
	Update(c *gin.Context)
}
type RecipeControllerImpl struct {
	recipeService services.RecipeService
}

func NewRecipeController(recipeService services.RecipeService) RecipeController {
	return &RecipeControllerImpl{
		recipeService: recipeService,
	}
}

func (c RecipeControllerImpl) Create(ctx *gin.Context) {
	var request requests.CreateRecipeRequest

	err := ctx.ShouldBind(&request)
	if err != nil {
		ctx.Error(err)
		return
	}

	log.Debug().Msg("File name: " + request.Image.Header.Get("Content-Type"))

	userValue, _ := ctx.Get("user")
	user, _ := userValue.(resources.UserResource)

	err = c.recipeService.Create(user.Username, request)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusCreated, resources.Resource[any]{
		Message: "Successfully created",
	})
}

func (c RecipeControllerImpl) Update(ctx *gin.Context) {
	var request requests.UpdateRecipeRequest

	err := ctx.ShouldBind(&request)
	if err != nil {
		ctx.Error(err)
		return
	}

	userValue, _ := ctx.Get("user")
	user, _ := userValue.(resources.UserResource)

	err = c.recipeService.Update(user.Username, ctx.Param("recipe_id"), request)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, resources.Resource[any]{
		Message: "Successfully updated",
	})
}
