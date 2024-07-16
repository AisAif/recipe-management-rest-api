package requests

import "mime/multipart"

type CreateRecipeRequest struct {
	Title   string                `form:"title" validate:"required,min=3,max=255"`
	Content string                `form:"content" validate:"required,min=1,max=1500"`
	Image   *multipart.FileHeader `form:"image" validate:"required,file_type=image/*,max_size=2048"`
}

type UpdateRecipeRequest struct {
	Title   string                `form:"title" validate:"omitempty,min=3,max=255"`
	Content string                `form:"content" validate:"omitempty,min=1,max=1500"`
	Image   *multipart.FileHeader `form:"image" validate:"omitempty,file_type=image/*,max_size=2048"`
}
