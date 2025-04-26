package profhndlr

import (
	"fmt"
	"html/template"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/SUT-technology/judgino/internal/domain/dto"
	"github.com/SUT-technology/judgino/internal/domain/model"
	"github.com/SUT-technology/judgino/internal/domain/service"
	"github.com/SUT-technology/judgino/internal/interface/htmltmp/serde"
	"github.com/SUT-technology/judgino/pkg/slogger"
	"github.com/labstack/echo/v4"
)

type ProfileHndlr struct {
	Services service.Service
}

func New(g *echo.Group, srvc service.Service) ProfileHndlr {
	handler := ProfileHndlr{
		Services: srvc,
	}

	g.GET("/:id", handler.HandleProfile)
	g.POST("/change-role", handler.HandleChangeRole)

	return handler
}

func (h ProfileHndlr) HandleChangeRole(c echo.Context) error {

	ctx := c.Request().Context()

	req, err := serde.BindRequestBody[dto.ChangeRoleRequest](c)
	if err != nil {
		slogger.Debug(ctx, "bad request", slogger.Err("error", err))
		return serde.Response(c, http.StatusBadRequest, model.BadRequestMessage, nil)
	}

	slogger.Debug(ctx, "received request", slog.Any("request", req))

	resp, err := h.Services.PrflSrvc.ChangeRole(ctx, req)
	if err != nil {
		slogger.Debug(ctx, "changerole ", slogger.Err("error", err))
		return serde.Response(c, http.StatusInternalServerError, model.InternalMessage, nil)
	}

	return serde.Response(c, http.StatusOK, model.OKMessage, resp)
}

func (h ProfileHndlr) HandleProfile(c echo.Context) error {
	ctx := c.Request().Context()

	slog.Info(fmt.Sprintf("test user id: %v", c.Get("user_id")))

	id := c.Param("id")
	userID64, _ := strconv.ParseUint(id, 10, 64)
	userId := int64(userID64)

	var currentUserId int64
	currentUser := serde.GetCurrentUser(c)
	currentUserId = currentUser.UserId

	slogger.Debug(ctx, "received request", slog.Any("request", userId))

	resp, err := h.Services.PrflSrvc.GetProfileById(ctx, currentUserId, userId)
	if err != nil {
		slogger.Debug(ctx, "profile", slogger.Err("error", err))
		return c.Render(http.StatusBadRequest, "test.html", dto.ProfileRespone{Error: model.InternalMessage})
	}

	tmpl := template.Must(template.New("profile.html").Funcs(template.FuncMap{
		"eqs": func(a, b string) bool {
			return a == b
		},
		"equi": func(a, b uint) bool {
			return a == b
		},
	}).ParseFiles("profile.html"))
	return c.Render(http.StatusOK, tmpl.Name(), resp)
}
