package authsrvc

import (
	"context"

	"github.com/SUT-technology/judgino/internal/domain/dto"
	"github.com/SUT-technology/judgino/internal/domain/repository"
)

type AuthSrvc struct {
	db repository.Pool
}

func NewAuthSrvc(db repository.Pool) AuthSrvc {
	return AuthSrvc{
		db: db,
	}
}

func (c AuthSrvc) Login(ctx context.Context, loginRequest dto.LoginRequest) (*dto.LoginResponse, error) {

	return &dto.LoginResponse{
		Token: "test",
	}, nil
}
func (c AuthSrvc) Signup(ctx context.Context, currentUserId int64, signupRequest dto.SignupRequest) (*dto.SignupResponse, error) {

	return &dto.SignupResponse{
		CurrentUserId: currentUserId,
		Username:      signupRequest.Username,
	}, nil
}
