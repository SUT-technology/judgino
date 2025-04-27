package config

import (
	"time"

	"github.com/go-playground/validator/v10"
)

type Config struct {
	DB         DB         `yaml:"db"`
	Logger     Logger     `yaml:"logger"`
	Env        string     `yaml:"env"`
	Server     Server     `yaml:"server"`
	CodeRunner CodeRunner `yaml:"code-runner"`
}

type CodeRunner struct {
	ApiKey string `yaml:"api-key"`
}
type RateLimiter struct {
	Enabled bool          `yaml:"enabled"`
	Rate    int           `yaml:"rate"`
	Burst   int           `yaml:"burst"`
	Expires time.Duration `yaml:"expires"`
}
type Server struct {
	Port        string      `yaml:port`
	SecretKey   string      `yaml:"secret_key" validate:"required"`
	Logger      bool        `yaml:"logger" validate:"required"`
	RateLimiter RateLimiter `yaml:"rate-limiter"`
	Addr        string      `yaml:"addr"`
}

type DB struct {
	Port     string `yaml:"port" validate:"required"`
	DBName   string `yaml:"db_name" validate:"required"`
	Host     string `yaml:"host" validate:"required"`
	Password string `yaml:"password" validate:"required"`
	Username string `yaml:"username" validate:"required"`
}

type Logger struct {
	Level string `yaml:"level" validate:"required,oneof=trace debug info warn error fatal"`
}

func (c Config) Validate() error {
	v := validator.New()
	return v.Struct(c)
}
