package config

import (
	"exercise/pkg/database"
	"exercise/pkg/grpc"
	"github.com/go-playground/validator/v10"
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/pkg/errors"
	"os"
)

const (
	confugPathEnvName     = "CONFIG_PATH"
	defaultConfigYAMLPath = "./config/application.yaml"
)

type Cfg struct {
	Env        string            `yaml:"env" validate:"required"`
	GRPCServer grpc.ServerConfig `yaml:"grpc_server" validate:"required"`
	PG         database.DBCfg    `yaml:"pg" env-prefix:"PG_" validate:"required"`
}

func Setup() (*Cfg, error) {
	configPath := os.Getenv(confugPathEnvName)

	if configPath == "" {
		configPath = defaultConfigYAMLPath
	}

	var cfg Cfg
	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		return nil, errors.Wrap(err, "failed to read application config")
	}

	if err := cfg.Validate(); err != nil {
		return nil, err
	}

	return &cfg, nil
}

func (cfg *Cfg) Validate() error {
	validate := validator.New()
	if err := validate.Struct(cfg); err != nil {
		return errors.Wrap(err, "fail validation")
	}

	return nil
}
