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

func (c PrflSrvc) GetProfileById(ctx context.Context, ProfileDTO dto.ProfileDTO) (*entity.User, error) {

	var (
		user *entity.User
		err  error
	)

	queryFuncFindCustomer := func(r *repository.Repo) error {
		user, err = r.Tables.Users.GetUserById(ctx, 1)
		if err != nil {
			return fmt.Errorf("find customer by id: %w", err)
		}

		return nil
	}

	err = c.db.Query(ctx, queryFuncFindCustomer)
	if err != nil {
		return nil, err
	}
	return user, nil
}
