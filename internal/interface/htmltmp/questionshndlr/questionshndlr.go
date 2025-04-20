package questionshndlr

import (
	"net/http"
	"fmt"

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
	// g.POST("/", handler.ShowQuestionsByFilter)
	// g.POST("", handler.ShowQuestionsByFilter)


	return handler
}



func (q *QuestionsHndlr) ShowQuestions(c echo.Context) error {

	userId := serde.GetCurrentUser(c).UserId

	ctx := c.Request().Context()
	
	// var questionsDto dto.QuestionRequest
	// if err := c.Bind(&questionsDto); err != nil {
	// 	return echo.NewHTTPError(http.StatusBadRequest, "Unable to bind form data")
	// }
	// fmt.Printf("%+v", questionsDto)

	req, err := serde.BindRequestBody[dto.QuestionRequest](c)
	if err != nil {
		slogger.Debug(ctx, "bad request", slogger.Err("error", err))
		return serde.Response(c, http.StatusBadRequest, model.BadRequestMessage, nil)
	}
	fmt.Printf("req: %+v", req)


	resp, err := q.Services.QuestionsSrvc.GetQuestions(ctx, req, uint(userId))
	if err != nil {
		slogger.Debug(ctx, "showQuestions", slogger.Err("error", err))
		// TODO: handle error
		return c.Render(http.StatusBadRequest, "questions.html", dto.QuestionsResponse{Error: err})
	}


	return c.Render(http.StatusOK, "questions.html", resp)
}

