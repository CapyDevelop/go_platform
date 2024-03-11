package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"log"
	"time"
)

type Config struct {
	Logger Logger
	GRPC   GRPC
}

type Logger struct {
	Level    string `example:"DEBUG|INFO|WARN|ERROR" env:"LOG_LEVEL" env-default:"INFO"`
	Encoding string `example:"console|json" env:"LOG_ENCODING" env-default:"json"`
}

type GRPC struct {
	Host    string        `env-required:"true" env:"GRPC_HOST"`
	Port    int           `env-required:"true" env:"GRPC_PORT"`
	Timeout time.Duration `env:"GRPC_TIMEOUT"`
}

func MustLoad() *Config {
	cfg := &Config{}

	err := cleanenv.ReadEnv(cfg)
	if err != nil {
		log.Fatalf("cannot read env: %v", err)
	}

	return cfg
}
