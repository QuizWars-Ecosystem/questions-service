package handler

import (
	"context"
	questionsv1 "github.com/QuizWars-Ecosystem/questions-service/gen/external/questions/v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (h *Handler) GetCategories(ctx context.Context, _ *emptypb.Empty) (*questionsv1.GetCategoriesResponse, error) {
	categories, err := h.service.GetCategories(ctx)
	if err != nil {
		return nil, err
	}

	categoriesList := make([]*questionsv1.Category, len(categories))
	for i, category := range categories {
		var c *questionsv1.Category
		if c, err = category.Response(); err != nil {
			return nil, err
		}

		categoriesList[i] = c
	}

	return &questionsv1.GetCategoriesResponse{
		Categories: categoriesList,
	}, nil
}
