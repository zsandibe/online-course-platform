package config

import (
	"fmt"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	logger "github.com/zsandibe/online-course-platform/pkg"
)

type Config struct {
	Server   serverConfig
	Postgres PostgresConfig
	Token    tokenConfig
	Smtp     smtpConfig
	Redis    redisConfig
	S3       s3Config
}

type s3Config struct {
	PartitionId   string `envconfig:"S3_PARTITION_ID" required:"true"`
	Url           string `envconfig:"S3_URL" required:"true"`
	SigningRegion string `envconfig:"S3_SIGNING_REGION" required:"true"`
}

type smtpConfig struct {
	Username string `envconfig:"SMTP_USERNAME" required:"true"`
	Password string `envconfig:"SMTP_PASSWORD" required:"true"`
	Port     string `envconfig:"SMTP_PORT" required:"true"`
	Server   string `envconfig:"SMTP_SERVER" required:"true"`
}

type PostgresConfig struct {
	User     string `envconfig:"DB_USER" required:"true"`
	Password string `envconfig:"DB_PASSWORD" required:"true"`
	Host     string `envconfig:"DB_HOST" required:"true"`
	Port     string `envconfig:"DB_PORT" required:"true"`
	Name     string `envconfig:"DB_NAME" required:"true"`
}

type serverConfig struct {
	Port string `envconfig:"SERVER_PORT" required:"true"`
}

type tokenConfig struct {
	SigningKey      string        `envconfig:"SIGNING_KEY" required:"true"`
	AccessTokenTTL  time.Duration `envconfig:"ACCESS_TOKEN_TTL" required:"true"`
	RefreshTokenTTL time.Duration `envconfig:"REFRESH_TOKEN_TTL" required:"true"`
}

type redisConfig struct {
	Host     string `envconfig:"REDIS_HOST" required:"true"`
	Port     int    `envconfig:"REDIS_PORT" required:"true"`
	Password string `envconfig:"REDIS_PASSWORD" required:"true"`
	Db       int    `envconfig:"REDIS_DB" required:"true"`
}

func NewConfig(path string) (*Config, error) {
	if err := godotenv.Load(path); err != nil {
		logger.Errorf("godotenv.Load(): %v", err)
		return nil, fmt.Errorf("error loading .env file: %v", err)
	}

	var config Config

	if err := envconfig.Process("", &config); err != nil {
		logger.Errorf("envconfig.Process(): %v", err)
		return nil, fmt.Errorf("error processing .env file: %v", err)
	}

	if os.Getenv("DOCKER") == "true" {
		config.Postgres.Host = "postgres"
	}

	return &config, nil
}
