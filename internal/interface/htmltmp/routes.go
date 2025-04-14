package htmltmp

import (
	"net/http"

	"github.com/SUT-technology/judgino/internal/domain/service"
)

type Group struct {
	prefix     string
	middleware Middleware
	mux        *http.ServeMux
}

// Function to create a new group
func NewGroup(prefix string, middleware Middleware, mux *http.ServeMux) *Group {
	return &Group{
		prefix:     prefix,
		middleware: middleware,
		mux:        mux,
	}
}

func (g *Group) Handle(method string, path string, handlerFunc http.HandlerFunc) {
	finalHandler := g.middleware(http.HandlerFunc(handlerFunc))
	g.mux.Handle(g.prefix+path, finalHandler)
}

func registerRoutes(mux *http.ServeMux, srvc service.Service, middlewares *middlewares) {
	// Create groups with middleware
	// swaggerGroup := NewGroup("/swagger", middlewares.loggingMiddleware, mux)
	authGroup := NewGroup("/auth", middlewares.JWTMiddleware, mux)

	// Register routes within groups
	NewAuthHndlr(authGroup, srvc)
}
