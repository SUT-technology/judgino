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

	g.GET("/login", handler.Login)
	g.GET("/signup", handler.Signup)

	return handler
}

func (h AuthHndlr) Login(c echo.Context) error {

	ctx := c.Request().Context()

	req, err := serde.BindRequestBody[dto.LoginRequest](c)
	if err != nil {
		slogger.Debug(ctx, "bad request", slogger.Err("error", err))
		return serde.Response(c, http.StatusBadRequest, model.BadRequestMessage, nil)
	}

	slogger.Debug(ctx, "received request", slog.Any("request", req))

	resp, err := h.Services.AuthSrvc.Login(ctx, req)
	if err != nil {
		slogger.Debug(ctx, "login ", slogger.Err("error", err))
		return serde.Response(c, http.StatusInternalServerError, model.InternalMessage, nil)
	}

	return serde.Response(c, http.StatusOK, model.OKMessage, resp)

}

func (h *AuthHndlr) Signup(c echo.Context) error {

	ctx := c.Request().Context()

	req, err := serde.BindRequestBody[dto.SignupRequest](c)
	if err != nil {
		slogger.Debug(ctx, "bad request", slogger.Err("error", err))
		return c.Render(http.StatusBadRequest, "test.html", dto.SignupResponse{Error: model.BadRequestMessage})
	}

	var currentUserId int64
	currentUser := serde.GetCurrentUser(c)
	if currentUser == nil {
		//example:  GET user id from path
		currentUserId = 2 // TEST
	} else {
		currentUserId = currentUser.UserId
	}

	slogger.Debug(ctx, "received request", slog.Any("request", req))

	resp, err := h.Services.AuthSrvc.Signup(ctx, currentUserId, req)
	if err != nil {
		slogger.Debug(ctx, "signup", slogger.Err("error", err))
		return c.Render(http.StatusBadRequest, "test.html", dto.SignupResponse{Error: model.InternalMessage})
	}

	return c.Render(http.StatusOK, "test.html", resp)

}
