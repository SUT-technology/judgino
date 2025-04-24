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

func (c usersTable) GetUserById(ctx context.Context, id int64) (*entity.User, error) {
	var user entity.User
	c.db.First(&user, id)
	return &user, nil
}

func (c usersTable) FindUserAndChangeRole(ctx context.Context, data dto.ChangeRoleRequest) (*entity.User, error) {
	var user entity.User

	if err := c.db.WithContext(ctx).First(&user, data.ID).Error; err != nil {
		return nil, err
	}

	if err := c.db.WithContext(ctx).Model(&user).Updates(data).Error; err != nil {
		return nil, err
	}

	c.db.First(&user, data.ID)
	return &user, nil
}


func (c usersTable) FindAndUpdateUser(ctx context.Context,userId int64, data entity.User) error {

	var user entity.User
	c.db.First(&user, userId)

	if err := c.db.WithContext(ctx).Model(&user).Updates(data).Error; err != nil {

func (c usersTable) GetUserByUsername(ctx context.Context, username string) (*entity.User, error) {
	var user entity.User
	if err := c.db.Where("username = ?", username).First(&user).Error; err != nil {
		return nil, err // User not found
	}

	return &user, nil
}
func (c usersTable) CreateUser(ctx context.Context, user entity.User) error {
	if err := c.db.Create(&user).Error; err != nil {

		return err
	}
	return nil
}
