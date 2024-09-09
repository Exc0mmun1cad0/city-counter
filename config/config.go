package config

import (
	"log"
	"os"
	"sync"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	HTTPServer `yaml:"http_server"`
	Redis      `yaml:"redis"`
	Postgres   `yaml:""`
}

// HTTP server configuration
type HTTPServer struct {
	Addr        string        `yaml:"addr" env-default:"localhost:8080"`
	ReadTimeout time.Duration `yaml:"read_timeout" env-default:"4s"`
	IdleTimeout time.Duration `yaml:"idle_timeout" env-default:"60s"`
}

// Redis configuraton
type Redis struct {
	Addr     string `yaml:"addr" env-default:"localhost:6379"`
	Password string `yaml:"password" env-default:""`
	DB       int    `yaml:"db" env-default:"0"`
}

// Postgres configuration
type Postgres struct {
	Host     string `yaml:"host" env-default:"localhost"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	DBName   string `yaml:"db_name" env-default:"postgres"`
	SSLMode  string `yaml:"ssl_mode" env-default:"false"`
	Port     string `yaml:"port" env-default:"5432"`
}

var (
	config Config
	once   sync.Once
)

func MustLoad() *Config {
	once.Do(func() {
		configPath := os.Getenv("CONFIG_PATH")
		if configPath == "" {
			log.Fatal("CONFIG_PATH is not set")
		}

		if _, err := os.Stat(configPath); os.IsNotExist(err) {
			log.Fatalf("config file doesn't exist: %s", err)
		}

		if err := cleanenv.ReadConfig(configPath, &config); err != nil {
			log.Fatalf("cannot read config: %s", err)
		}
	})

	return &config
}
