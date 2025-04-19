package service

import (
	"context"

	"github.com/SUT-technology/judgino/internal/domain/entity"
)

type UserService interface {
	GetUser(ctx context.Context, userId uint) (*entity.User, error)
}
