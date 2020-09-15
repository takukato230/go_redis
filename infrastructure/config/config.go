package config

import "github.com/kelseyhightower/envconfig"

type Config struct {
	RedisURL string `envconfig:"REDIS_URL" default"localhost:6379"`
	EchoURL  string `envconfig:"ECHO_URL" default:":8080"`
}

func NewConfig() (Config, error) {
	var config Config
	if err := envconfig.Process("", &config); err != nil {
		return Config{}, err
	}
	return config, nil
}
