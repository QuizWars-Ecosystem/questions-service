package service

import (
	"context"

	"github.com/QuizWars-Ecosystem/questions-service/internal/models/questions"
)

func (s *Service) GetCategories(ctx context.Context) ([]*questions.Category, error) {
	return s.store.GetCategories(ctx)
}
