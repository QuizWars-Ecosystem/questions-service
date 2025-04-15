package store

import (
	"github.com/QuizWars-Ecosystem/questions-service/internal/apis/store/cache"
	"github.com/QuizWars-Ecosystem/questions-service/internal/apis/store/db"
	"github.com/QuizWars-Ecosystem/questions-service/internal/config"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

type Store struct {
	db     IDatabase
	cache  ICache
	logger *zap.Logger
}

func NewService(pool *pgxpool.Pool, client *redis.ClusterClient, logger *zap.Logger, cfg *config.StoreConfig) *Store {
	var store Store
	store.db = db.NewDatabase(pool, logger)
	store.cache = cache.NewCache(client, logger)
	store.logger = logger

	if cfg.WarmUp {
		go func() {
			if err := store.warmUpStore(cfg.WarmUpAmount, cfg.WarmUpTimeout); err != nil {
				store.logger.Warn("failed to warm up store", zap.Error(err))
			}
		}()
	}

	return &store
}
