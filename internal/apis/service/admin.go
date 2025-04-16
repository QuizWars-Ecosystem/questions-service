package service

import (
	"context"

	"github.com/QuizWars-Ecosystem/questions-service/internal/models/admin"
	"github.com/QuizWars-Ecosystem/questions-service/internal/models/questions"
	uuid "github.com/jackc/pgx/pgtype/ext/gofrs-uuid"
)

func (s *Service) GetFilteredQuestions(ctx context.Context, filter *admin.QuestionsFilter) ([]*questions.Question, int, error) {
	return s.store.GetFilteredQuestions(ctx, filter)
}

func (s *Service) CreateQuestion(ctx context.Context, req *questions.CreateQuestionRequest) error {
	return s.store.SaveQuestion(ctx, req.Hashed)
}

func (s *Service) UpdateCategory(ctx context.Context, id int32, name string) error {
	return s.store.UpdateCategory(ctx, id, name)
}

func (s *Service) UpdateQuestion(ctx context.Context, id uuid.UUID, req *admin.UpdateQuestionRequest) error {
	return s.store.UpdateQuestion(ctx, id, req)
}

func (s *Service) UpdateQuestionOption(ctx context.Context, id uuid.UUID, req *admin.UpdateQuestionOptionRequest) error {
	return s.store.UpdateQuestionOption(ctx, id, req)
}

func (s *Service) DeleteQuestion(ctx context.Context, id uuid.UUID) error {
	return s.store.DeleteQuestion(ctx, id)
}

func (s *Service) DeleteQuestionOption(ctx context.Context, id uuid.UUID) error {
	return s.store.DeleteQuestionOption(ctx, id)
}
