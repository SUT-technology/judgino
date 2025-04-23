package submissionshndlr

import (
	"net/http"
	"strconv"

	"github.com/SUT-technology/judgino/internal/domain/dto"
	"github.com/SUT-technology/judgino/internal/domain/model"
	"github.com/SUT-technology/judgino/internal/domain/service"
	"github.com/SUT-technology/judgino/internal/interface/htmltmp/serde"
	"github.com/SUT-technology/judgino/pkg/slogger"
	"github.com/labstack/echo/v4"
)

type SubmissionsHndlr struct {
	Services service.Service
}

func New(g *echo.Group, srvc service.Service) SubmissionsHndlr {
	handler := SubmissionsHndlr{
		Services: srvc,
	}
	g.GET("/{question_id}", handler.ShowSubmissions)
	g.GET("/:question_id", handler.ShowSubmissions)

	return handler
}

func (q *SubmissionsHndlr) ShowSubmissions(c echo.Context) error {
	questionID := c.Param("question_id")
	questionIDInt, _ := strconv.Atoi(questionID)

	ctx := c.Request().Context()

	req, err := serde.BindRequestBody[dto.SubmissionRequest](c)
	if err != nil {
		slogger.Debug(ctx, "bad request", slogger.Err("error", err))
		return serde.Response(c, http.StatusBadRequest, model.BadRequestMessage, nil)
	}

	resp, err := q.Services.SubmissionSrvc.GetSubmissions(ctx, req, uint(serde.GetCurrentUser(c).UserId), serde.GetCurrentUser(c).IsAdmin, questionIDInt)
	if err != nil {
		slogger.Debug(ctx, "showSubmissions", slogger.Err("error", err))
		// TODO: handle error
		return c.Render(http.StatusBadRequest, "submissions.html", resp)
	}

	return c.Render(http.StatusOK, "submissions.html", resp)
}
