package controllers

import (
	"github.com/AisAif/recipe-management-rest-api/src/http/requests"
	"github.com/AisAif/recipe-management-rest-api/src/services"
	"github.com/gin-gonic/gin"
)

type AuthController interface {
	Register(c *gin.Context)
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

	ctx.JSON(201, gin.H{"message": "Successfully registered"})
}
