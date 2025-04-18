package modules

import (
	"context"
	"github.com/QuizWars-Ecosystem/go-common/pkg/jwt"
	"github.com/QuizWars-Ecosystem/questions-service/tests/integration_tests/config"
	"github.com/stretchr/testify/require"
	"testing"
)

func prepare(t *testing.T, cfg *config.TestConfig) {
	var err error
	auth := jwt.NewService(cfg.ServiceConfig.JWT.Secret, cfg.ServiceConfig.JWT.AccessExpiration, cfg.ServiceConfig.JWT.RefreshExpiration)

	adminToken, err = auth.GenerateToken("1", string(jwt.Admin))
	require.NoError(t, err)

	userToken, err = auth.GenerateToken("2", string(jwt.User))
	require.NoError(t, err)

	adminCtx = auth.SetTokenInContext(t.Context(), adminToken)
	userCtx = auth.SetTokenInContext(t.Context(), userToken)
	emptyCtx = auth.SetTokenInContext(t.Context(), "")
	invalidCtx = auth.SetTokenInContext(t.Context(), "invalid token")
}

var (
	emptyCtx   context.Context
	invalidCtx context.Context
)

var (
	adminToken string
	adminCtx   context.Context
)

var (
	userToken string
	userCtx   context.Context
)
