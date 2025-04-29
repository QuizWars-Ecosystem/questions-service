package store

import (
	"github.com/QuizWars-Ecosystem/go-common/pkg/abstractions"
	"github.com/QuizWars-Ecosystem/questions-service/internal/apis/store/cache"
	"github.com/QuizWars-Ecosystem/questions-service/internal/apis/store/db"
	"github.com/QuizWars-Ecosystem/questions-service/internal/config"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
)

type Store struct {
	db     IDatabase
	cache  ICache
	logger *zap.Logger
}

func NewStore(pool *pgxpool.Pool, client abstractions.RedisClient, logger *zap.Logger, cfg *config.StoreConfig) *Store {
	var store Store
	store.db = db.NewDatabase(pool, logger)
	store.cache = cache.NewCache(client, logger)
	store.logger = logger

	if cfg.WarmUp {
		go func() {
			if err := store.warmUpStore(cfg.WarmUpAmount, cfg.WarmUpTimeout); err != nil {
				logger.Warn("failed to warm up store", zap.Error(err))
			}
		}()
	}

	return &store
}
