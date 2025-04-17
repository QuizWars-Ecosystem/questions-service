package modules

import (
	"context"
	"github.com/QuizWars-Ecosystem/go-common/pkg/jwt"
	testerror "github.com/QuizWars-Ecosystem/go-common/pkg/testing/errors"
	questionsv1 "github.com/QuizWars-Ecosystem/questions-service/gen/external/questions/v1"
	"github.com/QuizWars-Ecosystem/questions-service/tests/integration_tests/config"
	"github.com/stretchr/testify/require"
	"testing"
)

func AdminServiceTest(t *testing.T, client questionsv1.QuestionsAdminServiceClient, cfg *config.TestConfig) {

	prepare(t, cfg)

	t.Run("admin.GetFilteredQuestions: access token not provided", func(t *testing.T) {
		_, err := client.GetFilteredQuestions(emptyCtx, &questionsv1.GetFilteredQuestionsRequest{
			Page: 1,
			Size: 10,
		})

		require.Error(t, err)
		testerror.RequireForbiddenError(t, err, jwt.AuthAccessTokenNotProvidedError)
	})

	t.Run("admin.GetFilteredQuestions: invalid token", func(t *testing.T) {
		_, err := client.GetFilteredQuestions(invalidCtx, &questionsv1.GetFilteredQuestionsRequest{
			Page: 1,
			Size: 10,
		})

		require.Error(t, err)
		testerror.RequireForbiddenError(t, err, jwt.AuthInvalidTokenError)
	})

	t.Run("admin.GetFilteredQuestions: permission denied", func(t *testing.T) {
		_, err := client.GetFilteredQuestions(userCtx, &questionsv1.GetFilteredQuestionsRequest{
			Page: 1,
			Size: 10,
		})

		require.Error(t, err)
		testerror.RequireForbiddenError(t, err, jwt.AuthPermissionDeniedError)
	})

	t.Run("admin.GetFilteredQuestions: empty", func(t *testing.T) {
		res, err := client.GetFilteredQuestions(adminCtx, &questionsv1.GetFilteredQuestionsRequest{
			Page: 1,
			Size: 10,
		})

		require.NoError(t, err)
		require.Empty(t, res.Questions)
	})
}

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
