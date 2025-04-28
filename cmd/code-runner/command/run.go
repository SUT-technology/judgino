package command

import (
	"flag"
	"fmt"


	"github.com/SUT-technology/judgino/coderunner"
	"github.com/SUT-technology/judgino/coderunner/config"
	"github.com/SUT-technology/judgino/coderunner/services"
)






func Run() error {
	var configPath string
	flag.StringVar(&configPath, "cfg", "assets/config/code-runner.yaml", "Configuration File")
	flag.Parse()
	c, err := config.Load(configPath)
	if err != nil {
		return fmt.Errorf("reading config: %w", err)
	}
	srvc := services.NewRunnerService(c)
	coderunner.Start(srvc)

	return nil
}
