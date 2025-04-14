package htmltmp

import (
	"fmt"
	"net/http"

	"github.com/SUT-technology/judgino/internal/domain/service"
	"github.com/SUT-technology/judgino/internal/interface/config"
)

type Server struct {
}

func NewServer(srvc service.Service, cfg config.Server) *Server {
	m := newMiddlewares(cfg)

	// Create your main handler (ServeMux)
	mux := http.NewServeMux()

	registerRoutes(mux, srvc, m)

	handlerWithMiddleware := use(mux, []Middleware{
		m.loggingMiddleware, // Apply the global logging middleware
	})
	// Start the server
	fmt.Println("Server running at :", cfg.Port)
	port := ":" + cfg.Port
	http.ListenAndServe(port, handlerWithMiddleware)

	return &Server{}
}
