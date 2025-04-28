package serde

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

type UserDataDto struct {
	UserId  int64
	IsAdmin bool
}

func GetCurrentUser(c echo.Context) *UserDataDto {
	userId, ok := c.Get("user_id").(int64)
	if !ok {
		return nil
	}

	isAdmin, ok := c.Get("is_admin").(bool)
	if !ok {
		return nil
	}

	return &UserDataDto{
		UserId:  userId,
		IsAdmin: isAdmin,
	}

}

func SetTokenCookie(c echo.Context, token string) {
	cookie := new(http.Cookie)
	cookie.Name = "token"
	cookie.Value = token
	cookie.Expires = time.Now().Add(72 * time.Hour) // Expiration time
	cookie.HttpOnly = true                          // Prevent JavaScript access
	cookie.Path = "/"
	c.SetCookie(cookie)
}
