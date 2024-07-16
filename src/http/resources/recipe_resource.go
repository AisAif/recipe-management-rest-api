package resources

import (
	"time"
)

type RecipeResource struct {
	ID        uint64       `json:"id"`
	Title     string       `json:"title"`
	Content   string       `json:"content"`
	ImageURL  string       `json:"image_url"`
	User      UserResource `json:"user"`
	IsPublic  bool         `json:"is_public"`
	CreatedAt time.Time    `json:"created_at"`
	UpdatedAt time.Time    `json:"updated_at"`
}
