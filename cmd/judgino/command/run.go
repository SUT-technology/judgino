package command

import (
	"flag"
	"fmt"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/SUT-technology/judgino/internal/application/services"
	"github.com/SUT-technology/judgino/internal/infrastructure/database/postgresql/pool"
	"github.com/SUT-technology/judgino/internal/interface/config"
	"github.com/SUT-technology/judgino/internal/interface/htmltmp"
	"github.com/SUT-technology/judgino/pkg/slogger"
)

func Run() error {
	var configPath string
	flag.StringVar(&configPath, "cfg", "assets/config/development.yaml", "Configuration File")
	flag.Parse()
	c, err := config.Load(configPath)
	if err != nil {
		return fmt.Errorf("reading config: %w", err)
	}
	logger := slogger.NewJSONLogger(c.Logger.Level, os.Stdout)
	slog.SetDefault(logger)

	db, err := pool.New(c.DB)
	if err != nil {
		fmt.Printf("Error initializing database connection: %v\n", err)
	}

	slog.Debug("initial db")

	srvc := services.New(db , c)

	httpSrv := htmltmp.NewServer(srvc, c.Server)
	defer func() {
		slog.Debug("gracefully stopping HTTP server")
		httpSrv.Stop()
	}()

	startErr := make(chan error)
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		slog.Info("starting HTTP server", slog.String("address", c.Server.Addr))
		err := httpSrv.Start(c.Server.Addr)
		startErr <- fmt.Errorf("HTTP server startup: %w", err)
	}()

	select {
	case err := <-startErr:
		slog.Error("failed to start server", slog.Any("error", err))
		return err
	case <-quit:
		slog.Info("received signal to stop server")
		return nil
	}

}
