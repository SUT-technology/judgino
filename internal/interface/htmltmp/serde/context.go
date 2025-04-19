package serde

import (
	"fmt"
	"log/slog"

	"github.com/labstack/echo/v4"
)

type UserDataDto struct {
	UserId  int64
	IsAdmin bool
}

func GetCurrentUser(c echo.Context) *UserDataDto {
	userId, ok := c.Get("userId").(int64)
	slog.Info(fmt.Sprintf("new user id: %v", userId))
	if !ok {
		return nil
	}

	isAdmin, ok := c.Get("isAdmin").(bool)
	if !ok {
		return nil
	}

	return &UserDataDto{
		UserId:  userId,
		IsAdmin: isAdmin,
	}

}
