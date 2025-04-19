package questions

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"strings"
	"time"

	"github.com/QuizWars-Ecosystem/go-common/pkg/abstractions"
	apperrors "github.com/QuizWars-Ecosystem/go-common/pkg/error"
	questionsv1 "github.com/QuizWars-Ecosystem/questions-service/gen/external/questions/v1"
	"github.com/google/uuid"
)

var _ abstractions.Requestable[Option, *questionsv1.Option] = (*Option)(nil)

func (o *Option) Request(req *questionsv1.Option) (*Option, error) {
	var id uuid.UUID
	var err error

	if req.Id == "" {
		id = uuid.New()
	} else {
		if id, err = uuid.Parse(req.Id); err != nil {
			return nil, apperrors.Internal(err)
		}
	}

	o.ID = id
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

	text := strings.TrimSpace(req.GetText())
	sum := md5.Sum([]byte(strings.ToLower(text)))
	hash := hex.EncodeToString(sum[:])

	q := Question{
		ID:         uuid.New(),
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

	var err error
	options := make([]*Option, len(req.Options))
	for i, o := range req.Options {
		option := &Option{}

		if option, err = option.Request(o); err != nil {
			return nil, err
		}

		options[i] = option
	}

	q.Options = options

	c.Hashed = &Hashed{
		Hash:     hash,
		Question: q,
	}

	return &c, nil
}
