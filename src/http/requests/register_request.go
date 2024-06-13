package requests

type RegisterRequest struct {
	Username string `validate:"required,min=3,max=255,username_exists" json:"username"`
	Name     string `validate:"required,min=3,max=255" json:"name"`
	Password string `validate:"required,min=8,max=255" json:"password"`
}
