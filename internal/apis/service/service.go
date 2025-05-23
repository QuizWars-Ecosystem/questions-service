package service

import (
	"context"

	"github.com/QuizWars-Ecosystem/questions-service/internal/apis/store"
	"github.com/QuizWars-Ecosystem/questions-service/internal/models/admin"
	"github.com/QuizWars-Ecosystem/questions-service/internal/models/filter"
	"github.com/QuizWars-Ecosystem/questions-service/internal/models/questions"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

type IService interface {
	GetCategories(ctx context.Context) ([]*questions.Category, error)
	GetQuestions(ctx context.Context, language string, difficulty questions.Difficulty, categoryID, amount int32) ([]*questions.Question, error)
	GetQuestionsBatch(ctx context.Context, filter *filter.QuestionsFilter) ([]*questions.Question, error)
	GetFilteredQuestions(ctx context.Context, filter *admin.QuestionsFilter) ([]*questions.Question, int, error)
	CreateCategory(ctx context.Context, name string) (int32, error)
	CreateQuestion(ctx context.Context, req *questions.CreateQuestionRequest) error
	CreateQuestionOption(ctx context.Context, questionID uuid.UUID, req *admin.CreateQuestionOptionRequest) error
	UpdateCategory(ctx context.Context, id int32, name string) error
	UpdateQuestion(ctx context.Context, id uuid.UUID, req *admin.UpdateQuestionRequest) error
	UpdateQuestionOption(ctx context.Context, id uuid.UUID, req *admin.UpdateQuestionOptionRequest) error
	DeleteQuestion(ctx context.Context, id uuid.UUID) error
	DeleteQuestionOption(ctx context.Context, id uuid.UUID) error
}

var _ IService = (*Service)(nil)

type Service struct {
	store  store.IStore
	logger *zap.Logger
}

func NewService(store store.IStore, logger *zap.Logger) *Service {
	return &Service{
		store:  store,
		logger: logger,
	}
}
