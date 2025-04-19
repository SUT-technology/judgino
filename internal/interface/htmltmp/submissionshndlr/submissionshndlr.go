package submissionshndlr

import (
	"net/http"
	"strconv"

	"github.com/SUT-technology/judgino/internal/domain/dto"
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
	g.POST("/{question_id}", handler.ShowSubmissions)
	g.POST("/:question_id", handler.ShowSubmissions)


	return handler
}


func (q *SubmissionsHndlr) ShowSubmissions(c echo.Context) error {
	questionID := c.Param("question_id")
	questionIDInt, _ := strconv.Atoi(questionID)
	
	submissionDto := dto.SubmissionRequest{
		UserId:       uint(serde.GetCurrentUser(c).UserId),
		IsAdmin: serde.GetCurrentUser(c).IsAdmin,
		QuestionId:    questionIDInt,
		SubmissionValue: "all",
		FinalValue: "final",
		PageParam: 1,
		PageAction: "",
	}
	ctx := c.Request().Context()
	resp, err := q.Services.SubmissionSrvc.GetSubmissions(ctx, submissionDto)
	if err != nil {
		slogger.Debug(ctx, "showSubmissions", slogger.Err("error", err))
		// TODO: handle error
		return c.Render(http.StatusBadRequest, "submissions.html", dto.SignupResponse{})
	}


	return c.Render(http.StatusOK, "submissions.html", resp)
}

func (q *SubmissionsHndlr) ShowSubmissionsWithFilter(c echo.Context) error {
	userId := serde.GetCurrentUser(c).UserId

	ctx := c.Request().Context()

	var submissionDto dto.SubmissionRequest
	if err := c.Bind(&submissionDto); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Unable to bind form data")
	}
	submissionDto.UserId = uint(userId)
	submissionDto.IsAdmin = serde.GetCurrentUser(c).IsAdmin
	questionID := c.Param("question_id")
	questionIDInt, _ := strconv.Atoi(questionID)
	submissionDto.QuestionId = questionIDInt
	
	resp, err := q.Services.SubmissionSrvc.GetSubmissions(ctx, submissionDto)
	if err != nil {
		slogger.Debug(ctx, "showSubmissions", slogger.Err("error", err))
	}
	return c.Render(http.StatusOK, "submissions.html", resp)
}