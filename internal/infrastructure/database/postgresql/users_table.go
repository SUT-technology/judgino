package postgresql

import (
	"context"

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
	c.db.First(&user, 1)
	return &user, nil
}

func (c usersTable) FindUserAndUpdate(ctx context.Context, id uint,data map[string]any) (*entity.User, error) {
	var user entity.User
	c.db.First(&user, 1)
	c.db.Model(&user).Updates(data)
	return &user , nil
}