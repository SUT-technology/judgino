package postgresql

import (
	"context"

	"github.com/SUT-technology/judgino/internal/domain/dto"
	"github.com/SUT-technology/judgino/internal/domain/entity"
	"gorm.io/gorm"
)

type usersTable struct {
	db *gorm.DB
}

func newUsersTable(db *gorm.DB) usersTable {
	return usersTable{db: db}
}

func (c usersTable) GetUserById(ctx context.Context, id uint) (*entity.User, error) {
	var user entity.User
	c.db.First(&user, id)
	return &user, nil
}

func (c usersTable) FindUserAndUpdate(ctx context.Context, data dto.UpdateUserDTO) (*entity.User, error) {
	var user entity.User

	if err := c.db.WithContext(ctx).First(&user, data.ID).Error; err != nil {
		return nil, err
	}

	if err := c.db.WithContext(ctx).Model(&user).Updates(data).Error; err != nil {
		return nil, err
	}

	return &user, nil
}
