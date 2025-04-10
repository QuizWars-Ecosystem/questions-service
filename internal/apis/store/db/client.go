package db

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
)

type Client struct {
	db     *pgxpool.Pool
	logger *zap.Logger
}

func NewClient(db *pgxpool.Pool, logger *zap.Logger) *Client {
	return &Client{
		db:     db,
		logger: logger,
	}
}
