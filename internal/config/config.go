package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Auth     Auth
	Postgres Postgres
	Redis    Redis
	Api      Api
}

type Auth struct {
	Secret string
}

type Postgres struct {
	DSN string
}

type Redis struct {
	Port     string
	Password string
}

type Api struct {
	RestPort string
	WSPort   string
}

func Init() (*Config, error) {
	var cfg Config

	if err := loadEnv(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}

func loadEnv(cfg *Config) error {
	if err := godotenv.Load(); err != nil {
		return err
	}

	cfg.Auth.Secret = os.Getenv("SECRET")
	cfg.Postgres.DSN = os.Getenv("DSN")
	cfg.Redis.Port = os.Getenv("REDIS_PORT")
	cfg.Redis.Password = os.Getenv("REDIS_PASSWORD")
	cfg.Api.RestPort = os.Getenv("REST_PORT")
	cfg.Api.WSPort = os.Getenv("WS_PORT")

	return nil
}
