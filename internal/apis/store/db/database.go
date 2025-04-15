package db

import (
	"errors"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
)

var (
	argumentsNotProvidedErr = errors.New("arguments not provided")
)

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
