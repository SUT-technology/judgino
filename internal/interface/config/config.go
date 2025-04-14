package config

import "github.com/go-playground/validator/v10"

type Config struct {
	DB     DB     `yaml:"db"`
	Logger Logger `yaml:"logger"`
	Env    string `yaml:"env"`
	Server Server `yaml:"server"`
}

type Server struct {
	Port string `yaml:port`
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
