package db

import (
	"context"
	"github.com/QuizWars-Ecosystem/go-common/pkg/dbx"
	apperrors "github.com/QuizWars-Ecosystem/go-common/pkg/error"
	"github.com/QuizWars-Ecosystem/questions-service/internal/models/questions"
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

func (c *Client) GetCategories(ctx context.Context) ([]*questions.Category, error) {
	builder := dbx.StatementBuilder.
		Select("id", "name").
		From("categories")

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, apperrors.Internal(err)
	}

	rows, err := c.db.Query(ctx, query, args...)
	if err != nil {
		return nil, apperrors.Internal(err)
	}

	defer rows.Close()

	var categories []*questions.Category
	for rows.Next() {
		var category questions.Category

		err = rows.Scan(&category.ID, &category.Name)
		if err != nil {
			return nil, apperrors.Internal(err)
		}

		categories = append(categories, &category)
	}

	if err = rows.Err(); err != nil {
		return nil, apperrors.Internal(err)
	}

	return categories, nil
}
