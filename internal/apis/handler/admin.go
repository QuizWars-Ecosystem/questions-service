package handler

import (
	"context"
	questionsv1 "github.com/QuizWars-Ecosystem/questions-service/gen/external/questions/v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (h *Handler) CreateQuestion(ctx context.Context, request *questionsv1.CreateQuestionRequest) (*emptypb.Empty, error) {
	return nil, nil
}

func (h *Handler) UpdateQuestion(ctx context.Context, request *questionsv1.UpdateQuestionRequest) (*emptypb.Empty, error) {
	return nil, nil
}

func (h *Handler) DeleteQuestion(ctx context.Context, request *questionsv1.DeleteQuestionRequest) (*emptypb.Empty, error) {
	return nil, nil
}

func (h *Handler) GetQuestions(ctx context.Context, request *questionsv1.GetQuestionsRequest) (*questionsv1.GetQuestionsResponse, error) {
	return nil, nil
}
