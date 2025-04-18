package modules

import (
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

	t.Run("admin.CreateCategory: access token not provided", func(t *testing.T) {
		_, err := client.CreateCategory(emptyCtx, &questionsv1.CreateCategoryRequest{})

		require.Error(t, err)
		testerror.RequireForbiddenError(t, err, jwt.AuthAccessTokenNotProvidedError)
	})

	t.Run("admin.CreateCategory: invalid token", func(t *testing.T) {
		_, err := client.CreateCategory(invalidCtx, &questionsv1.CreateCategoryRequest{})

		require.Error(t, err)
		testerror.RequireForbiddenError(t, err, jwt.AuthInvalidTokenError)
	})

	t.Run("admin.CreateCategory: permission denied", func(t *testing.T) {
		_, err := client.CreateCategory(userCtx, &questionsv1.CreateCategoryRequest{})

		require.Error(t, err)
		testerror.RequireForbiddenError(t, err, jwt.AuthPermissionDeniedError)
	})

	t.Run("admin.CreateCategory: successful", func(t *testing.T) {
		res, err := client.CreateCategory(adminCtx, &questionsv1.CreateCategoryRequest{
			Name: sportCategory.Name,
		})

		require.NoError(t, err)
		require.NotEqual(t, 0, res.Id)
		sportCategory.Id = res.Id
	})

	t.Run("admin.CreateCategory: list: successful", func(t *testing.T) {
		useCase := []struct {
			name  string
			owner *questionsv1.Category
		}{
			{
				name:  countriesCategory.Name,
				owner: countriesCategory,
			},
			{
				name:  starsCategory.Name,
				owner: starsCategory,
			},
		}

		for _, u := range useCase {
			res, err := client.CreateCategory(adminCtx, &questionsv1.CreateCategoryRequest{
				Name: u.name,
			})

			require.NoError(t, err)
			require.NotEqual(t, 0, res.Id)

			u.owner.Id = res.Id
		}
	})

	t.Run("admin.CreateQuestion: access token not provided", func(t *testing.T) {
		_, err := client.CreateQuestion(emptyCtx, &questionsv1.CreateQuestionRequest{})

		require.Error(t, err)
		testerror.RequireForbiddenError(t, err, jwt.AuthAccessTokenNotProvidedError)
	})

	t.Run("admin.CreateQuestion: invalid token", func(t *testing.T) {
		_, err := client.CreateQuestion(invalidCtx, &questionsv1.CreateQuestionRequest{})

		require.Error(t, err)
		testerror.RequireForbiddenError(t, err, jwt.AuthInvalidTokenError)
	})

	t.Run("admin.CreateQuestion: permission denied", func(t *testing.T) {
		_, err := client.CreateQuestion(userCtx, &questionsv1.CreateQuestionRequest{})

		require.Error(t, err)
		testerror.RequireForbiddenError(t, err, jwt.AuthPermissionDeniedError)
	})

	t.Run("admin.CreateQuestion: eng: successful", func(t *testing.T) {
		question := questionsMap[sportEngQuestionsKey]
		_, err := client.CreateQuestion(adminCtx, &questionsv1.CreateQuestionRequest{
			Type:       question.Type,
			Difficulty: question.Difficulty,
			CategoryId: question.Category.Id,
			Language:   question.Language,
			Text:       question.Text,
			Options:    question.Options,
		})

		require.NoError(t, err)
	})

	t.Run("admin.CreateQuestion: already exists", func(t *testing.T) {
		question := questionsMap[sportEngQuestionsKey]
		_, err := client.CreateQuestion(adminCtx, &questionsv1.CreateQuestionRequest{
			Type:       question.Type,
			Difficulty: question.Difficulty,
			CategoryId: question.Category.Id,
			Language:   question.Language,
			Text:       question.Text,
			Options:    question.Options,
		})

		require.NoError(t, err)
	})

	t.Run("admin.CreateQuestion: rus: successful", func(t *testing.T) {
		question := questionsMap[sportRusQuestionsKey]
		_, err := client.CreateQuestion(adminCtx, &questionsv1.CreateQuestionRequest{
			Type:       question.Type,
			Difficulty: question.Difficulty,
			CategoryId: question.Category.Id,
			Language:   question.Language,
			Text:       question.Text,
			Options:    question.Options,
		})

		require.NoError(t, err)
	})

	t.Run("admin.CreateQuestion: list: successful", func(t *testing.T) {
		for _, question := range questionsMap {
			_, err := client.CreateQuestion(adminCtx, &questionsv1.CreateQuestionRequest{
				Type:       question.Type,
				Difficulty: question.Difficulty,
				CategoryId: question.Category.Id,
				Language:   question.Language,
				Text:       question.Text,
				Options:    question.Options,
			})

			require.NoError(t, err)
		}
	})
}
