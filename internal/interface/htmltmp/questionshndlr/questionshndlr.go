package questionshndlr

import (
	"fmt"
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

func New(g *echo.Group, srvc service.Service) QuestionsHndlr {
	handler := QuestionsHndlr{
		Services: srvc,
	}
	g.GET("/", handler.ShowQuestions)
	g.GET("", handler.ShowQuestions)
	g.GET("/create",handler.createQuestion)
	g.POST("/draft",handler.draftQuestion)
	g.GET("/:question_id", handler.ShowQuestion)
	g.GET("/published/:question_id", handler.PublishQuestion)

	return handler
}



func (q *QuestionsHndlr) createQuestion(c echo.Context) error {
    slogger.Debug(c.Request().Context(), "Creating a new question...")
    // UserId := serde.GetCurrentUser(c).UserId
	UserId := int64(1)
	data:=dto.CreateQuestionResponse{UserID: UserId}
	fmt.Println(data.Title)
    return c.Render(http.StatusOK, "create-question.html", data)
}



func (q *QuestionsHndlr) draftQuestion(c echo.Context) error {
	// userId := serde.GetCurrentUser(c).UserId
	userId := int64(1)

	ctx := c.Request().Context()

	req, err := serde.BindRequestBody[dto.CreateQuestionRequest](c)
	// fmt.Printf("request: %v",req)
	if err != nil {
		slogger.Debug(ctx, "bad request", slogger.Err("error", err))
		return serde.Response(c, http.StatusBadRequest, model.BadRequestMessage, nil)
	}

	resp, err := q.Services.QuestionsSrvc.CreateQuestion(ctx, req, userId)
	if err != nil {
		slogger.Debug(ctx, "create_question_service_error", slogger.Err("error", err))
		return c.Render(http.StatusInternalServerError, "create-question.html", resp)
	}


	if resp.Error {
		return c.Render(http.StatusBadRequest, "create-question.html", resp)
	}

	fmt.Printf("question id: %v",resp.QuestionID)

	return c.Redirect(http.StatusSeeOther, fmt.Sprintf("/questions/%v", resp.QuestionID))
}
 

func (q *QuestionsHndlr) ShowQuestions(c echo.Context) error {

	// userId := serde.GetCurrentUser(c).UserId
	userId := 1

	ctx := c.Request().Context()

	req, err := serde.BindRequestBody[dto.QuestionSummeryRequest](c)
	if err != nil {
		slogger.Debug(ctx, "bad request", slogger.Err("error", err))
		return serde.Response(c, http.StatusBadRequest, model.BadRequestMessage, nil)
	}

    fmt.Printf("request: %+v",req)

	resp, err := q.Services.QuestionsSrvc.GetQuestions(ctx, req, uint(userId))
	if err != nil {
		slogger.Debug(ctx, "showQuestions", slogger.Err("error", err))
		// TODO: handle error
		return c.Render(http.StatusBadRequest, "questions.html", dto.QuestionsSummeryResponse{Error: err})
	}

	fmt.Printf("resp: %+v",resp)

	return c.Render(http.StatusOK, "questions.html", resp)
}

func (q *QuestionsHndlr) PublishQuestion(c echo.Context) error {
	// currentUser := serde.GetCurrentUser(c)
	userId := 1

	// if !currentUser.IsAdmin {
	// 	return c.Redirect(http.StatusMovedPermanently, "/auth")
	// }
	ctx := c.Request().Context()
	questionID := c.Param("question_id")
	questionIDInt, _ := strconv.Atoi(questionID)

	err := q.Services.QuestionsSrvc.PublishQuestion(ctx, uint(questionIDInt))
	if err != nil {
		slogger.Debug(ctx, "showQuestions", slogger.Err("error", err))
		return c.Render(http.StatusBadRequest, "questions.html", dto.PublishResponse{Msg: err.Error()})
	}

	resp, err := q.Services.QuestionsSrvc.GetQuestions(ctx, dto.QuestionSummeryRequest{}, uint(userId))
	if err != nil {
		slogger.Debug(ctx, "showQuestions", slogger.Err("error", err))
		// TODO: handle error
		return c.Render(http.StatusBadRequest, "questions.html", dto.QuestionsSummeryResponse{Error: err})
	}

	return c.Render(http.StatusOK, "questions.html", resp)
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
