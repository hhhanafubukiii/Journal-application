package config

import (
	"errors"
	"github.com/ilyakaznacheev/cleanenv"
	"os"
	"time"
)

type Config struct {
	Env         string     `yaml:"env" env-default:"local" env-required:"true"`
	databaseURL string     `yaml:"database_url" env-required:"true"`
	Server      HTTPServer `yaml:"http_server" env-required:"true"`
}

type HTTPServer struct {
	Address     string        `yaml:"address" env-default:"localhost:8080"`
	Timeout     time.Duration `yaml:"timeout" env-default:"4s"`
	IdleTimeout time.Duration `yaml:"idle_timeout" env-default:"60s"`
}

func NewConfig() (*Config, error) {
	var config Config
	configPath := "/Users/artyom/Projects/WiseNote/backend/config/local.yaml"

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		return nil, errors.New("CONFIG_PATH does not exist")
	}

	if err := cleanenv.ReadConfig(configPath, &config); err != nil {
		return nil, err
	}

	return &config, nil
}
