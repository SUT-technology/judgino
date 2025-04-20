package usersrvc

import (
	"context"
	"fmt"

	"github.com/SUT-technology/judgino/internal/domain/entity"
	"github.com/SUT-technology/judgino/internal/domain/repository"
)

type UserService struct {
	db repository.Pool
}

func NewUserSrvc(db repository.Pool) UserService {
	return UserService{
		db: db,
	}
}

func (c UserService) GetUserName(ctx context.Context, userId uint) (string, error) {
	var (
		user *entity.User
		err  error
	)

	queryFuncFindUser := func(r *repository.Repo) error {
		user, err = r.Tables.Users.GetUserById(ctx, int64(userId))
		if err != nil {
			return fmt.Errorf("find customer by id: %w", err)
		}
		return nil
	}
	err = c.db.Query(ctx, queryFuncFindUser)
	if err != nil {	
		return "", err
	}
	return user.FirstName, nil
}