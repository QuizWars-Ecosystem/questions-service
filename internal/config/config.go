package config

import (
	"time"

	"github.com/QuizWars-Ecosystem/go-common/pkg/config"
)

type Config struct {
	config.DefaultServiceConfig
	Postgres    PostgresConfig `envPrefix:"POSTGRES_"`
	Redis       RedisConfig    `envPrefix:"REDIS_"`
	StoreConfig StoreConfig    `envPrefix:"STORE_"`
	JWT         JWTConfig      `envPrefix:"JWT_"`
}

type PostgresConfig struct {
	URL string `env:"URL"`
}

type RedisConfig struct {
	URL string `env:"URL"`
}

type StoreConfig struct {
	WarmUp        bool          `env:"WARM_UP"`
	WarmUpAmount  int           `env:"WARM_UP_AMOUNT"`
	WarmUpTimeout time.Duration `env:"WARM_UP_TIMEOUT"`
}

type JWTConfig struct {
	Secret            string        `env:"SECRET"`
	AccessExpiration  time.Duration `env:"ACCESS_EXPIRATION"`
	RefreshExpiration time.Duration `env:"REFRESH_EXPIRATION"`
}
