package db

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
)

type Questions struct {
	db     *pgxpool.Pool
	logger *zap.Logger
}

func NewQuestions(db *pgxpool.Pool, logger *zap.Logger) *Questions {
	return &Questions{
		db:     db,
		logger: logger,
	}
}
