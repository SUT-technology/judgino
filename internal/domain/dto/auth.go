package dto

import "github.com/SUT-technology/judgino/internal/domain/model"

type LoginRequest struct {
	Username string `form:"username" validate:"required"`
	Password string `form:"password" validate:"required"`
}

type AuthResponse struct {
	Token string
	Error model.UserMessage
}

type SignupRequest struct {
	FirstName string `form:"first_name" validate:"required"`
	LastName  string `form:"last_name" validate:"required"`
	Phone     string `form:"phone" validate:"required"`
	Email     string `form:"email"`
	Password  string `form:"password" validate:"required"`
	Username  string `form:"username" validate:"required"`
}
