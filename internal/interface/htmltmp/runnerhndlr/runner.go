package runnerhndlr

import (
	"fmt"
	"net/http"

	"github.com/SUT-technology/judgino/internal/domain/dto"
	"github.com/SUT-technology/judgino/internal/domain/service"
	"github.com/SUT-technology/judgino/pkg/slogger"
	"github.com/labstack/echo/v4"
)

type RunnerHndlr struct {
	Services service.Service
}

func New(g *echo.Group, srvc service.Service) RunnerHndlr {
	handler := RunnerHndlr{
		Services: srvc,
	}

	g.GET("/get", handler.SendSubmissions)
	g.POST("/result", handler.SubmitResultHandler)

	return handler
}

func (q *RunnerHndlr) SendSubmissions(c echo.Context) error {
	ctx := c.Request().Context()
	resp, err := q.Services.RunnerSrvc.SendSubmissions(ctx)
	if err != nil {
		slogger.Debug(ctx, "showSubmissions", slogger.Err("error", err))
		return c.Redirect(http.StatusFound, c.Request().Referer())
	}
	fmt.Printf("SendSubmissions resp: %+v", resp)

	return c.JSON(http.StatusOK, resp)
}

func (q *RunnerHndlr) SubmitResultHandler(c echo.Context) error {
	var result dto.SubmissionResult
	if err := c.Bind(&result); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid result format"})
	}

	err := q.Services.RunnerSrvc.SubmitResult(c.Request().Context(), result)
	if err != nil {
		slogger.Debug(c.Request().Context(), "SubmitResultHandler", slogger.Err("error", err))
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to submit result"})
	}

	return c.JSON(http.StatusOK, map[string]string{"status": "result received"})
}
