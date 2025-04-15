package db

import (
	"context"

	"github.com/Masterminds/squirrel"
	"github.com/QuizWars-Ecosystem/go-common/pkg/dbx"
	apperrors "github.com/QuizWars-Ecosystem/go-common/pkg/error"
	"github.com/QuizWars-Ecosystem/questions-service/internal/models/filter"
	"github.com/QuizWars-Ecosystem/questions-service/internal/models/questions"
	uuid "github.com/jackc/pgx/pgtype/ext/gofrs-uuid"
)

func (db *Database) GetRandomQuestionMeta(ctx context.Context, amount int64) ([]*questions.Meta, error) {
	builder := dbx.StatementBuilder.
		Select("q.id", "q.language", "q.difficulty", "q.category_id").
		From("questions q").
		GroupBy("q.id").
		OrderBy("RANDOM()").
		Limit(uint64(amount))

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, apperrors.Internal(err)
	}

	rows, err := db.pool.Query(ctx, query, args...)
	if err != nil {
		return nil, apperrors.Internal(err)
	}

	defer rows.Close()

	var metas []*questions.Meta

	for rows.Next() {
		var meta questions.Meta

		if err = rows.Scan(&meta.ID, &meta.Language, &meta.Difficulty, &meta.CategoryID); err != nil {
			return nil, apperrors.Internal(err)
		}

		metas = append(metas, &meta)
	}

	if err = rows.Err(); err != nil {
		return nil, apperrors.Internal(err)
	}

	return metas, nil
}

func (db *Database) GetQuestionsByIDs(ctx context.Context, IDs []uuid.UUID) ([]*questions.Question, error) {
	builder := squirrel.StatementBuilder.
		Select("q.id", "q.text", "c.id", "c.name", "o.id", "o.text", "o.is_correct", "q.type", "q.source", "q.difficulty", "q.language", "q.created_at").
		From("questions q").
		Join("categories c ON c.id = q.category_id").
		LeftJoin("options o ON o.question_id = q.id").
		Where(squirrel.Eq{"q.id": IDs}).
		OrderBy("q.id")

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, apperrors.Internal(err)
	}

	rows, err := db.pool.Query(ctx, query, args...)
	if err != nil {
		return nil, apperrors.Internal(err)
	}

	defer rows.Close()

	questionsMap := make(map[uuid.UUID]*questions.Question, len(IDs))
	optionsMap := make(map[uuid.UUID][]*questions.Option, len(IDs)*4)

	for rows.Next() {
		var question questions.Question
		var option questions.Option

		if err = rows.Scan(
			&question.ID,
			&question.Text,
			&question.Category.ID,
			&question.Category.Name,
			&option.ID,
			&option.Text,
			&option.IsCorrect,
			&question.Type,
			&question.Source,
			&question.Difficulty,
			&question.Language,
			&question.CreatedAt,
		); err != nil {
			return nil, apperrors.Internal(err)
		}

		if _, ok := questionsMap[question.ID]; !ok {
			questionsMap[question.ID] = &question
		}

		optionsMap[question.ID] = append(optionsMap[question.ID], &option)
	}

	qs := make([]*questions.Question, 0, len(questionsMap))
	for questionID, question := range questionsMap {
		question.Options = optionsMap[questionID]
		qs = append(qs, question)
	}

	if err = rows.Err(); err != nil {
		return nil, apperrors.Internal(err)
	}

	return qs, nil
}

func (db *Database) GetFilteredRandomQuestions(ctx context.Context, filter *filter.QuestionsFilter) ([]*questions.Question, error) {
	cteBuilder := dbx.StatementBuilder.
		Select("q.id").
		From("questions q").
		Where(squirrel.And{
			squirrel.Eq{"q.type": filter.Types},
			squirrel.Eq{"q.source": filter.Sources},
			squirrel.Eq{"q.language": filter.Language},
			squirrel.Eq{"q.category_id": filter.Categories},
			squirrel.Eq{"q.difficulty": filter.Difficulties},
			squirrel.Eq{"q.language": filter.Language},
		}).
		OrderBy("RANDOM()").
		Limit(uint64(filter.Amount))

	cteQuery, cteArgs, err := cteBuilder.ToSql()
	if err != nil {
		return nil, apperrors.Internal(err)
	}

	builder := dbx.StatementBuilder.
		Select("q.id", "q.text", "c.id", "c.name", "o.id", "o.text", "o.is_correct", "q.type", "q.source", "q.difficulty", "q.language", "q.created_at").
		From("questions q").
		Join("random_questions r ON q.id = r.id").
		Join("categories c ON c.id = q.category_id").
		Prefix("WITH random_questions AS (" + cteQuery + ")")

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, apperrors.Internal(err)
	}

	rows, err := db.pool.Query(ctx, query, append(cteArgs, args...)...)
	if err != nil {
		return nil, apperrors.Internal(err)
	}

	defer rows.Close()

	questionsMap := make(map[uuid.UUID]*questions.Question, filter.Amount)
	optionsMap := make(map[uuid.UUID][]*questions.Option, filter.Amount*4)

	for rows.Next() {
		var question questions.Question
		var option questions.Option

		if err = rows.Scan(
			&question.ID,
			&question.Text,
			&question.Category.ID,
			&question.Category.Name,
			&option.ID,
			&option.Text,
			&option.IsCorrect,
			&question.Type,
			&question.Source,
			&question.Difficulty,
			&question.Language,
			&question.CreatedAt,
		); err != nil {
			return nil, apperrors.Internal(err)
		}

		if _, ok := questionsMap[question.ID]; !ok {
			questionsMap[question.ID] = &question
		}

		optionsMap[question.ID] = append(optionsMap[question.ID], &option)
	}

	qs := make([]*questions.Question, 0, len(questionsMap))
	for questionID, question := range questionsMap {
		question.Options = optionsMap[questionID]
		qs = append(qs, question)
	}

	if err = rows.Err(); err != nil {
		return nil, apperrors.Internal(err)
	}

	return qs, nil
}
