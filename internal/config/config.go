package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Env           string        `yaml:"env" env-default:"local"`
	StorageConfig StorageConfig `yaml:"storage"`
	HTTPServer    HTTPServer    `yaml:"http_server"`
	ReferalConfig ReferalConfig `yaml:"referals_config"`
}

type HTTPServer struct {
	Address     string        `yaml:"address" env-default:"localhost:8080"`
	Timeout     time.Duration `yaml:"timeout" env-default:"4s"`
	IdleTimeout time.Duration `yaml:"idle_timeout" env-default:"60s"`
	User        string        `yaml:"user" env-required:"true"`
	Password    string        `yaml:"password" env-required:"true" env:"HTTP_SERVER_PASSWORD"`
}

type StorageConfig struct {
	Host         string        `yaml:"host"`
	Port         string        `yaml:"port"`
	Database     string        `yaml:"database"`
	User         string        `yaml:"user"`
	Password     string        `yaml:"password"`
	MaxRetry     int           `yaml:"max_retry"`
	RetryTimeout time.Duration `yaml:"retry_timeout"`
}

type ReferalConfig struct {
	Winners   ReferalWinnersConfig   `yaml:"winners"`
	Statistic ReferalStatisticConfig `yaml:"statistic"`
}

type ReferalWinnersConfig struct {
	Limit    int `yaml:"default_limit"`
	Interval int `yaml:"default_interval"`
}

type ReferalStatisticConfig struct {
	Interval int `yaml:"default_interval"`
}

// NewConfig ctor
func NewConfig() *Config {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		log.Fatal("CONFIG_PATH is not set")
	}

	// check if file exists
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("config file does not exist: %s", configPath)
	}

	var cfg Config

	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		log.Fatalf("cannot read config: %s", err)
	}

	return &cfg
}
