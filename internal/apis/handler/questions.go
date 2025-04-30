package handler

import (
	"context"

	"github.com/QuizWars-Ecosystem/questions-service/internal/metrics"

	"github.com/QuizWars-Ecosystem/go-common/pkg/abstractions"
	questionsv1 "github.com/QuizWars-Ecosystem/questions-service/gen/external/questions/v1"
	"github.com/QuizWars-Ecosystem/questions-service/internal/models/filter"
	"github.com/QuizWars-Ecosystem/questions-service/internal/models/questions"
)

func (h *Handler) GetQuestions(ctx context.Context, request *questionsv1.GetQuestionsRequest) (*questionsv1.QuestionsResponse, error) {
	qs, err := h.service.GetQuestions(
		ctx,
		request.GetLanguage(),
		questions.DifficultyFromGRPCEnum(request.GetDifficulty()),
		request.GetCategoryId(),
		request.GetAmount(),
	)
	if err != nil {
		return nil, err
	}

	questionsList := make([]*questionsv1.Question, len(qs))
	for i, question := range qs {
		var q *questionsv1.Question
		if q, err = question.Response(); err != nil {
			return nil, err
		}

		questionsList[i] = q
	}

	metrics.QuestionsRequestsGauge.Set(float64(len(questionsList)))

	return &questionsv1.QuestionsResponse{
		Questions: questionsList,
	}, nil
}

func (h *Handler) GetQuestionBatch(ctx context.Context, request *questionsv1.GetQuestionBatchRequest) (*questionsv1.QuestionsResponse, error) {
	f, err := abstractions.MakeRequest[filter.QuestionsFilter](request)
	if err != nil {
		return nil, err
	}

	qs, err := h.service.GetQuestionsBatch(ctx, f)
	if err != nil {
		return nil, err
	}

	questionsList := make([]*questionsv1.Question, len(qs))
	for i, question := range qs {
		var q *questionsv1.Question
		if q, err = question.Response(); err != nil {
			return nil, err
		}

		questionsList[i] = q
	}

	metrics.QuestionsRequestsGauge.Set(float64(len(questionsList)))

	return &questionsv1.QuestionsResponse{
		Questions: questionsList,
	}, nil
}
