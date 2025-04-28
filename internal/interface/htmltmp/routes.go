package htmltmp

import (
	"github.com/SUT-technology/judgino/internal/domain/service"
	"github.com/SUT-technology/judgino/internal/interface/htmltmp/authhdnlr"
	"github.com/SUT-technology/judgino/internal/interface/htmltmp/profhndlr"
	"github.com/SUT-technology/judgino/internal/interface/htmltmp/questionshndlr"
	"github.com/SUT-technology/judgino/internal/interface/htmltmp/runnerhndlr"
	"github.com/SUT-technology/judgino/internal/interface/htmltmp/submissionshndlr"
	"github.com/SUT-technology/judgino/internal/interface/htmltmp/logouthndlr"
	"github.com/labstack/echo/v4"
)

func register(e *echo.Echo, srvc service.Service, m *middlewares) {
	// Create groups with middleware
	// swaggerGroup := NewGroup("/swagger", middlewares.loggingMiddleware, mux

	prof := e.Group("/profile", m.JWTMiddleware)
	profhndlr.New(prof, srvc)

	auth := e.Group("/auth")
	// Todo change middleware
	questions := e.Group("/questions", m.JWTMiddleware)
	submissions := e.Group("/submissions", m.JWTMiddleware)
	logout := e.Group("/logout", m.JWTMiddleware)

	runner := e.Group("/api/runner")

	authhdnlr.New(auth, srvc)
	questionshndlr.New(questions, srvc, m.JWTMiddleware)
	submissionshndlr.New(submissions, srvc)
	logouthndlr.New(logout, srvc)

	// prof := e.Group("/profile")
	// profhndlr.New(prof, srvc)

	// auth := e.Group("/auth")
	// // Todo change middleware
	// questions := e.Group("/questions")
	// submissions := e.Group("/submissions")

	// authhdnlr.New(auth, srvc)
	// questionshndlr.New(questions, srvc)
	// submissionshndlr.New(submissions, srvc)
	runnerhndlr.New(runner, srvc)
}
