package services

import (
	"time"

	"github.com/AisAif/recipe-management-rest-api/src/http/requests"
	"github.com/AisAif/recipe-management-rest-api/src/http/resources"
	"github.com/AisAif/recipe-management-rest-api/src/models"
	"github.com/AisAif/recipe-management-rest-api/src/utils"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type UserService interface {
	GetCurrent(username string) (resources.UserResource, error)
	UpdateCurrent(username string, request *requests.UpdateUserRequest) error
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
	err := s.DB.Find(&user, "username = ?", username).Error
	if err != nil {
		return resources.UserResource{}, err
	}
	return resources.UserResource{
		Name:     user.Name,
		Username: user.Username,
	}, nil
}

func (s *UserServiceImpl) UpdateCurrent(username string, request *requests.UpdateUserRequest) error {

	if err := s.Validate.Struct(request); err != nil {
		return err
	}

	var user *models.User
	err := s.DB.Find(&user, "username = ?", username).Error
	if err != nil {
		return err
	}

	if request.Password != "" {
		hash, err := utils.HashPassword(request.Password)
		if err != nil {
			return err
		}
		user.Password = hash
	}

	if request.Name != "" {
		user.Name = request.Name
	}

	user.UpdatedAt = time.Now()

	return s.DB.Save(&user).Error
}
