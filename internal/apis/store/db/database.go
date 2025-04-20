package db

import (
	"errors"

	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
)

var errArgumentsNotProvided = errors.New("arguments not provided")

type Database struct {
	pool   *pgxpool.Pool
	logger *zap.Logger
}

func NewDatabase(pool *pgxpool.Pool, logger *zap.Logger) *Database {
	return &Database{
		pool:   pool,
		logger: logger,
	}
}
