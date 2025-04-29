package store

import (
	"context"
	"errors"

	apperrors "github.com/QuizWars-Ecosystem/go-common/pkg/error"

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
	id, err := s.db.SaveCategory(ctx, name)
	if err != nil {
		s.logger.Error("failed to save category", zap.Error(err))
		return 0, err
	}

	err = s.cache.AddCategory(ctx, &questions.Category{ID: id, Name: name})
	if err != nil {
		s.logger.Error("failed to cache category", zap.Error(err))
	}

	return id, nil
}

func (s *Store) SaveQuestion(ctx context.Context, question *questions.Hashed) error {
	category, err := s.cache.GetCategory(ctx, question.Category.ID)

	var apperror *apperrors.Error
	if errors.Is(err, apperror) || category == nil {
		category, err = s.db.GetCategoryByID(ctx, question.Category.ID)
		if errors.Is(err, apperror) || category == nil {
			question.Category.ID, err = s.db.SaveCategory(ctx, question.Category.Name)
			if err != nil {
				s.logger.Error("failed to save category", zap.Error(err))
				return err
			}

			err = s.cache.AddCategory(ctx, &question.Category)
			if err != nil {
				s.logger.Error("failed to cache category", zap.Error(err))
			}
		} else {
			return err
		}
	}

	return s.db.SaveQuestion(ctx, question)
}

func (s *Store) SaveQuestionOption(ctx context.Context, questionID uuid.UUID, req *admin.CreateQuestionOptionRequest) error {
	return s.db.SaveQuestionOption(ctx, questionID, req)
}

func (s *Store) UpdateCategory(ctx context.Context, id int32, name string) error {
	err := s.db.UpdateCategory(ctx, id, name)
	if err != nil {
		s.logger.Error("failed to update category", zap.Error(err))
		return err
	}

	err = s.cache.AddCategory(ctx, &questions.Category{ID: id, Name: name})
	if err != nil {
		s.logger.Error("failed to cache category", zap.Error(err))
	}

	return nil
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
