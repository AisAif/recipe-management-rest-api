package controllers

import (
	"net/http"

	"github.com/AisAif/recipe-management-rest-api/src/http/resources"
	"github.com/AisAif/recipe-management-rest-api/src/services"
	"github.com/gin-gonic/gin"
)

type UserController interface {
	GetCurrent(c *gin.Context)
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
	username := ctx.GetString("username")

	user, err := c.UserService.GetCurrent(username)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusCreated, resources.Resource[resources.UserResource]{
		Message: "Success",
		Data:    user,
	})
}
