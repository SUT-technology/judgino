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
	g.GET("/create", handler.createQuestion)
	g.POST("/draft", handler.draftQuestion)
	g.GET("/edit/:question_id",handler.editQuestion)
	g.POST("/update/:question_id",handler.updateQuestion)
	g.GET("/:question_id", handler.ShowQuestion)
	g.GET("/published/:question_id", handler.StateQuestion)
	return handler
}




func (q *QuestionsHndlr) editQuestion(c echo.Context) error {
	questionID := c.Param("question_id")
	questionIDInt, _ := strconv.Atoi(questionID)
	response,err := q.Services.QuestionsSrvc.EditQuestion(c.Request().Context(),uint(questionIDInt))
	if err != nil {
		slogger.Debug(c.Request().Context(), "bad request", slogger.Err("error", err))
		return serde.Response(c, http.StatusBadRequest, model.BadRequestMessage, nil)
	}
	return c.Render(http.StatusOK,"edit-question.html",response)
}





func (q *QuestionsHndlr) updateQuestion(c echo.Context) error {
	
	ctx := c.Request().Context()

	questionID := c.Param("question_id")
	questionIDInt, _ := strconv.Atoi(questionID)
	currentUserId := serde.GetCurrentUser(c).UserId

	req, err := serde.BindRequestBody[dto.UpdateQuestionRequest](c)
	// fmt.Printf("request: %v",req)
	if err != nil {
		slogger.Debug(ctx, "bad request", slogger.Err("error", err))
		return serde.Response(c, http.StatusBadRequest, model.BadRequestMessage, nil)
	}

	err = q.Services.QuestionsSrvc.UpdateQuestion(ctx,uint(questionIDInt),req)
	if err != nil {
		slogger.Debug(c.Request().Context(), "bad request", slogger.Err("error", err))
		return serde.Response(c, http.StatusBadRequest, model.BadRequestMessage, nil)
	}
	data := dto.CreateQuestionResponse{UserID: currentUserId}
	return c.Render(http.StatusOK, "create-question.html", data)
}





func (q *QuestionsHndlr) createQuestion(c echo.Context) error {
	slogger.Debug(c.Request().Context(), "Creating a new question...")
	currentUserId := serde.GetCurrentUser(c).UserId
	// currentUserId := int64(1)
	data := dto.CreateQuestionResponse{UserID: currentUserId}
	return c.Render(http.StatusOK, "create-question.html", data)
}





func (q *QuestionsHndlr) draftQuestion(c echo.Context) error {
	currentUserId := serde.GetCurrentUser(c).UserId
	// currentUserId := int64(1)

	ctx := c.Request().Context()

	req, err := serde.BindRequestBody[dto.CreateQuestionRequest](c)
	// fmt.Printf("request: %v",req)
	if err != nil {
		slogger.Debug(ctx, "bad request", slogger.Err("error", err))
		return serde.Response(c, http.StatusBadRequest, model.BadRequestMessage, nil)
	}

	resp, err := q.Services.QuestionsSrvc.CreateQuestion(ctx, req, currentUserId)
	if err != nil {
		slogger.Debug(ctx, "create_question_service_error", slogger.Err("error", err))
		return c.Render(http.StatusInternalServerError, "create-question.html", resp)
	}

	if resp.Error {
		return c.Render(http.StatusOK, "create-question.html", resp)
	}

	resp2, err := q.Services.QuestionsSrvc.GetQuestion(ctx, currentUserId, uint(resp.QuestionID))
	if err != nil {
		slogger.Debug(ctx, "showQuestion", slogger.Err("error", err))
		return c.Render(http.StatusBadRequest, "question.html", nil)
	}

	return c.Render(http.StatusOK, "question.html", resp2)
}




func (q *QuestionsHndlr) ShowQuestions(c echo.Context) error {

	currentUserId := serde.GetCurrentUser(c).UserId
	// currentUserId := 1

	ctx := c.Request().Context()

	req, err := serde.BindRequestBody[dto.QuestionSummeryRequest](c)
	if err != nil {
		slogger.Debug(ctx, "bad request", slogger.Err("error", err))
		return serde.Response(c, http.StatusBadRequest, model.BadRequestMessage, nil)
	}

	resp, err := q.Services.QuestionsSrvc.GetQuestions(ctx, req, uint(currentUserId))
	if err != nil {
		slogger.Debug(ctx, "showQuestions", slogger.Err("error", err))
		// TODO: handle error
		return c.Render(http.StatusBadRequest, "questions.html", dto.QuestionsSummeryResponse{Error: err})
	}

	return c.Render(http.StatusOK, "questions.html", resp)
}





func (q *QuestionsHndlr) StateQuestion(c echo.Context) error {
	currentUser := serde.GetCurrentUser(c)
	// CurrentUserId := 1


	if !currentUser.IsAdmin {
		return c.Redirect(http.StatusMovedPermanently, "/auth")
	}
	ctx := c.Request().Context()
	questionID := c.Param("question_id")
	questionIDInt, _ := strconv.Atoi(questionID)

	err := q.Services.QuestionsSrvc.StateQuestion(ctx, uint(questionIDInt))
	if err != nil {
		slogger.Debug(ctx, "showQuestions", slogger.Err("error", err))
		return c.Render(http.StatusBadRequest, "questions.html", dto.PublishResponse{Msg: err.Error()})
	}

	resp, err := q.Services.QuestionsSrvc.GetQuestions(ctx, dto.QuestionSummeryRequest{}, uint(currentUser.UserId))
	if err != nil {
		slogger.Debug(ctx, "showQuestions", slogger.Err("error", err))
		// TODO: handle error
		return c.Render(http.StatusBadRequest, "questions.html", dto.QuestionsSummeryResponse{Error: err})
	}

	return c.Render(http.StatusOK, "questions.html", resp)
}




func (q *QuestionsHndlr) ShowQuestion(c echo.Context) error {

	currentUserId := serde.GetCurrentUser(c).UserId
	// userId := 1

	questionID := c.Param("question_id")
	questionIDInt, _ := strconv.Atoi(questionID)
	ctx := c.Request().Context()

	resp, err := q.Services.QuestionsSrvc.GetQuestion(ctx, int64(currentUserId), uint(questionIDInt))
	if err != nil {
		slogger.Debug(ctx, "showQuestion", slogger.Err("error", err))
		// TODO: handle error
		return c.Render(http.StatusBadRequest, "question.html", nil)
	}

	return c.Render(http.StatusOK, "question.html", resp)
}
