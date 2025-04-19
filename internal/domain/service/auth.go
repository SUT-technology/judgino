package service

import (
	"context"

	"github.com/SUT-technology/judgino/internal/domain/dto"
)

type AuthService interface {
	Login(ctx context.Context, loginDto dto.LoginDTO) (string, error)
}
