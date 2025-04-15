package cache

import (
	"fmt"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"time"
)

type Cache struct {
	db     *redis.ClusterClient
	logger *zap.Logger
}

func NewCache(db *redis.ClusterClient, logger *zap.Logger) *Cache {
	return &Cache{db: db, logger: logger}
}

func key(language, difficulty string, categoryID int32) string {
	return fmt.Sprintf("questions:%s:%s:%d", language, difficulty, categoryID)
}

func keyList(language string, difficulties []string, categoryIDs []int32) []string {
	var keys = make([]string, len(difficulties)+len(categoryIDs))

	for i, d := range difficulties {
		for j, id := range categoryIDs {
			keys[i+j] = key(language, d, id)
		}
	}

	return keys
}

func destinationKey() string {
	return fmt.Sprintf("temp:%d", time.Now().UnixMilli())
}
