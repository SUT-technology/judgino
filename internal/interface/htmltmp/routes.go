package htmltmp

import (
	"github.com/SUT-technology/judgino/internal/domain/service"
	"github.com/SUT-technology/judgino/internal/interface/htmltmp/authhdnlr"
	"github.com/labstack/echo/v4"
)

func register(e *echo.Echo, srvc service.Service, m *middlewares) {

func registerRoutes(mux *http.ServeMux, srvc service.Service, middlewares *middlewares) {
	// Create groups with middleware
	// swaggerGroup := NewGroup("/swagger", middlewares.loggingMiddleware, mux

	profileGroup := NewGroup("/profile", middlewares.JWTMiddleware, mux)
	NewProfileHndlr(profileGroup, srvc)
	auth := e.Group("/auth", m.JWTMiddleware)

	authhdnlr.New(auth, srvc)
}
