package service

import (
	"context"

	"github.com/SUT-technology/judgino/internal/domain/dto"
)

type AuthService interface {
	Login(ctx context.Context, loginRequest dto.LoginRequest) (*dto.LoginResponse, error)
	Signup(ctx context.Context, currentUserId int64, signupRequest dto.SignupRequest) (*dto.SignupResponse, error)
}
