package handler

import (
	"context"
	questionsv1 "github.com/QuizWars-Ecosystem/questions-service/gen/external/questions/v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (h *Handler) GetCategories(ctx context.Context, empty *emptypb.Empty) (*questionsv1.GetCategoriesResponse, error) {
	return nil, nil
}
