package controllers

import (
	"net/http"

	"github.com/AisAif/recipe-management-rest-api/src/http/requests"
	"github.com/AisAif/recipe-management-rest-api/src/http/resources"
	"github.com/AisAif/recipe-management-rest-api/src/services"
	"github.com/gin-gonic/gin"
)

type UserController interface {
	GetCurrent(c *gin.Context)
	UpdateCurrent(c *gin.Context)
}
type UserControllerImpl struct {
	UserService services.UserService
}

func NewUserController(userService services.UserService) UserController {
	return &UserControllerImpl{
		UserService: userService,
	}
}

func (c UserControllerImpl) GetCurrent(ctx *gin.Context) {
	userValue, _ := ctx.Get("user")
	user, _ := userValue.(resources.UserResource)

	user, err := c.UserService.GetCurrent(user.Username)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, resources.Resource[resources.UserResource]{
		Message: "Success",
		Data:    user,
	})
}

func (c UserControllerImpl) UpdateCurrent(ctx *gin.Context) {
	var request requests.UpdateUserRequest
	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		ctx.Error(err)
		return
	}

	userValue, _ := ctx.Get("user")
	user, _ := userValue.(resources.UserResource)

	err = c.UserService.UpdateCurrent(user.Username, &request)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, resources.Resource[any]{
		Message: "Updated successfully",
	})
}
