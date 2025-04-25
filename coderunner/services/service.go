package services

import "github.com/SUT-technology/judgino/coderunner/config"

type RunnerServices struct {
	cfg config.Config
}

func NewRunnerService(cfg config.Config) RunnerServices {
	return RunnerServices{
		cfg: cfg,
	}
}


