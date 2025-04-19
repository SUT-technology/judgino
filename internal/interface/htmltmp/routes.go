package htmltmp

import (
	"github.com/SUT-technology/judgino/internal/domain/service"
	"github.com/SUT-technology/judgino/internal/interface/htmltmp/authhdnlr"
	"github.com/SUT-technology/judgino/internal/interface/htmltmp/questionshndlr"
	"github.com/SUT-technology/judgino/internal/interface/htmltmp/submissionshndlr"
	"github.com/labstack/echo/v4"
)

func register(e *echo.Echo, srvc service.Service, m *middlewares) {

	auth := e.Group("/auth", m.JWTMiddleware)
	// Todo change middleware
	questions := e.Group("/questions", m.JWTMiddleware)
	submissions := e.Group("/submissions", m.JWTMiddleware)

	authhdnlr.New(auth, srvc)
	questionshndlr.New(questions, srvc)
	submissionshndlr.New(submissions, srvc)
}
