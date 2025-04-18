package integration_tests

import (
	"testing"

	questionsv1 "github.com/QuizWars-Ecosystem/questions-service/gen/external/questions/v1"
	"github.com/QuizWars-Ecosystem/questions-service/internal/server"
	"github.com/QuizWars-Ecosystem/questions-service/tests/integration_tests/config"
	"github.com/QuizWars-Ecosystem/questions-service/tests/integration_tests/modules"

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

	modules.AdminServiceTest(t, adminClient, cfg)
	modules.ClientServiceTest(t, clientClient, cfg)
	modules.QuestionsServiceTest(t, questionsClient, cfg)
}
