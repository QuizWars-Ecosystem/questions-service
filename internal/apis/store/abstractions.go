package store

import (
	"context"
	"time"

	"github.com/QuizWars-Ecosystem/questions-service/internal/models/admin"
	"github.com/QuizWars-Ecosystem/questions-service/internal/models/filter"
	"github.com/QuizWars-Ecosystem/questions-service/internal/models/questions"
	"github.com/google/uuid"
)

type IStore interface {
	GetQuestions(ctx context.Context, language string, difficulty questions.Difficulty, categoryID, amount int32) ([]*questions.Question, error)
	GetQuestionsBatch(ctx context.Context, filter *filter.QuestionsFilter) ([]*questions.Question, error)
	GetCategories(ctx context.Context) ([]*questions.Category, error)
	GetFilteredQuestions(ctx context.Context, filter *admin.QuestionsFilter) ([]*questions.Question, int, error)
	SaveCategory(ctx context.Context, name string) (int32, error)
	SaveQuestion(ctx context.Context, question *questions.Hashed) error
	SaveQuestionOption(ctx context.Context, questionID uuid.UUID, req *admin.CreateQuestionOptionRequest) error
	UpdateCategory(ctx context.Context, id int32, name string) error
	UpdateQuestion(ctx context.Context, id uuid.UUID, req *admin.UpdateQuestionRequest) error
	UpdateQuestionOption(ctx context.Context, id uuid.UUID, req *admin.UpdateQuestionOptionRequest) error
	DeleteQuestion(ctx context.Context, id uuid.UUID) error
	DeleteQuestionOption(ctx context.Context, id uuid.UUID) error
}

type IDatabase interface {
	IClientDatabase
	IAdminDatabase
	IQuestionsDatabase
}

type IClientDatabase interface {
	GetCategories(ctx context.Context) ([]*questions.Category, error)
}

type IAdminDatabase interface {
	GetFilteredQuestions(ctx context.Context, filter *admin.QuestionsFilter) ([]*questions.Question, int, error)
	GetCategoryByID(ctx context.Context, id int32) (*questions.Category, error)
	GetCategoryByName(ctx context.Context, name string) (*questions.Category, error)
	SaveCategory(ctx context.Context, name string) (int32, error)
	SaveQuestion(ctx context.Context, question *questions.Hashed) error
	SaveQuestionOption(ctx context.Context, questionID uuid.UUID, req *admin.CreateQuestionOptionRequest) error
	UpdateCategory(ctx context.Context, id int32, name string) error
	UpdateQuestion(ctx context.Context, id uuid.UUID, req *admin.UpdateQuestionRequest) error
	UpdateQuestionOption(ctx context.Context, id uuid.UUID, req *admin.UpdateQuestionOptionRequest) error
	DeleteQuestion(ctx context.Context, id uuid.UUID) error
	DeleteQuestionOption(ctx context.Context, id uuid.UUID) error
}

type IQuestionsDatabase interface {
	GetRandomQuestionMeta(ctx context.Context, amount int64) ([]*questions.Meta, error)
	GetQuestionsByIDs(ctx context.Context, IDs []uuid.UUID) ([]*questions.Question, error)
	GetFilteredRandomQuestions(ctx context.Context, filter *filter.QuestionsFilter) ([]*questions.Question, error)
}

type ICache interface {
	GetCachedIDs(ctx context.Context, language, difficulty string, categoryID, amount int32) ([]uuid.UUID, int, error)
	GetBatchCachedIDs(ctx context.Context, language string, difficulties []string, categoryIDs []int32, amount int32) ([]uuid.UUID, int, error)
	AddCachedIDs(ctx context.Context, metas []*questions.Meta, timeout time.Duration) error
	GetCategory(ctx context.Context, id int32) (*questions.Category, error)
	AddCategory(ctx context.Context, category *questions.Category) error
}
