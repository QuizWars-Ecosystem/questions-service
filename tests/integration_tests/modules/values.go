package modules

import (
	"context"
	"testing"

	"github.com/QuizWars-Ecosystem/go-common/pkg/jwt"
	questionsv1 "github.com/QuizWars-Ecosystem/questions-service/gen/external/questions/v1"
	"github.com/QuizWars-Ecosystem/questions-service/tests/integration_tests/config"
	"github.com/stretchr/testify/require"
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

// Contexts for auth (JWT)
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

// Categories
var (
	sportCategory = &questionsv1.Category{
		Name: "Sport",
	}

	countriesCategory = &questionsv1.Category{
		Name: "Countries",
	}

	starsCategory = &questionsv1.Category{
		Name: "Stars",
	}
)

// Questions
var (
	sportEngQuestionsKey = "sport:eng"
	sportRusQuestionsKey = "sport:rus"

	countriesEngQuestionsKey = "countries:eng"
	countriesRusQuestionsKey = "countries:rus"

	starsEngQuestionsKey = "stars:eng"
	starsRusQuestionsKey = "stars:rus"

	questionForDeleting = "sport:eng"

	questionsMap = map[string]*questionsv1.Question{
		sportEngQuestionsKey: {
			Type:       questionsv1.Type_TYPE_SINGLE,
			Difficulty: questionsv1.Difficulty_DIFFICULTY_EASY,
			Category:   sportCategory,
			Language:   "eng",
			Text:       "With which sport is Kenenisa Bekele associated?",
			Options: []*questionsv1.Option{
				{
					Text:      "Athletics",
					IsCorrect: true,
				},
				{
					Text:      "Boxing",
					IsCorrect: false,
				},
				{
					Text:      "Motor racing",
					IsCorrect: false,
				},
				{
					Text:      "Rowing",
					IsCorrect: false,
				},
			},
		},
		sportRusQuestionsKey: {
			Type:       questionsv1.Type_TYPE_SINGLE,
			Difficulty: questionsv1.Difficulty_DIFFICULTY_EASY,
			Category:   sportCategory,
			Language:   "rus",
			Text:       "С каким видом спорта связан Кенениса Бекеле?",
			Options: []*questionsv1.Option{
				{
					Text:      "Бокс",
					IsCorrect: false,
				},
				{
					Text:      "Автогонки",
					IsCorrect: false,
				},
				{
					Text:      "Легкая атлетика",
					IsCorrect: true,
				},
				{
					Text:      "Гребля",
					IsCorrect: false,
				},
			},
		},
		countriesEngQuestionsKey: {
			Type:       questionsv1.Type_TYPE_SINGLE,
			Difficulty: questionsv1.Difficulty_DIFFICULTY_EASY,
			Category:   countriesCategory,
			Language:   "eng",
			Text:       "Which country borders Italy, Switzerland, Germany, Czech Republic, Hungary, Slovenia, and Liechtenstein?",
			Options: []*questionsv1.Option{
				{
					Text:      "Austria",
					IsCorrect: true,
				},
				{
					Text:      "Bosnia and Herzegovina",
					IsCorrect: false,
				},
				{
					Text:      "Croatia",
					IsCorrect: true,
				},
				{
					Text:      "San Marino",
					IsCorrect: false,
				},
			},
		},
		countriesRusQuestionsKey: {
			Type:       questionsv1.Type_TYPE_SINGLE,
			Difficulty: questionsv1.Difficulty_DIFFICULTY_EASY,
			Category:   countriesCategory,
			Language:   "rus",
			Text:       "На них проходят границы Италии, Швейцарии, Германии, Швеции, Венгрии, Словении и Лихтенштейна?",
			Options: []*questionsv1.Option{
				{
					Text:      "Австрия",
					IsCorrect: true,
				},
				{
					Text:      "Босния и Герцеговина",
					IsCorrect: false,
				},
				{
					Text:      "Хорватия",
					IsCorrect: true,
				},
				{
					Text:      "Сан-Марино",
					IsCorrect: false,
				},
			},
		},
		starsEngQuestionsKey: {
			Type:       questionsv1.Type_TYPE_SINGLE,
			Difficulty: questionsv1.Difficulty_DIFFICULTY_EASY,
			Category:   starsCategory,
			Language:   "eng",
			Text:       "How old was Muhammad Ali when he died?",
			Options: []*questionsv1.Option{
				{
					Text:      "61",
					IsCorrect: false,
				},
				{
					Text:      "56",
					IsCorrect: true,
				},
				{
					Text:      "He is still alive",
					IsCorrect: false,
				},
				{
					Text:      "74",
					IsCorrect: true,
				},
			},
		},
		starsRusQuestionsKey: {
			Type:       questionsv1.Type_TYPE_SINGLE,
			Difficulty: questionsv1.Difficulty_DIFFICULTY_EASY,
			Category:   starsCategory,
			Language:   "eng",
			Text:       "Сколько лет было Мухаммеду Али, когда он умер?",
			Options: []*questionsv1.Option{
				{
					Text:      "61",
					IsCorrect: false,
				},
				{
					Text:      "56",
					IsCorrect: true,
				},
				{
					Text:      "Он все еще жив",
					IsCorrect: false,
				},
				{
					Text:      "74",
					IsCorrect: true,
				},
			},
		},
		questionForDeleting: {
			Type:       questionsv1.Type_TYPE_MULTI,
			Difficulty: questionsv1.Difficulty_DIFFICULTY_MEDIUM,
			Category:   sportCategory,
			Language:   "eng",
			Text:       "SOME NOT IMPORTANT TEXT",
			Options: []*questionsv1.Option{
				{
					Text:      "1",
					IsCorrect: true,
				},
				{
					Text:      "2",
					IsCorrect: false,
				},
				{
					Text:      "3",
					IsCorrect: true,
				},
				{
					Text:      "4",
					IsCorrect: false,
				},
			},
		},
	}
)
