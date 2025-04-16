package filter

import (
	"github.com/QuizWars-Ecosystem/go-common/pkg/abstractions"
	questionsv1 "github.com/QuizWars-Ecosystem/questions-service/gen/external/questions/v1"
	"github.com/QuizWars-Ecosystem/questions-service/internal/models/questions"
)

var _ abstractions.Requestable[QuestionsFilter, *questionsv1.GetQuestionBatchRequest] = (*QuestionsFilter)(nil)

func (q QuestionsFilter) Request(req *questionsv1.GetQuestionBatchRequest) (*QuestionsFilter, error) {
	var types = make([]questions.Type, len(req.Types))
	var sources = make([]questions.Source, len(req.Sources))
	var difficulties = make([]questions.Difficulty, len(req.Difficulties))

	for i, t := range req.Types {
		types[i] = questions.TypeFromGRPCEnum(t)
	}

	for i, s := range req.Sources {
		sources[i] = questions.SourceFromGRPCEnum(s)
	}

	for i, d := range req.Difficulties {
		difficulties[i] = questions.DifficultyFromGRPCEnum(d)
	}

	q.Categories = req.CategoriesIds
	q.Language = req.Language
	q.Amount = req.Amount

	return &q, nil
}
