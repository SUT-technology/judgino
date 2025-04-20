package service

import (
	"context"

)

type UserService interface {
	GetUserName(ctx context.Context, userId uint) (string, error)
}
