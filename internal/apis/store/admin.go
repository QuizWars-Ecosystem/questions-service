package store

import (
	"context"

	"go.uber.org/zap"

	"github.com/QuizWars-Ecosystem/questions-service/internal/models/admin"
	"github.com/QuizWars-Ecosystem/questions-service/internal/models/questions"
	"github.com/google/uuid"
)

func (s *Store) GetFilteredQuestions(ctx context.Context, filter *admin.QuestionsFilter) ([]*questions.Question, int, error) {
	qs, amount, err := s.db.GetFilteredQuestions(ctx, filter)
	if err != nil {
		s.logger.Error("error.GetFilteredQuestions", zap.Error(err))
		return nil, amount, err
	}

	return qs, amount, err
}

func (s *Store) SaveCategory(ctx context.Context, name string) (int32, error) {
	return s.db.SaveCategory(ctx, name)
}

func (s *Store) SaveQuestion(ctx context.Context, question *questions.Hashed) error {
	return s.db.SaveQuestion(ctx, question)
}

func (s *Store) UpdateCategory(ctx context.Context, id int32, name string) error {
	return s.db.UpdateCategory(ctx, id, name)
}

func (s *Store) UpdateQuestion(ctx context.Context, id uuid.UUID, req *admin.UpdateQuestionRequest) error {
	return s.db.UpdateQuestion(ctx, id, req)
}

func (s *Store) UpdateQuestionOption(ctx context.Context, id uuid.UUID, req *admin.UpdateQuestionOptionRequest) error {
	return s.db.UpdateQuestionOption(ctx, id, req)
}

func (s *Store) DeleteQuestion(ctx context.Context, id uuid.UUID) error {
	return s.db.DeleteQuestion(ctx, id)
}

func (s *Store) DeleteQuestionOption(ctx context.Context, id uuid.UUID) error {
	return s.db.DeleteQuestionOption(ctx, id)
}
