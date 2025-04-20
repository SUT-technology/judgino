package htmltmp

import (
	"html/template"
	"net/http"
	"log/slog"
	"strconv"
	"fmt"


	"github.com/SUT-technology/judgino/internal/domain/dto"
	"github.com/SUT-technology/judgino/internal/domain/service"
	"github.com/SUT-technology/judgino/pkg/slogger"
	"github.com/SUT-technology/judgino/internal/interface/htmltmp/serde"
	"github.com/SUT-technology/judgino/internal/domain/model"
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


func (h ProfileHndlr) HandleChangeRole(c echo.Context) error{

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


func (h ProfileHndlr) HandleProfile(c echo.Context)error {
	ctx := c.Request().Context()

	id := c.Param("id")
	userID64 , _ := strconv.ParseUint(id,10,64)
	userId:=int64(userID64)

	req:=dto.ProfileRequest{
		UserId: userId,
	}
	
	var currentUserId int64
	currentUser := serde.GetCurrentUser(c)
	if currentUser == nil {
		//example:  GET user id from path
		currentUserId = 1 // TEST
	} else {
		currentUserId = currentUser.UserId
	}

	slogger.Debug(ctx, "received request", slog.Any("request", req))

	resp,err:=h.Services.PrflSrvc.GetProfileById(ctx,currentUserId,req)
	fmt.Println(resp)
	if err != nil {
		slogger.Debug(ctx, "profile", slogger.Err("error", err))
		return c.Render(http.StatusBadRequest, "test.html", dto.SignupResponse{Error: model.InternalMessage})
	}

	tmpl := template.Must(template.New("profile.html").Funcs(template.FuncMap{
		"eqs": func(a, b string) bool {
			return a == b
		},
		"equi": func(a, b uint) bool {
			return a == b
		},
	}).ParseFiles("D:/GOprojects/practice/judgino/templates/profile.html"))
	return c.Render(http.StatusOK, tmpl.Name(), resp)
}
