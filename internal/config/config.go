package config

import "github.com/kelseyhightower/envconfig"

type Config struct {
	Port string `envconfig:"PORT" default:":8080"`
	Env  string `envconfig:"ENV" default:"dev"` //dev,staging,prod
}

func Load() (*Config, error) {
	var conf Config
	if err := envconfig.Process("", &conf); err != nil {
		return nil, err
	}
	return &conf, nil
}

func (c *Config) IsDev() bool {
	return c.Env == "dev"
}
