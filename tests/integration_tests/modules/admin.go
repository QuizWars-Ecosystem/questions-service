package modules

import (
	"github.com/google/uuid"
	"testing"

	"github.com/QuizWars-Ecosystem/go-common/pkg/jwt"
	testerror "github.com/QuizWars-Ecosystem/go-common/pkg/testing/errors"
	questionsv1 "github.com/QuizWars-Ecosystem/questions-service/gen/external/questions/v1"
	"github.com/QuizWars-Ecosystem/questions-service/tests/integration_tests/config"
	"github.com/stretchr/testify/require"
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

	t.Run("admin.UpdateCategory: access token not provided", func(t *testing.T) {
		_, err := client.UpdateCategory(emptyCtx, &questionsv1.UpdateCategoryRequest{})

		require.Error(t, err)
		testerror.RequireForbiddenError(t, err, jwt.AuthAccessTokenNotProvidedError)
	})

	t.Run("admin.UpdateCategory: invalid token", func(t *testing.T) {
		_, err := client.UpdateCategory(invalidCtx, &questionsv1.UpdateCategoryRequest{})

		require.Error(t, err)
		testerror.RequireForbiddenError(t, err, jwt.AuthInvalidTokenError)
	})

	t.Run("admin.UpdateCategory: permission denied", func(t *testing.T) {
		_, err := client.UpdateCategory(userCtx, &questionsv1.UpdateCategoryRequest{})

		require.Error(t, err)
		testerror.RequireForbiddenError(t, err, jwt.AuthPermissionDeniedError)
	})

	t.Run("admin.UpdateCategory: not found", func(t *testing.T) {
		testId, testName := int32(100), "new name"
		_, err := client.UpdateCategory(adminCtx, &questionsv1.UpdateCategoryRequest{
			Id:   testId,
			Name: testName,
		})

		require.Error(t, err)
		testerror.RequireNotFoundError(t, err, "category", "id", testId)
	})

	t.Run("admin.UpdateCategory: successful", func(t *testing.T) {
		_, err := client.UpdateCategory(adminCtx, &questionsv1.UpdateCategoryRequest{
			Id:   starsCategory.Id,
			Name: starsCategory.Name,
		})

		require.NoError(t, err)
	})

	t.Run("admin.GetFilteredQuestions: successful", func(t *testing.T) {
		res, err := client.GetFilteredQuestions(adminCtx, &questionsv1.GetFilteredQuestionsRequest{
			Page: 1,
			Size: 10,
		})

		require.NoError(t, err)
		require.Equal(t, int64(len(questionsMap)), res.Amount)
		require.Equal(t, len(questionsMap), len(res.Questions))

		for key, val := range questionsMap {
			for _, q := range res.Questions {
				if q.Text == val.Text {
					questionsMap[key] = q
				}
			}
		}
	})

	t.Run("admin.UpdateQuestion: access token not provided", func(t *testing.T) {
		_, err := client.UpdateQuestion(emptyCtx, &questionsv1.UpdateQuestionRequest{})

		require.Error(t, err)
		testerror.RequireForbiddenError(t, err, jwt.AuthAccessTokenNotProvidedError)
	})

	t.Run("admin.UpdateQuestion: invalid token", func(t *testing.T) {
		_, err := client.UpdateQuestion(invalidCtx, &questionsv1.UpdateQuestionRequest{})

		require.Error(t, err)
		testerror.RequireForbiddenError(t, err, jwt.AuthInvalidTokenError)
	})

	t.Run("admin.UpdateQuestion: permission denied", func(t *testing.T) {
		_, err := client.UpdateQuestion(userCtx, &questionsv1.UpdateQuestionRequest{})

		require.Error(t, err)
		testerror.RequireForbiddenError(t, err, jwt.AuthPermissionDeniedError)
	})

	t.Run("admin.UpdateQuestion: not found", func(t *testing.T) {
		testID := uuid.New()
		testData := questionsv1.Type_TYPE_SINGLE
		_, err := client.UpdateQuestion(adminCtx, &questionsv1.UpdateQuestionRequest{
			Id:   testID.String(),
			Type: &testData,
		})

		require.Error(t, err)
		testerror.RequireNotFoundError(t, err, "question", "id", testID)
	})

	t.Run("admin.UpdateQuestion: new type: successful", func(t *testing.T) {
		q := questionsMap[sportRusQuestionsKey]
		testData := questionsv1.Type_TYPE_SINGLE
		_, err := client.UpdateQuestion(adminCtx, &questionsv1.UpdateQuestionRequest{
			Id:   q.Id,
			Type: &testData,
		})

		require.NoError(t, err)
		q.Type = testData
	})

	t.Run("admin.UpdateQuestion: new difficulty: successful", func(t *testing.T) {
		q := questionsMap[sportRusQuestionsKey]
		testData := questionsv1.Difficulty_DIFFICULTY_HARD
		_, err := client.UpdateQuestion(adminCtx, &questionsv1.UpdateQuestionRequest{
			Id:         q.Id,
			Difficulty: &testData,
		})

		require.NoError(t, err)
		q.Difficulty = testData
	})

	t.Run("admin.UpdateQuestion: new category: specified category was not found", func(t *testing.T) {
		q := questionsMap[sportRusQuestionsKey]
		testData := int32(100)
		_, err := client.UpdateQuestion(adminCtx, &questionsv1.UpdateQuestionRequest{
			Id:         q.Id,
			CategoryId: &testData,
		})

		require.Error(t, err)
		testerror.RequireBadRequestError(t, err, "specified category was not found")
	})

	t.Run("admin.UpdateQuestion: new category: successful", func(t *testing.T) {
		q := questionsMap[sportRusQuestionsKey]
		testData := sportCategory.Id
		_, err := client.UpdateQuestion(adminCtx, &questionsv1.UpdateQuestionRequest{
			Id:         q.Id,
			CategoryId: &testData,
		})

		require.NoError(t, err)
		q.Category.Id = testData
	})

	t.Run("admin.UpdateQuestion: new text: successful", func(t *testing.T) {
		q := questionsMap[sportRusQuestionsKey]
		testData := q.Text
		_, err := client.UpdateQuestion(adminCtx, &questionsv1.UpdateQuestionRequest{
			Id:   q.Id,
			Text: &testData,
		})

		require.NoError(t, err)
		q.Text = testData
	})

	t.Run("admin.UpdateQuestion: new language: successful", func(t *testing.T) {
		q := questionsMap[sportRusQuestionsKey]
		testData := q.Language
		_, err := client.UpdateQuestion(adminCtx, &questionsv1.UpdateQuestionRequest{
			Id:       q.Id,
			Language: &testData,
		})

		require.NoError(t, err)
		q.Language = testData
	})

	t.Run("admin.UpdateQuestionOption: access token not provided", func(t *testing.T) {
		_, err := client.UpdateQuestionOption(emptyCtx, &questionsv1.UpdateQuestionOptionRequest{})

		require.Error(t, err)
		testerror.RequireForbiddenError(t, err, jwt.AuthAccessTokenNotProvidedError)
	})

	t.Run("admin.UpdateQuestionOption: invalid token", func(t *testing.T) {
		_, err := client.UpdateQuestionOption(invalidCtx, &questionsv1.UpdateQuestionOptionRequest{})

		require.Error(t, err)
		testerror.RequireForbiddenError(t, err, jwt.AuthInvalidTokenError)
	})

	t.Run("admin.UpdateQuestionOption: permission denied", func(t *testing.T) {
		_, err := client.UpdateQuestionOption(userCtx, &questionsv1.UpdateQuestionOptionRequest{})

		require.Error(t, err)
		testerror.RequireForbiddenError(t, err, jwt.AuthPermissionDeniedError)
	})

	t.Run("admin.UpdateQuestionOption: arguments not provided", func(t *testing.T) {
		testID := uuid.New()
		_, err := client.UpdateQuestionOption(adminCtx, &questionsv1.UpdateQuestionOptionRequest{
			Id: testID.String(),
		})

		require.Error(t, err)
		testerror.RequireBadRequestError(t, err, "arguments not provided")
	})

	t.Run("admin.UpdateQuestionOption: not found", func(t *testing.T) {
		testID := uuid.New()
		testData := true
		_, err := client.UpdateQuestionOption(adminCtx, &questionsv1.UpdateQuestionOptionRequest{
			Id:        testID.String(),
			IsCorrect: &testData,
		})

		require.Error(t, err)
		testerror.RequireNotFoundError(t, err, "option", "id", testID)
	})

	t.Run("admin.UpdateQuestionOption: successful", func(t *testing.T) {
		q := questionsMap[countriesEngQuestionsKey]
		o := q.Options[0]
		testData := o.IsCorrect

		_, err := client.UpdateQuestionOption(adminCtx, &questionsv1.UpdateQuestionOptionRequest{
			Id:        o.Id,
			IsCorrect: &testData,
		})

		require.NoError(t, err)
		o.IsCorrect = testData
	})

	t.Run("admin.DeleteQuestionOption: access token not provided", func(t *testing.T) {
		_, err := client.DeleteQuestionOption(emptyCtx, &questionsv1.DeleteQuestionOptionRequest{})

		require.Error(t, err)
		testerror.RequireForbiddenError(t, err, jwt.AuthAccessTokenNotProvidedError)
	})

	t.Run("admin.DeleteQuestionOption: invalid token", func(t *testing.T) {
		_, err := client.DeleteQuestionOption(invalidCtx, &questionsv1.DeleteQuestionOptionRequest{})

		require.Error(t, err)
		testerror.RequireForbiddenError(t, err, jwt.AuthInvalidTokenError)
	})

	t.Run("admin.DeleteQuestionOption: permission denied", func(t *testing.T) {
		_, err := client.DeleteQuestionOption(userCtx, &questionsv1.DeleteQuestionOptionRequest{})

		require.Error(t, err)
		testerror.RequireForbiddenError(t, err, jwt.AuthPermissionDeniedError)
	})

	t.Run("admin.DeleteQuestionOption: not found", func(t *testing.T) {
		testID := uuid.New()
		_, err := client.DeleteQuestionOption(adminCtx, &questionsv1.DeleteQuestionOptionRequest{
			Id: testID.String(),
		})

		require.Error(t, err)
		testerror.RequireNotFoundError(t, err, "option", "id", testID)
	})

	t.Run("admin.DeleteQuestionOption: successful", func(t *testing.T) {
		q := questionsMap[countriesRusQuestionsKey]
		o := q.Options[0]

		_, err := client.DeleteQuestionOption(adminCtx, &questionsv1.DeleteQuestionOptionRequest{
			Id: o.Id,
		})

		require.NoError(t, err)
		q.Options = append(q.Options[1:])
	})

	t.Run("admin.DeleteQuestion: access token not provided", func(t *testing.T) {
		_, err := client.DeleteQuestion(emptyCtx, &questionsv1.DeleteQuestionRequest{})

		require.Error(t, err)
		testerror.RequireForbiddenError(t, err, jwt.AuthAccessTokenNotProvidedError)
	})

	t.Run("admin.DeleteQuestion: invalid token", func(t *testing.T) {
		_, err := client.DeleteQuestion(invalidCtx, &questionsv1.DeleteQuestionRequest{})

		require.Error(t, err)
		testerror.RequireForbiddenError(t, err, jwt.AuthInvalidTokenError)
	})

	t.Run("admin.DeleteQuestion: permission denied", func(t *testing.T) {
		_, err := client.DeleteQuestion(userCtx, &questionsv1.DeleteQuestionRequest{})

		require.Error(t, err)
		testerror.RequireForbiddenError(t, err, jwt.AuthPermissionDeniedError)
	})

	t.Run("admin.DeleteQuestion: not found", func(t *testing.T) {
		testID := uuid.New().String()
		_, err := client.DeleteQuestion(adminCtx, &questionsv1.DeleteQuestionRequest{
			Id: testID,
		})

		require.Error(t, err)
		testerror.RequireNotFoundError(t, err, "question", "id", testID)
	})

	t.Run("admin.DeleteQuestion: successful", func(t *testing.T) {
		q := questionsMap[questionForDeleting]
		_, err := client.DeleteQuestion(adminCtx, &questionsv1.DeleteQuestionRequest{
			Id: q.Id,
		})

		require.NoError(t, err)
		delete(questionsMap, questionForDeleting)
	})
}
