package integration_tests

import (
	questionsv1 "github.com/QuizWars-Ecosystem/questions-service/gen/external/questions/v1"
	"github.com/QuizWars-Ecosystem/questions-service/internal/server"
	"github.com/QuizWars-Ecosystem/questions-service/tests/integration_tests/config"
	"testing"

	test "github.com/QuizWars-Ecosystem/go-common/pkg/testing/server"
	"github.com/stretchr/testify/require"
)

func Test(t *testing.T) {
	testCtx := t.Context()
	cfg := config.NewTestConfig()

	prepareInfrastructure(testCtx, t, cfg, runServer)
}

func runServer(t *testing.T, cfg *config.TestConfig) {
	srv, err := server.NewTestServer(t.Context(), cfg.ServiceConfig)
	require.NoError(t, err)

	conn, stop := test.RunServer(t, srv, cfg.ServiceConfig.GRPCPort)
	defer stop()

	adminClient := questionsv1.NewQuestionsAdminServiceClient(conn)
	clientClient := questionsv1.NewQuestionsClientServiceClient(conn)
	questionsClient := questionsv1.NewQuestionsServiceClient(conn)
}
