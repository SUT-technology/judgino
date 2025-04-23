package service

import (
	"context"

	"github.com/SUT-technology/judgino/internal/domain/dto"
)

type AuthService interface {
	Login(ctx context.Context, loginRequest dto.LoginRequest) (*dto.AuthResponse, error)
	Signup(ctx context.Context, signupRequest dto.SignupRequest) (*dto.AuthResponse, error)
}
