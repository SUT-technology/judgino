package repository

import (
	"context"

	"github.com/SUT-technology/judgino/internal/domain/dto"
	"github.com/SUT-technology/judgino/internal/domain/entity"
)

type UserRepository interface {
	GetUserById(ctx context.Context, id int64)(*entity.User , error)
	FindUserAndChangeRole(ctx context.Context,data dto.ChangeRoleRequest) (*entity.User,error)
}