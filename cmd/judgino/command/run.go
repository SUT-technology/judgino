package command

import (
	"flag"
	"fmt"

	"github.com/SUT-technology/judgino/internal/application/services"
	"github.com/SUT-technology/judgino/internal/infrastructure/database/postgresql/pool"
	"github.com/SUT-technology/judgino/internal/interface/config"
	"github.com/SUT-technology/judgino/internal/interface/htmltmp"
)

func Run() error {
	var configPath string
	flag.StringVar(&configPath, "cfg", "assets/config/development.yaml", "Configuration File")
	flag.Parse()
	c, err := config.Load(configPath)
	if err != nil {
		return fmt.Errorf("reading config: %w", err)
	}

	db, err := pool.New(c.DB)
	if err != nil {
		fmt.Printf("Error initializing database connection: %v\n", err)
	}

	srvc := services.New(db)

	_ = htmltmp.NewServer(srvc, c.Server)

	return nil
}
