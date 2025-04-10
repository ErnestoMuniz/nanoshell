package dto

type CreateUserDto struct {
	Username string `json:"username" validate:"required,min=3,max=32"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
	Admin    bool   `json:"admin" validate:"required"`
}
