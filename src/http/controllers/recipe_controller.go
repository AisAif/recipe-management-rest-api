package controllers

import (
	"net/http"

	"github.com/AisAif/recipe-management-rest-api/src/http/requests"
	"github.com/AisAif/recipe-management-rest-api/src/http/resources"
	"github.com/AisAif/recipe-management-rest-api/src/services"
	"github.com/gin-gonic/gin"
)

type RecipeController interface {
	Create(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
	TogglePublish(c *gin.Context)
	Current(c *gin.Context)
	PublishedList(c *gin.Context)
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

func (c RecipeControllerImpl) Delete(ctx *gin.Context) {
	userValue, _ := ctx.Get("user")
	user, _ := userValue.(resources.UserResource)

	err := c.recipeService.Delete(user.Username, ctx.Param("recipe_id"))
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, resources.Resource[any]{
		Message: "Successfully deleted",
	})
}

func (c RecipeControllerImpl) TogglePublish(ctx *gin.Context) {
	userValue, _ := ctx.Get("user")
	user, _ := userValue.(resources.UserResource)

	err := c.recipeService.TogglePublish(user.Username, ctx.Param("recipe_id"))
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, resources.Resource[any]{
		Message: "Successfully toggled",
	})
}

func (c RecipeControllerImpl) Current(ctx *gin.Context) {
	userValue, _ := ctx.Get("user")
	user, _ := userValue.(resources.UserResource)

	recipes, pageInfo, err := c.recipeService.List(ctx, user.Username, false)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, resources.Resource[[]resources.RecipeResource]{
		Message:  "Successfully fetched",
		Data:     recipes,
		PageInfo: &pageInfo,
	})
}

func (c RecipeControllerImpl) PublishedList(ctx *gin.Context) {
	recipes, pageInfo, err := c.recipeService.List(ctx, "", true)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, resources.Resource[[]resources.RecipeResource]{
		Message:  "Successfully fetched",
		Data:     recipes,
		PageInfo: &pageInfo,
	})
}
