package config

import "github.com/kelseyhightower/envconfig"

type Config struct {
	RedisURL   string `envconfig:"REDIS_URL" default:"localhost:6379"`
	EchoURL    string `envconfig:"ECHO_URL" default:":8080"`
	DBHost     string `envconfig:"DB_HOST" default:"127.0.0.1"`
	DBUser     string `envconfig:"DB_USER" default:"sa"`
	DBPort     int    `envconfig:"DB_PORT" default:"3306"`
	DBPassword string `envconfig:"DB_PASSWORD" default:"sa"`
	DBName     string `envconfig:"DB_NAME" default:"test_database"`
	DBSslMode  string `envconfig:"DB_SSL_MODE" default:"false"`
}

func NewConfig() (Config, error) {
	var config Config
	if err := envconfig.Process("", &config); err != nil {
		return Config{}, err
	}
	return config, nil
}
