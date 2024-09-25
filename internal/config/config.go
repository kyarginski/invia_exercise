package config

import (
	"log"
	"os"
	"strconv"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Env            string `yaml:"env" env-default:"local"`
	Version        string `yaml:"version" env-default:"unknown"`
	Port           int    `yaml:"port" env-default:""`
	DBConnect      string `yaml:"db_connect" env-default:""`
	UseTracing     bool   `yaml:"use_tracing"`
	TracingAddress string `yaml:"tracing_address" env-default:""`
}

func MustLoad() *Config {
	configPath := os.Getenv("INVIA_CONFIG_PATH")
	if configPath == "" {
		configPath = "config/local.yaml"
	}
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("config file does not exist: %s", configPath)
	}

	var cfg Config

	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		log.Fatalf("cannot read config: %s", err)
	}

	portEnv := os.Getenv("INVIA_PORT")
	if portEnv != "" {
		newPort, err := strconv.Atoi(portEnv)
		if err == nil {
			cfg.Port = newPort
		}
	}

	return &cfg
}
