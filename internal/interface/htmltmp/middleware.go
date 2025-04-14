package htmltmp

import (
	"context"
	"log"
	"net/http"

	"github.com/SUT-technology/judgino/internal/interface/config"
)

type middlewares struct {
	cfg config.Server
}

func newMiddlewares(cfg config.Server) *middlewares {
	return &middlewares{cfg: cfg}
}

type Middleware func(http.Handler) http.Handler

func use(handler http.Handler, middleware []Middleware) http.Handler {
	for _, m := range middleware {
		handler = m(handler)
	}
	return handler
}

func (m *middlewares) loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Request: %s %s", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}

func (m *middlewares) JWTMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Add data to the request context
		// Example data
		ctx := context.WithValue(r.Context(), "userId", 1)

		// Pass the request with the updated context to the next handler
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
