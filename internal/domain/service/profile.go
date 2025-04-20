package service

import (
	"context"

	"github.com/SUT-technology/judgino/internal/domain/dto"
)

type ProfileService interface {
	GetProfileById(ctx context.Context,currentUserId int64, profileDto dto.ProfileRequest) (*dto.ProfileRespone, error)
	ChangeRole(ctx context.Context,updateUserDTO dto.ChangeRoleRequest) (*dto.ChangeRoleResponse,error)
}