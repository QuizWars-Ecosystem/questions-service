package handler

import (
	"context"

	questionsv1 "github.com/QuizWars-Ecosystem/questions-service/gen/external/questions/v1"
)

func (h *Handler) GetQuestionBatch(ctx context.Context, request *questionsv1.GetQuestionBatchRequest) (*questionsv1.GetQuestionsBatchResponse, error) {
	return nil, nil
}
