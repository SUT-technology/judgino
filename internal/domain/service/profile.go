package service

import (
	"context"

	"github.com/SUT-technology/judgino/internal/domain/dto"
	"github.com/SUT-technology/judgino/internal/domain/entity"
)

type ProfileService interface {
	GetProfileById(ctx context.Context, profileDto dto.ProfileDTO) (*dto.ProfileResponeDTO, error)
	ChangeRole(ctx context.Context,updateUserDTO dto.UpdateUserDTO) (*entity.User,error)
}