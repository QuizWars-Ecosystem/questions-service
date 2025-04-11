package db

import (
	"context"
	"github.com/Masterminds/squirrel"
	"github.com/QuizWars-Ecosystem/go-common/pkg/dbx"
	apperrors "github.com/QuizWars-Ecosystem/go-common/pkg/error"
	"github.com/QuizWars-Ecosystem/questions-service/internal/models/filter"
	"github.com/QuizWars-Ecosystem/questions-service/internal/models/questions"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/lib/pq"
	"go.uber.org/zap"
)

type Questions struct {
	db     *pgxpool.Pool
	logger *zap.Logger
}

func NewQuestions(db *pgxpool.Pool, logger *zap.Logger) *Questions {
	return &Questions{
		db:     db,
		logger: logger,
	}
}

func (q *Questions) GetQuestionsByIDs(ctx context.Context, IDs []uuid.UUID) ([]*questions.Question, error) {
	builder := squirrel.StatementBuilder.
		Select("q.id", "q.text", "q.type", "q.source", "q.difficulty", "q.language", "q.created_at",
			"c.id AS category_id", "c.name AS category_name").
		From("questions q").
		Join("question_categories qc ON qc.question_id = q.id").
		Join("categories c ON c.id = qc.category_id").
		Where(squirrel.Eq{"q.id": IDs}).
		OrderBy("q.id")

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, apperrors.Internal(err)
	}

	rows, err := q.db.Query(ctx, query, args...)
	if err != nil {
		return nil, apperrors.Internal(err)
	}

	defer rows.Close()

	var qs []*questions.Question

	if err = rows.Err(); err != nil {
		return nil, apperrors.Internal(err)
	}

	return qs, nil
}

func (q *Questions) GetRandomQuestionMeta(ctx context.Context, amount int64) ([]*questions.Meta, error) {
	builder := dbx.StatementBuilder.
		Select("q.id", "q.language", "q.difficulty", "ARRAY_AGG(qc.category_id) AS categories_ids").
		From("questions q").
		LeftJoin("question_categories qc ON q.id = qc.question_id").
		GroupBy("q.id").
		OrderBy("RANDOM()").
		Limit(uint64(amount))

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, apperrors.Internal(err)
	}

	rows, err := q.db.Query(ctx, query, args...)
	if err != nil {
		return nil, apperrors.Internal(err)
	}

	defer rows.Close()

	var metas []*questions.Meta

	for rows.Next() {
		var meta questions.Meta
		var ids pq.Int32Array

		if err = rows.Scan(&meta.ID, &meta.Language, &meta.Difficulty, &ids); err != nil {
			return nil, apperrors.Internal(err)
		}

		meta.Categories = ids
		metas = append(metas, &meta)
	}

	if err = rows.Err(); err != nil {
		return nil, apperrors.Internal(err)
	}

	return metas, nil
}

func (q *Questions) GetFilteredRandomQuestions(ctx context.Context, filter *filter.QuestionsFilter) ([]*questions.Question, error) {
	cteBuilder := dbx.StatementBuilder.
		Select("q.id").
		From("qs q").
		Join("question_categories qc ON q.id = qc.question_id").
		Where(squirrel.And{
			squirrel.Eq{"q.type": filter.Types},
			squirrel.Eq{"q.source": filter.Sources},
			squirrel.Eq{"q.difficulty": filter.Difficulties},
			squirrel.Eq{"q.language": filter.Languages},
			squirrel.Eq{"qc.category_id": filter.Categories},
		}).
		OrderBy("RANDOM()").
		Limit(uint64(filter.Amount))

	cteQuery, cteArgs, err := cteBuilder.ToSql()
	if err != nil {
		return nil, apperrors.Internal(err)
	}

	builder := dbx.StatementBuilder.
		Select("q.id", "q.text", "q.type", "q.source", "q.difficulty", "q.language", "q.created_at",
			"c.id AS category_id", "c.name AS category_name").
		From("qs q").
		Join("random_questions r ON q.id = r.id").
		Join("question_categories qc ON q.id = qc.question_id").
		Join("categories c ON c.id = qc.category_id").
		Prefix("WITH random_questions AS (" + cteQuery + ")")

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, apperrors.Internal(err)
	}

	rows, err := q.db.Query(ctx, query, append(cteArgs, args...)...)
	if err != nil {
		return nil, apperrors.Internal(err)
	}

	defer rows.Close()

	var qs []*questions.Question
	qs, err = scanQuestionsWithCategories(rows)
	if err != nil {
		return nil, apperrors.Internal(err)
	}

	return qs, nil
}

func scanQuestionsWithCategories(rows pgx.Rows) ([]*questions.Question, error) {
	var qs []*questions.Question
	var currentQuestion *questions.Question
	var categoriesMap = make(map[uuid.UUID][]*questions.Category)

	for rows.Next() {
		var q questions.Question
		var categoryID int32
		var categoryName string
		err := rows.Scan(&q.ID, &q.Text, &q.Type, &q.Source, &q.Difficulty, &q.Language, &q.CreatedAt,
			&categoryID, &categoryName)
		if err != nil {
			return nil, err
		}

		if currentQuestion == nil || currentQuestion.ID != q.ID {
			if currentQuestion != nil {
				currentQuestion.Categories = categoriesMap[currentQuestion.ID]
				qs = append(qs, currentQuestion)
			}

			currentQuestion = &q
			categoriesMap = make(map[uuid.UUID][]*questions.Category)
		}

		categoriesMap[q.ID] = append(categoriesMap[q.ID], &questions.Category{
			ID:   categoryID,
			Name: categoryName,
		})
	}

	if currentQuestion != nil {
		currentQuestion.Categories = categoriesMap[currentQuestion.ID]
		qs = append(qs, currentQuestion)
	}

	return qs, nil
}
