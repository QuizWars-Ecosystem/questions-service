package db

import (
	"context"
	"github.com/Masterminds/squirrel"
	"github.com/QuizWars-Ecosystem/go-common/pkg/dbx"
	apperrors "github.com/QuizWars-Ecosystem/go-common/pkg/error"
	"github.com/QuizWars-Ecosystem/questions-service/internal/models/admin"
	"github.com/QuizWars-Ecosystem/questions-service/internal/models/questions"
	uuid "github.com/jackc/pgx/pgtype/ext/gofrs-uuid"
	"github.com/jackc/pgx/v5"
)

func (db *Database) GetFilteredQuestions(ctx context.Context, filter *admin.QuestionsFilter) ([]*questions.Question, int, error) {
	builder := squirrel.StatementBuilder.
		Select("q.id", "q.text", "c.id", "c.name", "o.id", "o.text", "o.is_correct", "q.type", "q.source", "q.difficulty", "q.language", "q.created_at").
		From("questions q").
		Join("categories c ON c.id = q.category_id").
		LeftJoin("question_options o ON o.question_id = q.id").
		OrderBy(filter.Order.String() + " " + filter.Sort.String()).
		Limit(filter.Limit).
		Offset(filter.Offset)

	if filter.TypesFilter != nil {
		builder = builder.
			Where(squirrel.Eq{"q.type": filter.TypesFilter.Array})
	}

	if filter.DifficultiesFilter != nil {
		builder = builder.
			Where(squirrel.Eq{"q.difficulty": filter.DifficultiesFilter.Array})
	}

	if filter.CategoriesFilter != nil {
		builder = builder.
			Where(squirrel.Eq{"q.category_id": filter.CategoriesFilter.Array})
	}

	if filter.LanguagesFilter != nil {
		builder = builder.
			Where(squirrel.Eq{"q.language": filter.LanguagesFilter.Array})
	}

	if filter.CreatedAtFilter != nil {
		builder = builder.
			Where(squirrel.GtOrEq{"q.created_at": filter.CreatedAtFilter.From}).
			Where(squirrel.LtOrEq{"q.created_at": filter.CreatedAtFilter.To})
	}

	countQuery := dbx.StatementBuilder.Select("COUNT(*)").From("questions")

	b := &pgx.Batch{}

	if err := dbx.QueryBatch(b, builder); err != nil {
		return nil, 0, apperrors.Internal(err)
	}

	if err := dbx.QueryBatch(b, countQuery); err != nil {
		return nil, 0, apperrors.Internal(err)
	}

	br := db.pool.SendBatch(ctx, b)
	defer func() {
		_ = br.Close()
	}()

	rows, err := br.Query()
	if err != nil {
		return nil, 0, apperrors.Internal(err)
	}

	defer rows.Close()

	questionsMap := make(map[uuid.UUID]*questions.Question, filter.Limit)
	optionsMap := make(map[uuid.UUID][]*questions.Option, filter.Limit*4)

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
			return nil, 0, apperrors.Internal(err)
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

	if rows.Err() != nil {
		return nil, 0, apperrors.Internal(err)
	}

	var total int
	if err = br.QueryRow().Scan(&total); err != nil {
		return nil, 0, apperrors.Internal(err)
	}

	return qs, total, nil
}

func (db *Database) SaveQuestion(ctx context.Context, question *questions.Hashed) error {
	builder := dbx.StatementBuilder.
		Insert("questions").
		Columns("id", "text", "text_hash", "category_id", "type", "source", "difficulty", "language", "created_at").
		Values(question.ID, question.Text, question.Hash, question.Category.ID, question.Type, question.Source, question.Difficulty, question.Language, question.CreatedAt).
		Suffix("ON CONFLICT DO NOTHING")

	query, args, err := builder.ToSql()
	if err != nil {
		return apperrors.Internal(err)
	}

	_, err = db.pool.Exec(ctx, query, args...)
	switch {
	case dbx.IsForeignKeyViolation(err, "category_id"):
		return apperrors.BadRequestHidden(err, "specified category was not found")
	case dbx.NotValidEnumType(err, "difficulty"):
		return apperrors.BadRequestHidden(err, "invalid question difficulty")
	case dbx.NotValidEnumType(err, "type"):
		return apperrors.BadRequestHidden(err, "invalid question type")
	case dbx.NotValidEnumType(err, "source"):
		return apperrors.BadRequestHidden(err, "invalid question source")
	case dbx.NotValidEnumType(err, "language"):
		return apperrors.BadRequestHidden(err, "invalid question language")
	case err != nil:
		return apperrors.Internal(err)
	}

	return nil
}

func (db *Database) UpdateCategory(ctx context.Context, id int32, name string) error {
	builder := dbx.StatementBuilder.
		Update("categories").
		Set("name", name).
		Where(squirrel.Eq{"id": id})

	query, args, err := builder.ToSql()
	if err != nil {
		return apperrors.Internal(err)
	}

	cmd, err := db.pool.Exec(ctx, query, args...)
	switch {
	case err == nil && cmd.RowsAffected() == 0:
		return apperrors.NotFound("category", "id", id)
	case err != nil:
		return apperrors.Internal(err)
	}

	return nil
}

func (db *Database) UpdateQuestion(ctx context.Context, id uuid.UUID, req *admin.UpdateQuestionRequest) error {
	builder := dbx.StatementBuilder.
		Update("questions").
		Where(squirrel.Eq{"id": id})

	flag := false

	if req.Type != nil {
		builder = builder.Set("type", req.Type)
		flag = true
	}

	if req.Difficulty != nil {
		builder = builder.Set("difficulty", req.Difficulty)
		flag = true
	}

	if req.Source != nil {
		builder = builder.Set("source", req.Source)
		flag = true
	}

	if req.CategoryID != nil {
		builder = builder.Set("category_id", req.CategoryID)
		flag = true
	}

	if req.Text != nil {
		builder = builder.
			Set("text", req.Text).
			Set("text_hash", req.Hash)
		flag = true
	}

	if req.Language != nil {
		builder = builder.Set("language", req.Language)
		flag = true
	}

	if !flag {
		return apperrors.BadRequest(argumentsNotProvidedErr)
	}

	query, args, err := builder.ToSql()
	if err != nil {
		return apperrors.Internal(err)
	}

	cmd, err := db.pool.Exec(ctx, query, args...)
	switch {
	case err == nil && cmd.RowsAffected() == 0:
		return apperrors.NotFound("question", "id", id)
	case dbx.IsForeignKeyViolation(err, "category_id"):
		return apperrors.BadRequestHidden(err, "specified category was not found")
	case dbx.NotValidEnumType(err, "difficulty"):
		return apperrors.BadRequestHidden(err, "invalid question difficulty")
	case dbx.NotValidEnumType(err, "type"):
		return apperrors.BadRequestHidden(err, "invalid question type")
	case dbx.NotValidEnumType(err, "source"):
		return apperrors.BadRequestHidden(err, "invalid question source")
	case dbx.NotValidEnumType(err, "language"):
		return apperrors.BadRequestHidden(err, "invalid question language")
	case err != nil:
		return apperrors.Internal(err)
	}

	return nil
}

func (db *Database) UpdateQuestionOption(ctx context.Context, id uuid.UUID, req *admin.UpdateQuestionOptionRequest) error {
	builder := dbx.StatementBuilder.
		Update("question_options").
		Where(squirrel.Eq{"id": id})

	flag := false

	if req.Text != nil {
		builder = builder.Set("text", req.Text)
		flag = true
	}

	if req.IsCorrect != nil {
		builder = builder.Set("is_correct", req.IsCorrect)
		flag = true
	}

	if !flag {
		return apperrors.BadRequest(argumentsNotProvidedErr)
	}

	query, args, err := builder.ToSql()
	if err != nil {
		return apperrors.Internal(err)
	}

	cmd, err := db.pool.Exec(ctx, query, args...)
	switch {
	case err == nil && cmd.RowsAffected() == 0:
		return apperrors.NotFound("option", "id", id)
	case err != nil:
		return apperrors.Internal(err)
	}

	return nil
}

func (db *Database) DeleteQuestion(ctx context.Context, id uuid.UUID) error {
	builder := dbx.StatementBuilder.
		Delete("questions").
		Where(squirrel.Eq{"id": id})

	query, args, err := builder.ToSql()
	if err != nil {
		return apperrors.Internal(err)
	}

	cmd, err := db.pool.Exec(ctx, query, args...)

	switch {
	case err == nil && cmd.RowsAffected() == 0:
		return apperrors.NotFound("question", "id", id)
	case err != nil:
		return apperrors.Internal(err)
	}

	return nil
}

func (db *Database) DeleteQuestionOption(ctx context.Context, id uuid.UUID) error {
	builder := dbx.StatementBuilder.
		Delete("question_options").
		Where(squirrel.Eq{"id": id})

	query, args, err := builder.ToSql()
	if err != nil {
		return apperrors.Internal(err)
	}

	cmd, err := db.pool.Exec(ctx, query, args...)
	switch {
	case err == nil && cmd.RowsAffected() == 0:
		return apperrors.NotFound("option", "id", id)
	case err != nil:
		return apperrors.Internal(err)
	}

	return nil
}
