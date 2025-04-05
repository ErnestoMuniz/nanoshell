package dto

type UpdateUserDto struct {
	Username string `json:"username" validate:"required,min=3,max=32"`
	Email    string `json:"email" validate:"required,email"`
	Admin    bool   `json:"admin" validate:"required"`
	Active   bool   `json:"active" validate:"required"`
	Password string `json:"password" validate:"min=8"`
}
