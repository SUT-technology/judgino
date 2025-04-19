package authsrvc

import (
	"context"
	"fmt"

	"github.com/SUT-technology/judgino/internal/domain/dto"
	"github.com/SUT-technology/judgino/internal/domain/entity"
	"github.com/SUT-technology/judgino/internal/domain/repository"
)

type PrflSrvc struct {
	db repository.Pool
}

func NewPrflSrvc(db repository.Pool) PrflSrvc {
	return PrflSrvc{
		db: db,
	}
}

func (c PrflSrvc) GetProfileById(ctx context.Context, ProfileDTO dto.ProfileDTO) (*dto.ProfileResponeDTO, error) {

	var (
		resp *dto.ProfileResponeDTO
		user *entity.User
		currentUser *entity.User
		err  error
	)

	queryFuncFindUser := func(r *repository.Repo) error {
		user, err = r.Tables.Users.GetUserById(ctx, ProfileDTO.UserId)
		if err != nil {
			return fmt.Errorf("find user by id: %w", err)
		}
		return nil
	}
	err = c.db.Query(ctx, queryFuncFindUser)
	if err != nil {
		return nil ,err
	}
	

	queryFuncFindCurrentUser := func(r *repository.Repo) error {
		currentUser, err = r.Tables.Users.GetUserById(ctx, 1)
		if err != nil {
			return fmt.Errorf("find current user by id: %w", err)
		}
		return nil
	}
	err = c.db.Query(ctx, queryFuncFindCurrentUser)
	if err != nil {
		return nil ,err
	}

	var p int

	if user.SubmissionsCount == 0 {
		p=0
	} else {
		p=100*int(user.SolvedQuestionsCount/user.SubmissionsCount)
	}

	resp = &dto.ProfileResponeDTO {
		UserId: ProfileDTO.UserId,    
		CurrentUserId: currentUser.ID,         
		Username: user.Username,
		Phone: user.Phone,
		Email: user.Email,
		Role: user.Role,
		Accepted: user.SolvedQuestionsCount,
		NotAccepted: user.SubmissionsCount-user.SolvedQuestionsCount,
		Total: user.SubmissionsCount,
		SolvedPercentage:  p,
		IsCurrentUserAdmin: currentUser.Role=="admin",
	}
	return resp, nil
}

func (c PrflSrvc) ChangeRole(ctx context.Context, UpdateUserDTO dto.ChangeRoleDTO) (*entity.User, error) {

	var (
		user *entity.User
		err  error
	)

	queryFuncUpdateUser := func(r *repository.Repo) error {
		user, err = r.Tables.Users.FindUserAndChangeRole(ctx,UpdateUserDTO)
		if err != nil {
			return fmt.Errorf("find customer by id: %w", err)
		}

		return nil
	}

	err = c.db.Query(ctx, queryFuncUpdateUser)
	if err != nil {
		return nil, err
	}
	return user, nil
}
