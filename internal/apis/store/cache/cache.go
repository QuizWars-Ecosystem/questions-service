package cache

import (
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

type Cache struct {
	db     *redis.ClusterClient
	logger *zap.Logger
}

func NewCache(db *redis.ClusterClient, logger *zap.Logger) *Cache {
	return &Cache{db: db, logger: logger}
}
