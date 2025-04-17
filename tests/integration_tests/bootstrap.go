package integration_tests

import (
	"context"
	"github.com/QuizWars-Ecosystem/questions-service/tests/integration_tests/config"
	"strings"
	"testing"

	"github.com/QuizWars-Ecosystem/go-common/pkg/testing/migrations"

	"github.com/QuizWars-Ecosystem/go-common/pkg/testing/containers"
	"github.com/stretchr/testify/require"
	"github.com/testcontainers/testcontainers-go"
)

type runServerFn func(t *testing.T, cfg *config.TestConfig)

func prepareInfrastructure(
	ctx context.Context,
	t *testing.T,
	cfg *config.TestConfig,
	runServerFn runServerFn,
) {
	postgres, err := containers.NewPostgresContainer(ctx, cfg.Postgres)
	require.NoError(t, err)

	defer testcontainers.CleanupContainer(t, postgres)

	postgresUrl, err := postgres.ConnectionString(ctx)
	require.NoError(t, err)

	cfg.ServiceConfig.Postgres.URL = postgresUrl

	migrations.RunMigrations(t, postgresUrl, "../../migrations")

	redis, err := containers.NewRedisContainer(ctx, cfg.Redis)
	require.NoError(t, err)

	defer testcontainers.CleanupContainer(t, redis)

	redisUrl, err := redis.ConnectionString(ctx)
	require.NoError(t, err)

	cfg.ServiceConfig.Redis.URL = strings.TrimPrefix(redisUrl, "redis://")

	runServerFn(t, cfg)
}
