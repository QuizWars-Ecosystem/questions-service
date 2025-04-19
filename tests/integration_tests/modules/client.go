package modules

import (
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/types/known/emptypb"
	"testing"

	questionsv1 "github.com/QuizWars-Ecosystem/questions-service/gen/external/questions/v1"
	"github.com/QuizWars-Ecosystem/questions-service/tests/integration_tests/config"
)

func ClientServiceTest(t *testing.T, client questionsv1.QuestionsClientServiceClient, _ *config.TestConfig) {

	t.Run("client.GetCategories: successful", func(t *testing.T) {
		res, err := client.GetCategories(t.Context(), &emptypb.Empty{})

		require.NoError(t, err)
		require.Equal(t, len(categoriesList), len(res.Categories))

		for _, category := range res.Categories {
			for _, c := range categoriesList {
				if category.Name == c.Name {
					c = category
				}
			}
		}
	})

}
