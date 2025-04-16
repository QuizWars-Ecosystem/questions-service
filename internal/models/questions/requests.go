package questions

import (
	"crypto/md5"
	"errors"
	"github.com/QuizWars-Ecosystem/go-common/pkg/abstractions"
	apperrors "github.com/QuizWars-Ecosystem/go-common/pkg/error"
	questionsv1 "github.com/QuizWars-Ecosystem/questions-service/gen/external/questions/v1"
	"github.com/gofrs/uuid"
	"strings"
	"time"

	pgxuuid "github.com/jackc/pgx/pgtype/ext/gofrs-uuid"
)

var _ abstractions.Requestable[Option, *questionsv1.Option] = (*Option)(nil)

func (o *Option) Request(req *questionsv1.Option) (*Option, error) {
	id, err := uuid.FromString(req.Id)
	if err != nil {
		return nil, err
	}

	o.ID = pgxuuid.UUID{UUID: id}
	o.Text = req.Text
	o.IsCorrect = req.IsCorrect

	return o, nil
}

var _ abstractions.Requestable[CreateQuestionRequest, *questionsv1.CreateQuestionRequest] = (*CreateQuestionRequest)(nil)

type CreateQuestionRequest struct {
	*Hashed
}

func (c CreateQuestionRequest) Request(req *questionsv1.CreateQuestionRequest) (*CreateQuestionRequest, error) {
	if req.Text == "" {
		return nil, apperrors.BadRequest(errors.New("text is required"))
	}

	text := strings.ToLower(strings.Trim(strings.TrimSpace(req.GetText()), " "))
	hash := md5.Sum([]byte(text))

	id, err := uuid.NewV4()
	if err != nil {
		return nil, apperrors.Internal(err)
	}

	var q = Question{
		ID:         pgxuuid.UUID{UUID: id},
		Type:       TypeFromGRPCEnum(req.Type),
		Source:     Text,
		Difficulty: DifficultyFromGRPCEnum(req.Difficulty),
		Category:   Category{ID: req.CategoryId},
		Text:       text,
		Language:   req.Language,
		CreatedAt:  time.Now(),
	}

	if len(req.Options) == 0 {
		return nil, apperrors.BadRequest(errors.New("question must contain at least one option"))
	}

	var options = make([]*Option, len(req.Options))
	for i, o := range req.Options {
		var option = &Option{}

		if option, err = option.Request(o); err != nil {
			return nil, err
		}

		options[i] = option
	}

	q.Options = options

	c.Hashed = &Hashed{
		Hash:     string(hash[:]),
		Question: q,
	}

	return &c, nil
}
