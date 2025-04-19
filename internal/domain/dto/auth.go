package dto

import "github.com/SUT-technology/judgino/internal/domain/model"

type LoginRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type LoginResponse struct {
	Token string
}

type SignupRequest struct {
	Username string `query:"username" validate:"required"`
}

type SignupResponse struct {
	CurrentUserId int64
	Error         model.UserMessage
	Username      string
}
