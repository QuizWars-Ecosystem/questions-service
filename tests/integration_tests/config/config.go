package config

import (
	"time"

	"github.com/QuizWars-Ecosystem/go-common/pkg/jwt"

	def "github.com/QuizWars-Ecosystem/go-common/pkg/config"
	test "github.com/QuizWars-Ecosystem/go-common/pkg/testing/config"
	"github.com/QuizWars-Ecosystem/questions-service/internal/config"
)

type TestConfig struct {
	ServiceConfig *config.Config
	Postgres      *test.PostgresConfig
	Redis         *test.RedisConfig
}

func NewTestConfig() *TestConfig {
	postgresCfg := test.DefaultPostgresConfig()
	redisCfg := test.DefaultRedisConfig()

	return &TestConfig{
		ServiceConfig: &config.Config{
			ServiceConfig: &def.ServiceConfig{
				Name:         "questions-service",
				Address:      "questions_address",
				Local:        true,
				GRPCPort:     50051,
				StartTimeout: time.Second * 30,
				StopTimeout:  time.Second * 30,
				ConsulURL:    "consul:8500",
			},
			JWT: &jwt.Config{
				Secret:            "secret",
				AccessExpiration:  time.Hour,
				RefreshExpiration: time.Hour,
			},
			StoreConfig: &config.StoreConfig{
				WarmUp:        true,
				WarmUpAmount:  100,
				WarmUpTimeout: time.Minute,
			},
		},
		Postgres: &postgresCfg,
		Redis:    &redisCfg,
	}
}
