package questionshndlr

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

type QuestionsHndlr struct {
	Services service.Service
}

func New(g *echo.Group, srvc service.Service, m echo.MiddlewareFunc) QuestionsHndlr {
	handler := QuestionsHndlr{
		Services: srvc,
	}

	g.GET("/", handler.ShowQuestions)
	g.GET("", handler.ShowQuestions)
	g.GET("/:question_id", handler.ShowQuestion)

	g.POST("/published/:question_id", handler.PublishQuestion, m)

	// g.POST("/", handler.ShowQuestionsByFilter)
	// g.POST("", handler.ShowQuestionsByFilter)

	return handler
}

func (q *QuestionsHndlr) ShowQuestions(c echo.Context) error {

	userId := serde.GetCurrentUser(c).UserId

	ctx := c.Request().Context()

	req, err := serde.BindRequestBody[dto.QuestionRequest](c)
	if err != nil {
		slogger.Debug(ctx, "bad request", slogger.Err("error", err))
		return serde.Response(c, http.StatusBadRequest, model.BadRequestMessage, nil)
	}

	resp, err := q.Services.QuestionsSrvc.GetQuestions(ctx, req, uint(userId))
	if err != nil {
		slogger.Debug(ctx, "showQuestions", slogger.Err("error", err))
		// TODO: handle error
		return c.Render(http.StatusBadRequest, "questions.html", dto.QuestionsResponse{Error: err})
	}

	return c.Render(http.StatusOK, "questions.html", resp)
}

func (q *QuestionsHndlr) PublishQuestion(c echo.Context) error {
	userData := serde.GetCurrentUser(c)

	if !userData.IsAdmin {
		return c.Redirect(http.StatusMovedPermanently, "/auth")
	}
	ctx := c.Request().Context()
	questionID := c.Param("question_id")
	questionIDInt, _ := strconv.Atoi(questionID)

	err := q.Services.QuestionsSrvc.PublishQuestion(ctx, uint(questionIDInt))
	if err != nil {
		slogger.Debug(ctx, "showQuestions", slogger.Err("error", err))
		return c.Render(http.StatusBadRequest, "questions.html", dto.PublishResponse{Msg: err.Error()})
	}

	return c.Render(http.StatusOK, "questions.html", nil)
}

func (q *QuestionsHndlr) ShowQuestion(c echo.Context) error {
	questionID := c.Param("question_id")
	questionIDInt, _ := strconv.Atoi(questionID)
	ctx := c.Request().Context()

	resp, err := q.Services.QuestionsSrvc.GetQuestion(ctx, uint(questionIDInt))
	if err != nil {
		slogger.Debug(ctx, "showQuestion", slogger.Err("error", err))
		// TODO: handle error
		return c.Render(http.StatusBadRequest, "question.html", nil)
	}

	return c.Render(http.StatusOK, "question.html", resp)
}
