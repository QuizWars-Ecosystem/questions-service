package store

import (
	"context"
	"github.com/QuizWars-Ecosystem/questions-service/internal/models/questions"
)

func (s *Store) GetCategories(ctx context.Context) ([]*questions.Category, error) {
	return s.db.GetCategories(ctx)
}
