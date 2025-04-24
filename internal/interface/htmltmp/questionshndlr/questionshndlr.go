package questionshndlr

import (
	"net/http"

	"github.com/SUT-technology/judgino/internal/domain/dto"
	"github.com/SUT-technology/judgino/internal/domain/service"
	"github.com/SUT-technology/judgino/internal/interface/htmltmp/serde"
	"github.com/SUT-technology/judgino/pkg/slogger"
	"github.com/labstack/echo/v4"
	"github.com/SUT-technology/judgino/internal/domain/model"
)

type QuestionsHndlr struct {
	Services service.Service
}

func New(g *echo.Group, srvc service.Service) QuestionsHndlr {
	handler := QuestionsHndlr{
		Services: srvc,
	}

	g.GET("/", handler.ShowQuestions)
	g.GET("", handler.ShowQuestions)
	g.POST("/create",handler.createQuestions)


	return handler
}

func (q *QuestionsHndlr) createQuestions(c echo.Context) error {

	userId := serde.GetCurrentUser(c).UserId
	
	ctx := c.Request().Context()

	req, err := serde.BindRequestBody[dto.CreateQuestionRequest](c)
	if err != nil {
		slogger.Debug(ctx, "bad request", slogger.Err("error", err))
		return serde.Response(c, http.StatusBadRequest, model.BadRequestMessage, nil)
	}


	resp, err := q.Services.QuestionsSrvc.CreateQuestion(ctx, req,userId)
	if err != nil {
		slogger.Debug(ctx, "showQuestions", slogger.Err("error", err))
		// TODO: handle error
		return c.Render(http.StatusBadRequest, "create-question", resp)
	}
	
}


func (q *QuestionsHndlr) ShowQuestions(c echo.Context) error {

	userId := serde.GetCurrentUser(c).UserId

	ctx := c.Request().Context()

	req, err := serde.BindRequestBody[dto.QuestionSummeryRequest](c)
	if err != nil {
		slogger.Debug(ctx, "bad request", slogger.Err("error", err))
		return serde.Response(c, http.StatusBadRequest, model.BadRequestMessage, nil)
	}


	resp, err := q.Services.QuestionsSrvc.GetQuestions(ctx, req, uint(userId))
	if err != nil {
		slogger.Debug(ctx, "showQuestions", slogger.Err("error", err))
		// TODO: handle error
		return c.Render(http.StatusBadRequest, "questions.html", dto.QuestionsSummeryResponse{Error: err})
	}


	return c.Render(http.StatusOK, "questions.html", resp)
}

