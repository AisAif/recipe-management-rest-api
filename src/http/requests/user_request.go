package requests

type UpdateUserRequest struct {
	Name     string `validate:"omitempty,min=3,max=255" json:"name"`
	Password string `validate:"omitempty,min=8,max=255" json:"password"`
}
