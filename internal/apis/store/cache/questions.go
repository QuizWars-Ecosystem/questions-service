package cache

import (
	"context"
	"time"

	apperrors "github.com/QuizWars-Ecosystem/go-common/pkg/error"
	"github.com/QuizWars-Ecosystem/questions-service/internal/models/questions"
	uuid "github.com/jackc/pgx/pgtype/ext/gofrs-uuid"
)

func (c *Cache) GetCachedIDs(ctx context.Context, language, difficulty string, categoryID, amount int32) ([]uuid.UUID, int, error) {
	ids := make([]uuid.UUID, amount)
	var count int

	values, err := c.db.SRandMemberN(ctx, key(language, difficulty, categoryID), int64(amount)).Result()
	if err != nil {
		return nil, 0, apperrors.Internal(err)
	}

	for i, v := range values {
		if err = ids[i].Scan(v); err != nil {
			return nil, 0, apperrors.Internal(err)
		}
		count++
	}

	return ids, count, nil
}

func (c *Cache) GetBatchCachedIDs(ctx context.Context, language string, difficulties []string, categoryIDs []int32, amount int32) ([]uuid.UUID, int, error) {
	keys := keyList(language, difficulties, categoryIDs)
	destination := destinationKey()

	count, err := c.db.SUnionStore(ctx, destination, keys...).Result()
	if err != nil {
		return nil, 0, apperrors.Internal(err)
	}

	go c.db.Del(ctx, destination)

	if count == 0 {
		return nil, 0, apperrors.NotFound("questions", "no questions found", destination)
	}

	var values []string
	var size int

	if count > int64(amount) {
		size = int(amount)

		values, err = c.db.SRandMemberN(ctx, destination, int64(size)).Result()
		if err != nil {
			return nil, 0, apperrors.Internal(err)
		}
	} else {
		size = int(count)

		values, err = c.db.SMembers(ctx, destination).Result()
		if err != nil {
			return nil, 0, apperrors.Internal(err)
		}
	}

	ids := make([]uuid.UUID, size)
	for i, v := range values {
		if err = ids[i].Scan(v); err != nil {
			return nil, 0, apperrors.Internal(err)
		}
		count++
	}

	return ids, size, nil
}

func (c *Cache) AddCachedIDs(ctx context.Context, metas []*questions.Meta, timeout time.Duration) error {
	metasMap := make(map[string][]uuid.UUID, len(metas)/4)

	for _, meta := range metas {
		k := meta.Key()
		metasMap[k] = append(metasMap[k], meta.ID)
	}

	for k, values := range metasMap {
		if err := c.db.SAdd(ctx, k, values).Err(); err != nil {
			return apperrors.Internal(err)
		}

		c.db.Expire(ctx, k, timeout)
	}

	return nil
}
