package config

import (
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
)

type DATABASE struct {
	Username       string `env:"DATABASE_USERNAME" env-default:"postgres"`
	Password       string `env:"DATABASE_PASSWORD" env-default:"postgres"`
	Host           string `env:"DATABASE_HOST" env-default:"localhost"`
	Port           string `env:"DATABASE_PORT" env-default:"5432"`
	DbName         string `env:"DATABASE_DB_NAME" env-default:"person_enrichment"`
	MigrationsPath string `env:"MIGRATIONS_PATH" env-default:"./migrations"`
}

type ExternalApi struct {
	AgifyURL       string `env:"AGIFY_API_URL" env-default:"https://api.agify.io"`
	GenderizeURL   string `env:"GENDERIZE_API_URL" env-default:"https://api.genderize.io"`
	NationalizeURL string `env:"NATIONALIZE_API_URL" env-default:"https://api.nationalize.io"`
}

type Config struct {
	Env         string `env:"ENV" env-default:"local"`
	HTTPServer  `env-required:"true"`
	DATABASE    `env-required:"true"`
	ExternalApi `env-required:"true"`
}

type HTTPServer struct {
	Address     string        `env:"HTTP_SERVER_ADDRESS" env-default:"localhost:8080"`
	Timeout     time.Duration `env:"HTTP_SERVER_TIMEOUT" env-default:"5s"`
	IdleTimeout time.Duration `env:"HTTP_SERVER_IDLE_TIMEOUT" env-default:"60s"`
}

func MustLoad() *Config {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}

	config := &Config{
		Env: os.Getenv("ENV"),
		HTTPServer: HTTPServer{
			Address:     os.Getenv("HTTP_SERVER_ADDRESS"),
			Timeout:     parseDuration(os.Getenv("HTTP_SERVER_TIMEOUT"), "5s"),
			IdleTimeout: parseDuration(os.Getenv("HTTP_SERVER_IDLE_TIMEOUT"), "60s"),
		},
		DATABASE: DATABASE{
			Username:       os.Getenv("DATABASE_USERNAME"),
			Password:       os.Getenv("DATABASE_PASSWORD"),
			Host:           os.Getenv("DATABASE_HOST"),
			Port:           os.Getenv("DATABASE_PORT"),
			DbName:         os.Getenv("DATABASE_DB_NAME"),
			MigrationsPath: os.Getenv("MIGRATIONS_PATH"),
		},
		ExternalApi: ExternalApi{
			AgifyURL:       os.Getenv("AGIFY_API_URL"),
			GenderizeURL:   os.Getenv("GENDERIZE_API_URL"),
			NationalizeURL: os.Getenv("NATIONALIZE_API_URL"),
		},
	}

	return config
}

func parseDuration(value, defaultValue string) time.Duration {
	if value == "" {
		value = defaultValue
	}
	duration, err := time.ParseDuration(value)
	if err != nil {
		log.Fatalf("Error parsing duration: %s", err)
	}
	return duration
}
