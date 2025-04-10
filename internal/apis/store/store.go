package store

import (
	"github.com/QuizWars-Ecosystem/questions-service/internal/apis/store/db"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
)

type Store struct {
	Questions *db.Questions
	Admin     *db.Admin
	Client    *db.Client
	logger    *zap.Logger
}

func NewService(postgres *pgxpool.Pool, logger *zap.Logger) *Store {
	return &Store{
		Questions: db.NewQuestions(postgres, logger),
		Admin:     db.NewAdmin(postgres, logger),
		Client:    db.NewClient(postgres, logger),
		logger:    logger,
	}
}
