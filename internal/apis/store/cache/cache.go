package cache

import (
	"github.com/QuizWars-Ecosystem/go-common/pkg/abstractions"
	"go.uber.org/zap"
)

type Cache struct {
	db     abstractions.RedisClient
	logger *zap.Logger
}

func NewCache(db abstractions.RedisClient, logger *zap.Logger) *Cache {
	return &Cache{db: db, logger: logger}
}
