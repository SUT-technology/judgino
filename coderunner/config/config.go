package config

import (
	"fmt"
	"os"
	"time"

	"gopkg.in/yaml.v3"
)

type Config struct {
	ApiKey       string `yaml:"api-key"`
	ApiUrl       string `yaml:"api-url"`
	TimeInterval time.Duration    `yaml:"time-interval"`
}

func Load(path string) (Config, error) {
	f, err := os.ReadFile(path)
	if err != nil {
		return Config{}, fmt.Errorf("reading file: %w", err)
	}

	c, err := Parse(f)
	if err != nil {
		return Config{}, fmt.Errorf("parsing configs: %w", err)
	}

	return c, nil
}

// Parse reads the yaml data into a Config struct. It does not perform any validations on the configurations themselves.
func Parse(data []byte) (Config, error) {
	c := Config{}
	err := yaml.Unmarshal(data, &c)
	if err != nil {
		return Config{}, fmt.Errorf("parsing yaml file: %w", err)
	}
	return c, nil
}
