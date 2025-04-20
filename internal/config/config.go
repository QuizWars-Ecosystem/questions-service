package config

import (
	"time"

	"github.com/QuizWars-Ecosystem/go-common/pkg/jwt"
	"github.com/QuizWars-Ecosystem/go-common/pkg/log"

	"github.com/QuizWars-Ecosystem/go-common/pkg/config"
)

type Config struct {
	*config.ServiceConfig `mapstructure:"service"`
	Logger                *log.Config     `mapstructure:"logger"`
	JWT                   *jwt.Config     `mapstructure:"jwt"`
	Postgres              *PostgresConfig `mapstructure:"postgres"`
	Redis                 *RedisConfig    `mapstructure:"redis"`
	StoreConfig           *StoreConfig    `mapstructure:"store"`
}

type PostgresConfig struct {
	URL string `mapstructure:"url"`
}

type RedisConfig struct {
	URL string `mapstructure:"url"`
}

type StoreConfig struct {
	WarmUp        bool          `mapstructure:"warm_up"`
	WarmUpAmount  int           `mapstructure:"warm_up_amount"`
	WarmUpTimeout time.Duration `mapstructure:"warm_up_timeout"`
}
