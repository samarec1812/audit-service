package config

import (
	"log"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"

	ll "github.com/samarec1812/audit-service/internal/pkg/logger"
)

type Config struct {
	Env        ll.LogLevel `yaml:"env" env-default:"local"`
	HTTPServer `yaml:"http_server"`
	DBConfig   `yaml:"db"`
}

type HTTPServer struct {
	Address     string        `yaml:"address" env-default:"localhost:8080"`
	Timeout     time.Duration `yaml:"timeout" env-default:"6s"`
	IdleTimeout time.Duration `yaml:"idle_timeout" env-default:"60s"`
}

type DBConfig struct {
	DatabaseURL string `yaml:"url" env-required:"true"`
}

func MustLoad() *Config {
	configPath := os.Getenv("CONFIG_PATH")

	if configPath == "" {
		log.Fatal("CONFIG_PATH is not set")
	}

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("config file does not exist: %s", configPath)
	}

	var cfg Config

	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		log.Fatalf("cannot read config: %s", err)
	}

	return &cfg
}
