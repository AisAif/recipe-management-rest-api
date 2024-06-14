package controllers

import (
	"net/http"

	"github.com/AisAif/recipe-management-rest-api/src/http/requests"
	"github.com/AisAif/recipe-management-rest-api/src/http/resources"
	"github.com/AisAif/recipe-management-rest-api/src/services"
	"github.com/gin-gonic/gin"
)

type AuthController interface {
	Register(c *gin.Context)
	Login(c *gin.Context)
}
type AuthControllerImpl struct {
	AuthService services.AuthService
}

func NewAuthController(authService services.AuthService) AuthController {
	return &AuthControllerImpl{
		AuthService: authService,
	}
}

func (c AuthControllerImpl) Register(ctx *gin.Context) {
	var request requests.RegisterRequest

	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		ctx.Error(err)
		return
	}

	err = c.AuthService.Register(ctx, &request)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusCreated, resources.Resource[any]{
		Message: "Successfully registered",
	})
}

func (c AuthControllerImpl) Login(ctx *gin.Context) {
	var request requests.LoginRequest

	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		ctx.Error(err)
		return
	}

	token, err := c.AuthService.Login(ctx, &request)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, resources.Resource[any]{
		Message: "Login succesfully",
		Data: map[string]string{
			"token": token,
		},
	})
}
