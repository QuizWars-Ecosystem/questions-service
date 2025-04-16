package service

import (
	"context"
	"github.com/QuizWars-Ecosystem/questions-service/internal/models/filter"
	"github.com/QuizWars-Ecosystem/questions-service/internal/models/questions"
)

func (s *Service) GetQuestions(ctx context.Context, language string, difficulty questions.Difficulty, categoryID, amount int32) ([]*questions.Question, error) {
	return s.store.GetQuestions(ctx, language, difficulty, categoryID, amount)
}

func (s *Service) GetQuestionsBatch(ctx context.Context, filter *filter.QuestionsFilter) ([]*questions.Question, error) {
	return s.store.GetQuestionsBatch(ctx, filter)
}
