package services

import (
	"time"

	"github.com/AisAif/recipe-management-rest-api/src/http/requests"
	"github.com/AisAif/recipe-management-rest-api/src/http/resources"
	"github.com/AisAif/recipe-management-rest-api/src/models"
	"github.com/AisAif/recipe-management-rest-api/src/storage"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/rosberry/go-pagination"
	"gorm.io/gorm"
)

type RecipeService interface {
	Create(username string, request requests.CreateRecipeRequest) error
	Update(username string, id string, request requests.UpdateRecipeRequest) error
	Delete(username string, id string) error
	TogglePublish(username string, id string) error
	List(ctx *gin.Context, username string, isPublish bool) ([]resources.RecipeResource, pagination.PageInfo, error)
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

func (s *RecipeServiceImpl) TogglePublish(username string, id string) error {
	var recipe *models.Recipe
	result := s.DB.Find(&recipe, "id = ?", id)
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	recipe.IsPublic = !recipe.IsPublic

	return s.DB.Save(&recipe).Error
}

func (s *RecipeServiceImpl) List(ctx *gin.Context, username string, isPublish bool) ([]resources.RecipeResource, pagination.PageInfo, error) {
	paginator, err := pagination.New(pagination.Options{
		GinContext:    ctx,
		DB:            s.DB,
		Model:         &models.Recipe{},
		Limit:         5,
		DefaultCursor: nil,
	})
	if err != nil {
		return nil, pagination.PageInfo{}, err
	}

	var result *gorm.DB

	if username == "" {
		result = s.DB.Preload("User").Find(&[]models.Recipe{})
	} else {
		result = s.DB.Preload("User").Where("username = ?", username).Find(&[]models.Recipe{})
	}

	if isPublish {
		result = result.Where("is_public = ?", true)
	}

	var recipes []models.Recipe
	err = paginator.Find(result, &recipes)
	if err != nil {
		return nil, pagination.PageInfo{}, err
	}

	if len(recipes) == 0 {
		return []resources.RecipeResource{}, pagination.PageInfo{}, nil
	}

	var recipeResources []resources.RecipeResource
	for _, recipe := range recipes {

		imageUrl, err := storage.Storage.GetURL(recipe.ImageURL)
		if err != nil {
			return nil, pagination.PageInfo{}, err
		}

		recipeResources = append(recipeResources, resources.RecipeResource{
			ID:        recipe.ID,
			Title:     recipe.Title,
			Content:   recipe.Content,
			ImageURL:  imageUrl,
			IsPublic:  recipe.IsPublic,
			CreatedAt: recipe.CreatedAt,
			UpdatedAt: recipe.UpdatedAt,
			User: resources.UserResource{
				Name:     recipe.User.Name,
				Username: recipe.User.Username,
			},
		})
	}

	return recipeResources, *paginator.PageInfo, nil
}
