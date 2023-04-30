package config

import (
	"context"

	"github.com/BurntSushi/toml"
	"github.com/pkg/errors"
)

type Config struct {
	DataBase DbConfig `toml:"database"`
}

type DbConfig struct {
	Host     string `toml:"db_host"`
	Name     string `toml:"db_name"`
	User     string `toml:"db_user"`
	Password string `toml:"db_password"`
	Port     string `toml:"db_port"`
}

func NewConfig(ctx context.Context) (*Config, error) {
	var conf Config
	if _, err := toml.DecodeFile("config.toml", &conf); err != nil {
		return nil, errors.Wrap(err, "No 'config.toml' file loaded")
	}

	return &conf, nil
}
