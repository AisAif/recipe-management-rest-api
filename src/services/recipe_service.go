package services

import (
	"time"

	"github.com/AisAif/recipe-management-rest-api/src/http/requests"
	"github.com/AisAif/recipe-management-rest-api/src/models"
	"github.com/AisAif/recipe-management-rest-api/src/storage"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type RecipeService interface {
	Create(username string, request requests.CreateRecipeRequest) error
	Update(username string, id string, request requests.UpdateRecipeRequest) error
	Delete(username string, id string) error
}

type RecipeServiceImpl struct {
	DB       *gorm.DB
	Validate *validator.Validate
}

func NewRecipeService(db *gorm.DB, validate *validator.Validate) RecipeService {
	return &RecipeServiceImpl{
		DB:       db,
		Validate: validate,
	}
}

func (s *RecipeServiceImpl) Create(username string, request requests.CreateRecipeRequest) error {

	if err := s.Validate.Struct(request); err != nil {
		return err
	}

	imageUrl, err := storage.Storage.Store("recipe/"+username, request.Image)
	if err != nil {
		return err
	}

	recipe := models.Recipe{
		Title:     request.Title,
		Content:   request.Content,
		ImageURL:  imageUrl,
		Username:  username,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if err := s.DB.Create(&recipe).Error; err != nil {
		return err
	}

	return nil
}

func (s *RecipeServiceImpl) Update(username string, id string, request requests.UpdateRecipeRequest) error {
	if err := s.Validate.Struct(request); err != nil {
		return err
	}

	var recipe *models.Recipe
	result := s.DB.Find(&recipe, "id = ?", id)
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	if request.Title != "" {
		recipe.Title = request.Title
	}

	if request.Content != "" {
		recipe.Content = request.Content
	}

	if request.Image != nil {

		err := storage.Storage.Delete(recipe.ImageURL)
		if err != nil {
			return err
		}
		imageUrl, err := storage.Storage.Store("recipe/"+username, request.Image)
		if err != nil {
			return err
		}

		recipe.ImageURL = imageUrl
	}

	recipe.UpdatedAt = time.Now()

	return s.DB.Save(&recipe).Error
}

func (s *RecipeServiceImpl) Delete(username string, id string) error {
	var recipe *models.Recipe
	result := s.DB.Find(&recipe, "id = ?", id)
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	err := storage.Storage.Delete(recipe.ImageURL)
	if err != nil {
		return err
	}

	return s.DB.Delete(&recipe).Error
}
