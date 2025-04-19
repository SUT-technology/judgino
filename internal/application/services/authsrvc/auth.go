package authsrvc

import (
	"context"
	"fmt"

	"github.com/SUT-technology/judgino/internal/domain/dto"
	"github.com/SUT-technology/judgino/internal/domain/entity"
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

func (c AuthSrvc) Login(ctx context.Context, loginDto dto.LoginDTO) (string, error) {

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
		return "", err
	}
	return user.FirstName, nil
}
func (c AuthSrvc) GetUser(ctx context.Context, userId uint) (*entity.User, error) {
	var (
		user *entity.User
		err  error
	)

	queryFuncFindUser := func(r *repository.Repo) error {
		user, err = r.Tables.Users.GetUserById(ctx, userId)
		if err != nil {
			return fmt.Errorf("find customer by id: %w", err)
		}
		return nil
	}
	err = c.db.Query(ctx, queryFuncFindUser)
	if err != nil {	
		return nil, err
	}
	return user, nil
}
