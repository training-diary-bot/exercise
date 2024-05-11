package grpc

import "time"

type ServerConfig struct {
	Port    int           `yaml:"port" validate:"required"`
	Timeout time.Duration `yaml:"timeout" validate:"required"`
}
