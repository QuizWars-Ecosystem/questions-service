package db

import (
	"context"
	"github.com/QuizWars-Ecosystem/go-common/pkg/dbx"
	apperrors "github.com/QuizWars-Ecosystem/go-common/pkg/error"
	"github.com/QuizWars-Ecosystem/questions-service/internal/models/questions"
)

func (db *Database) GetCategories(ctx context.Context) ([]*questions.Category, error) {
	builder := dbx.StatementBuilder.
		Select("id", "name").
		From("categories")

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, apperrors.Internal(err)
	}

	rows, err := db.pool.Query(ctx, query, args...)
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
