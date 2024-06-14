package services

import (
	"time"

	"github.com/AisAif/recipe-management-rest-api/src/http/requests"
	"github.com/AisAif/recipe-management-rest-api/src/http/resources"
	"github.com/AisAif/recipe-management-rest-api/src/models"
	"github.com/AisAif/recipe-management-rest-api/src/utils"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type AuthService interface {
	Register(ctx *gin.Context, request *requests.RegisterRequest) error
	Login(ctx *gin.Context, request *requests.LoginRequest) (string, error)
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
	if err := s.Validate.Struct(request); err != nil {
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

	if err := s.DB.Create(&userRegisterData).Error; err != nil {
		return err
	}

	return nil
}

func (s *AuthServiceImpl) Login(ctx *gin.Context, request *requests.LoginRequest) (string, error) {
	if err := s.Validate.Struct(request); err != nil {
		return "", err
	}

	var user models.User
	if err := s.DB.Where("username = ?", request.Username).First(&user).Error; err != nil {
		return "", err
	}

	if err := utils.CheckPasswordHash(request.Password, user.Password); err != nil {
		return "", err
	}

	tokenString, err := utils.GenerateJWT(resources.UserResource{
		Name:     user.Name,
		Username: user.Username,
	})

	if err != nil {
		return "", err
	}

	return tokenString, nil
}
