package cache

import (
	"context"
	"encoding/json"
	"errors"
	apperrors "github.com/QuizWars-Ecosystem/go-common/pkg/error"
	"github.com/QuizWars-Ecosystem/questions-service/internal/models/questions"
	"github.com/redis/go-redis/v9"
	"time"
)

func (c *Cache) GetCategory(ctx context.Context, id int32) (*questions.Category, error) {
	data, err := c.db.Get(ctx, categoryKey(id)).Bytes()

	switch {
	case errors.Is(err, redis.Nil):
		return nil, apperrors.NotFound("category", "id", id)
	case err != nil:
		return nil, apperrors.Internal(err)
	}

	var category questions.Category
	if err = json.Unmarshal(data, &category); err != nil {
		return nil, apperrors.Internal(err)
	}

	return &category, nil
}

func (c *Cache) AddCategory(ctx context.Context, category *questions.Category) error {
	data, err := json.Marshal(category)
	if err != nil {
		return apperrors.Internal(err)
	}

	if err = c.db.Set(ctx, categoryKey(category.ID), data, time.Hour*24).Err(); err != nil {
		return apperrors.Internal(err)
	}

	return nil
}
