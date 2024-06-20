package services

import (
	"github.com/AisAif/recipe-management-rest-api/src/http/resources"
	"github.com/AisAif/recipe-management-rest-api/src/models"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type UserService interface {
	GetCurrent(username string) (resources.UserResource, error)
}

type UserServiceImpl struct {
	DB       *gorm.DB
	Validate *validator.Validate
}

func NewUserService(db *gorm.DB, validate *validator.Validate) UserService {
	return &UserServiceImpl{
		DB:       db,
		Validate: validate,
	}
}

func (s *UserServiceImpl) GetCurrent(username string) (resources.UserResource, error) {

	var user models.User
	err := s.DB.Find(&user, username).Error
	if err != nil {
		return resources.UserResource{}, err
	}
	return resources.UserResource{
		Name:     user.Name,
		Username: user.Username,
	}, nil
}
