package authhdnlr

import (
	"log/slog"
	"net/http"

	"github.com/SUT-technology/judgino/internal/domain/dto"
	"github.com/SUT-technology/judgino/internal/domain/model"
	"github.com/SUT-technology/judgino/internal/domain/service"
	"github.com/SUT-technology/judgino/internal/interface/htmltmp/serde"
	"github.com/SUT-technology/judgino/pkg/slogger"
	"github.com/labstack/echo/v4"
)

type AuthHndlr struct {
	Services service.Service
}

func New(g *echo.Group, srvc service.Service) AuthHndlr {
	handler := AuthHndlr{
		Services: srvc,
	}

	g.GET("", handler.LoginPage)
	g.POST("/signup", handler.Signup)
	g.POST("/login", handler.Login)

	return handler
}

func (h AuthHndlr) Login(c echo.Context) error {

	ctx := c.Request().Context()

	req, err := serde.BindRequestBody[dto.LoginRequest](c)
	if err != nil {
		slogger.Debug(ctx, "bad request", slogger.Err("error", err))
		return c.Render(http.StatusBadRequest, "login.html", dto.AuthResponse{Error: model.BadRequestMessage})
	}

	slogger.Debug(ctx, "received request", slog.Any("request", req))

	resp, err := h.Services.AuthSrvc.Login(ctx, req)
	if err != nil {
		slogger.Debug(ctx, "login ", slogger.Err("error", err))
		return c.Render(http.StatusBadRequest, "login.html", resp)
	}

	serde.SetTokenCookie(c, resp.Token)

	return c.Redirect(http.StatusMovedPermanently, "/questions")

}

func (h *AuthHndlr) Signup(c echo.Context) error {

	ctx := c.Request().Context()

	req, err := serde.BindRequestBody[dto.SignupRequest](c)
	if err != nil {
		slogger.Debug(ctx, "bad request", slogger.Err("error", err))
		return c.Render(http.StatusBadRequest, "login.html", dto.AuthResponse{Error: model.BadRequestMessage})
	}

	slogger.Debug(ctx, "received request", slog.Any("request", req))

	resp, err := h.Services.AuthSrvc.Signup(ctx, req)
	if err != nil {
		slogger.Debug(ctx, "signup", slogger.Err("error", err))
		return c.Render(http.StatusBadRequest, "test.html", resp)
	}

	serde.SetTokenCookie(c, resp.Token)

	return c.Redirect(http.StatusMovedPermanently, "/questions")
}

func (h AuthHndlr) LoginPage(c echo.Context) error {

	// ctx := c.Request().Context()
	return c.Render(http.StatusOK, "login.html", dto.AuthResponse{})
}
