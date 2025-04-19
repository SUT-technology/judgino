package service

import (
	"context"

	"github.com/SUT-technology/judgino/internal/domain/dto"
	"github.com/SUT-technology/judgino/internal/domain/entity"
)

type AuthService interface {
	Login(ctx context.Context, loginDto dto.LoginDTO) (string, error)
	GetUser(ctx context.Context, userId uint) (*entity.User, error)
}
