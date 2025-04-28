package logouthndlr

import (
	"fmt"
	"net/http"

	"github.com/SUT-technology/judgino/internal/domain/service"
	"github.com/SUT-technology/judgino/internal/interface/htmltmp/serde"
	"github.com/labstack/echo/v4"
)

type LogoutHndlr struct {
	Services service.Service
}

func New(g *echo.Group, srvc service.Service) LogoutHndlr {
	handler := LogoutHndlr{
		Services: srvc,
	}
	g.GET("/", handler.ClearCookies)
	g.GET("", handler.ClearCookies)
	return handler
}

func (l *LogoutHndlr) ClearCookies (c echo.Context) error {
	serde.SetTokenCookie(c, "")
	fmt.Println("hoooooo")

	return c.Redirect(http.StatusFound, "/auth")
}