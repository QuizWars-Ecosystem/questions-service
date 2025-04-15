package store

import (
	"context"
	"github.com/QuizWars-Ecosystem/questions-service/internal/models/admin"
	"github.com/QuizWars-Ecosystem/questions-service/internal/models/filter"
	"github.com/QuizWars-Ecosystem/questions-service/internal/models/questions"
	uuid "github.com/jackc/pgx/pgtype/ext/gofrs-uuid"
	"time"
)

type IStore interface {
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
	SaveQuestion(ctx context.Context, question *questions.Hashed) error
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
}
