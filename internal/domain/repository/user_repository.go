package repository

import (
	"context"

	"github.com/SUT-technology/judgino/internal/domain/entity"
)

type UserRepository interface {
	GetUserById(ctx context.Context, id uint)(*entity.User , error)
}