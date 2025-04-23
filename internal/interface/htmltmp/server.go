package htmltmp

import (
	"context"
	"errors"
	"fmt"
	"html/template"
	"io"
	"log/slog"
	"slices"
	"strings"
	"time"

	"github.com/SUT-technology/judgino/internal/domain/service"
	"github.com/SUT-technology/judgino/internal/interface/config"
	"github.com/SUT-technology/judgino/pkg/slogger"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type Server struct {
	srv    *echo.Echo
	defers []func()
}

func NewServer(srvc service.Service, cfg config.Server) *Server {
	var dfrs []func()

	e := echo.New()
	e.Debug = true
	e.HideBanner = true
	e.HidePort = true
	e.Static("/static", "static")
	e.Validator = &Validator{validator: validator.New()}

	closer := func() {
		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()
		err := e.Shutdown(ctx)
		if err != nil {
			slog.Error("close HTTP server", slogger.Err("error", err))
		}
	}
	dfrs = append(dfrs, closer)

	// manage middlewares
	var middleware []echo.MiddlewareFunc
	m := newMiddlewares(cfg)

	if cfg.Logger {
		middleware = append(middleware, m.loggerMiddleware)
	}

	// application specific middlewares
	middleware = append(middleware, m.corsMiddleware())

	// default recover middleware
	middleware = append(middleware, m.recoverMiddleware)

	// applying middlewares and create a new server
	e.Use(middleware...)

	e.Renderer = &TemplateRenderer{
		templates: template.Must(template.ParseGlob("./templates/*.html")),
	}

	register(e, srvc, m)

	return &Server{srv: e, defers: dfrs}
}

func (s *Server) Start(addr string) error {
	return s.srv.Start(addr)
}

func (s *Server) Stop() {
	for _, f := range slices.Backward(s.defers) {
		f()
	}
}

type Validator struct {
	validator *validator.Validate
}

func (v *Validator) Validate(i interface{}) error {
	if err := v.validator.Struct(i); err != nil {
		if ve, ok := err.(validator.ValidationErrors); ok {
			var errMsgs []string
			for _, fieldErr := range ve {
				errMsgs = append(errMsgs, fmt.Sprintf("%s failed on '%s'", fieldErr.Field(), fieldErr.Tag()))
			}
			return errors.New(strings.Join(errMsgs, ", "))
		}
		return errors.New("validation failed with unknown error")
	}
	return nil
}

type TemplateRenderer struct {
	templates *template.Template
}

func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}
