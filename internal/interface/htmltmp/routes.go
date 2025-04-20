package htmltmp

import (
	"github.com/SUT-technology/judgino/internal/domain/service"
	"github.com/SUT-technology/judgino/internal/interface/htmltmp/authhdnlr"
	profhndlr "github.com/SUT-technology/judgino/internal/interface/htmltmp/profhndlr"
	"github.com/labstack/echo/v4"
)

func register(e *echo.Echo, srvc service.Service, m *middlewares) {
	// Create groups with middleware
	// swaggerGroup := NewGroup("/swagger", middlewares.loggingMiddleware, mux

	prof := e.Group("/profile", m.JWTMiddleware)
	profhndlr.New(prof, srvc)

	auth := e.Group("/auth", m.JWTMiddleware)
	authhdnlr.New(auth, srvc)
}
