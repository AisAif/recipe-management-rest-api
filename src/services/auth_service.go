package services

import (
	"time"

	"github.com/AisAif/recipe-management-rest-api/src/http/requests"
	"github.com/AisAif/recipe-management-rest-api/src/models"
	"github.com/AisAif/recipe-management-rest-api/src/utils"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type AuthService interface {
	Register(ctx *gin.Context, request *requests.RegisterRequest) error
}

type AuthServiceImpl struct {
	DB       *gorm.DB
	Validate *validator.Validate
}

func NewAuthService(db *gorm.DB, validate *validator.Validate) AuthService {
	return &AuthServiceImpl{
		DB:       db,
		Validate: validate,
	}
}

func (s *AuthServiceImpl) Register(ctx *gin.Context, request *requests.RegisterRequest) error {
	err := s.Validate.Struct(request)
	if err != nil {
		return err
	}

	hash, err := utils.HashPassword(request.Password)
	if err != nil {
		return err
	}

	userRegisterData := models.User{
		Username:  request.Username,
		Name:      request.Name,
		Password:  hash,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err = s.DB.Create(&userRegisterData).Error
	if err != nil {
		return err
	}

	return nil
}
