package questionshndlr

import (
	"net/http"

	"github.com/SUT-technology/judgino/internal/domain/dto"
	"github.com/SUT-technology/judgino/internal/domain/service"
	"github.com/SUT-technology/judgino/internal/interface/htmltmp/serde"
	"github.com/SUT-technology/judgino/pkg/slogger"
	"github.com/labstack/echo/v4"
)

type QuestionsHndlr struct {
	Services service.Service
}

func New(g *echo.Group, srvc service.Service) QuestionsHndlr {
	handler := QuestionsHndlr{
		Services: srvc,
	}

	g.GET("/", handler.ShowQuestions)
	g.POST("/", handler.ShowQuestionsByFilter)
	g.POST("", handler.ShowQuestionsByFilter)


	return handler
}



func (q *QuestionsHndlr) ShowQuestions(c echo.Context) error {

	userId := serde.GetCurrentUser(c).UserId

	ctx := c.Request().Context()
	
	questionsDto := dto.QuestionRequest{
		UserId:       uint(userId),
		SearchFilter:  "",
		QuestionValue: "all",
		SortValue:    "publish_date",
		PageParam: 1,
	}

	resp, err := q.Services.QuestionsSrvc.GetQuestions(ctx, questionsDto)
	if err != nil {
		slogger.Debug(ctx, "showQuestions", slogger.Err("error", err))
		// TODO: handle error
		return c.Render(http.StatusBadRequest, "questions.html", dto.SignupResponse{})
	}


	return c.Render(http.StatusOK, "questions.html", resp)
}

func (q *QuestionsHndlr) ShowQuestionsByFilter(c echo.Context) error {
	userId := serde.GetCurrentUser(c).UserId

	ctx := c.Request().Context()

	var questionsDto dto.QuestionRequest
	if err := c.Bind(&questionsDto); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Unable to bind form data")
	}
	questionsDto.UserId = uint(userId)
	
	resp, err := q.Services.QuestionsSrvc.GetQuestions(ctx, questionsDto)
	if err != nil {
		slogger.Debug(ctx, "showQuestions", slogger.Err("error", err))
	}
	return c.Render(http.StatusOK, "questions.html", resp)
}

