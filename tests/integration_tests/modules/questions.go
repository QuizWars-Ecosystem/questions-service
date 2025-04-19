package modules

import (
	"github.com/stretchr/testify/require"
	"testing"

	questionsv1 "github.com/QuizWars-Ecosystem/questions-service/gen/external/questions/v1"
	"github.com/QuizWars-Ecosystem/questions-service/tests/integration_tests/config"
)

func QuestionsServiceTest(t *testing.T, client questionsv1.QuestionsServiceClient, _ *config.TestConfig) {

	t.Run("questions.GetQuestions: easy | eng | sport | <=10", func(t *testing.T) {
		difficulty := questionsv1.Difficulty_DIFFICULTY_EASY
		language := "eng"
		categoryID := sportCategory.Id
		amount := int32(10)

		res, err := client.GetQuestions(t.Context(), &questionsv1.GetQuestionsRequest{
			Difficulty: difficulty,
			Language:   language,
			CategoryId: categoryID,
			Amount:     amount,
		})

		require.NoError(t, err)
		require.LessOrEqual(t, len(res.Questions), int(amount))

		for _, q := range res.Questions {
			require.Equal(t, difficulty, q.Difficulty)
			require.Equal(t, language, q.Language)
			require.Equal(t, categoryID, q.Category.Id)
		}
	})

	t.Run("questions.GetQuestions: easy | rus | stars | <=10", func(t *testing.T) {
		difficulty := questionsv1.Difficulty_DIFFICULTY_EASY
		language := "rus"
		categoryID := starsCategory.Id
		amount := int32(10)

		res, err := client.GetQuestions(t.Context(), &questionsv1.GetQuestionsRequest{
			Difficulty: difficulty,
			Language:   language,
			CategoryId: categoryID,
			Amount:     amount,
		})

		require.NoError(t, err)
		require.LessOrEqual(t, len(res.Questions), int(amount))

		for _, q := range res.Questions {
			require.Equal(t, difficulty, q.Difficulty)
			require.Equal(t, language, q.Language)
			require.Equal(t, categoryID, q.Category.Id)
		}
	})

	t.Run("questions.GetQuestions: medium | rus | countries | <=10", func(t *testing.T) {
		difficulty := questionsv1.Difficulty_DIFFICULTY_MEDIUM
		language := "rus"
		categoryID := countriesCategory.Id
		amount := int32(10)

		res, err := client.GetQuestions(t.Context(), &questionsv1.GetQuestionsRequest{
			Difficulty: difficulty,
			Language:   language,
			CategoryId: categoryID,
			Amount:     amount,
		})

		require.NoError(t, err)
		require.LessOrEqual(t, len(res.Questions), int(amount))

		for _, q := range res.Questions {
			require.Equal(t, difficulty, q.Difficulty)
			require.Equal(t, language, q.Language)
			require.Equal(t, categoryID, q.Category.Id)
		}
	})

	t.Run("questions.GetQuestions: hard | eng | countries | <=10", func(t *testing.T) {
		difficulty := questionsv1.Difficulty_DIFFICULTY_HARD
		language := "eng"
		categoryID := countriesCategory.Id
		amount := int32(10)

		res, err := client.GetQuestions(t.Context(), &questionsv1.GetQuestionsRequest{
			Difficulty: difficulty,
			Language:   language,
			CategoryId: categoryID,
			Amount:     amount,
		})

		require.NoError(t, err)
		require.LessOrEqual(t, len(res.Questions), int(amount))

		for _, q := range res.Questions {
			require.Equal(t, difficulty, q.Difficulty)
			require.Equal(t, language, q.Language)
			require.Equal(t, categoryID, q.Category.Id)
		}
	})

	t.Run("questions.GetQuestionsBatch: (single | multi) & (easy | hard) & (sport | countries) & eng & <=10", func(t *testing.T) {
		req := &questionsv1.GetQuestionBatchRequest{
			Types: []questionsv1.Type{
				questionsv1.Type_TYPE_SINGLE,
				questionsv1.Type_TYPE_MULTI,
			},
			Difficulties: []questionsv1.Difficulty{
				questionsv1.Difficulty_DIFFICULTY_EASY,
				questionsv1.Difficulty_DIFFICULTY_HARD,
			},
			CategoriesIds: []int32{
				sportCategory.Id,
				countriesCategory.Id,
			},
			Language: "eng",
			Amount:   10,
		}

		res, err := client.GetQuestionBatch(t.Context(), req)

		require.NoError(t, err)
		require.LessOrEqual(t, len(res.Questions), int(req.Amount))

		for _, q := range res.Questions {
			require.Contains(t, req.Types, q.Type)
			require.Contains(t, req.Difficulties, q.Difficulty)
			require.Contains(t, req.CategoriesIds, q.Category.Id)
			require.Equal(t, req.Language, q.Language)
		}
	})

	t.Run("questions.GetQuestionsBatch: (multi) & (hard) & (countries) & eng & <=10", func(t *testing.T) {
		req := &questionsv1.GetQuestionBatchRequest{
			Types: []questionsv1.Type{
				questionsv1.Type_TYPE_MULTI,
			},
			Difficulties: []questionsv1.Difficulty{
				questionsv1.Difficulty_DIFFICULTY_HARD,
			},
			CategoriesIds: []int32{
				countriesCategory.Id,
			},
			Language: "eng",
			Amount:   10,
		}

		res, err := client.GetQuestionBatch(t.Context(), req)

		require.NoError(t, err)
		require.LessOrEqual(t, len(res.Questions), int(req.Amount))

		for _, q := range res.Questions {
			require.Contains(t, req.Types, q.Type)
			require.Contains(t, req.Difficulties, q.Difficulty)
			require.Contains(t, req.CategoriesIds, q.Category.Id)
			require.Equal(t, req.Language, q.Language)
		}
	})

	t.Run("questions.GetQuestionsBatch: (easy) & rus & <=10", func(t *testing.T) {
		req := &questionsv1.GetQuestionBatchRequest{
			Difficulties: []questionsv1.Difficulty{
				questionsv1.Difficulty_DIFFICULTY_EASY,
			},
			Language: "rus",
			Amount:   10,
		}

		res, err := client.GetQuestionBatch(t.Context(), req)

		require.NoError(t, err)
		require.LessOrEqual(t, len(res.Questions), int(req.Amount))

		for _, q := range res.Questions {
			require.Contains(t, req.Difficulties, q.Difficulty)
			require.Equal(t, req.Language, q.Language)
		}
	})

	t.Run("questions.GetQuestionsBatch: (stars | sport) & rus & <=10", func(t *testing.T) {
		req := &questionsv1.GetQuestionBatchRequest{
			CategoriesIds: []int32{
				starsCategory.Id,
				sportCategory.Id,
			},
			Language: "rus",
			Amount:   10,
		}

		res, err := client.GetQuestionBatch(t.Context(), req)

		require.NoError(t, err)
		require.LessOrEqual(t, len(res.Questions), int(req.Amount))

		for _, q := range res.Questions {
			require.Contains(t, req.CategoriesIds, q.Category.Id)
			require.Equal(t, req.Language, q.Language)
		}
	})
}
